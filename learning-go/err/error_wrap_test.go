package err

import (
	"errors"
	"fmt"
	"testing"
)

func TestCommon(t *testing.T) {
	err := errors.New("param illegal")
	newErr := fmt.Errorf("param input err: %v", err)
	fmt.Println(newErr)
}

type MyError struct {
	Err error
	Msg string
}

func (e MyError) Error() string {
	return e.Err.Error() + e.Msg
}

func dealLogic() error {
	err := errors.New("param illegal")
	newErr := MyError{err, "param input err"}
	return newErr
}

func TestErrStruct(t *testing.T) {
	err := dealLogic()
	if e, ok := err.(MyError); ok {
		fmt.Println(e)
	}
}

func TestWrapError(t *testing.T) {
	e := errors.New("original err")
	w := fmt.Errorf("wrap err: %w", e)
	fmt.Println(w)
}

func TestUnWrapError(t *testing.T) {
	e := errors.New("original err")
	w := fmt.Errorf("wrap err: %w", e)
	fmt.Println(errors.Unwrap(w)) // wrap err: original err
}

func TestIsError(t *testing.T) {
	e := errors.New("original err")
	e1 := errors.New("myError")
	e2 := MyError{Err: e1}
	w := fmt.Errorf("wrap err: %w", e)
	fmt.Println(errors.Is(w, e))  // true
	fmt.Println(errors.Is(w, e1)) // false
	fmt.Println(errors.Is(w, e2)) // false
}

func TestAsError(t *testing.T) {
	e := errors.New("original err")
	e1 := errors.New("myError")
	e2 := MyError{Err: e1}
	w := fmt.Errorf("wrap err: %w", e)
	fmt.Println(errors.As(w, &e))  // true
	fmt.Println(errors.As(w, &e1)) // true
	fmt.Println(errors.As(w, &e2)) // false
}
