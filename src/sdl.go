package garlic

/*
#cgo pkg-config: sdl2

#include <SDL2/SDL.h>


Uint32 get_sdl_init_flags() {
  return SDL_INIT_VIDEO;
}

Uint32 get_sdl_window_flags() {
  return SDL_WINDOW_OPENGL;
}

int get_sdl_windowpos_undefined() {
  return SDL_WINDOWPOS_UNDEFINED;
}

Uint32 get_sdl_event_type(SDL_Event * ev) {
  return ev->type;
}

*/
import "C" 

import (
  "errors"
  "log"
  "time"
)

//
// sdl imports
//
type SDL struct {
  conf *DisplayConfig
  mainwin *C.SDL_Window
  eventChnl chan C.SDL_Event
  gl C.SDL_GLContext
  
  running bool
}

// make an sdl error into a go error
func sdlError() (err error) {
  cstr := C.SDL_GetError()
  err = errors.New(C.GoString(cstr))
  return
}

func (sdl *SDL) Init() (err error) {
  flags := C.get_sdl_init_flags()
  err = sdl.InitFlags(flags)
  
  if err == nil {
    // sdl init success
    // create window
    title := C.CString(sdl.conf.Name)
    w, h  := C.int(sdl.conf.Width), C.int(sdl.conf.Height)
    undef := C.get_sdl_windowpos_undefined()
    x, y := undef, undef
    flags = C.get_sdl_window_flags()
    sdl.mainwin = C.SDL_CreateWindow(title, x, y, w, h, flags)
    if sdl.mainwin == nil {
      err = sdlError()
    } else {
      // intialize gl
      C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MAJOR_VERSION, 3)
      C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MINOR_VERSION, 0)
      sdl.gl = C.SDL_GL_CreateContext(sdl.mainwin)
      if sdl.gl == nil {
        err = sdlError()
      } else {
        sdl.eventChnl = make(chan C.SDL_Event, 32)
        sdl.running = true
      }
    }
  }
  return
}

func (sdl *SDL) Continue() bool {
  return sdl.running
}

func (sdl *SDL) InitFlags(flags C.Uint32) (err error) {
  ret := C.SDL_Init(flags)
  if ret == C.int(0) {
    // we gud
  } else {
    // error
    err = sdlError()
  }
  return
}

// defer quit
func (sdl *SDL) Quit() {
  log.Println("SDL_Quit()")
  close(sdl.eventChnl)
  sdl.running = false
  // give time to everyone to die
  time.Sleep(time.Second)
  C.SDL_Quit();
}

// sleep for given milliseconds
func (sdl *SDL) Sleep(milli int) {
  C.SDL_Delay(C.Uint32(milli))
}


func (sdl *SDL) PollEvents() {
  for {
    var ev C.SDL_Event
    res := C.SDL_PollEvent(&ev)
    if res == C.int(0) {
      // no more events
      // sleep for a bit
      sdl.Sleep(50)
    } else {
      if C.get_sdl_event_type(&ev) == C.SDL_QUIT {
        // sdl quit action
        log.Println("sdl quit event")
        sdl.Quit()
        break
      } else {
        // non quit action
        sdl.eventChnl <- ev
      }
    }
  }
}

// flip buffers
func (sdl *SDL) VSync() {
  if sdl.running {
    // additional check
    glClear()
    C.SDL_GL_SwapWindow(sdl.mainwin)
  }
}
