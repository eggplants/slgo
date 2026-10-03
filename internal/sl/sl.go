// Package sl renders the SL (steam locomotive) animation, ported from sl.c
// (SL version 5.03, Copyright 1993,1998,2014-2015 Toyoda Masashi).
package sl

import "errors"

// Screen is the drawing surface the animation is rendered onto.
// Cells that are not overwritten must keep their previous contents.
type Screen interface {
	Size() (cols, lines int)
	SetCell(x, y int, r rune)
}

// Options selects the train and its behavior.
type Options struct {
	Accident bool // -a: people cry for help
	Fly      bool // -F: the train flies
	Logo     bool // -l: draw the small SL logo train
	C51      bool // -c: draw C51 instead of D51
	Patch    bool // -p: sl5-1.patch mode (level crossing, coaches, round trip)
	Help     bool // -h, --help: print usage and exit
	Version  bool // -v, --version: print version and exit
}

// ParseArgs interprets command line arguments the same way as sl.c:
// every argument starting with '-' is scanned for option letters,
// unknown letters and other arguments are ignored.
// --help and --version are also accepted as long forms of -h and -v.
// -p cannot be combined with -a, -F, -l or -c unless -h or -v is given.
func ParseArgs(args []string) (Options, error) {
	var o Options
	for _, arg := range args {
		switch arg {
		case "--help":
			o.Help = true
			continue
		case "--version":
			o.Version = true
			continue
		}
		if len(arg) == 0 || arg[0] != '-' {
			continue
		}
		for _, c := range arg[1:] {
			switch c {
			case 'a':
				o.Accident = true
			case 'F':
				o.Fly = true
			case 'l':
				o.Logo = true
			case 'c':
				o.C51 = true
			case 'p':
				o.Patch = true
			case 'h':
				o.Help = true
			case 'v':
				o.Version = true
			}
		}
	}
	if !o.Help && !o.Version && o.Patch && (o.Accident || o.Fly || o.Logo || o.C51) {
		return Options{}, errors.New("-p cannot be combined with -a, -F, -l or -c")
	}
	return o, nil
}

// canvas provides the curses-like drawing primitives of sl.c.
type canvas struct {
	scr Screen
}

// addStr mirrors my_mvaddstr: characters left of the screen are skipped and
// drawing stops at the first character that does not fit.
func (c canvas) addStr(y, x int, s string) bool {
	cols, lines := c.scr.Size()
	i := 0
	for ; x < 0; x, i = x+1, i+1 {
		if i >= len(s) {
			return false
		}
	}
	for ; i < len(s); i, x = i+1, x+1 {
		if y < 0 || y >= lines || x >= cols {
			return false
		}
		c.scr.SetCell(x, y, rune(s[i]))
	}
	return true
}

// addStrR mirrors my_mvaddstr_r of sl5-1.patch: s is drawn right to left
// starting at column x, with brackets and slashes flipped.
func (c canvas) addStrR(y, x int, s string) bool {
	cols, lines := c.scr.Size()
	i := 0
	for ; x >= cols; x, i = x-1, i+1 {
		if i >= len(s) {
			return false
		}
	}
	for ; i < len(s); i, x = i+1, x-1 {
		if y < 0 || y >= lines || x < 0 {
			return false
		}
		c.scr.SetCell(x, y, mirror(s[i]))
	}
	return true
}

func mirror(b byte) rune {
	switch b {
	case '\\':
		return '/'
	case '/':
		return '\\'
	case '(':
		return ')'
	case ')':
		return '('
	case '[':
		return ']'
	case ']':
		return '['
	}
	return rune(b)
}

type smokeParticle struct {
	y, x       int
	ptrn, kind int
}

type smokeTrail struct {
	particles []smokeParticle
}

// add mirrors add_smoke: every 4 columns existing puffs drift and a new one
// is emitted at (y, x).
func (s *smokeTrail) add(c canvas, y, x int) {
	if x%4 != 0 {
		return
	}
	for i := range s.particles {
		p := &s.particles[i]
		c.addStr(p.y, p.x, smokeEraser[p.ptrn])
		p.y -= smokeDY[p.ptrn]
		p.x += smokeDX[p.ptrn]
		if p.ptrn < smokePatterns-1 {
			p.ptrn++
		}
		c.addStr(p.y, p.x, smoke[p.kind][p.ptrn])
	}
	s.emit(c, y, x)
}

