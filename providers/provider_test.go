// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package providers

import (
	"testing"

	"github.com/gcloud/identity"
)

var ProviderName = "MockProvider"

type MockSize struct {
	id   string
	name string
}

func (m *MockSize) Id() string {
	return m.id
}
func (m *MockSize) Name() string {
	return m.name
}

type MockSizes struct{}

func (s *MockSizes) List() ([]Size, error) {
	return []Size{
		&MockSize{name: "The Size", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"},
	}, nil
}
func (s *MockSizes) Show(id string) (Size, error) {
	return &MockSize{name: "The Size", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"}, nil
}

func Test_GetProviderSetAccount(t *testing.T) {
	p := GetProvider(ProviderName)
	p.SetAccount(&identity.Account{Key: "test"})
	expected := "test"
	result := p.Account.Key
	if expected != result {
		t.Errorf("Expected %s, but is %s.", expected, result)
	}
}

func Test_RegisterSizes(t *testing.T) {
	RegisterSizes(ProviderName, &MockSizes{})
	p := GetProvider(ProviderName)
	r, err := p.Sizes.Show("1")
	if err != nil {
		t.Error("Error in Register Sizes.")
		return
	}
	expected := "The Size"
	result := r.Name()
	if expected != result {
		t.Errorf("Expected %s, but is %s.", expected, result)
	}
}
