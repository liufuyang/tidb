// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package types

// http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
func abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

// uintSizeTable is used as a table to do comparison to get uint length is faster than doing loop on division with 10
var uintSizeTable = [21]uint64{
	0, // redundant 0 here, so to make function lenOfUint64Fast to count from 1 and return i directly
	9, 99, 999, 9999, 99999,
	999999, 9999999, 99999999, 999999999, 9999999999,
	99999999999, 999999999999, 9999999999999, 99999999999999, 999999999999999,
	9999999999999999, 99999999999999999, 999999999999999999, 9999999999999999999,
	maxUint,
} // maxUint is 18446744073709551615 and it has 20 digits

// LenOfUint64Fast efficiently calculate the string character lengths of an uint64 as input
// Optimized for input x mostly range between 0 to 99999
func LenOfUint64Fast(x uint64) int {
	if x > uintSizeTable[3] { // 4 - 20
		if x > uintSizeTable[5] { // 6 - 20, mid at 12
			if x > uintSizeTable[12] { // 13 - 20, mid at 16
				if x > uintSizeTable[16] { // 17 - 20, mid at 18
					if x > uintSizeTable[18] { // 19 - 20
						if x > uintSizeTable[19] {
							return 20
						}
						return 19
					}
					// 17 - 18
					if x > uintSizeTable[17] {
						return 18
					}
					return 17
				}
				// 13 - 16, mid at 14
				if x > uintSizeTable[14] { // 15 - 16
					if x > uintSizeTable[15] {
						return 16
					}
					return 15
				}
				// 13 - 14
				if x > uintSizeTable[13] {
					return 14
				}
				return 13
			}
			// 6 - 12, mid at 9
			if x > uintSizeTable[9] { // 10 - 12, mid at 11
				if x > uintSizeTable[11] { // 12 - 12
					return 12
				}
				// 10 - 11
				if x > uintSizeTable[10] {
					return 11
				}
				return 10
			}
			// 6 - 9, mid at 7
			if x > uintSizeTable[7] { // 8 - 9
				if x > uintSizeTable[8] {
					return 9
				}
				return 8
			}
			// 6 - 7
			if x > uintSizeTable[6] {
				return 7
			}
			return 6
		}
		// 4 - 5
		if x > uintSizeTable[4] {
			return 5
		}
		return 4
	}
	// 1 - 3
	if x > uintSizeTable[1] {
		if x > uintSizeTable[2] {
			return 3
		}
		return 2
	}
	return 1
}

// LenOfInt64Fast efficiently calculate the string character lengths of an int64 as input
func LenOfInt64Fast(x int64) int {
	size := 0
	if x < 0 {
		size = 1 // add "-" sign on the length count
	}
	return size + LenOfUint64Fast(uint64(abs(x)))
}
