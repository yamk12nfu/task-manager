package utils

import "fmt"

func PanicToError(fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			e, ok := r.(error)
			if !ok {
				e = fmt.Errorf("%v", r)
			}
			err = e
		}
	}()

	err = fn()

	return
}
