/*
Copyright 2021 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_LinearRetryMax(t *testing.T) {
	t.Parallel()

	cases := []struct {
		desc              string
		config            LinearConfig
		previousCompareFn require.ComparisonAssertionFunc
	}{
		{
			desc: "HalfJitter",
			config: LinearConfig{
				First:  time.Second * 45,
				Step:   time.Second * 30,
				Max:    time.Minute,
				Jitter: NewJitter(),
			},
			previousCompareFn: require.NotEqual,
		},
		{
			desc: "SeventhJitter",
			config: LinearConfig{
				First:  time.Second * 45,
				Step:   time.Second * 30,
				Max:    time.Minute,
				Jitter: NewSeventhJitter(),
			},
			previousCompareFn: require.NotEqual,
		},
		{
			desc: "NoJitter",
			config: LinearConfig{
				First: time.Second * 45,
				Step:  time.Second * 30,
				Max:   time.Minute,
			},
			previousCompareFn: require.Equal,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			linear, err := NewLinear(tc.config)
			require.NoError(t, err)

			// artificially spike the attempts to get to max
			linear.attempt = 100

			// get the initial previous value to compare with
			previous := linear.Duration()
			linear.Inc()

			for i := 0; i < 50; i++ {
				duration := linear.Duration()
				linear.Inc()

				// ensure duration does not exceed maximum
				require.LessOrEqual(t, duration, tc.config.Max)

				// ensure duration comparison to previous is satisfied
				tc.previousCompareFn(t, duration, previous)

			}
		})
	}
}

// BenchmarkJitter is an attempt to check the effect of concurrency on the performance
// of a global jitter instance. I'm a bit skeptical of how "true to life" this benchmark
// really is, but the results would seem to indicate that >100k concurrent jitters would
// still all complete in <1s, which is very good for our purposes.
func BenchmarkSingleJitter(b *testing.B) {
	benchmarkSharedJitter(b, NewHalfJitter())
}

func BenchmarkShardedJitter(b *testing.B) {
	benchmarkSharedJitter(b, NewShardedHalfJitter())
}

func benchmarkSharedJitter(b *testing.B, jitter Jitter) {
	benchmarkJitter(b, func() Jitter { return jitter })
}

func benchmarkJitter(b *testing.B, mkjitter func() Jitter) {
	procs := runtime.GOMAXPROCS(0)
	for n := procs; n < 200_000; n = n * 2 {
		b.Run(fmt.Sprintf("n%d", n), func(b *testing.B) {
			b.SetParallelism(n / procs)
			b.RunParallel(func(pb *testing.PB) {
				jitter := mkjitter()
				for pb.Next() {
					jitter(time.Hour)
				}
			})
		})
	}
}
