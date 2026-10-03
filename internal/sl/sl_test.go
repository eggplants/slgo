package sl

import (
	"strings"
	"testing"
)

type fakeScreen struct {
	cols, lines int
	cells       [][]rune
}

func newFakeScreen(cols, lines int) *fakeScreen {
	cells := make([][]rune, lines)
	for y := range cells {
		cells[y] = []rune(strings.Repeat(" ", cols))
	}
	return &fakeScreen{cols: cols, lines: lines, cells: cells}
}

func (f *fakeScreen) Size() (int, int) { return f.cols, f.lines }

func (f *fakeScreen) SetCell(x, y int, r rune) {
	if x < 0 || y < 0 || x >= f.cols || y >= f.lines {
		panic("SetCell out of range")
	}
	f.cells[y][x] = r
}

func (f *fakeScreen) row(y int) string { return string(f.cells[y]) }

func TestParseArgs(t *testing.T) {
	tests := []struct {
		args []string
		want Options
	}{
		{nil, Options{}},
		{[]string{"-a"}, Options{Accident: true}},
		{[]string{"-alFc"}, Options{Accident: true, Logo: true, Fly: true, C51: true}},
		{[]string{"-F", "l", "-x", ""}, Options{Fly: true}},
		{[]string{"--c"}, Options{C51: true}},
		{[]string{"-p"}, Options{Patch: true}},
		{[]string{"-pp", "x"}, Options{Patch: true}},
	}
	for _, tt := range tests {
		got, err := ParseArgs(tt.args)
		if err != nil {
			t.Errorf("ParseArgs(%q) returned error: %v", tt.args, err)
		} else if got != tt.want {
			t.Errorf("ParseArgs(%q) = %+v, want %+v", tt.args, got, tt.want)
		}
	}
}

func TestParseArgsPatchExclusive(t *testing.T) {
	for _, args := range [][]string{
		{"-pa"},
		{"-p", "-F"},
		{"-l", "-p"},
		{"-cp"},
		{"-alFcp"},
	} {
		if _, err := ParseArgs(args); err == nil {
			t.Errorf("ParseArgs(%q) should return error", args)
		}
	}
}

func TestAddStrClipping(t *testing.T) {
	scr := newFakeScreen(5, 2)
	tr := New(scr, Options{})

	if !tr.addStr(0, -2, "abcd") {
		t.Error("addStr partially left of screen should succeed")
	}
	if got := scr.row(0); got != "cd   " {
		t.Errorf("row 0 = %q", got)
	}
	if tr.addStr(1, 3, "xyz") {
		t.Error("addStr overflowing right edge should fail")
	}
	if got := scr.row(1); got != "   xy" {
		t.Errorf("row 1 = %q", got)
	}
	if tr.addStr(0, -5, "abc") {
		t.Error("addStr entirely left of screen should fail")
	}
	if tr.addStr(-1, 0, "a") || tr.addStr(2, 0, "a") {
		t.Error("addStr outside vertical range should fail")
	}
}

func TestDrawD51Frame(t *testing.T) {
	scr := newFakeScreen(120, 30)
	tr := New(scr, Options{})
	const x = 1 // x+d51Funnel is a multiple of 4, so smoke is emitted
	if !tr.Draw(x) {
		t.Fatal("Draw returned false")
	}
	y := 30/2 - 5
	for i := 0; i <= d51Height; i++ {
		row := scr.row(y + i)
		if want := d51[(d51Length+x)%d51Patterns][i][:53]; row[x:x+53] != want {
			t.Errorf("engine row %d = %q, want %q", i, row[x:x+53], want)
		}
		if want := d51Coal[i]; row[x+53:x+53+len(want)] != want {
			t.Errorf("coal row %d = %q, want %q", i, row[x+53:x+53+len(want)], want)
		}
	}
	if got := scr.row(y - 1)[x+d51Funnel : x+d51Funnel+5]; got != "(   )" {
		t.Errorf("smoke = %q", got)
	}
}

func TestDrawFinishes(t *testing.T) {
	for _, opts := range []Options{
		{},
		{C51: true},
		{Logo: true},
		{Accident: true, Fly: true},
		{Accident: true, Fly: true, Logo: true},
		{Accident: true, Fly: true, C51: true},
	} {
		scr := newFakeScreen(80, 24)
		tr := New(scr, opts)
		frames := 0
		for x := tr.StartX(); tr.Draw(x); x-- {
			frames++
		}
		length := d51Length
		switch {
		case opts.Logo:
			length = logoLength
		case opts.C51:
			length = c51Length
		}
		if want := 80 + length; frames != want {
			t.Errorf("%+v: %d frames, want %d", opts, frames, want)
		}
	}
}

func TestAccident(t *testing.T) {
	scr := newFakeScreen(120, 30)
	tr := New(scr, Options{Accident: true})
	tr.Draw(0)
	var found bool
	for y := range scr.lines {
		if strings.Contains(scr.row(y), "(O)") || strings.Contains(scr.row(y), "\\O/") {
			found = true
		}
	}
	if !found {
		t.Error("no passengers drawn with Accident option")
	}
}
