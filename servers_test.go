// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	"encoding/json"
	"testing"
)

type MockServer struct {
	id   string
	name string
}

func (m *MockServer) Id() string {
	return m.id
}
func (m *MockServer) Name() string {
	return m.name
}
func (m *MockServer) State() string {
	return "running"
}
func (m *MockServer) Ips(t string) []string {
	return []string{}
}
func (m *MockServer) Size() string {
	return ""
}
func (m *MockServer) Image() string {
	return ""
}
func (m *MockServer) String() string {
	b, _ := m.MarshalJSON()
	return string(b)
}
func (m *MockServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(Map{
		"id": m.Id(), "name": m.Name(),
	})
}
func (m *MockServer) Map() Map {
	return Map{
		"id": m.Id(), "name": m.Name(),
	}
}

type MockServers struct{}

func (s *MockServers) New(m Map) Server {
	return &MockServer{name: "My Server", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"}
}
func (s *MockServers) List() ([]Server, error) {
	results := make([]Server, 0)
	r := s.New(nil)
	return append(results, r), nil
}
func (s *MockServers) Show(Server) (Server, error) {
	r := s.New(nil)
	return r, nil
}
func (s *MockServers) Create(Server) (Server, error) {
	r := s.New(nil)
	return r, nil
}
func (s *MockServers) Destroy(Server) (bool, error) {
	return true, nil
}
func (s *MockServers) Reboot(Server) (bool, error) {
	return true, nil
}
func (s *MockServers) Start(Server) (bool, error) {
	return true, nil
}
func (s *MockServers) Stop(Server) (bool, error) {
	return true, nil
}

func init() {
	RegisterServers("mock", &MockServers{})
}

func Test_ServersList(t *testing.T) {
	servers := GetServers("mock", nil)
	results, err := servers.List()
	if err != nil {
		t.Error("Servers List failed with %s.", err)
	}
	if results == nil {
		t.Error("Results should not be nil.")
		return
	}
	for _, v := range results {
		if v.Id() != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
			t.Error("Wrong value for Id.")
		}
		if v.Name() != "My Server" {
			t.Error("Wrong value for Name.")
		}
	}
}

func Test_ServersShow(t *testing.T) {
	servers := GetServers("mock", nil)
	result, err := servers.Show(&MockServer{id: "616fb98f-46ca-475e-917e-2563e5a8cd19"})
	if err != nil {
		t.Errorf("Servers Show failed with %s.", err)
	}
	if result == nil {
		t.Error("Results should not be nil.")
		return
	}
	r := result
	if r.Id() != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		t.Error("Wrong value for Id.")
	}
	if r.Name() != "My Server" {
		t.Error("Wrong value for Name.")
	}
}

func Test_ServersCreate(t *testing.T) {
	servers := GetServers("mock", nil)
	result, err := servers.Create(&MockServer{
		name: "My Server",
	})
	if err != nil {
		t.Errorf("Servers Create failed with %s.", err)
	}
	if result == nil {
		t.Error("Results should not be nil.")
		return
	}
	r := result
	if r.Id() != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		t.Error("Wrong value for Id.")
	}
	if r.Name() != "My Server" {
		t.Error("Wrong value for Name.")
	}
}

func Test_ServersDestroy(t *testing.T) {
	servers := GetServers("mock", nil)
	ok, err := servers.Destroy(&MockServer{id: "616fb98f-46ca-475e-917e-2563e5a8cd19"})
	if !ok {
		t.Error("Servers Destroy failed.")
	}
	if err != nil {
		t.Errorf("Servers Destroy failed with %s.", err)
	}
}

func Test_ServersReboot(t *testing.T) {
	servers := GetServers("mock", nil)
	ok, err := servers.Reboot(&MockServer{id: "616fb98f-46ca-475e-917e-2563e5a8cd19"})
	if !ok {
		t.Error("Servers Reboot failed.")
	}
	if err != nil {
		t.Errorf("Servers Reboot failed with %s.", err)
	}
}

func Test_ServersStart(t *testing.T) {
	servers := GetServers("mock", nil)
	ok, err := servers.Start(&MockServer{id: "616fb98f-46ca-475e-917e-2563e5a8cd19"})
	if !ok {
		t.Error("Servers Start failed.")
	}
	if err != nil {
		t.Errorf("Servers Start failed with %s.", err)
	}
}

func Test_ServersStop(t *testing.T) {
	servers := GetServers("mock", nil)
	ok, err := servers.Stop(&MockServer{id: "616fb98f-46ca-475e-917e-2563e5a8cd19"})
	if !ok {
		t.Error("Servers Stop failed.")
	}
	if err != nil {
		t.Errorf("Servers Stop failed with %s.", err)
	}
}
