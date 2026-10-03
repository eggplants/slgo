package sl

import (
	"math/rand/v2"
	"strconv"
	"strings"
	"time"
)

// Port of sl5-1.patch (2000) by Izumi (Tokyo Metropolitan Institute of
// Technology), with random parameters by Ueda and Akiyama (Osaka University):
// a level crossing, a D51 triple-header pulling passenger coaches, and an
// optional round trip.

// patchedD51Length is D51LENGTH as redefined by sl5-1.patch. It is longer
// than the drawn train, so the animation keeps running for a while after
// the train has passed.
const patchedD51Length = 172

const (
	// coalWidth is the width of the tender as shortened by sl5-1.patch.
	coalWidth = 29
	// engineWidth is how much of each D51 row is used before its tender.
	engineWidth = 53
	// unitWidth is the width of one D51 with its tender.
	unitWidth = engineWidth + coalWidth
)

// PatchConfig holds the parameters sl5-1.patch picks at random.
type PatchConfig struct {
	Coaches   int           // number of passenger coaches (PASSNUM)
	RoundTrip bool          // the train comes back after passing (ONEDIREC)
	Wait      time.Duration // delay between frames (WAIT_TIME)
}

// RandomPatchConfig picks parameters the same way as sl5-1.patch:
// 10-29 coaches, a round trip half of the time and a delay under 50ms.
func RandomPatchConfig() PatchConfig {
	return PatchConfig{
		Coaches:   rand.IntN(20) + 10,
		RoundTrip: rand.Float64() > 0.5,
		Wait:      time.Duration(rand.Float64()*50000) * time.Microsecond,
	}
}

// PatchTrain draws the sl5-1.patch animation.
type PatchTrain struct {
	t         *Train
	cfg       PatchConfig
	allLength int
	rows      [d51Patterns][d51Height + 1]string
	smokeR    smokeTrail
	crossX    int  // column of the level crossing
	crossTick int  // lamp blink counter while a train passes
	ltr       bool // the arrow on the crossing points right
}

// NewPatch returns a PatchTrain that draws onto scr. opts.C51 is ignored.
func NewPatch(scr Screen, opts Options, cfg PatchConfig) *PatchTrain {
	cfg.Coaches = max(cfg.Coaches, 1)
	cols, _ := scr.Size()
	h := &PatchTrain{
		t:         &Train{canvas: canvas{scr}, opts: opts, patched: true},
		cfg:       cfg,
		allLength: 3*patchedD51Length + passLength*(cfg.Coaches-1) + lPassLength,
		crossX:    3 * cols / 10,
	}
	for j := range d51Patterns {
		for i := range d51Height + 1 {
			h.rows[j][i] = h.buildRow(j, i)
		}
	}
	return h
}

// buildRow composes one row of the whole train running right to left.
func (h *PatchTrain) buildRow(pattern, i int) string {
	var b []byte
	for range 3 {
		b = append(b, d51[pattern][i][:engineWidth]...)
		b = append(b, d51Coal[i][:coalWidth]...)
	}
	for k := range h.cfg.Coaches {
		start := len(b)
		if k < h.cfg.Coaches-1 {
			b = append(b, coach[i]...)
		} else {
			b = append(b, lastCoach[i]...)
		}
		if i == 3 {
			copy(b[start+8:], strconv.Itoa(k+1))
		}
	}
	return string(b)
}

// Run plays the whole animation, calling frame after each frame is drawn.
func (h *PatchTrain) Run(frame func()) {
	opts := h.t.opts
	if !opts.Fly {
		h.beginGate(frame)
	}
	for x := h.t.StartX(); ; x-- {
		var ok bool
		if opts.Logo {
			ok = h.t.addSL(x)
		} else {
			ok = h.drawForward(x)
		}
		if !ok {
			break
		}
		if !opts.Fly {
			h.addCross()
		}
		frame()
	}
	if !opts.Fly && !opts.Logo && h.cfg.RoundTrip {
		h.reverseGate(frame)
		for x := 0; h.drawReverse(x); x++ {
			h.addCross()
			frame()
		}
	}
	if !opts.Fly {
		h.endGate(frame)
	}
}

// drawForward mirrors add_D51_coach.
func (h *PatchTrain) drawForward(x int) bool {
	if x < -h.allLength+4 {
		return false
	}
	cols, lines := h.t.scr.Size()
	y := lines/2 - 5
	if h.t.opts.Fly {
		y = x/7 + lines - cols/7 - d51Height
	}
	rows := &h.rows[(h.allLength+x)%d51Patterns]
	for i, row := range rows {
		// The original draws the row from column 0, padded with spaces.
		if x > 0 {
			h.t.addStr(y+i, 0, strings.Repeat(" ", x))
		}
		h.t.addStr(y+i, x, row)
	}
	if h.t.opts.Accident {
		for _, dx := range []int{43, 47, 125, 129, 207, 211} {
			h.t.addMan(y+2, x+dx)
		}
	}
	h.t.addSmoke(y-1, x+d51Funnel)
	h.t.addSmoke(y-1, x+d51Funnel+81)
	h.t.addSmoke(y-1, x+d51Funnel+162)
	return true
}

