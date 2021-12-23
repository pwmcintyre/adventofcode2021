package load

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func FromFile(path string) (io.ReadCloser, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	return os.Open(path)
}

func FromURL(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, err
}

func FromAdventOfCode(day string, fn SessionCodeGetter) (io.ReadCloser, error) {

	// get auth
	key, err := fn()
	if err != nil {
		return nil, fmt.Errorf("failed to get session code: %w", err)
	}

	// auth cookie
	authCookie := &http.Cookie{Name: "session", Value: key}

	// make request
	url := fmt.Sprintf("https://adventofcode.com/2021/day/%s/input", day)
	req, err := http.NewRequest("GET", url, nil)
	req.AddCookie(authCookie)

	// do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, err
}

func AsStrings(r io.ReadCloser, err error) ([]string, error) {

	if r != nil {
		defer r.Close()
	}

	var lines []string
	if err != nil {
		return lines, err
	}

	// scan
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func AsBytes2D(r io.ReadCloser, err error) ([][]byte, error) {

	if r != nil {
		defer r.Close()
	}

	var lines [][]byte
	if err != nil {
		return lines, err
	}

	// scan
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func AsInts(r io.ReadCloser, err error) ([]int, error) {

	if r != nil {
		defer r.Close()
	}

	var lines []int
	if err != nil {
		return lines, err
	}

	// scan
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return lines, err
		}
		lines = append(lines, n)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
