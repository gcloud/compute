// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package digitalocean

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	p "github.com/gcloud/compute/providers"
	"github.com/gcloud/identity"
)

var DOServerName = "GCloudServerOnDO"

var account = &identity.Account{
	Id:  "id",
	Key: "key",
}

var provider *p.Provider

func init() {
	provider = p.GetProvider("digitalocean")
	provider.SetAccount(account)
}

func Test_ServersCreate(t *testing.T) {
	r := []byte(`{
		"action_status": "done",
		"status": "OK",
		"droplet": {
			"id": 100824,
			"name": "GCloudServerOnDO",
			"image_id": 419,
			"size_id": 32,
			"event_id": 7499
		}
	}`)
	ts := newTestResponse(r)
	defer ts.Close()
	servers := &Servers{provider}
	result, err := servers.Create(p.Map{
		"name":      DOServerName,
		"image_id":  1505447,
		"size_id":   66,
		"region_id": 1,
		"ssh_keys":  18420,
	})
	if err != nil {
		t.Error("Servers Create failed with " + err.Error() + ".")
	}
	if result == nil {
		t.Error("Results should not be nil.")
	}
	if len(result.Id()) <= 0 {
		t.Error("Wrong value for id.")
	}
	if len(result.Name()) <= 0 {
		t.Error("Wrong value for name.")
	}
}

func Test_ServersList(t *testing.T) {
	r := []byte(`{
		"status": "OK",
		"droplets": [
			{
			  "id": 100823,
			  "name": "test222",
			  "image_id": 420,
			  "size_id":33,
			  "region_id": 1,
			  "backups_active": false,
			  "ip_address": "127.0.0.1",
			  "private_ip_address": null,
			  "locked": false,
			  "status": "active",
			  "created_at": "2013-01-01T09:30:00Z"
			}
		]
	}`)
	ts := newTestResponse(r)
	defer ts.Close()
	servers := &Servers{provider}
	results, err := servers.List()
	if err != nil {
		t.Error("Servers List failed with " + err.Error())
	}
	if results == nil {
		t.Error("Results should not be nil.")
	}
	for _, server := range results {
		if len(server.Id()) <= 0 {
			t.Error("Wrong value for id.")
		}
		if len(server.Name()) <= 0 {
			t.Error("Wrong value for name.")
		}
	}
}

func Test_ServersShow(t *testing.T) {
	r := []byte(`{
		"status": "OK",
		"droplet": {
			"id": 100823,
			"image_id": 420,
			"name": "test222",
			"region_id": 1,
			"size_id": 33,
			"backups_active": false,
			"backups": [],
			"snapshots": [],
			"ip_address": "127.0.0.1",
			"private_ip_address": null,
			"locked": false,
			"status": "active"
		}
	}`)
	ts := newTestResponse(r)
	defer ts.Close()
	servers := &Servers{provider}
	result, err := servers.Show("100823")
	if err != nil {
		t.Error("Servers Show failed with " + err.Error() + ".")
	}
	if result == nil {
		t.Error("Results should not be nil.")
	}
	if len(result.Id()) <= 0 {
		t.Error("Wrong value for id.")
	}
	if len(result.Name()) <= 0 {
		t.Error("Wrong value for name.")
	}
}

func Test_ServersStart(t *testing.T) {
	r := []byte(`{
		"status": "OK",
		"event_id": 7504
	}`)
	ts := newTestResponse(r)
	defer ts.Close()
	servers := &Servers{provider}
	ok, err := servers.Start("100823")
	if !ok {
		t.Error("Servers Start failed.")
	}
	if err != nil {
		t.Error("Servers Start failed with " + err.Error() + ".")
	}
}

func Test_ServersStop(t *testing.T) {
	r := []byte(`{
		"status": "OK",
		"event_id": 7504
	}`)
	ts := newTestResponse(r)
	defer ts.Close()
	servers := &Servers{provider}
	ok, err := servers.Stop("100823")
	if !ok {
		t.Error("Servers Stop failed.")
	}
	if err != nil {
		t.Error("Servers Stop failed with " + err.Error() + ".")
	}
}

func Test_ServersReboot(t *testing.T) {
	r := []byte(`{
		"status": "OK",
		"event_id": 7504
	}`)
	ts := newTestResponse(r)
	defer ts.Close()
	servers := &Servers{provider}
	ok, err := servers.Reboot("100823")
	if !ok {
		t.Error("Servers Reboot failed.")
	}
	if err != nil {
		t.Error("Servers Reboot failed with " + err.Error() + ".")
	}
}

func Test_ServersDestroy(t *testing.T) {
	r := []byte(`{
		"status": "OK",
		"event_id": 7504
	}`)
	ts := newTestResponse(r)
	defer ts.Close()
	servers := &Servers{provider}
	ok, err := servers.Destroy("100823")
	if !ok {
		t.Error("Servers Destroy failed.")
	}
	if err != nil {
		t.Error("Servers Destroy failed with " + err.Error() + ".")
	}
}

func newTestResponse(response []byte) *httptest.Server {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", string(response))
	}))
	provider.Endpoint = testServer.URL
	return testServer
}
