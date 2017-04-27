package sentry

import "fmt"

type Release struct {
	Version  string        `json:"version"`
	Ref      string        `json:"ref"`
}

type Deploy struct {
	Name        string        `json:"name"`
	Environment string        `json:"environment"`
}

type Error struct {
	Code int
	Body string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Sentry Error: %d %s", e.Code, e.Body)
}
