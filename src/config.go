package garlic

import (
  "encoding/json"
  "io/ioutil"
)

// a config for a certain resource
type ResourceConfig struct {
  Filename string
  Props map[string]string
}

// config for all assets
type AssetConfig struct {
  // all resources to load
  Resources []ResourceConfig
  // root directory for all resources
  RootDir string
}

// config for our window
type DisplayConfig struct {
  Width int
  Height int
  Depth int
  Name string
}

//
// main configuration
//
type MainConfig struct {
  Assets AssetConfig
  Display DisplayConfig
}

func LoadConfig(fname string) (cfg *MainConfig, err error) {
  var d []byte
  d, err = ioutil.ReadFile(fname)
  if err == nil {
    cfg = new(MainConfig)
    err = json.Unmarshal(d, cfg)
  }
  return
}
