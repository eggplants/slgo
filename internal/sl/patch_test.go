package sl

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestAddStrR(t *testing.T) {
	scr := newFakeScreen(6, 1)
	c := canvas{scr}

	if !c.addStrR(0, 7, "ab(/[x") {
		t.Error("addStrR partially right of screen should succeed")
	}
	if got := scr.row(0); got != `  x]\)` {
		t.Errorf("row = %q", got)
	}
	if c.addStrR(0, 1, "123") {
		t.Error("addStrR overflowing left edge should fail")
	}
	if got := scr.row(0); got != `21x]\)` {
		t.Errorf("row = %q", got)
	}
}

func TestBuildRowNumbersCoaches(t *testing.T) {
	scr := newFakeScreen(80, 24)
	h := NewPatch(scr, PatchConfig{Coaches: 12})
	row := h.rows[0][3]
	if want := 3*unitWidth + passLength*11 + len(lastCoach[3]); len(row) != want {
		t.Fatalf("row length = %d, want %d", len(row), want)
	}
	for k := range 12 {
		start := 3*unitWidth + passLength*k + 8
		got := strings.TrimSpace(row[start : start+2])
		if want := strconv.Itoa(k + 1); got != want {
			t.Errorf("coach %d number = %q, want %q", k+1, got, want)
		}
	}
}

func TestPatchFrames(t *testing.T) {
	const cols, lines, coaches = 100, 30, 10
	allLength := 3*patchedD51Length + passLength*(coaches-1) + lPassLength
	gates := 80 + 16 + 21 + 21 + 16 + 80
	forward := cols + allLength - 4
	reverse := 21 + 11 + 11 + 21 + allLength + cols + 1

	tests := []struct {
		name string
		trip bool
		want int
	}{
		{"one way", false, gates + forward},
		{"round trip", true, gates + forward + reverse},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scr := newFakeScreen(cols, lines)
			h := NewPatch(scr, PatchConfig{Coaches: coaches, RoundTrip: tt.trip})
			frames := 0
			h.Run(func() { frames++ })
			if frames != tt.want {
				t.Errorf("%d frames, want %d", frames, tt.want)
			}
		})
	}
}

func TestPatchCrossing(t *testing.T) {
	scr := newFakeScreen(100, 30)
	h := NewPatch(scr, PatchConfig{Coaches: 10})
	frames := 0
	h.Run(func() {
		frames++
		if frames == 1 {
			y := 30/2 - 5
			if got := scr.row(y + 9)[h.crossX : h.crossX+8]; got != `\&||~|||` {
				t.Errorf("crossing at start = %q", got)
			}
			if got := scr.row(y + 4)[h.crossX : h.crossX+6]; got != "-O||X-" {
				t.Errorf("lamps at start = %q", got)
			}
		}
	})
	y := 30/2 - 5
	if got := scr.row(y + 4)[h.crossX : h.crossX+6]; got != "-X||X-" {
		t.Errorf("lamps at end = %q", got)
	}
	// After the train has gone only the crossing remains.
	for i := 0; i <= d51Height; i++ {
		row := scr.row(y + i)
		rest := row[:h.crossX] + row[h.crossX+25:]
		if strings.TrimSpace(rest) != "" {
			t.Errorf("row %d not cleared: %q", i, row)
		}
	}
}

func TestRandomPatchConfig(t *testing.T) {
	for range 1000 {
		c := RandomPatchConfig()
		if c.Coaches < 10 || c.Coaches > 29 {
			t.Fatalf("Coaches = %d", c.Coaches)
		}
		if c.Wait < 0 || c.Wait >= 50*time.Millisecond {
			t.Fatalf("Wait = %v", c.Wait)
		}
	}
}
