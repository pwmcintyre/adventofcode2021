package load

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/zalando/go-keyring"
)

type SessionCodeGetter = func() (string, error)

var ErrNotFound = errors.New("session code not found")

func DefaultAuthChain() (string, error) {
	for _, fn := range []SessionCodeGetter{
		SessionCodeFromEnvVar,
		SessionCodeFromKeyring,
	} {

		// get key
		key, err := fn()
		if err != nil {
			if errors.Is(err, ErrNotFound) {
				continue
			}
			return "", err
		}

		return key, err
	}

	return "", fmt.Errorf("%w; exhausted all methods", ErrNotFound)
}

func SessionCodeFromEnvVar() (string, error) {
	key := os.Getenv("AOC_SESSION_KEY")
	if key == "" {
		return "", ErrNotFound
	}
	return key, nil
}

func SessionCodeFromUser() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Missing required session code!")
	fmt.Println("To find this, use Web Browser to load the page and authenticate ...")
	fmt.Println("https://adventofcode.com")
	fmt.Println("Then check dev console, under Cookie header. ")
	fmt.Print("Session code: ")
	key, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	key = strings.Trim(key, " \n")
	return key, nil
}

const (
	service = "AdventOfCode"
	user    = ""
)

func SessionCodeFromKeyring() (string, error) {

	// attempt to get from keyring
	key, err := keyring.Get(service, user)
	if err != nil {

		// not found in keyring
		if !errors.Is(err, keyring.ErrNotFound) {
			return "", err
		}

		// get from user
		key, err = SessionCodeFromUser()
		if err != nil {
			return "", err
		}

		// remember it
		err = keyring.Set(service, user, key)
		if err != nil {
			return "", err
		}

	}
	return key, nil
}
