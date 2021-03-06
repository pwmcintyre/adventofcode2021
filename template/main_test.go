package main

import (
	"testing"

	"github.com/pwmcintyre/adventofcode2021/load"
)

const day = "2"

func Test_Sample1(t *testing.T) {

	lines, err := load.AsStrings(load.FromStrings(
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	))
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}

	want := 1
	if answer := part1(lines); answer != want {
		t.Errorf("wrong; got %d want %d", answer, want)
	}

}

func Test_part1(t *testing.T) {

	lines, err := load.AsStrings(load.Input(day))
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}

	want := 1
	if answer := part1(lines); answer != want {
		t.Errorf("wrong; got %d want %d", answer, want)
	}

}

func Test_Sample2(t *testing.T) {

	lines, err := load.AsStrings(load.FromStrings(
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	))
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}

	want := 1
	if answer := part2(lines); answer != want {
		t.Errorf("wrong; got %d want %d", answer, want)
	}
}

func Test_part2(t *testing.T) {

	lines, err := load.AsStrings(load.Input(day))
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}

	want := 1
	if answer := part2(lines); answer != want {
		t.Errorf("wrong; got %d want %d", answer, want)
	}

}
