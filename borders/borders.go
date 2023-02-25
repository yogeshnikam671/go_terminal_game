package borders

import (
	"github.com/JoelOtter/termloop"
)

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

func RenderBorders(level *termloop.BaseLevel) {
    renderTopBorder(level)
    renderDeathBorder(level)
    renderLeftBorder(level)
    renderRightBorder(level)
}

