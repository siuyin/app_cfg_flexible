package main

import (
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v2"
)

//Cfg is a system configuration.
type Cfg struct {
	FileWatchersCfg *struct {
		Board     string `yaml:"Board"`
		Subscribe bool
		Dirs      []struct {
			Dir    string `yaml:"Dir"`
			Output []Msg  `yaml:"Output"`
		} `yaml:"Dirs"`
	} `yaml:"FileWatchersCfg"`

	FileMoversCfg *struct {
		Board     string `yaml:"Board"`
		Subscribe bool
	} `yaml:"FileMoversCfg"`

	FileMoversOutCfg *struct {
		Board     string `yaml:"Board"`
		Subscribe bool
		Status    struct {
			OK struct {
				Output []Msg `yaml:"Output"`
			} `yaml:"OK"`
			Fail struct {
				Output []Msg `yaml:"Output"`
			} `yaml:"Fail"`
		} `yaml:"Status"`
	} `yaml:"FileMoversOutCfg"`

	EmailersCfg *struct {
		Board     string
		Subscribe bool
	}

	EmailersOutCfg *struct {
		Board     string
		Subscribe bool
		Status    struct {
			OK struct {
				Output []Msg
			}
		}
	}
}

//Msg is a generic message.
type Msg struct {
	Board string            `yaml:"Board"`
	ID    string            `yaml:"ID"`
	Op    string            `yaml:"Op"`
	Data  map[string]string `yaml:"Data"`
}

func loadConfig(fn string) Cfg {
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	var c Cfg
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

// subs(c, true) => list of subscribed boards
// subs(c, false) => list of boards published to.
// This uses fancy reflection stuff.
func subs(c Cfg, dir bool) []string {
	bs := []string{}
	v := reflect.ValueOf(c)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Elem().FieldByName("Subscribe").Bool() == dir {
			bs = append(bs, v.Field(i).Elem().FieldByName("Board").String())
		}
	}
	return bs
}

// These are just some silly example processing routines.
func procWFMsg(c *Cfg, d string) string {
	for i := 0; i < len(c.FileWatchersCfg.Dirs); i++ {
		if d == c.FileWatchersCfg.Dirs[i].Dir {
			return c.FileWatchersCfg.Dirs[i].Output[0].Data["Tgt"]
		}
	}
	return ""
}

func procFMFail(c *Cfg, s, d string) string {
	return c.FileMoversOutCfg.Status.Fail.Output[1].ID
}

func main() {
}
