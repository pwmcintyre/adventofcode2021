package main

import (
	"testing"

	"github.com/pwmcintyre/adventofcode2021/load"
)

func Test_Sample1(t *testing.T) {

	lines := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	want := 7
	if answer := part1(lines); answer != want {
		t.Errorf("wrong; got %d want %d", answer, want)
	}

}

func Test_part1(t *testing.T) {

	lines, err := load.AsInts(load.Input("1"))
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}

	want := 1393
	if answer := part1(lines); answer != want {
		t.Errorf("wrong; got %d want %d", answer, want)
	}

}

func Test_Sample2(t *testing.T) {

	lines := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	want := 5
	if answer := part2(lines); answer != want {
		t.Errorf("wrong; got %d want %d", answer, want)
	}
}

func Test_part2(t *testing.T) {

	lines, err := load.AsInts(load.Input("1"))
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}

	want := 1359
	if answer := part2(lines); answer != want {
		t.Errorf("wrong; got %d want %d", answer, want)
	}

}
