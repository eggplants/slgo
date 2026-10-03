// Command sl is a Go port of sl(1): a steam locomotive runs across your
// terminal when you mistype ls.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gdamore/tcell/v2"

	"github.com/eggplants/slgo/internal/sl"
)

const frameInterval = 40 * time.Millisecond

type tcellScreen struct{ s tcell.Screen }

func (t tcellScreen) Size() (int, int) { return t.s.Size() }

func (t tcellScreen) SetCell(x, y int, r rune) {
	t.s.SetContent(x, y, r, nil, tcell.StyleDefault)
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "sl:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	opts, err := sl.ParseArgs(args)
	if err != nil {
		return err
	}

	s, err := tcell.NewScreen()
	if err != nil {
		return err
	}
	if err := s.Init(); err != nil {
		return err
	}
	defer s.Fini()

	// Like the original, you cannot stop the train.
	signal.Ignore(os.Interrupt)
	s.HideCursor()

	// Drain and discard input so key presses do not pile up.
	go func() {
		for {
			switch s.PollEvent().(type) {
			case nil, *tcell.EventError:
				return
			}
		}
	}()

	scr := tcellScreen{s}
	if opts.Patch {
		cfg := sl.RandomPatchConfig()
		sl.NewPatch(scr, opts, cfg).Run(func() {
			s.Show()
			time.Sleep(cfg.Wait)
		})
		return nil
	}

	train := sl.New(scr, opts)
	for x := train.StartX(); train.Draw(x); x-- {
		s.Show()
		time.Sleep(frameInterval)
	}
	return nil
}
