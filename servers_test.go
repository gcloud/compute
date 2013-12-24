// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	"testing"

	p "github.com/gcloud/compute/providers"
)

type MockServers struct{}

// List servers available on the account.
func (s *MockServers) List() ([]byte, error) {
	return []byte(`[{"Name":"My Server","Id": "616fb98f-46ca-475e-917e-2563e5a8cd19"}]`), nil
}
func (s *MockServers) Show(id string) ([]byte, error) {
	return []byte(`{"Name":"My Server","Id": "616fb98f-46ca-475e-917e-2563e5a8cd19"}`), nil
}
func (s *MockServers) Create(n *p.Server) ([]byte, error) {
	return []byte(`{"Name":"My Server","Id": "616fb98f-46ca-475e-917e-2563e5a8cd19"}`), nil
}
func (s *MockServers) Destroy(id string) (bool, error) {
	return true, nil
}
func (s *MockServers) Reboot(id string) (bool, error) {
	return true, nil
}
func (s *MockServers) Start(id string) (bool, error) {
	return true, nil
}
func (s *MockServers) Stop(id string) (bool, error) {
	return true, nil
}

func init() {
	p.RegisterServers("mock", &MockServers{})
}

func Test_ServersList(t *testing.T) {
	servers := &Servers{Provider: "mock"}
	results, err := servers.List()
	if err != nil {
		t.Error("Servers List failed with " + err.Error() + "(bool, error).")
	}
	if results == nil {
		t.Error("Results should not be nil.")
	}
	for _, v := range *results {
		if v.Id != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
			t.Error("Wrong value for Id.")
		}
		if v.Name != "My Server" {
			t.Error("Wrong value for Name.")
		}
	}
}

func Test_ServersShow(t *testing.T) {
	servers := &Servers{Provider: "mock"}
	result, err := servers.Show("616fb98f-46ca-475e-917e-2563e5a8cd19")
	if err != nil {
		t.Error("Servers Show failed with " + err.Error() + ".")
	}
	if result == nil {
		t.Error("Results should not be nil.")
	}
	r := *result
	if r.Id != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		t.Error("Wrong value for Id.")
	}
	if r.Name != "My Server" {
		t.Error("Wrong value for Name.")
	}
}

func Test_ServersCreate(t *testing.T) {
	servers := &Servers{Provider: "mock"}
	result, err := servers.Create(&p.Server{
		Name:  "My Server",
		Image: "70a599e0-31e7-49b7-b260-868f441e862b",
		Size:  "1"})
	if err != nil {
		t.Error("Servers Create failed with " + err.Error() + ".")
	}
	if result == nil {
		t.Error("Results should not be nil.")
	}
	r := *result
	if r.Id != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		t.Error("Wrong value for Id.")
	}
	if r.Name != "My Server" {
		t.Error("Wrong value for Name.")
	}
}

func Test_ServersDestroy(t *testing.T) {
	servers := &Servers{Provider: "mock"}
	ok, err := servers.Destroy("616fb98f-46ca-475e-917e-2563e5a8cd19")
	if !ok {
		t.Error("Servers Destroy failed.")
	}
	if err != nil {
		t.Error("Servers Destroy failed with " + err.Error() + ".")
	}
}

func Test_ServersReboot(t *testing.T) {
	servers := &Servers{Provider: "mock"}
	ok, err := servers.Reboot("616fb98f-46ca-475e-917e-2563e5a8cd19")
	if !ok {
		t.Error("Servers Reboot failed.")
	}
	if err != nil {
		t.Error("Servers Reboot failed with " + err.Error() + ".")
	}
}

func Test_ServersStart(t *testing.T) {
	servers := &Servers{Provider: "mock"}
	ok, err := servers.Start("616fb98f-46ca-475e-917e-2563e5a8cd19")
	if !ok {
		t.Error("Servers Start failed.")
	}
	if err != nil {
		t.Error("Servers Start failed with " + err.Error() + ".")
	}
}

func Test_ServersStop(t *testing.T) {
	servers := &Servers{Provider: "mock"}
	ok, err := servers.Stop("616fb98f-46ca-475e-917e-2563e5a8cd19")
	if !ok {
		t.Error("Servers Stop failed.")
	}
	if err != nil {
		t.Error("Servers Stop failed with " + err.Error() + ".")
	}
}
