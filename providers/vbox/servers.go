// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package vbox

import (
	"encoding/json"
	p "github.com/gcloud/compute/providers"
	"os/exec"
	"regexp"
	"strings"
)

type Servers struct{}

// List servers available on the account.
func (s *Servers) List() ([]byte, error) {
	c := exec.Command("VBoxManage", "list", "vms")
	output, err := c.CombinedOutput()
	re := regexp.MustCompile("(?P<name>[A-z -]+)(?P<id>[A-z0-9-]+)")
	r := strings.NewReplacer("\r", "")
	results := strings.Split(strings.Trim(r.Replace(string(output)), "\n"), "\n")
	responses := make([]interface{}, 0)
	for _, v := range results {
		matches := re.FindAllString(v, -1)
		responses = append(responses, map[string]interface{}{
			"id":   matches[1],
			"name": matches[0],
		})
	}
	b, err := json.Marshal(responses)
	return b, err
}

// Show server information for a given id.
func (s *Servers) Show(id string) ([]byte, error) {
	return []byte(`{"id": "1", "name": "test"}`), nil
}

// Create a server.
func (s *Servers) Create(n interface{}) ([]byte, error) {
	return []byte(`{"id": "1", "name": "test"}`), nil
}

// Destroy a server.
func (s *Servers) Destroy(id string) (bool, error) {
	return false, nil
}

// Reboot a server.
func (s *Servers) Reboot(id string) (bool, error) {
	return false, nil
}

// Start a server that is stopped.
func (s *Servers) Start(id string) (bool, error) {
	return false, nil
}

// Stop a server that is running.
func (s *Servers) Stop(id string) (bool, error) {
	return false, nil
}

func init() {
	p.RegisterServers("vbox", &Servers{})
}
