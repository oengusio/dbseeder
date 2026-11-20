/*
 *     Yummy seed for the database 😋
 *     Copyright (C) 2026 DHG Media
 *     Copyright (C) 2026 duncte123
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Affero General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Affero General Public License for more details.
 *
 *     You should have received a copy of the GNU Affero General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 */

package queue

// source: https://stackoverflow.com/a/55214816/4807235

type Queue []interface{}

func (self *Queue) Push(x interface{}) {
	*self = append(*self, x)
}

func (self *Queue) Pop() interface{} {
	h := *self
	var el interface{}
	l := len(h)
	//el, *self = h[0], h[1:l]
	// Or use this instead for a Stack
	el, *self = h[l-1], h[0:l-1]
	return el
}

func New() *Queue {
	return &Queue{}
}
