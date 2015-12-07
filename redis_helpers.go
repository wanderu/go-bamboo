/* Copyright 2015 Christopher Kirkos

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

package bamboo

import (
	"gopkg.in/redis.v3"
	"strconv"
)

type ZScanResult struct {
	Result   string
	Priority float64
	Error    error
}

// ZScanIter. Iterates through the ZSet, making successsive ZSCAN calls until
// the redis cursor has been exhausted (end of the ZSet).
func ZScanIter(client *redis.Client, key string, match string) chan ZScanResult {
	ch := make(chan ZScanResult)

	go func() {
		var cursor int64 = 0
		defer close(ch)
		for {
			cursor, keys, err := client.ZScan(key, cursor, match, 0).Result()
			// Keys is an array of Queue keys and priorities
			// IE. [ job1 3 job2 5 ... ]

			if err != nil {
				ch <- ZScanResult{"", 0, err}
				return
			}

			for i := 1; i < len(keys); i += 2 {
				s := keys[i-1]
				p, err := strconv.ParseFloat(keys[i], 64)
				if err != nil {
					ch <- ZScanResult{"", 0, err}
					return
				}
				ch <- ZScanResult{s, p, nil}
			}

			if cursor == 0 {
				return
			}
		}
	}()

	return ch
}
