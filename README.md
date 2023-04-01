# Retry [![Build Status](https://github.com/xgfone/go-retry/actions/workflows/go.yml/badge.svg)](https://github.com/xgfone/go-retry/actions/workflows/go.yml) [![GoDoc](https://pkg.go.dev/badge/github.com/xgfone/go-retry)](https://pkg.go.dev/github.com/xgfone/go-retry) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](https://raw.githubusercontent.com/xgfone/go-retry/master/LICENSE)

Provide some retry policies to call a function, supporting `Go1.7+`.


## Installation
```shell
$ go get -u github.com/xgfone/go-retry
```

## Example
```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/xgfone/go-retry"
)

func main() {
	num1, num2 := 1, 2
	var result int

	retry1 := retry.NewPeriodicIntervalRetry(1, time.Second)
	err := retry1.Run(context.TODO(), func(ctx context.Context) (success bool, err error) {
		result = num1 + num2
		return true, nil
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d + %d = %v\n", num1, num2, result)
	}

	// Output:
	// 1 + 2 = 3
}
```
