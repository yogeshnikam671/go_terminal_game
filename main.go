package main

import (
	"github.com/JoelOtter/termloop"
)

// :set tabstop=4
// :set shiftwidth=4
// :set expandtab

type Bar struct {
    x int;
    y int;
    width int;
    height int;
    color termloop.Attr;
}

func renderBar(level *termloop.BaseLevel) {
    bar := Bar { 56, 43, 30, 1, termloop.ColorRed }
    level.AddEntity(termloop.NewRectangle(bar.x, bar.y, bar.width, bar.height, bar.color))
}

func main() {
    game := termloop.NewGame()
    level := termloop.NewBaseLevel(termloop.Cell{
        Bg: termloop.ColorBlack,
        Fg: termloop.ColorBlack,
    })
    renderBar(level)

    game.Screen().SetLevel(level)
    game.Start()
}
