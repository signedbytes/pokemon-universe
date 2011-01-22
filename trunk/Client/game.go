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
	"sdl"
)

const (
	NUMTILES_X = 23
	NUMTILES_Y = 17
	
	MID_X = 11
	MID_Y = 8
)

const (
	GAMESTATE_UNLOADING = iota
	GAMESTATE_LOADING
	GAMESTATE_LOGIN
	GAMESTATE_WORLD
	GAMESTATE_BATTLE_INIT
	GAMESTATE_BATTLE
)

type PU_Game struct {
	state int
	tileImageMap map[uint16]*PU_Image
	guiImageMap map[uint16]*PU_Image
	creatureImageMap map[uint32]*PU_Image
	
	screenOffsetX int
	screenOffsetY int
	
	self *PU_Player
	
	lastDirKey int
}

func NewGame() *PU_Game {
	return &PU_Game{state : GAMESTATE_LOADING,
					tileImageMap : make(map[uint16]*PU_Image),
					guiImageMap : make(map[uint16]*PU_Image),
					creatureImageMap : make(map[uint32]*PU_Image)}
}

func (g *PU_Game) SetState(_state int) {
	g.state = _state
}

func (g *PU_Game) Draw() {
	switch g.state {
		case GAMESTATE_LOADING:
			g_engine.GetFont(FONT_PURITANBOLD_14).DrawText("Loading, please wait...", 10, 10)
			
		case GAMESTATE_LOGIN:
			g.GetGuiImage(IMG_GUI_INTROBG).Draw(0, 0)
			g_gui.Draw()
			
		case GAMESTATE_WORLD:
			g.DrawWorld()
	}
}

func (g *PU_Game) DrawWorld() {
	g.screenOffsetX, g.screenOffsetY = g.GetScreenOffset()
			
	layer2tiles := make([]*PU_Tile, (NUMTILES_X*NUMTILES_Y))
	layer2tilesCount := 0
	
	walkers := make([]ICreature, g_map.creatureList.Len())
	walkerCount := 0
	
	//draw tile layer 0 and 1 
	for x := 0; x < NUMTILES_X; x++ {
		for y := 0; y < NUMTILES_Y; y++ {
			mapx := int(g.self.GetX()-MID_X)+x
			mapy := int(g.self.GetY()-MID_Y)+y
			
			tile := g_map.GetTile(mapx, mapy)
			if tile != nil {
				for i := 0; i < 2; i++ {
					tile.DrawLayer(i, x, y)
				}
				if tile.layers[2] != nil {
					layer2tiles[layer2tilesCount] = tile
					layer2tilesCount++
				}
			}
		}
	}
	
	//draw creatures
	for e := g_map.creatureList.Front(); e != nil; e = e.Next() {
		if creature, is_type := e.Value.(ICreature); is_type {
			screenx := MID_X-(int(g.self.GetX())-int(creature.GetX()))
			screeny := MID_Y-(int(g.self.GetY())-int(creature.GetY()))
			g.DrawCreature(creature, screenx, screeny)
			
			if creature.IsWalking() {
				walkers[walkerCount] = creature
				walkerCount++
			}
		}
	}
	
	//draw tile layer 2
	for i := 0; i < layer2tilesCount; i++ {
		tile := layer2tiles[i]
		
		screenx := MID_X-(int(g.self.GetX())-tile.position.X)
		screeny := MID_Y-(int(g.self.GetY())-tile.position.Y)
		
		tile.DrawLayer(2, screenx, screeny)
	}
	
	//draw player names
	
	//draw extra info
	info := fmt.Sprintf("x: %v y: %v", g.self.x, g.self.y)
	if font := g_engine.GetFont(FONT_PURITANBOLD_14); font != nil {
		font.SetColor(255, 242, 0)
		font.DrawText(info, 5, 5)
	}
	
	//after all the drawing, update all walking creatures
	if walkerCount > 0 {
		for i := 0; i < walkerCount; i++ {
			walkers[i].UpdateWalk()
		}
	}
}

