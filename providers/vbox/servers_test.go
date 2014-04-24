// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package vbox

import (
	"testing"
)

var ServerName = "GCloudServer"

func Test_ServersCreate(t *testing.T) {
	servers := &Servers{}
	result, err := servers.Create(&Server{name: ServerName})
	if err != nil {
		t.Error("Servers Create failed with " + err.Error() + ".")
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
		t.Error("Servers List failed with " + err.Error() + "(bool, error).")
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
	result, err := servers.Show(ServerName)
	if err != nil {
		t.Error("Servers Show failed with " + err.Error() + ".")
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
	ok, err := servers.Start(ServerName)
	if !ok {
		t.Error("Servers Start failed.")
	}
	if err != nil {
		t.Error("Servers Start failed with " + err.Error() + ".")
	}
}

func Test_ServersReboot(t *testing.T) {
	servers := &Servers{}
	ok, err := servers.Reboot(ServerName)
	if !ok {
		t.Error("Servers Reboot failed.")
	}
	if err != nil {
		t.Error("Servers Reboot failed with " + err.Error() + ".")
	}
}

func Test_ServersStop(t *testing.T) {
	servers := &Servers{}
	ok, err := servers.Stop(ServerName)
	if !ok {
		t.Error("Servers Stop failed.")
	}
	if err != nil {
		t.Error("Servers Stop failed with " + err.Error() + ".")
	}
}

func Test_ServersDestroy(t *testing.T) {
	servers := &Servers{}
	ok, err := servers.Destroy(ServerName)
	if !ok {
		t.Error("Servers Destroy failed.")
	}
	if err != nil {
		t.Error("Servers Destroy failed with " + err.Error() + ".")
	}
}
