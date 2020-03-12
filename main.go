package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func f1() error {
	return errors.New("f1 error")
}

func f2() error {
	err := f1()
	return errors.WithMessage(err, "f2 error")
}

func main()  {
	err := f2()

	fmt.Println(errors.WithStack(err))
	fmt.Println(errors.Cause(err))
	fmt.Println(errors.Wrap(err, "wrap error"))
}
