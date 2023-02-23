package main

import (
	"github.com/JoelOtter/termloop"
)

// :set tabstop=4
// :set shiftwidth=4
// :set expandtab
func main() {
    game := termloop.NewGame()
    level := termloop.NewBaseLevel(termloop.Cell{
        Bg: termloop.ColorBlack,
        Fg: termloop.ColorBlack,
    })
    level.AddEntity(termloop.NewRectangle(56, 43, 30, 1, termloop.ColorRed))

    game.Screen().SetLevel(level)
    game.Start()
}
