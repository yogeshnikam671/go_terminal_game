package main

import (
	"math/rand"
	"time"
	"github.com/JoelOtter/termloop"
	"github.com/yogeshnikam671/go_term_game/borders"
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

// Tick method is executed continuously by termloop for the provided pointer receiver type
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
    ballStartX := getRandomTravelValue()
    ballStartY := getRandomTravelValue()
    ball := Ball {
        Entity: termloop.NewEntity(ballStartX, ballStartY, 2, 1),
        level: level,
    }
    ball.Fill(&termloop.Cell { Bg: termloop.ColorWhite})
    level.AddEntity(&ball)
}

// Tick method is executed continuously by termloop for the provided pointer receiver type
// Here we are using this method to make the ball travel in the desired directions
func (ball *Ball) Tick(event termloop.Event) {
    if(isDead) {
        displayDeathScreen()
        return
    }
    ball.prevX, ball.prevY = ball.Position()
    x := getNextXPositionForTravel(ball.prevX)
    y := getNextYPositionForTravel(ball.prevY)
    ball.SetPosition(x, y)
}

func displayDeathScreen() {
    deathLevel := termloop.NewBaseLevel(termloop.Cell{
            Ch: ':',
    })
    game.Screen().SetLevel(deathLevel)
}

// collision handler methods
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
        x := getNextXPositionForCollisions(ball.prevX)
        y := getNextYPositionForCollisions(ball.prevY)
        ball.SetPosition(x, y)
    }
}

func (ball *Ball) handleTopBorderCollision(collision termloop.Physical) {
    ball.prevX, ball.prevY = ball.Position()
    if _, ok := collision.(*borders.TopBorder); ok {
        isBarCollided = false
        x := getNextXPositionForCollisions(ball.prevX)
        y := getNextYPositionForTravel(ball.prevY)
        ball.SetPosition(x, y)
    }
}

func (ball *Ball) handleLeftBorderCollision(collision termloop.Physical) {
    ball.prevX, ball.prevY = ball.Position()
    if _, ok := collision.(*borders.LeftBorder); ok {
        isRightCollided = false
        x := getNextXPositionForCollisions(ball.prevX)
        y := getNextYPositionForCollisions(ball.prevY)
        ball.SetPosition(x, y)
    }
}

func (ball *Ball) handleRightBorderCollision(collision termloop.Physical) {
    ball.prevX, ball.prevY = ball.Position()
    // execute this only if the collision is with RightBorder entity
    // This is essentially checking if collision is of type RightBorder. Refer - https://go.dev/tour/methods/15
    if _, ok := collision.(*borders.RightBorder); ok {
        isRightCollided = true
        x := getNextXPositionForCollisions(ball.prevX)
        y := getNextYPositionForCollisions(ball.prevY)
        ball.SetPosition(x, y)
    }
}

// This method is executed by termloop whenever there is any collision of 
// the provided pointer receiver variable (Ball in this case) with anything rendered on screen
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
        y = currentY - 2
    } else {
        y = currentY + 2
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

func getNextXPositionForCollisions(currentX int) int {
    var x int
    if(isRightCollided) {
        x = currentX - getRandomTravelValue()
    } else {
        x = currentX + getRandomTravelValue()
    }
    return x
}

func getNextXPositionForTravel(currentX int) int {
    var x int
    if(isRightCollided) {
        x = currentX - 2
    } else {
        x = currentX + 2
    }
    return x
}

func getRandomTravelValue() int {
    min := 1
    max := 10
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min) + min
}

func main() {
    // Turn this true if you want to add any game.Log("..") for debugging
    //game.SetDebugOn(true)
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
