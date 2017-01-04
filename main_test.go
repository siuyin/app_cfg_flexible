package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadConfigSimple2(t *testing.T) {
	c := loadConfig("managerCfg.yaml")

	if rs := c.FileMoversOutCfg.Status.Fail.Output[1].Data["Strategy"]; rs != "ExponentialBackoff" {
		t.Errorf("Bad retry strategy: %s", rs)
	}

	if ec := c.EmailersCfg.Board; ec != "EmailersA" {
		t.Errorf("Bad emailers board: %s", ec)
	}

	if ec := c.EmailersOutCfg.Status.OK.Output[0].Data["Msg"]; ec != "Email {{ EmailDetails }} sucessfully sent" {
		t.Errorf("Bad response: %s", ec)
	}
}

func TestSubs(t *testing.T) {
	c := loadConfig("managerCfg.yaml")
	s := subs(c, true)
	if ls := fmt.Sprintf("%s", s); ls != "[FileStableA FileMoversAOut EmailersAOut]" {
		t.Errorf("Bad list: %s", ls)
	}
}

func TestPubs(t *testing.T) {
	c := loadConfig("managerCfg.yaml")
	s := subs(c, false) // <-- note false == list of boards to publish to.
	if ls := fmt.Sprintf("%s", s); ls != "[FileMoversA EmailersA]" {
		t.Errorf("Bad list: %s", ls)
	}
}

func TestProcWFMsg(t *testing.T) {
	c := loadConfig("managerCfg.yaml")
	dat := []struct {
		i, e string
	}{{"Turner", "XCodeA"}, {"Diva", "XCodeB"}}

	for _, d := range dat {
		if s := procWFMsg(&c, d.i); s != d.e {
			t.Errorf("Bad dest: %s,expected %s", s, d.e)
		}
	}
}

func TestProcFMFail(t *testing.T) {
	c := loadConfig("managerCfg.yaml")
	if s := procFMFail(&c, "fileA", "destA"); s != "FileMoveRetry" {
		t.Errorf("bad response ID: %s", s)
	}
}

func TestA(t *testing.T) {
	c := Cfg{}
	t.Errorf("Val: %v", reflect.ValueOf(c).Kind())

}
