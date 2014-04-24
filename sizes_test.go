// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	"testing"

	p "github.com/gcloud/compute/providers"
)

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

// List sizes available on the account.
func (s *MockSizes) List() ([]p.Size, error) {
	return []p.Size{
		&MockSize{name: "The Size", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"},
	}, nil
}
func (s *MockSizes) Show(id string) (p.Size, error) {
	return &MockSize{name: "The Size", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"}, nil
}

func init() {
	p.RegisterSizes("mock", &MockSizes{})
}

func Test_SizesList(t *testing.T) {
	sizes := &Sizes{Provider: "mock"}
	results, err := sizes.List()
	if err != nil {
		t.Error("Sizes List failed with " + err.Error() + "(bool, error).")
	}
	if results == nil {
		t.Error("Results should not be nil.")
	}
	for _, v := range results {
		if v.Id() != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
			t.Error("Wrong value for Id.")
		}
		if v.Name() != "The Size" {
			t.Error("Wrong value for Name.")
		}
	}
}

func Test_SizesShow(t *testing.T) {
	sizes := &Sizes{Provider: "mock"}
	result, err := sizes.Show("616fb98f-46ca-475e-917e-2563e5a8cd19")
	if err != nil {
		t.Error("Sizes Show failed with " + err.Error() + ".")
	}
	if result == nil {
		t.Error("Results should not be nil.")
	}
	r := result
	if r.Id() != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		t.Error("Wrong value for Id.")
	}
	if r.Name() != "The Size" {
		t.Error("Wrong value for Name.")
	}
}
