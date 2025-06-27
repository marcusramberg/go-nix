package internal

import "fmt"

type Error string

func Panic(v string) {
	panic(Error(v))
}

func Panicf(format string, a ...any) {
	Panic(fmt.Sprintf(format, a...))
}

func Check(err error) {
	if err != nil {
		Panic(err.Error())
	}
}