// drawReverse mirrors add_D51_coach_r: the train runs left to right with
// its front at column x.
func (h *PatchTrain) drawReverse(x int) bool {
	cols, lines := h.t.scr.Size()
	if x > h.allLength+cols {
		return false
	}
	y := lines/2 - 5
	// Wheel patterns run backwards so the wheels turn the other way.
	engine := &d51[(d51Patterns-(h.allLength+x)%d51Patterns)%d51Patterns]
	last := h.cfg.Coaches - 1
	for i := range d51Height + 1 {
		for u := range 3 {
			h.t.addStrR(y+i, x-unitWidth*u, engine[i])
			h.t.addStrR(y+i, x-unitWidth*u-engineWidth, d51Coal[i][:coalWidth])
		}
		for k := range h.cfg.Coaches {
			cx := x - 3*unitWidth - passLength*k
			if k < last {
				h.t.addStrR(y+i, cx, coach[i])
			} else {
				h.t.addStrR(y+i, cx, lastCoach[i])
			}
			if i == 3 {
				h.t.addStr(y+i, cx-9, strconv.Itoa(k+1))
			}
		}
	}
	if h.t.opts.Accident {
		for _, dx := range []int{45, 49, 127, 131, 209, 213} {
			h.t.addMan(y+2, x-dx)
		}
	}
	for _, dx := range []int{3, 84, 167} {
		if fx := x - d51Funnel - dx; fx <= 2*cols {
			h.smokeR.addR(h.t.canvas, y-1, fx)
		}
	}
	return true
}

func (h *PatchTrain) crossingY() int {
	_, lines := h.t.scr.Size()
	if h.t.opts.Logo {
		return lines/2 - 7
	}
	return lines/2 - 5
}

func (h *PatchTrain) drawPost(y int) {
	for i := 2; i < d51Height; i++ {
		h.t.addStr(y+i, crossingPostX[i]+h.crossX, crossingPost[i])
	}
}

func (h *PatchTrain) drawGate(y int, gate *[d51Height]string, from int) {
	for i := from; i < d51Height; i++ {
		h.t.addStr(y+i, h.crossX+5, gate[i])
	}
}

// drawLamps blinks the crossing lamps and shows the train's direction.
func (h *PatchTrain) drawLamps(y, tick int) {
	lx := crossingPostX[5] + h.crossX
	if tick%20 < 10 {
		h.t.addStr(y+4, lx-1, "O")
		h.t.addStr(y+4, lx+2, "X")
		h.t.addStr(y+5, lx-1, " || ")
		return
	}
	h.t.addStr(y+4, lx-1, "X")
	h.t.addStr(y+4, lx+2, "O")
	if h.ltr {
		h.t.addStr(y+5, lx-1, " -->")
	} else {
		h.t.addStr(y+5, lx-1, "<-- ")
	}
}

// addCross redraws the closed crossing on top of the passing train.
func (h *PatchTrain) addCross() {
	y := h.crossingY()
	h.drawPost(y)
	h.drawGate(y, &gateDown, 8)
	h.drawLamps(y, h.crossTick)
	h.crossTick++
}

func (h *PatchTrain) gatePhase(gate *[d51Height]string, frames int, frame func()) {
	y := h.crossingY()
	for tick := range frames {
		h.drawGate(y, gate, 0)
		h.drawPost(y)
		h.drawLamps(y, tick)
		frame()
	}
}

// beginGate mirrors begin_gate: the lamps start blinking and the gate closes.
func (h *PatchTrain) beginGate(frame func()) {
	h.gatePhase(&gateUp, 80, frame)
	h.gatePhase(&gateHalf, 16, frame)
	h.gatePhase(&gateDown, 21, frame)
}

// reverseGate mirrors x_gate: the gate starts to rise, then closes again
// for the train coming back.
func (h *PatchTrain) reverseGate(frame func()) {
	h.gatePhase(&gateDown, 21, frame)
	h.gatePhase(&gateHalf, 11, frame)
	h.ltr = !h.ltr
	h.gatePhase(&gateHalf, 11, frame)
	h.gatePhase(&gateDown, 21, frame)
}

// endGate mirrors end_gate: the gate opens and the lamps go off.
func (h *PatchTrain) endGate(frame func()) {
	h.gatePhase(&gateDown, 21, frame)
	h.gatePhase(&gateHalf, 16, frame)
	y := h.crossingY()
	lx := crossingPostX[5] + h.crossX
	for range 80 {
		h.drawGate(y, &gateUp, 0)
		h.drawPost(y)
		h.t.addStr(y+4, lx-1, "X")
		h.t.addStr(y+4, lx+2, "X")
		h.t.addStr(y+5, lx-1, " || ")
		frame()
	}
}
