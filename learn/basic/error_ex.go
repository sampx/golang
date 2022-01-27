package basic

import (
	"errors"
	"fmt"
	"time"
)

//MyError struct
type MyError struct {
	When  time.Time
	Cause error
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.Cause)
}

func run() error {
	return &MyError{
		time.Now(),
		errors.New("emit macho dwarf: elf header corrupted"),
	}
}

func ErrTest() {
	err := run()
	if err != nil {
		fmt.Print(err)
	}
}
