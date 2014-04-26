// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package vbox

import (
	"testing"

	"github.com/gcloud/compute"
)

var NewServer = &Server{name: "GCloudServer"}
var TestServer compute.Server

func Test_ServersCreate(t *testing.T) {
	servers := &Servers{}
	result, err := servers.Create(NewServer)
	TestServer = result
	if err != nil {
		t.Errorf("Servers Create failed with %s.", err)
	}
	if result == nil {
		t.Error("Servers Create failed with no result.")
		return
	}
	if len(result.Id()) <= 0 {
		t.Error("Wrong value for id.")
	}
	if len(result.Name()) <= 0 {
		t.Error("Wrong value for name.")
	}
}

func Test_ServersList(t *testing.T) {
	servers := &Servers{}
	results, err := servers.List()
	if err != nil {
		t.Errorf("Servers List failed with %s.", err)
		return
	}
	if results == nil {
		t.Error("Servers List failed with no result.")
		return
	}
	for _, result := range results {
		if len(result.Id()) <= 0 {
			t.Error("Wrong value for id.")
		}
		if len(result.Name()) <= 0 {
			t.Error("Wrong value for name.")
		}
	}
}

func Test_ServersShow(t *testing.T) {
	servers := &Servers{}
	result, err := servers.Show(TestServer)
	if err != nil {
		t.Errorf("Servers Show failed with %s.", err)
		return
	}
	if result == nil {
		t.Error("Servers Show failed with no result.")
		return
	}
	if len(result.Id()) <= 0 {
		t.Error("Wrong value for id.")
	}
	if len(result.Name()) <= 0 {
		t.Error("Wrong value for name.")
	}
}

func Test_ServersStart(t *testing.T) {
	servers := &Servers{}
	ok, err := servers.Start(TestServer)
	if !ok {
		t.Error("Servers Start failed.")
	}
	if err != nil {
		t.Errorf("Servers Start failed with %s.", err)
	}
}

func Test_ServersReboot(t *testing.T) {
	servers := &Servers{}
	ok, err := servers.Reboot(TestServer)
	if !ok {
		t.Error("Servers Reboot failed.")
	}
	if err != nil {
		t.Errorf("Servers Reboot failed with %s.", err)
	}
}

func Test_ServersStop(t *testing.T) {
	servers := &Servers{}
	ok, err := servers.Stop(TestServer)
	if !ok {
		t.Error("Servers Stop failed.")
	}
	if err != nil {
		t.Errorf("Servers Stop failed with %s.", err)
	}
}

func Test_ServersDestroy(t *testing.T) {
	servers := &Servers{}
	ok, err := servers.Destroy(TestServer)
	if !ok {
		t.Error("Servers Destroy failed.")
	}
	if err != nil {
		t.Errorf("Servers Destroy failed with %s.", err)
	}
}
