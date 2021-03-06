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

const filename = "./input.txt"

func Input(day string) (io.ReadCloser, error) {

	// attempt to load local file
	lines, err := FromFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to load input: %w", err)
	}
	if err == nil {
		return lines, nil
	}

	// attempt to load from internet
	lines, err = FromAdventOfCode(day, DefaultAuthChain)
	if err != nil {
		return nil, fmt.Errorf("failed to get input from internet: %w", err)
	}

	// cache local
	outFile, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to save input from internet: %w", err)
	}

	// readCloser helps us bind the Reader interface of tee, and the Closer interface of File
	type readCloser struct {
		io.Reader
		io.Closer
	}
	tee := io.TeeReader(lines, outFile)
	return readCloser{tee, outFile}, nil
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
	if err != nil {
		return nil, fmt.Errorf("failed to get input: %w", err)
	}
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

func AsIntsFromBinary(r io.ReadCloser, err error) ([]int, error) {

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
		n, err := strconv.ParseInt(scanner.Text(), 2, 64)
		if err != nil {
			return lines, err
		}
		lines = append(lines, int(n))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// FromStrings is a utility to aid testsing, allows you to convert []string into io.ReadCloser
// to mimic data streaming from a file or URL
func FromStrings(data ...string) (*stringReadCloser, error) {
	return &stringReadCloser{data, 0}, nil
}

type stringReadCloser struct {
	data      []string
	readIndex int
}

func (r *stringReadCloser) Read(p []byte) (n int, err error) {
	if r.readIndex >= len(r.data) {
		err = io.EOF
		return
	}

	n = copy(p, r.data[r.readIndex]+"\n")
	r.readIndex += 1
	return
}
func (r *stringReadCloser) Close() error {
	return nil
}
