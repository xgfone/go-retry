# Retry [![Build Status](https://api.travis-ci.com/xgfone/go-retry.svg?branch=master)](https://travis-ci.com/github/xgfone/go-retry) [![GoDoc](https://pkg.go.dev/badge/github.com/xgfone/go-retry)](https://pkg.go.dev/github.com/xgfone/go-retry) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](https://raw.githubusercontent.com/xgfone/go-retry/master/LICENSE)

The package supporting `Go1.7+` supplies some retry policies to call a function.


## Installation
```shell
$ go get -u github.com/xgfone/go-retry
```

## Example
```go
package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/xgfone/go-retry"
)

var errnum int

func callErr(ctx context.Context, args ...interface{}) (interface{}, error) {
	errnum++
	return nil, errors.New("error")
}

func add(ctx context.Context, args ...interface{}) (interface{}, error) {
	return args[0].(int) + args[1].(int), nil
}

func sub(ctx context.Context, args ...interface{}) (interface{}, error) {
	return args[0].(int) - args[1].(int), nil
}

func main() {
	num1, num2 := 1, 2

	// Retry once if failing to call the add function.
	retry1 := retry.NewIntervalRetry(1, time.Second)
	result, err := retry1.Call(context.TODO(), add, num1, num2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d + %d = %v\n", num1, num2, result)
	}

	// Retry twice if failing to call the sub function.
	retry2 := retry.NewDoubleDelayRetry(2, time.Second, 0)
	result, err = retry2.Call(context.TODO(), sub, num2, num1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d - %d = %v\n", num2, num1, result)
	}

	// Retry three times if failing to call the function returning the error.
	retry3 := retry.NewIntervalRetry(3, time.Second)
	_, err = retry3.Call(context.Background(), callErr)
	fmt.Println(errnum, err)

	// Output:
	// 1 + 2 = 3
	// 2 - 1 = 1
	// 4 error
}
```
