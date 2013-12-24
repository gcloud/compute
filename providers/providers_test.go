// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package providers

import (
	"encoding/json"
	"testing"

	"github.com/gcloud/identity"
)

var ProviderName = "MockProvider"

type MockSizes struct{}

func (s *MockSizes) List() ([]byte, error) {
	return []byte(`[{"Name":"The Size","Id": "616fb98f-46ca-475e-917e-2563e5a8cd19"}]`), nil
}
func (s *MockSizes) Show(id string) ([]byte, error) {
	return []byte(`{"Name":"The Size","Id": "616fb98f-46ca-475e-917e-2563e5a8cd19"}`), nil
}

func Test_GetProvider(t *testing.T) {
	p := GetProvider(ProviderName, &identity.Account{Key: "test"})
	expected := "test"
	result := p.Account.Key
	if expected != result {
		t.Errorf("Expected %s, but is %s.", expected, result)
	}
}

func Test_RegisterSizes(t *testing.T) {
	RegisterSizes(ProviderName, &MockSizes{})
	p := GetProvider(ProviderName, &identity.Account{})
	r, err := p.Sizes.Show("1")
	if err != nil {
		t.Error("Error in Register Sizes.")
		return
	}
	var size Size
	err = json.Unmarshal(r, &size)
	if err != nil {
		t.Error("Error in Register Sizes.")
		return
	}
	expected := "The Size"
	result := size.Name
	if expected != result {
		t.Errorf("Expected %s, but is %s.", expected, result)
	}
}
