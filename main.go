package main

import (
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

type TopBorder struct {
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

type LeftBorder struct {
    *termloop.Entity
    prevX int
    prevY int
    level *termloop.BaseLevel
}

type RightBorder struct {
    *termloop.Entity
    prevX int
    prevY int
    level *termloop.BaseLevel
}

func renderTopBorder(level *termloop.BaseLevel) {
    topBorder := TopBorder {
        Entity: termloop.NewEntity(0, 0, 500, 1),
        level: level,
    }
    topBorder.Fill(&termloop.Cell { Bg: termloop.ColorBlue })
    level.AddEntity(&topBorder)
}

func renderDeathBorder(level *termloop.BaseLevel) {
    deathBorder := DeathBorder {
        Entity: termloop.NewEntity(0, 50, 500, 1),
        level: level,
    }
    deathBorder.Fill(&termloop.Cell { Bg: termloop.ColorRed })
    level.AddEntity(&deathBorder)
}

func renderLeftBorder(level *termloop.BaseLevel) {
    leftBorder := LeftBorder {
        Entity: termloop.NewEntity(0, 0, 1, 500),
        level: level,
    }
    leftBorder.Fill(&termloop.Cell{ Bg: termloop.ColorBlue })
    level.AddEntity(&leftBorder)
}

func renderRightBorder(level *termloop.BaseLevel) {
    rightBorder := RightBorder {
        Entity: termloop.NewEntity(161, 0, 1, 500),
        level: level,
    }
    rightBorder.Fill(&termloop.Cell{ Bg: termloop.ColorBlue })
    level.AddEntity(&rightBorder)
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
        deathLevel := termloop.NewBaseLevel(termloop.Cell{
            Ch: ':',
        })
        game.Screen().SetLevel(deathLevel)
        return
    }
    ball.prevX, ball.prevY = ball.Position()
    x := getNextXPosition(ball.prevX)
    y := getNextYPositionForTravel(ball.prevY)
    ball.SetPosition(x, y)
}

func (ball *Ball) handleDeathBorderCollision(collision termloop.Physical) {
    if _, entityOk := collision.(*DeathBorder); entityOk {
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
    if _, ok := collision.(*TopBorder); ok {
        isBarCollided = false
        x := getNextXPosition(ball.prevX)
        ball.SetPosition(x, ball.prevY + 1)
    }
}

func (ball *Ball) handleLeftBorderCollision(collision termloop.Physical) {
    ball.prevX, ball.prevY = ball.Position()
    if _, ok := collision.(*LeftBorder); ok {
        isRightCollided = false
        y := getNextYPositionForCollisions(ball.prevY)
        ball.SetPosition(ball.prevX + 2, y)
    }
}

func (ball *Ball) handleRightBorderCollision(collision termloop.Physical) {
    ball.prevX, ball.prevY = ball.Position()
    if _, ok := collision.(*RightBorder); ok {
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
    renderTopBorder(level)
    renderDeathBorder(level)
    renderLeftBorder(level)
    renderRightBorder(level)
    renderBar(level)
    renderBall(level)
    
    game.Screen().SetLevel(level)
    game.Start()
}
