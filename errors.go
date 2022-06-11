package main

import (
	"errors"
	"fmt"
)

/*
	it's idiomatice to communicate errors via an explicit, separate return value. This contrasts with exception
	used in languagees like Java and Rube and the overloaded single result / error value sometimes used in C. Go's approach makes it easy to see which functions
	return errors and to handle them using the same language constructs employed for any other, non-error taks
*/

//by convention errors are the last return value and have type error
func f1(arg int) (int, error) {
	if arg == 42 {

		//errors.New constructs a basic error value with the given error message.
		return -1, errors.New("can't work with 42")
	}

	//nil value in error indicates no error
	return arg + 3, nil
}

/*
possible to create custom error types
*/

type argError struct {
	arg  int
	prob string
}

//this allows for the use of custom types to explicitly represent an argument error
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		//in this case we use &argError syntax to build a new struct
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	/*
		these two loops test out error returning functions
	*/
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		//common idiom for an inline error check
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	//if you programmatically use the data in custom error
	//you need to get the instance via type asserstion
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
