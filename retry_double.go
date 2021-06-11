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

// NewDoubleDelayRetry returns a new Retry to call it until the caller returns
// nil or is called for number times.
//
// If number is equal to 0, it won't retry it.
// If greater than 0, it is the maximum delay duration to retry it.
// start is the start delay to retry it, which must not be equal to 0.
// But, end may be equal to 0.
func NewDoubleDelayRetry(number int, start, end time.Duration) Retry {
	return doubleDelayRetry{num: number, start: start, end: end}
}

type doubleDelayRetry struct {
	start time.Duration
	end   time.Duration
	num   int
}

func (r doubleDelayRetry) Call(c context.Context, f Caller, a ...interface{}) (interface{}, error) {
	result, err := f(c, a...)

	for number, start := r.num, r.start; err != nil && number > 0; number-- {
		if err == ErrEndRetry {
			break
		}

		timer := time.NewTimer(start)
		select {
		case <-timer.C:
		case <-c.Done():
			timer.Stop()
			return result, err
		}

		if result, err = f(c, a...); err != nil {
			if start = start * 2; r.end > 0 && start > r.end {
				start = r.end
			}
		}
	}

	return result, err
}
