// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	"testing"
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
func (s *MockSizes) List() ([]Size, error) {
	return []Size{
		&MockSize{name: "The Size", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"},
	}, nil
}
func (s *MockSizes) Show(id string) (Size, error) {
	return &MockSize{name: "The Size", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"}, nil
}

func init() {
	RegisterSizes("mock", &MockSizes{})
}

func Test_SizesList(t *testing.T) {
	sizes := GetSizes("mock", nil)
	results, err := sizes.List()
	if err != nil {
		t.Errorf("Sizes List failed with %s", err)
	}
	if results == nil {
		t.Error("Sizes List results should not be nil.")
		return
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
	sizes := GetSizes("mock", nil)
	result, err := sizes.Show("616fb98f-46ca-475e-917e-2563e5a8cd19")
	if err != nil {
		t.Errorf("Sizes Show failed with %s", err)
	}
	if result == nil {
		t.Error("Sizes Show results should not be nil.")
		return
	}
	r := result
	if r.Id() != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		t.Error("Wrong value for Id.")
	}
	if r.Name() != "The Size" {
		t.Error("Wrong value for Name.")
	}
}
