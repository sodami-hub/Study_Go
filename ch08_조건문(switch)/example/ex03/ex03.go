package main

import "fmt"

type Direction int

const (
	EAST Direction = iota + 1
	WEST
	SOUTH
	NORTH
	NONE
)

func GetDirection(angle float64) Direction {
	switch true {
	case angle >= 315:
		return NORTH
	case angle >= 0 && angle < 45:
		return NORTH
	case angle >= 45 && angle < 135:
		return EAST
	case angle >= 135 && angle < 225:
		return SOUTH
	case angle >= 225 && angle < 315:
		return WEST
	default:
		return NONE
	}
}

func DirectionToString(direction Direction) string {
	switch direction {
	case EAST:
		return "East"
	case WEST:
		return "West"
	case SOUTH:
		return "South"
	case NORTH:
		return "North"
	case NONE:
		return "None"
	}
	return ""
}

func main() {
	fmt.Println(DirectionToString(GetDirection(38.3)))
	fmt.Println(DirectionToString(GetDirection(235.8)))
	fmt.Println(DirectionToString(GetDirection(94.2)))
	fmt.Println(DirectionToString(GetDirection(-30)))
}
