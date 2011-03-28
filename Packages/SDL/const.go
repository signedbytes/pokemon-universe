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

/*
Simple SDL 1.3 wrapper
The goal is not to make a complete SDL wrapper, but to wrap only the SDL functions that the PU client needs.
*/
package sdl

//#include <SDL.h>
import "C"

const (
	SDL_INIT_VIDEO = C.SDL_INIT_VIDEO
	SDL_INIT_TIMER = C.SDL_INIT_TIMER

	SDL_WINDOWPOS_CENTERED    = C.SDL_WINDOWPOS_CENTERED
	SDL_WINDOW_SHOWN          = C.SDL_WINDOW_SHOWN
	SDL_RENDERER_PRESENTVSYNC = C.SDL_RENDERER_PRESENTVSYNC
	SDL_RENDERER_ACCELERATED  = C.SDL_RENDERER_ACCELERATED

	BLENDMODE_NONE  = C.SDL_BLENDMODE_NONE
	BLENDMODE_BLEND = C.SDL_BLENDMODE_BLEND
	BLENDMODE_ADD   = C.SDL_BLENDMODE_ADD
	BLENDMODE_MOD   = C.SDL_BLENDMODE_MOD

	/* Window events */
	SDL_WINDOWEVENT_NONE         = C.SDL_WINDOWEVENT_NONE
	SDL_WINDOWEVENT_SHOWN        = C.SDL_WINDOWEVENT_SHOWN
	SDL_WINDOWEVENT_HIDDEN       = C.SDL_WINDOWEVENT_HIDDEN
	SDL_WINDOWEVENT_EXPOSED      = C.SDL_WINDOWEVENT_EXPOSED
	SDL_WINDOWEVENT_MOVED        = C.SDL_WINDOWEVENT_MOVED
	SDL_WINDOWEVENT_RESIZED      = C.SDL_WINDOWEVENT_RESIZED
	SDL_WINDOWEVENT_SIZE_CHANGED = C.SDL_WINDOWEVENT_SIZE_CHANGED
	SDL_WINDOWEVENT_MINIMIZED    = C.SDL_WINDOWEVENT_MINIMIZED
	SDL_WINDOWEVENT_MAXIMIZED    = C.SDL_WINDOWEVENT_MAXIMIZED
	SDL_WINDOWEVENT_RESTORED     = C.SDL_WINDOWEVENT_RESTORED
	SDL_WINDOWEVENT_ENTER        = C.SDL_WINDOWEVENT_ENTER
	SDL_WINDOWEVENT_LEAVE        = C.SDL_WINDOWEVENT_LEAVE
	SDL_WINDOWEVENT_FOCUS_GAINED = C.SDL_WINDOWEVENT_FOCUS_GAINED
	SDL_WINDOWEVENT_FOCUS_LOST   = C.SDL_WINDOWEVENT_FOCUS_LOST
	SDL_WINDOWEVENT_CLOSE        = C.SDL_WINDOWEVENT_CLOSE

	SDL_QUIT        = C.SDL_QUIT
	SDL_WINDOWEVENT = C.SDL_WINDOWEVENT
	SDL_SYSWMEVENT  = C.SDL_SYSWMEVENT

	/* Keyboard events */
	SDL_KEYDOWN     = C.SDL_KEYDOWN
	SDL_KEYUP       = C.SDL_KEYUP
	SDL_TEXTEDITING = C.SDL_TEXTEDITING
	SDL_TEXTINPUT   = C.SDL_TEXTINPUT

	/* Mouse events */
	SDL_MOUSEMOTION     = C.SDL_MOUSEMOTION
	SDL_MOUSEBUTTONDOWN = C.SDL_MOUSEBUTTONDOWN
	SDL_MOUSEBUTTONUP   = C.SDL_MOUSEBUTTONUP
	SDL_MOUSEWHEEL      = C.SDL_MOUSEWHEEL
	SCROLL_UP           = 0
	SCROLL_DOWN         = 1
)

