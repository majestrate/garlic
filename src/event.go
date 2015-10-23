package garlic

/*
 #cgo pkg-config: sdl2
 #include <SDL2/SDL.h>

Uint32 sdl_event_type(SDL_Event * ev) {
  return ev->type;
}

SDL_KeyboardEvent* GetKeyEvent(SDL_Event * ev) {
  return &ev->key;
}

*/
import "C"

import (
  "log"
)

// event handler engine
type EventEngine struct {
  input chan C.SDL_Event
}


// run events mainloop
func (ev *EventEngine) RunLoop() {
  for {
    e, ok := <- ev.input
    if ok {
      ev_type := C.sdl_event_type(&e)
      switch ev_type {
      case C.SDL_KEYUP:
        ev.keyUp(C.GetKeyEvent(&e))
        break
      case C.SDL_KEYDOWN:
        ev.keyDown(C.GetKeyEvent(&e))
      }
    } else {
      return
    }
  }
}

func (ev EventEngine) logKey(evname string, sym C.SDL_Keycode) {
  cstr := C.SDL_GetKeyName(sym)
  keyname := C.GoString(cstr)
  log.Println(evname, keyname)
}

func (ev *EventEngine) keyUp(key *C.SDL_KeyboardEvent) {
  keysym := key.keysym
  ev.logKey("keyup", keysym.sym)
}
func (ev *EventEngine) keyDown(key *C.SDL_KeyboardEvent) {
  keysym := key.keysym
  ev.logKey("keydown", keysym.sym)
}
