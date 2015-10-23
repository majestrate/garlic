package garlic

/*
#cgo pkg-config: sdl2 gl

#include <SDL2/SDL_opengles2.h>

*/
import "C" 

//
// initialize opengl
//
func initGL() {
  C.glClearColor(0.3, 0.2, 0.1, 1.0)
}

func glClear() {
  C.glClear(C.GL_COLOR_BUFFER_BIT)
}