// addR mirrors add_smoke_r of sl5-1.patch for a train running left to right.
func (s *smokeTrail) addR(c canvas, y, x int) {
	if x%4 != 0 {
		return
	}
	for i := range s.particles {
		p := &s.particles[i]
		c.addStrR(p.y, p.x, smokeEraser[p.ptrn])
		p.y -= smokeDY[p.ptrn]
		p.x -= smokeDX[p.ptrn]
		if p.ptrn < smokePatterns-1 {
			p.ptrn++
		}
		c.addStrR(p.y, p.x, smoke[p.kind][p.ptrn])
	}
	// The original draws the new puff unmirrored here.
	s.emit(c, y, x)
}

func (s *smokeTrail) emit(c canvas, y, x int) {
	kind := len(s.particles) % 2
	c.addStr(y, x, smoke[kind][0])
	s.particles = append(s.particles, smokeParticle{y: y, x: x, ptrn: 0, kind: kind})
}

// Train draws successive frames of the animation.
type Train struct {
	canvas
	opts  Options
	smoke smokeTrail

	// patched enables the tweak sl5-1.patch made to add_smoke.
	patched bool
}

// New returns a Train that draws onto scr.
func New(scr Screen, opts Options) *Train {
	return &Train{canvas: canvas{scr}, opts: opts}
}

// StartX returns the column of the first frame.
func (t *Train) StartX() int {
	cols, _ := t.scr.Size()
	return cols - 1
}

// Draw renders the frame for column x. It returns false once the train
// has completely left the screen.
func (t *Train) Draw(x int) bool {
	switch {
	case t.opts.Logo:
		return t.addSL(x)
	case t.opts.C51:
		return t.addC51(x)
	default:
		return t.addD51(x)
	}
}

func (t *Train) addSL(x int) bool {
	if x < -logoLength {
		return false
	}
	cols, lines := t.scr.Size()
	y := lines/2 - 3
	py1, py2, py3 := 0, 0, 0

	if t.opts.Fly {
		y = x/6 + lines - cols/6 - logoHeight
		py1, py2, py3 = 2, 4, 6
	}
	for i := 0; i <= logoHeight; i++ {
		t.addStr(y+i, x, logo[(logoLength+x)/3%logoPatterns][i])
		t.addStr(y+i+py1, x+21, logoCoal[i])
		t.addStr(y+i+py2, x+42, logoCar[i])
		t.addStr(y+i+py3, x+63, logoCar[i])
	}
	if t.opts.Accident {
		t.addMan(y+1, x+14)
		t.addMan(y+1+py2, x+45)
		t.addMan(y+1+py2, x+53)
		t.addMan(y+1+py3, x+66)
		t.addMan(y+1+py3, x+74)
	}
	t.addSmoke(y-1, x+logoFunnel)
	return true
}

func (t *Train) addD51(x int) bool {
	if x < -d51Length {
		return false
	}
	cols, lines := t.scr.Size()
	y := lines/2 - 5
	dy := 0

	if t.opts.Fly {
		y = x/7 + lines - cols/7 - d51Height
		dy = 1
	}
	for i := 0; i <= d51Height; i++ {
		t.addStr(y+i, x, d51[(d51Length+x)%d51Patterns][i])
		t.addStr(y+i+dy, x+53, d51Coal[i])
	}
	if t.opts.Accident {
		t.addMan(y+2, x+43)
		t.addMan(y+2, x+47)
	}
	t.addSmoke(y-1, x+d51Funnel)
	return true
}

func (t *Train) addC51(x int) bool {
	if x < -c51Length {
		return false
	}
	cols, lines := t.scr.Size()
	y := lines/2 - 5
	dy := 0

	if t.opts.Fly {
		y = x/7 + lines - cols/7 - c51Height
		dy = 1
	}
	for i := 0; i <= c51Height; i++ {
		t.addStr(y+i, x, c51[(c51Length+x)%c51Patterns][i])
		t.addStr(y+i+dy, x+55, c51Coal[i])
	}
	if t.opts.Accident {
		t.addMan(y+3, x+45)
		t.addMan(y+3, x+49)
	}
	t.addSmoke(y-1, x+c51Funnel)
	return true
}

func (t *Train) addMan(y, x int) {
	for i := 0; i < 2; i++ {
		t.addStr(y+i, x, man[(logoLength+x)/12%2][i])
	}
}

func (t *Train) addSmoke(y, x int) {
	if t.patched {
		if cols, _ := t.scr.Size(); x < -cols {
			return
		}
	}
	t.smoke.add(t.canvas, y, x)
}
