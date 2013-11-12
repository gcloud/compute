// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package vbox

import (
	"encoding/json"
	p "github.com/gcloud/compute/providers"
	"testing"
)

var ServerName = "GCloudServer"

func Test_ServersCreate(t *testing.T) {
	servers := &Servers{}
	result, err := servers.Create(&p.Server{Name: ServerName})
	if err != nil {
		t.Error("Servers Create failed with " + err.Error() + ".")
	}
	if result == nil {
		t.Error("Results should not be nil.")
	}
	var server p.Server
	err = json.Unmarshal(result, &server)
	if err != nil {
		t.Error("Servers Create failed with " + err.Error() + ".")
	}
	if len(server.Id) <= 0 {
		t.Error("Wrong value for id.")
	}
	if len(server.Name) <= 0 {
		t.Error("Wrong value for name.")
	}
}

func Test_ServersList(t *testing.T) {
	servers := &Servers{}
	results, err := servers.List()
	if err != nil {
		t.Error("Servers List failed with " + err.Error() + "(bool, error).")
	}
	if results == nil {
		t.Error("Results should not be nil.")
	}
	var response []p.Server
	err = json.Unmarshal(results, &response)

	if err != nil {
		t.Error("Servers List failed with " + err.Error() + "(bool, error).")
	}
	for _, server := range response {
		if len(server.Id) <= 0 {
			t.Error("Wrong value for id.")
		}
		if len(server.Name) <= 0 {
			t.Error("Wrong value for name.")
		}
	}
}

func Test_ServersShow(t *testing.T) {
	servers := &Servers{}
	result, err := servers.Show(ServerName)
	if err != nil {
		t.Error("Servers Show failed with " + err.Error() + ".")
	}
	if result == nil {
		t.Error("Results should not be nil.")
	}
	var server p.Server
	err = json.Unmarshal(result, &server)
	if err != nil {
		t.Error("Servers Show failed with " + err.Error() + ".")
	}
	if len(server.Id) <= 0 {
		t.Error("Wrong value for id.")
	}
	if len(server.Name) <= 0 {
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
