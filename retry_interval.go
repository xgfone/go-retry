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

package retry

import (
	"context"
	"time"
)

// NewIntervalRetry returns a new Retry to call it until the caller returns nil
// or is called for number times.
//
// If number is equal to 0, it won't retry it.
// If interval is 0, the caller will be called to retry immediately.
func NewIntervalRetry(number int, interval time.Duration) Retry {
	return intervalRetry{number: number, interval: interval}
}

type intervalRetry struct {
	interval time.Duration
	number   int
}

func (r intervalRetry) Call(c context.Context, f Caller, a ...interface{}) (interface{}, error) {
	result, err := f(c, a...)
	for number := r.number; err != nil && number > 0; number-- {
		if err == ErrEndRetry || waitForExit(c, r.interval) {
			break
		}
		result, err = f(c, a...)
	}
	return result, err
}

func waitForExit(ctx context.Context, sleep time.Duration) (exit bool) {
	if sleep > 0 {
		timer := time.NewTimer(sleep)
		select {
		case <-timer.C:
		case <-ctx.Done():
			timer.Stop()
			return true
		}
	} else {
		select {
		case <-ctx.Done():
			return true
		default:
		}
	}

	return false
}
