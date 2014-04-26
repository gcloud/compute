// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package digitalocean

import (
	"testing"

	"github.com/gcloud/compute"
	"github.com/gcloud/identity"
)

var TestServer compute.Server

var account = &identity.Account{
	Id:  "id",
	Key: "key",
}

func init() {
	provider = compute.GetProvider("digitalocean")
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
	result, err := servers.Create(servers.New(compute.Map{
		"name":        "GCloudServer",
		"image_id":    3101045,
		"size_id":     66,
		"region_id":   1,
		"ssh_key_ids": 18420,
	}))
	TestServer = result
	if err != nil {
		t.Errorf("Servers Create failed with %s", err)
	}
	if result == nil {
		t.Error("Servers Create result should not be nil.")
		return
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
		t.Errorf("Servers List failed with %s", err)
	}
	if results == nil {
		t.Error("Servers List results should not be nil.")
		return
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
	result, err := servers.Show(TestServer)
	if err != nil {
		t.Errorf("Servers Show failed with %s.", err)
	}
	if result == nil {
		t.Error("Servers Show result should not be nil.")
		return
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
	ok, err := servers.Start(TestServer)
	if !ok {
		t.Error("Servers Start failed.")
	}
	if err != nil {
		t.Errorf("Servers Start failed with %s.", err)
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
	ok, err := servers.Stop(TestServer)
	if !ok {
		t.Error("Servers Stop failed.")
	}
	if err != nil {
		t.Errorf("Servers Stop failed with %s.", err)
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
	ok, err := servers.Reboot(TestServer)
	if !ok {
		t.Error("Servers Reboot failed.")
	}
	if err != nil {
		t.Errorf("Servers Reboot failed with %s.", err)
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
	ok, err := servers.Destroy(TestServer)
	if !ok {
		t.Error("Servers Destroy failed.")
	}
	if err != nil {
		t.Errorf("Servers Destroy failed with %s.", err)
	}
}
