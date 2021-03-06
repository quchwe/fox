// Copyright 2019 Molander.
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

package helper

import (
	"errors"
	"strconv"
)

func PageParser(page string, limit string) (p int, l int, err error) {
	p = -1
	l = -1
	if page != "" {
		p, err = strconv.Atoi(page)
		if err != nil {
			return
		}
		if limit != "" {
			l, err = strconv.Atoi(limit)
			if err != nil {
				return
			}
		} else {
			err = errors.New("You set page but skip limit params, please check your input")
			return
		}
		if p <= 0 || l <= 0 {
			err = errors.New("limit or page can not set to 0 or less than 0")
			return
		}
		if p == 1 {
			p = 0
		} else if p != 1 {
			p = ((p - 1) * l)
		}
	}
	return
}
