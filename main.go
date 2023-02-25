package main

import (
        "github.com/yogeshnikam671/go_term_game/borders" 
	"github.com/JoelOtter/termloop"
)

// :set tabstop=4
// :set shiftwidth=4
// :set expandtab

var game *termloop.Game = termloop.NewGame()
var isRightCollided bool = false
var isBarCollided bool = false
var isDead bool = false 

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

func renderBar(level *termloop.BaseLevel) {
    bar := Bar {
       Rectangle: termloop.NewRectangle(65, 48, 30, 1, termloop.ColorYellow), 
       level: level,
    }
    level.AddEntity(&bar)
}

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

func renderBall(level *termloop.BaseLevel) {
    ball := Ball {
        Entity: termloop.NewEntity(1, 1, 2, 1),
        level: level,
    }
    ball.Fill(&termloop.Cell { Bg: termloop.ColorWhite})
    level.AddEntity(&ball)
}

func (ball *Ball) Tick(event termloop.Event) {
    if(isDead) {
        displayDeathScreen()
        return
    }
    ball.prevX, ball.prevY = ball.Position()
    x := getNextXPosition(ball.prevX)
    y := getNextYPositionForTravel(ball.prevY)
    ball.SetPosition(x, y)
}

func displayDeathScreen() {
    deathLevel := termloop.NewBaseLevel(termloop.Cell{
            Ch: ':',
    })
    game.Screen().SetLevel(deathLevel)
}

func (ball *Ball) handleDeathBorderCollision(collision termloop.Physical) {
    if _, entityOk := collision.(*borders.DeathBorder); entityOk {
        isDead = true
        return
    }
}

func (ball *Ball) handleBarCollision(collision termloop.Physical) {
    if _, entityOk := collision.(*Bar); entityOk {
        isBarCollided = true
        ball.prevX, ball.prevY = ball.Position()
        x := getNextXPosition(ball.prevX)
        ball.SetPosition(x, ball.prevY - 5)
    }
}

func (ball *Ball) handleTopBorderCollision(collision termloop.Physical) {
    ball.prevX, ball.prevY = ball.Position()
    if _, ok := collision.(*borders.TopBorder); ok {
        isBarCollided = false
        x := getNextXPosition(ball.prevX)
        ball.SetPosition(x, ball.prevY + 1)
    }
}

func (ball *Ball) handleLeftBorderCollision(collision termloop.Physical) {
    ball.prevX, ball.prevY = ball.Position()
    if _, ok := collision.(*borders.LeftBorder); ok {
        isRightCollided = false
        y := getNextYPositionForCollisions(ball.prevY)
        ball.SetPosition(ball.prevX + 2, y)
    }
}

func (ball *Ball) handleRightBorderCollision(collision termloop.Physical) {
    ball.prevX, ball.prevY = ball.Position()
    if _, ok := collision.(*borders.RightBorder); ok {
        isRightCollided = true
        y := getNextYPositionForCollisions(ball.prevY)
        ball.SetPosition(ball.prevX - 2, y)
    }
}

func (ball *Ball) Collide(collision termloop.Physical) {
    ball.handleBarCollision(collision) 
    ball.handleTopBorderCollision(collision)
    ball.handleDeathBorderCollision(collision)
    ball.handleLeftBorderCollision(collision)
    ball.handleRightBorderCollision(collision)
}

// helper methods
func getNextYPositionForCollisions(currentY int) int {
    var y int
    if(isBarCollided) {
        y = currentY - 5
    } else {
        y = currentY + 5
    }
    return y
}

func getNextYPositionForTravel(currentY int) int {
    var y int
    if(isBarCollided) {
        y = currentY - 1
    } else {
        y = currentY + 1
    }
    return y
}

func getNextXPosition(currentX int) int {
    var x int
    if(isRightCollided) {
        x = currentX - 2
    } else {
        x = currentX + 2
    }
    return x
}

func main() {
    game.SetDebugOn(true)
    game.Screen().SetFps(30)
    level := termloop.NewBaseLevel(termloop.Cell{
        Bg: termloop.ColorBlack,
        Fg: termloop.ColorBlack,
    })
    borders.RenderBorders(level)
    renderBar(level)
    renderBall(level)
    
    game.Screen().SetLevel(level)
    game.Start()
}
