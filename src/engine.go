package garlic

import (
  "errors"
  "log"
)

//
// main game engine
//
type Engine struct {
  Config *MainConfig
  sdl *SDL
  event *EventEngine
}

// intialize engine
func (eng *Engine) Init() (err error) {
  eng.sdl = new(SDL)
  eng.sdl.conf = &eng.Config.Display
  err = eng.sdl.Init()
  if err == nil {
    log.Println("registering event channel")
    eng.event = new(EventEngine)
    eng.event.input = eng.sdl.eventChnl
  }
  return
}

func (eng *Engine) LoadConfig(fname string) (err error) {
  eng.Config, err = LoadConfig(fname)
  return
}

func (eng *Engine) LoadResources() (err error) {
  if eng.Config == nil {
    err = errors.New("no config loaded")
  }
  return
}

func (eng *Engine) RunMain() {
  log.Println("enter mainloop")
  go eng.event.RunLoop()
  go eng.sdl.PollEvents()
  for eng.sdl.Continue() {
    eng.sdl.VSync()
    eng.sdl.Sleep(10)
  }
}
