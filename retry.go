// Copyright 2021 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package retry implements some retry policies to call a function.
package retry

import (
	"context"
	"errors"
)

// ErrEndRetry is used by the callee to end to retry.
var ErrEndRetry = errors.New("end to retry")

// Caller is the caller function.
type Caller func(ctx context.Context, args ...interface{}) (result interface{}, err error)

// Retry is used to retry a function call when it returns an error.
type Retry interface {
	// Call calls the callee and returns its result, which will retry it
	// when the caller returns the error.
	//
	// Notice: the callee maybe return ErrEndRetry to end to retry,
	// and the implementation should support it.
	Call(ctx context.Context, callee Caller, args ...interface{}) (result interface{}, err error)
}
