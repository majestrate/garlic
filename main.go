package main

import (
  "github.com/majestrate/garlic/src"
  "log"
  "os"
)

//
// main function for game
//

func main() {
  eng := new(garlic.Engine)
  cfgFname := "game.json"
  if len(os.Args) > 1 {
    cfgFname = os.Args[1]
  }
  log.Println("using config file", cfgFname)
  err := eng.LoadConfig(cfgFname)
  if err == nil {
    err = eng.Init()
    if err == nil {
      log.Println("loading resources...")
      err = eng.LoadResources()
      if err == nil {
        log.Println("run main")
        eng.RunMain()
      } else {
        log.Println("failed to load resources", err)
      }
    } else {
      log.Println("failed to initialize engine", err)
    }
  } else {
    log.Println("failed to load config", err)
  }
  log.Println("end")
}
