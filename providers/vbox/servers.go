// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package vbox

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	p "github.com/gcloud/compute/providers"
	"os/exec"
	"regexp"
	"strings"
)

func init() {
	p.RegisterServers("vbox", &Servers{})
}

type Servers struct{}

// List servers available on the account.
func (s *Servers) List() ([]byte, error) {
	c := exec.Command("VBoxManage", "list", "vms")
	output, err := c.CombinedOutput()
	if err != nil {
		print(string(output))
		return []byte(`[]`), err
	}
	if string(output) == "" {
		return []byte(`[]`), err
	}
	r := strings.NewReplacer("\r", "")
	results := strings.Split(strings.Trim(r.Replace(string(output)), "\n"), "\n")
	re := regexp.MustCompile("(?P<name>[A-z -]+)(?P<id>[A-z0-9-]+)")
	responses := make([]interface{}, 0)
	for _, v := range results {
		matches := re.FindAllString(v, -1)
		responses = append(responses, map[string]interface{}{
			"Id":   matches[1],
			"Name": matches[0],
		})
	}
	b, err := json.Marshal(responses)
	return b, err
}

// Show server information for a given id.
func (s *Servers) Show(id string) ([]byte, error) {
	c := exec.Command("VBoxManage", "showvminfo", id, "--machinereadable")
	output, err := c.CombinedOutput()
	if err != nil {
		print(string(output))
		return []byte(`{}`), err
	}
	config := strings.Split(string(output), "Time offset=0")
	if len(config) != 2 {
		return []byte(`{}`), nil
	}
	var response struct {
		Id   string `toml:"UUID"`
		Name string `toml:"name"`
	}
	if _, err := toml.Decode(config[0], &response); err != nil {
		return nil, err
	}
	b, err := json.Marshal(response)
	return b, err
}

// Create a server.
func (s *Servers) Create(n *p.Server) ([]byte, error) {
	c := exec.Command("VBoxManage", "createvm", "--name", n.Name, "--register")
	output, err := c.CombinedOutput()
	if err != nil {
		return output, err
	}
	re := regexp.MustCompile("([A-z0-9]{8}-[A-z0-9]{4}-[A-z0-9]{4}-[A-z0-9]{4}-[A-z0-9]{12})")
	matches := re.FindAllString(string(output), -1)
	if len(matches) < 1 {
		return output, err
	}
	n.Id = matches[0]
	b, err := json.Marshal(n)
	return b, err
}

// Destroy a server.
func (s *Servers) Destroy(id string) (bool, error) {
	c := exec.Command("VBoxManage", "unregistervm", id, "--delete")
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
func (s *Servers) Start(id string) (bool, error) {
	c := exec.Command("VBoxManage", "startvm", id, "--type", "headless")
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
func (s *Servers) Reboot(id string) (bool, error) {
	c := exec.Command("VBoxManage", "controlvm", id, "reset")
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
func (s *Servers) Stop(id string) (bool, error) {
	c := exec.Command("VBoxManage", "controlvm", id, "poweroff")
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
