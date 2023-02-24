package main

import (
	//"math/rand"

	"github.com/JoelOtter/termloop"
)

// :set tabstop=4
// :set shiftwidth=4
// :set expandtab

type Bar struct {
    *termloop.Rectangle
    prevX int
    prevY int
    level *termloop.BaseLevel
}

type Ball struct {
    *termloop.Entity
    prevX int
    prevY int
    level *termloop.BaseLevel
}

func renderBall(level *termloop.BaseLevel) {
    ball := Ball {
        Entity: termloop.NewEntity(0, 0, 2, 1),
        level: level,
    }
    ball.Fill(&termloop.Cell { Bg: termloop.ColorWhite})
    level.AddEntity(&ball)
}

func (ball *Ball) Tick(event termloop.Event) {
    ball.prevX, ball.prevY = ball.Position()
    ball.SetPosition(ball.prevX + 1, ball.prevY)
}

func renderBar(level *termloop.BaseLevel) {
    bar := Bar {
       Rectangle: termloop.NewRectangle(65, 48, 30, 1, termloop.ColorRed), 
       level: level,
    }
    level.AddEntity(&bar)
}

// Learn what this bar *Bar stands for in go lang
func (bar *Bar) Tick(event termloop.Event) {
   if(event.Type == termloop.EventKey) {  // If it is a keyboard event
        bar.prevX, bar.prevY = bar.Position()
        switch(event.Key) {
            case termloop.KeyArrowLeft : bar.SetPosition(bar.prevX - 2, bar.prevY)   
            case termloop.KeyArrowRight : bar.SetPosition(bar.prevX + 2, bar.prevY)   
            default: return
        }
   } 
}

func main() {
    game := termloop.NewGame()
    game.Screen().SetFps(30)
    level := termloop.NewBaseLevel(termloop.Cell{
        Bg: termloop.ColorBlack,
        Fg: termloop.ColorBlack,
    })
    renderBar(level)
    renderBall(level)

    game.Screen().SetLevel(level)
    game.Start()
}