const (
	KEY_UNKNOWN            = C.SDL_SCANCODE_UNKNOWN
	KEY_A                  = C.SDL_SCANCODE_A
	KEY_B                  = C.SDL_SCANCODE_B
	KEY_C                  = C.SDL_SCANCODE_C
	KEY_D                  = C.SDL_SCANCODE_D
	KEY_E                  = C.SDL_SCANCODE_E
	KEY_F                  = C.SDL_SCANCODE_F
	KEY_G                  = C.SDL_SCANCODE_G
	KEY_H                  = C.SDL_SCANCODE_H
	KEY_I                  = C.SDL_SCANCODE_I
	KEY_J                  = C.SDL_SCANCODE_J
	KEY_K                  = C.SDL_SCANCODE_K
	KEY_L                  = C.SDL_SCANCODE_L
	KEY_M                  = C.SDL_SCANCODE_M
	KEY_N                  = C.SDL_SCANCODE_N
	KEY_O                  = C.SDL_SCANCODE_O
	KEY_P                  = C.SDL_SCANCODE_P
	KEY_Q                  = C.SDL_SCANCODE_Q
	KEY_R                  = C.SDL_SCANCODE_R
	KEY_S                  = C.SDL_SCANCODE_S
	KEY_T                  = C.SDL_SCANCODE_T
	KEY_U                  = C.SDL_SCANCODE_U
	KEY_V                  = C.SDL_SCANCODE_V
	KEY_W                  = C.SDL_SCANCODE_W
	KEY_X                  = C.SDL_SCANCODE_X
	KEY_Y                  = C.SDL_SCANCODE_Y
	KEY_Z                  = C.SDL_SCANCODE_Z
	KEY_1                  = C.SDL_SCANCODE_1
	KEY_2                  = C.SDL_SCANCODE_2
	KEY_3                  = C.SDL_SCANCODE_3
	KEY_4                  = C.SDL_SCANCODE_4
	KEY_5                  = C.SDL_SCANCODE_5
	KEY_6                  = C.SDL_SCANCODE_6
	KEY_7                  = C.SDL_SCANCODE_7
	KEY_8                  = C.SDL_SCANCODE_8
	KEY_9                  = C.SDL_SCANCODE_9
	KEY_0                  = C.SDL_SCANCODE_0
	KEY_RETURN             = C.SDL_SCANCODE_RETURN
	KEY_ESCAPE             = C.SDL_SCANCODE_ESCAPE
	KEY_BACKSPACE          = C.SDL_SCANCODE_BACKSPACE
	KEY_TAB                = C.SDL_SCANCODE_TAB
	KEY_SPACE              = C.SDL_SCANCODE_SPACE
	KEY_MINUS              = C.SDL_SCANCODE_MINUS
	KEY_EQUALS             = C.SDL_SCANCODE_EQUALS
	KEY_LEFTBRACKET        = C.SDL_SCANCODE_LEFTBRACKET
	KEY_RIGHTBRACKET       = C.SDL_SCANCODE_RIGHTBRACKET
	KEY_BACKSLASH          = C.SDL_SCANCODE_BACKSLASH
	KEY_NONUSHASH          = C.SDL_SCANCODE_NONUSHASH
	KEY_SEMICOLON          = C.SDL_SCANCODE_SEMICOLON
	KEY_APOSTROPHE         = C.SDL_SCANCODE_APOSTROPHE
	KEY_GRAVE              = C.SDL_SCANCODE_GRAVE
	KEY_COMMA              = C.SDL_SCANCODE_COMMA
	KEY_PERIOD             = C.SDL_SCANCODE_PERIOD
	KEY_SLASH              = C.SDL_SCANCODE_SLASH
	KEY_CAPSLOCK           = C.SDL_SCANCODE_CAPSLOCK
	KEY_F1                 = C.SDL_SCANCODE_F1
	KEY_F2                 = C.SDL_SCANCODE_F2
	KEY_F3                 = C.SDL_SCANCODE_F3
	KEY_F4                 = C.SDL_SCANCODE_F4
	KEY_F5                 = C.SDL_SCANCODE_F5
	KEY_F6                 = C.SDL_SCANCODE_F6
	KEY_F7                 = C.SDL_SCANCODE_F7
	KEY_F8                 = C.SDL_SCANCODE_F8
	KEY_F9                 = C.SDL_SCANCODE_F9
	KEY_F10                = C.SDL_SCANCODE_F10
	KEY_F11                = C.SDL_SCANCODE_F11
	KEY_F12                = C.SDL_SCANCODE_F12
	KEY_PRINTSCREEN        = C.SDL_SCANCODE_PRINTSCREEN
	KEY_SCROLLLOCK         = C.SDL_SCANCODE_SCROLLLOCK
	KEY_PAUSE              = C.SDL_SCANCODE_PAUSE
	KEY_INSERT             = C.SDL_SCANCODE_INSERT
	KEY_HOME               = C.SDL_SCANCODE_HOME
	KEY_PAGEUP             = C.SDL_SCANCODE_PAGEUP
	KEY_DELETE             = C.SDL_SCANCODE_DELETE
	KEY_END                = C.SDL_SCANCODE_END
	KEY_PAGEDOWN           = C.SDL_SCANCODE_PAGEDOWN
	KEY_RIGHT              = C.SDL_SCANCODE_RIGHT
	KEY_LEFT               = C.SDL_SCANCODE_LEFT
	KEY_DOWN               = C.SDL_SCANCODE_DOWN
	KEY_UP                 = C.SDL_SCANCODE_UP
	KEY_NUMLOCKCLEAR       = C.SDL_SCANCODE_NUMLOCKCLEAR
	KEY_KP_DIVIDE          = C.SDL_SCANCODE_KP_DIVIDE
	KEY_KP_MULTIPLY        = C.SDL_SCANCODE_KP_MULTIPLY
	KEY_KP_MINUS           = C.SDL_SCANCODE_KP_MINUS
	KEY_KP_PLUS            = C.SDL_SCANCODE_KP_PLUS
	KEY_KP_ENTER           = C.SDL_SCANCODE_KP_ENTER
	KEY_KP_1               = C.SDL_SCANCODE_KP_1
	KEY_KP_2               = C.SDL_SCANCODE_KP_2
	KEY_KP_3               = C.SDL_SCANCODE_KP_3
	KEY_KP_4               = C.SDL_SCANCODE_KP_4
	KEY_KP_5               = C.SDL_SCANCODE_KP_5
	KEY_KP_6               = C.SDL_SCANCODE_KP_6
	KEY_KP_7               = C.SDL_SCANCODE_KP_7
	KEY_KP_8               = C.SDL_SCANCODE_KP_8
	KEY_KP_9               = C.SDL_SCANCODE_KP_9
	KEY_KP_0               = C.SDL_SCANCODE_KP_0
	KEY_KP_PERIOD          = C.SDL_SCANCODE_KP_PERIOD
	KEY_NONUSBACKSLASH     = C.SDL_SCANCODE_NONUSBACKSLASH
	KEY_APPLICATION        = C.SDL_SCANCODE_APPLICATION
	KEY_POWER              = C.SDL_SCANCODE_POWER
	KEY_KP_EQUALS          = C.SDL_SCANCODE_KP_EQUALS
	KEY_F13                = C.SDL_SCANCODE_F13
	KEY_F14                = C.SDL_SCANCODE_F14
	KEY_F15                = C.SDL_SCANCODE_F15
	KEY_F16                = C.SDL_SCANCODE_F16
	KEY_F17                = C.SDL_SCANCODE_F17
	KEY_F18                = C.SDL_SCANCODE_F18
	KEY_F19                = C.SDL_SCANCODE_F19
	KEY_F20                = C.SDL_SCANCODE_F20
	KEY_F21                = C.SDL_SCANCODE_F21
	KEY_F22                = C.SDL_SCANCODE_F22
	KEY_F23                = C.SDL_SCANCODE_F23
	KEY_F24                = C.SDL_SCANCODE_F24
	KEY_EXECUTE            = C.SDL_SCANCODE_EXECUTE
	KEY_HELP               = C.SDL_SCANCODE_HELP
	KEY_MENU               = C.SDL_SCANCODE_MENU
	KEY_SELECT             = C.SDL_SCANCODE_SELECT
	KEY_STOP               = C.SDL_SCANCODE_STOP
	KEY_AGAIN              = C.SDL_SCANCODE_AGAIN
	KEY_UNDO               = C.SDL_SCANCODE_UNDO
	KEY_CUT                = C.SDL_SCANCODE_CUT
	KEY_COPY               = C.SDL_SCANCODE_COPY
	KEY_PASTE              = C.SDL_SCANCODE_PASTE
	KEY_FIND               = C.SDL_SCANCODE_FIND
	KEY_MUTE               = C.SDL_SCANCODE_MUTE
	KEY_VOLUMEUP           = C.SDL_SCANCODE_VOLUMEUP
	KEY_VOLUMEDOWN         = C.SDL_SCANCODE_VOLUMEDOWN
	KEY_KP_COMMA           = C.SDL_SCANCODE_KP_COMMA
	KEY_KP_EQUALSAS400     = C.SDL_SCANCODE_KP_EQUALSAS400
	KEY_INTERNATIONAL1     = C.SDL_SCANCODE_INTERNATIONAL1
	KEY_INTERNATIONAL2     = C.SDL_SCANCODE_INTERNATIONAL2
	KEY_INTERNATIONAL3     = C.SDL_SCANCODE_INTERNATIONAL3
	KEY_INTERNATIONAL4     = C.SDL_SCANCODE_INTERNATIONAL4
	KEY_INTERNATIONAL5     = C.SDL_SCANCODE_INTERNATIONAL5
	KEY_INTERNATIONAL6     = C.SDL_SCANCODE_INTERNATIONAL6
	KEY_INTERNATIONAL7     = C.SDL_SCANCODE_INTERNATIONAL7
	KEY_INTERNATIONAL8     = C.SDL_SCANCODE_INTERNATIONAL8
	KEY_INTERNATIONAL9     = C.SDL_SCANCODE_INTERNATIONAL9
	KEY_LANG1              = C.SDL_SCANCODE_LANG1
	KEY_LANG2              = C.SDL_SCANCODE_LANG2
	KEY_LANG3              = C.SDL_SCANCODE_LANG3
	KEY_LANG4              = C.SDL_SCANCODE_LANG4
	KEY_LANG5              = C.SDL_SCANCODE_LANG5
	KEY_LANG6              = C.SDL_SCANCODE_LANG6
	KEY_LANG7              = C.SDL_SCANCODE_LANG7
	KEY_LANG8              = C.SDL_SCANCODE_LANG8
	KEY_LANG9              = C.SDL_SCANCODE_LANG9
	KEY_ALTERASE           = C.SDL_SCANCODE_ALTERASE
	KEY_SYSREQ             = C.SDL_SCANCODE_SYSREQ
	KEY_CANCEL             = C.SDL_SCANCODE_CANCEL
	KEY_CLEAR              = C.SDL_SCANCODE_CLEAR
	KEY_PRIOR              = C.SDL_SCANCODE_PRIOR
	KEY_RETURN2            = C.SDL_SCANCODE_RETURN2
	KEY_SEPARATOR          = C.SDL_SCANCODE_SEPARATOR
	KEY_OUT                = C.SDL_SCANCODE_OUT
	KEY_OPER               = C.SDL_SCANCODE_OPER
	KEY_CLEARAGAIN         = C.SDL_SCANCODE_CLEARAGAIN
	KEY_CRSEL              = C.SDL_SCANCODE_CRSEL
	KEY_EXSEL              = C.SDL_SCANCODE_EXSEL
	KEY_KP_00              = C.SDL_SCANCODE_KP_00
	KEY_KP_000             = C.SDL_SCANCODE_KP_000
	KEY_THOUSANDSSEPARATOR = C.SDL_SCANCODE_THOUSANDSSEPARATOR
	KEY_DECIMALSEPARATOR   = C.SDL_SCANCODE_DECIMALSEPARATOR
	KEY_CURRENCYUNIT       = C.SDL_SCANCODE_CURRENCYUNIT
	KEY_CURRENCYSUBUNIT    = C.SDL_SCANCODE_CURRENCYSUBUNIT
	KEY_KP_LEFTPAREN       = C.SDL_SCANCODE_KP_LEFTPAREN
	KEY_KP_RIGHTPAREN      = C.SDL_SCANCODE_KP_RIGHTPAREN
	KEY_KP_LEFTBRACE       = C.SDL_SCANCODE_KP_LEFTBRACE
	KEY_KP_RIGHTBRACE      = C.SDL_SCANCODE_KP_RIGHTBRACE
	KEY_KP_TAB             = C.SDL_SCANCODE_KP_TAB
	KEY_KP_BACKSPACE       = C.SDL_SCANCODE_KP_BACKSPACE
	KEY_KP_A               = C.SDL_SCANCODE_KP_A
	KEY_KP_B               = C.SDL_SCANCODE_KP_B
	KEY_KP_C               = C.SDL_SCANCODE_KP_C
	KEY_KP_D               = C.SDL_SCANCODE_KP_D
	KEY_KP_E               = C.SDL_SCANCODE_KP_E
	KEY_KP_F               = C.SDL_SCANCODE_KP_F
	KEY_KP_XOR             = C.SDL_SCANCODE_KP_XOR
	KEY_KP_POWER           = C.SDL_SCANCODE_KP_POWER
	KEY_KP_PERCENT         = C.SDL_SCANCODE_KP_PERCENT
	KEY_KP_LESS            = C.SDL_SCANCODE_KP_LESS
	KEY_KP_GREATER         = C.SDL_SCANCODE_KP_GREATER
	KEY_KP_AMPERSAND       = C.SDL_SCANCODE_KP_AMPERSAND
	KEY_KP_DBLAMPERSAND    = C.SDL_SCANCODE_KP_DBLAMPERSAND
	KEY_KP_VERTICALBAR     = C.SDL_SCANCODE_KP_VERTICALBAR
	KEY_KP_DBLVERTICALBAR  = C.SDL_SCANCODE_KP_DBLVERTICALBAR
	KEY_KP_COLON           = C.SDL_SCANCODE_KP_COLON
	KEY_KP_HASH            = C.SDL_SCANCODE_KP_HASH
	KEY_KP_SPACE           = C.SDL_SCANCODE_KP_SPACE
	KEY_KP_AT              = C.SDL_SCANCODE_KP_AT
	KEY_KP_EXCLAM          = C.SDL_SCANCODE_KP_EXCLAM
	KEY_KP_MEMSTORE        = C.SDL_SCANCODE_KP_MEMSTORE
	KEY_KP_MEMRECALL       = C.SDL_SCANCODE_KP_MEMRECALL
	KEY_KP_MEMCLEAR        = C.SDL_SCANCODE_KP_MEMCLEAR
	KEY_KP_MEMADD          = C.SDL_SCANCODE_KP_MEMADD
	KEY_KP_MEMSUBTRACT     = C.SDL_SCANCODE_KP_MEMSUBTRACT
	KEY_KP_MEMMULTIPLY     = C.SDL_SCANCODE_KP_MEMMULTIPLY
	KEY_KP_MEMDIVIDE       = C.SDL_SCANCODE_KP_MEMDIVIDE
	KEY_KP_PLUSMINUS       = C.SDL_SCANCODE_KP_PLUSMINUS
	KEY_KP_CLEAR           = C.SDL_SCANCODE_KP_CLEAR
	KEY_KP_CLEARENTRY      = C.SDL_SCANCODE_KP_CLEARENTRY
	KEY_KP_BINARY          = C.SDL_SCANCODE_KP_BINARY
	KEY_KP_OCTAL           = C.SDL_SCANCODE_KP_OCTAL
	KEY_KP_DECIMAL         = C.SDL_SCANCODE_KP_DECIMAL
	KEY_KP_HEXADECIMAL     = C.SDL_SCANCODE_KP_HEXADECIMAL
	KEY_LCTRL              = C.SDL_SCANCODE_LCTRL
	KEY_LSHIFT             = C.SDL_SCANCODE_LSHIFT
	KEY_LALT               = C.SDL_SCANCODE_LALT
	KEY_LGUI               = C.SDL_SCANCODE_LGUI
	KEY_RCTRL              = C.SDL_SCANCODE_RCTRL
	KEY_RSHIFT             = C.SDL_SCANCODE_RSHIFT
	KEY_RALT               = C.SDL_SCANCODE_RALT
	KEY_RGUI               = C.SDL_SCANCODE_RGUI
	KEY_MODE               = C.SDL_SCANCODE_MODE
	KEY_AUDIONEXT          = C.SDL_SCANCODE_AUDIONEXT
	KEY_AUDIOPREV          = C.SDL_SCANCODE_AUDIOPREV
	KEY_AUDIOSTOP          = C.SDL_SCANCODE_AUDIOSTOP
	KEY_AUDIOPLAY          = C.SDL_SCANCODE_AUDIOPLAY
	KEY_AUDIOMUTE          = C.SDL_SCANCODE_AUDIOMUTE
	KEY_MEDIASELECT        = C.SDL_SCANCODE_MEDIASELECT
	KEY_WWW                = C.SDL_SCANCODE_WWW
	KEY_MAIL               = C.SDL_SCANCODE_MAIL
	KEY_CALCULATOR         = C.SDL_SCANCODE_CALCULATOR
	KEY_COMPUTER           = C.SDL_SCANCODE_COMPUTER
	KEY_AC_SEARCH          = C.SDL_SCANCODE_AC_SEARCH
	KEY_AC_HOME            = C.SDL_SCANCODE_AC_HOME
	KEY_AC_BACK            = C.SDL_SCANCODE_AC_BACK
	KEY_AC_FORWARD         = C.SDL_SCANCODE_AC_FORWARD
	KEY_AC_STOP            = C.SDL_SCANCODE_AC_STOP
	KEY_AC_REFRESH         = C.SDL_SCANCODE_AC_REFRESH
	KEY_AC_BOOKMARKS       = C.SDL_SCANCODE_AC_BOOKMARKS
	KEY_BRIGHTNESSDOWN     = C.SDL_SCANCODE_BRIGHTNESSDOWN
	KEY_BRIGHTNESSUP       = C.SDL_SCANCODE_BRIGHTNESSUP
	KEY_DISPLAYSWITCH      = C.SDL_SCANCODE_DISPLAYSWITCH
	KEY_KBDILLUMTOGGLE     = C.SDL_SCANCODE_KBDILLUMTOGGLE
	KEY_KBDILLUMDOWN       = C.SDL_SCANCODE_KBDILLUMDOWN
	KEY_KBDILLUMUP         = C.SDL_SCANCODE_KBDILLUMUP
	KEY_EJECT              = C.SDL_SCANCODE_EJECT
	KEY_SLEEP              = C.SDL_SCANCODE_SLEEP
)
