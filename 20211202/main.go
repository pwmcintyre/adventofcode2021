package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Submarine struct {
	depth    int
	position int
}

func (s *Submarine) MoveForward(amount int) {
	s.position += amount
}

func (s *Submarine) MoveDown(amount int) {
	s.depth += amount
}

func (s *Submarine) MoveUp(amount int) {
	s.depth -= amount
}

func part1(lines []string) int {
	submarine := Submarine{}
	for _, line := range lines {

		// split
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			panic("should be 2 parts")
		}

		// parse amount
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		// parse command
		switch parts[0] {
		case "forward":
			submarine.MoveForward(n)
		case "down":
			submarine.MoveDown(n)
		case "up":
			submarine.MoveUp(n)
		default:
			panic(fmt.Errorf("unknown command: %s", parts[0]))

		}
	}
	return submarine.depth * submarine.position
}

type Submarine2 struct {
	depth    int
	position int
	aim      int
}

func (s *Submarine2) MoveForward(amount int) {
	s.position += amount
	s.depth += s.aim * amount
}

func (s *Submarine2) AimDown(amount int) {
	s.aim += amount
}

func (s *Submarine2) AimUp(amount int) {
	s.aim -= amount
}

func part2(lines []string) int {
	submarine := Submarine2{}
	for _, line := range lines {

		// split
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			panic("should be 2 parts")
		}

		// parse amount
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		// parse command
		switch parts[0] {
		case "forward":
			submarine.MoveForward(n)
		case "down":
			submarine.AimDown(n)
		case "up":
			submarine.AimUp(n)
		default:
			panic(fmt.Errorf("unknown command: %s", parts[0]))

		}
	}
	return submarine.depth * submarine.position
}
