// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package vbox

import (
	"encoding/json"
	"os/exec"
	"regexp"
	"strings"

	"github.com/gcloud/compute"
	"github.com/mitchellh/mapstructure"
)

func init() {
	compute.RegisterServers("vbox", &Servers{})
}

type Server struct {
	id    string
	name  string
	state string
}

func (s *Server) Id() string {
	return s.id
}
func (s *Server) Name() string {
	return s.name
}
func (s *Server) State() string {
	return s.state
}
func (s *Server) Ips(t string) []string {
	return []string{}
}
func (s *Server) Size() string {
	return ""
}
func (s *Server) Image() string {
	return ""
}
func (s *Server) String() string {
	b, err := s.MarshalJSON()
	if err != nil {
		return ""
	}
	return string(b)
}
func (s *Server) MarshalJSON() ([]byte, error) {
	return json.Marshal(compute.Map{
		"id": s.Id(), "name": s.Name(),
	})
}

type Servers struct{}

func (s *Servers) New(m compute.Map) compute.Server {
	var server *Server
	err := mapstructure.Decode(m, &server)
	if err != nil {
		return nil
	}
	return server
}

// List servers available on the account.
func (s *Servers) List() ([]compute.Server, error) {
	var r []compute.Server
	c := exec.Command("VBoxManage", "list", "vms")
	output, err := c.CombinedOutput()
	if err != nil {
		print(string(output))
		return r, err
	}
	if string(output) == "" {
		return r, err
	}
	print(string(output))
	replacer := strings.NewReplacer("\r", "")
	results := strings.Split(strings.Trim(replacer.Replace(string(output)), "\n"), "\n")
	re := regexp.MustCompile("(?P<name>[A-z -]+)(?P<id>[A-z0-9-]+)")
	responses := make([]compute.Server, 0)
	for _, v := range results {
		matches := re.FindAllString(v, -1)
		if len(matches) < 2 {
			continue
		}
		responses = append(responses, &Server{
			id:   matches[1],
			name: matches[0],
		})
	}
	return responses, nil
}

// Show server information for a given id.
func (s *Servers) Show(server compute.Server) (compute.Server, error) {
	var r compute.Server
	c := exec.Command("VBoxManage", "showvminfo", server.Id(), "--machinereadable")
	output, err := c.CombinedOutput()
	if err != nil {
		print(string(output))
		return r, err
	}
	config := strings.Split(string(output), "Time offset=0")
	if len(config) != 2 {
		return r, nil
	}
	result := make(map[string]string, 0)
	for _, s := range strings.Split(strings.TrimSpace(string(config[0])), "\n") {
		v := strings.Split(s, "=")
		if len(v) > 1 {
			result[v[0]] = v[1]
		}
	}
	return &Server{result["UUID"], result["name"], ""}, nil
}

// Create a server.
func (s *Servers) Create(server compute.Server) (compute.Server, error) {
	c := exec.Command("VBoxManage", "createvm", "--name", server.Name(), "--register")
	output, err := c.CombinedOutput()
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile("([A-z0-9]{8}-[A-z0-9]{4}-[A-z0-9]{4}-[A-z0-9]{4}-[A-z0-9]{12})")
	matches := re.FindAllString(string(output), -1)
	if len(matches) < 1 {
		return nil, err
	}
	return &Server{matches[0], server.Name(), ""}, nil
}

// Destroy a server.
func (s *Servers) Destroy(server compute.Server) (bool, error) {
	c := exec.Command("VBoxManage", "unregistervm", server.Id(), "--delete")
	output, err := c.CombinedOutput()
	if err != nil {
		print(string(output))
		return false, err
	}
	if output != nil {
		return true, nil
	}
	return false, err
}

// Start a server that is stopped.
func (s *Servers) Start(server compute.Server) (bool, error) {
	c := exec.Command("VBoxManage", "startvm", server.Id(), "--type", "headless")
	output, err := c.CombinedOutput()
	if err != nil {
		print(string(output))
		return false, err
	}
	if output != nil {
		return true, nil
	}
	return false, err
}

// Reboot a server.
func (s *Servers) Reboot(server compute.Server) (bool, error) {
	c := exec.Command("VBoxManage", "controlvm", server.Id(), "reset")
	output, err := c.CombinedOutput()
	if err != nil {
		print(string(output))
		return false, err
	}
	if output != nil {
		return true, nil
	}
	return false, err
}

// Stop a server that is running.
func (s *Servers) Stop(server compute.Server) (bool, error) {
	c := exec.Command("VBoxManage", "controlvm", server.Id(), "poweroff")
	output, err := c.CombinedOutput()
	if err != nil {
		print(string(output))
		return false, err
	}
	if output != nil {
		return true, nil
	}
	return false, err
}
