// Copyright (C) 2017, Beijing Bochen Technology Co.,Ltd.  All rights reserved.
//
// This file is part of L0
// 
// The L0 is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// 
// The L0 is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// 
// GNU General Public License for more details.
// 
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package vote

import (
	"fmt"
	"testing"
)

type INT int

func (e INT) Serialize() []byte {
	str := fmt.Sprintf("%d", e)
	return []byte(str)
}

func TestVote(t *testing.T) {
	v := NewVote()
	v.Add("a", INT(1))
	v.Add("a", INT(2))
	fmt.Println(v.Size())
	v.Add("b", INT(1))
	fmt.Println(v.Voter())
	fmt.Println(v.VoterByTicket(INT(1)))
}
