package digitalocean

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gcloud/compute"
)

var provider *compute.Provider

// Get event results
func event(provider *compute.Provider, url string, id string) (bool, error) {
	r := false
	response, err := request(provider, "GET", fmt.Sprintf(url, id), nil)
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
func request(provider *compute.Provider, method string, path string, data interface{}) ([]byte, error) {
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
	case compute.Map:
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

func newTestResponse(response []byte) *httptest.Server {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", string(response))
	}))
	provider.Endpoint = testServer.URL
	return testServer
}
