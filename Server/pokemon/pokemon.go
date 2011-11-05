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

import (	
	"fmt"
	"container/list"
)

type Pokemon struct {
	PokemonId				int
	Species					*PokemonSpecies
	Height					int
	Weight					int
	BaseExperience			int
	Order					int
	IsDefault				int
	
	Stats					PokemonStatArray // Size = 6
	
	Abilities				PokemonAbilityList
	Forms					*list.List
	Moves					PokemonMoveList
	Types					PokemonTypeArray
}

func NewPokemon() *Pokemon {
	pokemon := &Pokemon{ Stats: make(PokemonStatArray, 6),
					 Abilities: make(PokemonAbilityList),
					 Moves: make(PokemonMoveList),
					 Types: make(PokemonTypeArray, 2) }
	pokemon.Forms.Init()
	
	return pokemon
}

func (p *Pokemon) loadStats() {
	var query string = "SELECT stat_id, base_stat, effort FROM pokemon_stats WHERE pokemon_id='%d'"
	result, err := DBQuerySelect(fmt.Sprintf(query, p.PokemonId))
	if err != nil {
		return
	}
	
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		stat := NewPokemonStat()
		stat.StatType = row[0].(int)
		stat.BaseStat = row[1].(int)
		stat.Effort = row[2].(int)
		p.Stats[stat.StatType] = stat
	}
}

func (p *Pokemon) loadAbilities() {
	var query string = "SELECT ability_id, is_dream, slot FROM pokemon_abilities WHERE pokemon_id='%d'"
	result, err := DBQuerySelect(fmt.Sprintf(query, p.PokemonId))
	if err != nil {
		return
	}
	
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		ability := NewPokemonAbility()
		id := row[0].(int)
		ability.Ability = g_PokemonManager.GetAbilityById(id)
		ability.IsDream = row[1].(int)
		ability.Slot = row[2].(int)

		if ability.Ability != nil {
			p.Abilities[id] = ability
		}
	}		
}

func (p *Pokemon) loadForms() {
	var query string = "SELECT id, form_identifier, is_default, is_battle_only, order FROM pokemon_forms WHERE pokemon_id='%d'"
	result, err := DBQuerySelect(fmt.Sprintf(query, p.PokemonId))
	if err != nil {
		return
	}
	
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		form := NewPokemonForm()
		form.Id = row[0].(int)
		form.Identifier = row[1].(string)
		form.IsDefault = row[2].(int)
		form.IsBattleOnly = row[3].(int)
		form.Order = row[4].(int)
		
		p.Forms.PushBack(form)
	}
}

func (p *Pokemon) loadMoves() {
	var query string = "SELECT verion_group_id, move_id, pokemon_move_method_id, level, order FROM pokemon_moves" + 
						" WHERE pokemon_id='%d' AND verion_group_id=11"
	result, err := DBQuerySelect(fmt.Sprintf(query, p.PokemonId))
	if err != nil {
		return
	}
	
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		pmove := NewPokemonMove()
		pmove.Pokemon = p
		pmove.VersionGroup = row[0].(int)
		moveId := row[1].(int)
		pmove.Move = g_PokemonManager.GetMoveById(moveId)
		pmove.PokemonMoveMethod = row[2].(int)
		pmove.Level = row[3].(int)
		pmove.Order = row[4].(int)
		
		if pmove.Move != nil {
			p.Moves[moveId] = pmove
		}
	}
}

func (p *Pokemon) loadTypes() {
	var query string = "SELECT type_id, slot FROM pokemon_types WHERE pokemon_id='%d' ORDER BY slot"
	result, err := DBQuerySelect(fmt.Sprintf(query, p.PokemonId))
	if err != nil {
		return
	}
	
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		slot := row[1].(int)
		p.Types[slot - 1] = row[0].(int)
	}
}