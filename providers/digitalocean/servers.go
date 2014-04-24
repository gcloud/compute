package digitalocean

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	//"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	p "github.com/gcloud/compute/providers"
	"github.com/mitchellh/mapstructure"
)

var name = "digitalocean"

func init() {
	provider := p.GetProvider(name)
	p.RegisterServers(name, &Servers{provider: provider})
}

type result struct {
	Status        string
	Droplet       *droplet
	Error_message string `json:",omitempty"`
}

type results struct {
	Status        string
	Droplets      []*droplet
	Error_message string `json:",omitempty"`
}

type droplet struct {
	Id                 int
	Name               string
	Image_id           int
	Size_id            int
	Event_id           int
	Region_id          int
	Backups_active     bool
	Ip_address         string
	Private_ip_address string
	Locked             bool
	Status             string
	Created_at         time.Time
}

func (d *droplet) toServer() *Server {
	return &Server{data: d}
}

type Server struct {
	data *droplet
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

type Servers struct {
	provider *p.Provider
}

func (s *Servers) NewServer(m p.Map) p.Server {
	var server *Server
	err := mapstructure.Decode(m, &server)
	if err != nil {
		return nil
	}
	return server
}

// List droplets available on the account.
func (s *Servers) List() ([]p.Server, error) {
	var r []p.Server
	response, err := request(s.provider, "GET", "/droplets", nil)
	if err != nil {
		return r, err
	}
	var results results
	err = json.Unmarshal(response, &results)
	if err != nil {
		return r, err
	}
	if results.Status != "OK" {
		return r, errors.New(results.Error_message)
	}
	servers := make([]p.Server, 0)
	for _, v := range results.Droplets {
		servers = append(servers, v.toServer())
	}
	return servers, nil
}

// Show droplet information for a given id.
func (s *Servers) Show(id string) (p.Server, error) {
	var r p.Server
	response, err := request(s.provider, "GET", "/droplets/"+id, nil)
	if err != nil {
		return r, err
	}
	var result result
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
func (s *Servers) Create(n interface{}) (p.Server, error) {
	var r p.Server
	response, err := request(s.provider, "GET", "/droplets/new", n)
	if err != nil {
		return r, err
	}
	var result result
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
func (s *Servers) Destroy(id string) (bool, error) {
	return event(s, "/droplets/%s/destroy", id)
}

// Reboot a droplet.
func (s *Servers) Reboot(id string) (bool, error) {
	return event(s, "/droplets/%s/reboot", id)
}

// Start a droplet.
func (s *Servers) Start(id string) (bool, error) {
	return event(s, "/droplets/%s/power_on", id)
}

// Stop a droplet.
func (s *Servers) Stop(id string) (bool, error) {
	return event(s, "/droplets/%s/power_off", id)
}

func event(s *Servers, url string, id string) (bool, error) {
	r := false
	response, err := request(s.provider, "GET", fmt.Sprintf(url, id), nil)
	if err != nil {
		return r, err
	}
	var result struct {
		Status        string
		Event_id      int
		Error_message string
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return r, err
	}
	if result.Status != "OK" {
		return r, errors.New(result.Error_message)
	}
	return true, nil
}

// Send HTTP Request, return data
func request(provider *p.Provider, method string, path string, data interface{}) ([]byte, error) {
	url := provider.Endpoint + path
	url += fmt.Sprintf("?client_id=%s&api_key=%s", provider.Account.Id, provider.Account.Key)
	url += fmt.Sprintf("&%s", urlencode(data))
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func urlencode(data interface{}) string {
	var q bytes.Buffer
	switch d := data.(type) {
	case map[string]interface{}:
		for k, value := range d {
			q.WriteString(url.QueryEscape(k))
			q.WriteByte('=')
			switch v := value.(type) {
			default:
				q.WriteString(url.QueryEscape(fmt.Sprintf("%v", v)))
			}
			q.WriteByte('&')
		}
		s := q.String()
		return s[0 : len(s)-1]
	}
	return ""
}
