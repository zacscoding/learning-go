package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type Data struct {
	retryable bool
	content   string
}

func main() {
	err := errors.New("origin error")

	errWrap := errors.Wrap(err, "wrap error")
	errWithMessage := errors.WithMessage(err, "message error")

	fmt.Println("> errors.Cause(errWrap) == err", errors.Cause(errWrap) == err)
	fmt.Println("> errWrap.Error() :" + errWrap.Error())
	fmt.Println("> errors.Cause(errWithMessage) == err", errors.Cause(errWithMessage) == err)
	fmt.Println("> errWithMessage.Error() :" + errWithMessage.Error())
	// Output
	//> errors.Cause(errWrap) == err true
	//> errWrap.Error() :wrap error: origin error
	//> errors.Cause(errWithMessage) == err true
	//> errWithMessage.Error() :message error: origin error
}
