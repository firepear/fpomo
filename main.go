package main

import (
	//"flag"
	"fmt"
	"log"
	"math"

	"github.com/gdamore/tcell/v2"
)

var (
	// screen dimensions, x and y
	dimx, dimy    int
	// currently active cell
	curx, cury    int
	// number of cells on screen
	cells         int
	// the time we'll be counting down, in seconds
	time          int
	// elapsed and remaining time
	telap, trem   int
	// frames pre second
	fps           int
	// frames per cell: how many color transitions for each cell
	// to fade from fgc to bgc
	fpc           int
	// sleep per frame
	spf           float64
	// foreground, background, and time display colors
	fgc, bgc, tdc []int32
	// foreground gradient: a slice of `fpc` precomputed styles
	// with foreground colors that transition from fgc to bgc
	fgg           []tcell.Style
	fpcStyle = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.NewRGBColor(230, 230, 200))

)

func calcScreenParams(s tcell.Screen) {
	// get our dimensions
	dimx, dimy = s.Size()
	// total number of cells: dimentions minus 6 for the time
	// display (` mm:ss`)
	cells = dimx * dimy - 6
	// total time divided by cells, divided by (60 divided by fps)
	// is frames per cell
	fpc = int(math.Floor(float64(cells) / float64(time) / (60.0 / float64(fps))))
	// use fpc to calculate fgg, based on the distances between r,
	// g, and b of fgc and bgc
	rdiff := float64(fgc[0] - bgc[0]) / float64(fpc)
	gdiff := float64(fgc[1] - bgc[1]) / float64(fpc)
	bdiff := float64(fgc[2] - bgc[2]) / float64(fpc)
	for i := 0; i < fpc; i++ {
		r := int32(math.Floor(float64(fgc[0]) + float64(i) * rdiff))
		g := int32(math.Floor(float64(fgc[1]) + float64(i) * gdiff))
		b := int32(math.Floor(float64(fgc[2]) + float64(i) * bdiff))
		style := tcell.StyleDefault.Foreground(tcell.NewRGBColor(r, g, b)).Background(tcell.NewRGBColor(bgc[0], bgc[1], bgc[2]))
		fgg = append(fgg, style)
	}
}

func main() {
	fps = 10
	fgc = []int32{0,255,0}
	bgc = []int32{128, 0, 0}
	tdc = []int32{0, 255, 255}
	time = 120
	curx = dimx
	cury = dimy

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.HideCursor()
	calcScreenParams(s)
	//s.SetStyle(fpcStyle)

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can die
		// without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	s.SetContent(curx, cury, 'X', nil, fpcStyle)
	s.SetContent(curx + 1, cury + 1, 'X', nil, fpcStyle)
	cury = cury + 2
	for _, r := range []rune(fmt.Sprintf("(%v %v) (%v %v) %v %v %v", dimx, dimy, curx, cury, cells, fpc, fgg[0])) {
		s.SetContent(curx, cury, r, nil, fpcStyle)
		curx++
	}
	cury++
	curx = 0
	for _, r := range []rune(fmt.Sprintf("%v %v", curx, cury)) {
		s.SetContent(curx, cury, r, nil, fpcStyle)
		curx++
	}
	cury++
	curx = 0
	s.Show()

	for _, style := range fgg {
		for _, r := range []rune(fmt.Sprintf("%v", style)) {
			s.SetContent(curx, cury, r, nil, style)
			curx++
		}
		cury++
		curx = 0
	}
	s.Show()

	for {
		// Update screen
		s.Show()
		// Poll event
		ev := s.PollEvent()
		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			calcScreenParams(s)
			// TODO 
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Rune() == 'Q' || ev.Rune() == 'q' {
				goto THEEND
			}
		}
	}
THEEND:
}