func (g *PU_Game) DrawCreature(_creature ICreature, _x int, _y int) {
	var drawX, drawY int
	offsetX, offsetY := g.GetScreenOffset()
	
	if _creature.IsWalking() {
		switch _creature.GetDirection() {
			case DIR_NORTH:
				drawX = (_x*TILE_WIDTH)-TILE_WIDTH-22+offsetX
				drawY = ((_y*TILE_HEIGHT)+(TILE_HEIGHT-_creature.GetOffset()))-TILE_HEIGHT+offsetY
				
			case DIR_EAST:
				drawX = ((_x*TILE_WIDTH)-(TILE_WIDTH-_creature.GetOffset()))-TILE_WIDTH-22+offsetX
				drawY = (_y*TILE_HEIGHT)-TILE_HEIGHT+offsetY
				
			case DIR_SOUTH:
				drawX = (_x*TILE_WIDTH)-TILE_WIDTH-22+offsetX
				drawY = ((_y*TILE_HEIGHT)-(TILE_HEIGHT-_creature.GetOffset()))-TILE_HEIGHT+offsetY
				
			case DIR_WEST:			
				drawX = ((_x*TILE_WIDTH)+(TILE_WIDTH-_creature.GetOffset()))-TILE_WIDTH-22+offsetX
				drawY = (_y*TILE_HEIGHT)-TILE_HEIGHT+offsetY
		}
	} else {
		drawX = (_x*TILE_WIDTH)-TILE_WIDTH-22+offsetX
		drawY = (_y*TILE_HEIGHT)-TILE_HEIGHT+offsetY
	}
	
	_creature.Draw(drawX, drawY)
}

func (g *PU_Game) GetScreenOffset() (x int, y int) {
	x = 0
	y = 0
	if g.self != nil && g.self.walking {
		switch g.self.direction {
			case DIR_NORTH:
				y = 0-(TILE_HEIGHT-g.self.offset)

			case DIR_EAST:
				x = (TILE_WIDTH-g.self.offset)

			case DIR_SOUTH:
				y = (TILE_HEIGHT-g.self.offset)

			case DIR_WEST:
				x = 0-(TILE_WIDTH-g.self.offset)
		}
	}
	return
}

func (g *PU_Game) KeyDown(_keysym int, _scancode int) {
	if g.state == GAMESTATE_WORLD {
		if g.self == nil {
			return
		}
		
		ctrlDown := sdl.KeyDown(sdl.SDL_SCANCODE_LCTRL) || sdl.KeyDown(sdl.SDL_SCANCODE_RCTRL)
		
		switch _scancode {
			case sdl.SDL_SCANCODE_LEFT:
				if ctrlDown {
					if !g.self.walking {
						g.self.Turn(DIR_WEST, true)
					}
				} else {
					g.self.Walk(DIR_WEST)
					g.lastDirKey = sdl.SDL_SCANCODE_LEFT
				}
				
			case sdl.SDL_SCANCODE_UP:
				if ctrlDown {
					if !g.self.walking {
						g.self.Turn(DIR_NORTH, true)
					}
				} else {
					g.self.Walk(DIR_NORTH)
					g.lastDirKey = sdl.SDL_SCANCODE_UP
				}
				
			case sdl.SDL_SCANCODE_RIGHT:
				if ctrlDown {
					if !g.self.walking {
						g.self.Turn(DIR_EAST, true)
					}
				} else {
					g.self.Walk(DIR_EAST)
					g.lastDirKey = sdl.SDL_SCANCODE_RIGHT
				}
				
			case sdl.SDL_SCANCODE_DOWN:
				if ctrlDown {
					if !g.self.walking {
						g.self.Turn(DIR_SOUTH, true)
					}
				} else {
					g.self.Walk(DIR_SOUTH)
					g.lastDirKey = sdl.SDL_SCANCODE_DOWN
				}
		}
	}
}

