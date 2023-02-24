package main

import (
	//"math/rand"

	"github.com/JoelOtter/termloop"
)

// :set tabstop=4
// :set shiftwidth=4
// :set expandtab

var isRightCollided bool = false

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

type DeathBorder struct {
    *termloop.Entity
    prevX int
    prevY int
    level *termloop.BaseLevel
}

func renderBorders(level *termloop.BaseLevel) {
    level.AddEntity(termloop.NewRectangle(0, 0, 500, 1, termloop.ColorBlue))
    level.AddEntity(termloop.NewRectangle(0, 0, 1, 500, termloop.ColorBlue))
    level.AddEntity(termloop.NewRectangle(161, 0, 1, 500, termloop.ColorBlue))
}

func renderDeathBorder(level *termloop.BaseLevel) {
    deathBorder := DeathBorder {
        Entity: termloop.NewEntity(1, 50, 500, 1),
        level: level,
    }
    deathBorder.Fill(&termloop.Cell { Bg: termloop.ColorRed})
    level.AddEntity(&deathBorder)
}

func renderBall(level *termloop.BaseLevel) {
    ball := Ball {
        Entity: termloop.NewEntity(1, 1, 2, 1),
        level: level,
    }
    ball.Fill(&termloop.Cell { Bg: termloop.ColorWhite})
    level.AddEntity(&ball)
}

func (ball *Ball) Tick(event termloop.Event) {
    ball.prevX, ball.prevY = ball.Position()
    var x int
    if(isRightCollided) {
        x = ball.prevX - 2 
    } else {
        x = ball.prevX + 2 
    }
    ball.SetPosition(x, ball.prevY + 1)
}

func (ball *Ball) Collide(collision termloop.Physical) {
    // Check if it's a Rectangle we're colliding with
    ball.prevX, ball.prevY = ball.Position()
    if _, ok := collision.(*termloop.Rectangle); ok {
        var x int
        if(ball.prevX > 55) {
            isRightCollided = true
            x = ball.prevX - 2 
        } else {
            isRightCollided = false
            x = ball.prevX + 2 
        }
        
        if(ball.prevY > 20) {
            ball.SetPosition(1,1)
            return
        }

        ball.SetPosition(x, ball.prevY + 5)
    }
}

func renderBar(level *termloop.BaseLevel) {
    bar := Bar {
       Rectangle: termloop.NewRectangle(65, 48, 30, 1, termloop.ColorYellow), 
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
    renderBorders(level)
    renderDeathBorder(level)
    renderBar(level)
    renderBall(level)
    
    game.Screen().SetLevel(level)
    game.Start()
}
