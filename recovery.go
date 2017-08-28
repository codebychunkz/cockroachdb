package main

import "fmt"

func SimpleErrorRecovery(err *error) {
	if res := recover(); res != nil {
		*err = fmt.Errorf("Error occured: %s", res)
	}
}
