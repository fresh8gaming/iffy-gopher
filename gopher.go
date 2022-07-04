package gopher

import "fmt"

type Gopher struct {
	Position    [2]int64
	Field       [2]int64
	Orientation string
}

func (g *Gopher) Move(command string) {
	if command == "F" {
		g.MoveForwards()
	} else if command == "T" {
		g.TurnClockwise()
	}
}

func (g *Gopher) MoveForwards() error {
	if g.Orientation == "N" {
		g.Position = [2]int64{g.Position[0], g.Position[1] + 1}
	} else if g.Orientation == "E" {
		g.Position = [2]int64{g.Position[0] + 1, g.Position[1]}
	} else if g.Orientation == "S" {
		g.Position = [2]int64{g.Position[0], g.Position[1] - 1}
	} else if g.Orientation == "W" {
		g.Position = [2]int64{g.Position[0] - 1, g.Position[1]}
	}
	if g.Position[0] > g.Field[0] || g.Position[1] > g.Field[1] {
		return fmt.Errorf("gopher in ditch")
	}

	return nil
}

func (g *Gopher) TurnClockwise() {
	if g.Orientation == "N" {
		g.Orientation = "E"
	} else if g.Orientation == "E" {
		g.Orientation = "S"
	} else if g.Orientation == "S" {
		g.Orientation = "W"
	} else if g.Orientation == "W" {
		g.Orientation = "N"
	}
}

func (g *Gopher) GetPositionX() int64 {
	return g.Position[0]
}

func (g *Gopher) GetPositionY() int64 {
	return g.Position[1]
}

func (g *Gopher) GetOrientation() string {
	return g.Orientation
}
