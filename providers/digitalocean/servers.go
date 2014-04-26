package digitalocean

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gcloud/compute"
	"github.com/mitchellh/mapstructure"
)

func init() {
	name := "digitalocean"
	provider := compute.GetProvider(name)
	provider.Endpoint = "https://api.digitalocean.com"
	compute.RegisterServers(name, &Servers{provider: provider})
}

type sResult struct {
	Status        string
	Droplet       *droplet
	Error_message string `json:",omitempty"`
}

type sResults struct {
	Status        string
	Droplets      []*droplet
	Error_message string `json:",omitempty"`
}

type droplet struct {
	Id                 int
	Name               string
	Image_id           int
	Size_id            int
	Region_id          int
	Ssh_key_ids        string
	Event_id           int       `json:",omitempty"`
	Backups_active     bool      `json:",omitempty"`
	Ip_address         string    `json:",omitempty"`
	Private_ip_address string    `json:",omitempty"`
	Locked             bool      `json:",omitempty"`
	Status             string    `json:",omitempty"`
	Created_at         time.Time `json:",omitempty"`
}

func (d *droplet) toServer() *Server {
	return &Server{data: d}
}

type Server struct {
	data    *droplet
	generic compute.Map
}

func (s *Server) Id() string {
	return fmt.Sprintf("%d", s.data.Id)
}
func (s *Server) Name() string {
	return s.data.Name
}
func (s *Server) State() string {
	return s.data.Status
}
func (s *Server) Ips(t string) []string {
	if t == "private" {
		return []string{s.data.Private_ip_address}
	}
	return []string{s.data.Ip_address}
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
	return json.Marshal(s.data)
}
func (s *Server) Map() compute.Map {
	return s.generic
}

type Servers struct {
	provider *compute.Provider
}

func (s *Servers) New(m compute.Map) compute.Server {
	var drop droplet
	err := mapstructure.Decode(m, &drop)
	if err != nil {
		return nil
	}
	server := drop.toServer()
	server.generic = m
	return server
}

// List droplets available on the account.
func (s *Servers) List() ([]compute.Server, error) {
	var r []compute.Server
	response, err := request(s.provider, "GET", "/droplets", nil)
	if err != nil {
		return r, err
	}
	var results sResults
	err = json.Unmarshal(response, &results)
	if err != nil {
		return r, err
	}
	if results.Status != "OK" {
		return r, errors.New(results.Error_message)
	}
	servers := make([]compute.Server, 0)
	for _, v := range results.Droplets {
		servers = append(servers, v.toServer())
	}
	return servers, nil
}

// Show droplet information for a given id.
func (s *Servers) Show(server compute.Server) (compute.Server, error) {
	var r compute.Server
	response, err := request(s.provider, "GET", "/droplets/"+server.Id(), nil)
	if err != nil {
		return r, err
	}
	var result sResult
	err = json.Unmarshal(response, &result)
	if err != nil {
		return r, err
	}
	if result.Status != "OK" {
		return r, errors.New(result.Error_message)
	}
	return result.Droplet.toServer(), nil
}

// Create a droplet.
func (s *Servers) Create(server compute.Server) (compute.Server, error) {
	var r compute.Server
	fmt.Printf("%#v", server)
	response, err := request(s.provider, "GET", "/droplets/new", server.Map())
	if err != nil {
		return r, err
	}
	var result sResult
	err = json.Unmarshal(response, &result)
	if err != nil {
		return r, err
	}
	if result.Status != "OK" {
		return r, errors.New(result.Error_message)
	}
	var i int
	var event struct {
		Id            string
		Action_status string
		Droplet_id    int
		Event_type_id int
		Percentage    string
	}
	for i < 5 && result.Droplet.Status != "active" && event.Action_status != "done" {
		uri := fmt.Sprintf("/events/%d", result.Droplet.Event_id)
		response, err := request(s.provider, "GET", uri, nil)
		if err == nil {
			err = json.Unmarshal(response, &event)
			if err == nil && event.Action_status == "done" {
				return result.Droplet.toServer(), nil
			}
		}
		i++
		time.Sleep(time.Second * 30)
	}
	return result.Droplet.toServer(), nil
}

// Delete a droplet.
func (s *Servers) Destroy(server compute.Server) (bool, error) {
	return event(s.provider, "/droplets/%s/destroy", server.Id())
}

// Reboot a droplet.
func (s *Servers) Reboot(server compute.Server) (bool, error) {
	return event(s.provider, "/droplets/%s/reboot", server.Id())
}

// Start a droplet.
func (s *Servers) Start(server compute.Server) (bool, error) {
	return event(s.provider, "/droplets/%s/power_on", server.Id())
}

// Stop a droplet.
func (s *Servers) Stop(server compute.Server) (bool, error) {
	return event(s.provider, "/droplets/%s/power_off", server.Id())
}
