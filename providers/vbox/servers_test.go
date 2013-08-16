// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package vbox

import (
	"encoding/json"
	"testing"
)

type MockServer struct {
	Id         string
	Name       string
	State      string
	PublicIps  []string
	PrivateIps []string
	Size       string
	Image      string
}

type MockResponse struct {
	id   string
	name string
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
	var responses []map[string]string
	err = json.Unmarshal(results, &responses)

	if err != nil {
		t.Error("Servers List failed with " + err.Error() + "(bool, error).")
	}
	for _, r := range responses {
		if len(r["id"]) <= 0 {
			t.Error("Wrong value for id.")
		}
		if len(r["name"]) <= 0 {
			t.Error("Wrong value for name.")
		}
	}
}

func Test_ServersShow(t *testing.T) {
	servers := &Servers{}
	result, err := servers.Show("616fb98f-46ca-475e-917e-2563e5a8cd19")
	if err != nil {
		t.Error("Servers Show failed with " + err.Error() + ".")
	}
	if result == nil {
		t.Error("Results should not be nil.")
	}
	var r map[string]string
	err = json.Unmarshal(result, &r)
	if err != nil {
		t.Error("Servers Show failed with " + err.Error() + ".")
	}
	if len(r["id"]) <= 0 {
		t.Error("Wrong value for id.")
	}
	if len(r["name"]) <= 0 {
		t.Error("Wrong value for name.")
	}
}

func Test_ServersCreate(t *testing.T) {
	servers := &Servers{}
	result, err := servers.Create(&MockServer{Name: "My Server"})
	if err != nil {
		t.Error("Servers Create failed with " + err.Error() + ".")
	}
	if result == nil {
		t.Error("Results should not be nil.")
	}
	var r map[string]string
	err = json.Unmarshal(result, &r)
	if err != nil {
		t.Error("Servers Create failed with " + err.Error() + ".")
	}
	if len(r["id"]) <= 0 {
		t.Error("Wrong value for id.")
	}
	if len(r["name"]) <= 0 {
		t.Error("Wrong value for name.")
	}
}

func Test_ServersDestroy(t *testing.T) {
	servers := &Servers{}
	ok, err := servers.Destroy("616fb98f-46ca-475e-917e-2563e5a8cd19")
	if !ok {
		t.Error("Servers Destroy failed.")
	}
	if err != nil {
		t.Error("Servers Destroy failed with " + err.Error() + ".")
	}
}

func Test_ServersReboot(t *testing.T) {
	servers := &Servers{}
	ok, err := servers.Reboot("616fb98f-46ca-475e-917e-2563e5a8cd19")
	if !ok {
		t.Error("Servers Reboot failed.")
	}
	if err != nil {
		t.Error("Servers Reboot failed with " + err.Error() + ".")
	}
}

func Test_ServersStart(t *testing.T) {
	servers := &Servers{}
	ok, err := servers.Start("616fb98f-46ca-475e-917e-2563e5a8cd19")
	if !ok {
		t.Error("Servers Start failed.")
	}
	if err != nil {
		t.Error("Servers Start failed with " + err.Error() + ".")
	}
}

func Test_ServersStop(t *testing.T) {
	servers := &Servers{}
	ok, err := servers.Stop("616fb98f-46ca-475e-917e-2563e5a8cd19")
	if !ok {
		t.Error("Servers Stop failed.")
	}
	if err != nil {
		t.Error("Servers Stop failed with " + err.Error() + ".")
	}
}
