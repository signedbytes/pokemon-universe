/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/
package main

type BattleConfiguration struct {
	gen		uint8
	mode 	uint8
	ids		[]int32
	clauses	uint32
}

func NewBattleConfiguration() *BattleConfiguration {
	return &BattleConfiguration{ ids: make([]int32, 2) }
}

func (b *BattleConfiguration) slot(spot int32, poke int32) int32 {
	return spot + (poke*2)
}

func (b *BattleConfiguration) spot(id int32) (ret int32) {
	ret = 1
	if b.ids[0] == id {
		ret = 0
	}
	return
}