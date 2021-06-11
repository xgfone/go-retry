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
	"testing"
)

func bench(context.Context, ...interface{}) (interface{}, error) { return true, nil }

func BenchmarkIntervalRetryWithoutArgs(b *testing.B) {
	retry := NewIntervalRetry(1, 0)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		retry.Call(context.Background(), bench)
	}
}

func BenchmarkIntervalRetryWithOneArg(b *testing.B) {
	retry := NewIntervalRetry(1, 0)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		retry.Call(context.Background(), bench, 1)
	}
}

func BenchmarkIntervalRetryWithTwoArgs(b *testing.B) {
	retry := NewIntervalRetry(1, 0)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		retry.Call(context.Background(), bench, 1, 2)
	}
}
