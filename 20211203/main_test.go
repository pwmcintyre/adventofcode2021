package main

import (
	"testing"

	"github.com/pwmcintyre/adventofcode2021/load"
)

const day = "3"

func Test_Sample1(t *testing.T) {

	lines, err := load.AsIntsFromBinary(load.FromStrings(
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	))
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}

	want := 198
	if answer := part1(lines); answer != want {
		t.Errorf("wrong; got %d want %d", answer, want)
	}

}

func Test_part1(t *testing.T) {

	lines, err := load.AsIntsFromBinary(load.Input(day))
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}

	want := 1393
	if answer := part1(lines); answer != want {
		t.Errorf("wrong; got %d want %d", answer, want)
	}

}

func Test_Sample2(t *testing.T) {

	lines, err := load.AsIntsFromBinary(load.FromStrings(
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	))
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}

	want := 5
	if answer := part2(lines); answer != want {
		t.Errorf("wrong; got %d want %d", answer, want)
	}
}

func Test_part2(t *testing.T) {

	lines, err := load.AsIntsFromBinary(load.Input(day))
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}

	want := 1359
	if answer := part2(lines); answer != want {
		t.Errorf("wrong; got %d want %d", answer, want)
	}

}
