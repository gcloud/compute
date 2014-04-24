// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package vbox

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	p "github.com/gcloud/compute/providers"
)

func init() {
	p.RegisterImages("vbox", &Images{})
}

type Image struct {
	id   string
	name string
	file string
}

func (i *Image) Id() string {
	return i.id
}
func (i *Image) Name() string {
	return i.name
}
func (i *Image) File() string {
	return i.file
}

func (i *Image) String() string {
	b, err := json.Marshal(p.Map{
		"id": i.id, "name": i.name,
	})
	if err != nil {
		return ""
	}
	return string(b)
}

type Images struct {
	Path string
}

func NewImages() *Images {
	return &Images{}
}

// List images available to the account.
func (i *Images) List() ([]p.Image, error) {
	results, err := filepath.Glob(fmt.Sprintf("%s/*.ovf", i.path()))
	if err != nil {
		return nil, err
	}
	responses := make([]p.Image, 0)
	for _, r := range results {
		id, err := i.id(r)
		if err != nil {
			return nil, err
		}
		responses = append(responses, &Image{
			id:   id,
			name: strings.Replace(path.Base(r), path.Ext(r), "", 1),
			file: r,
		})
	}
	return responses, nil
}

// Show image information for a given id.
func (i *Images) Show(name string) (p.Image, error) {
	file := fmt.Sprintf("%s/%s.ovf", i.path(), name)
	id, err := i.id(file)
	if err != nil {
		return nil, err
	}
	return &Image{id, name, file}, nil
}

// Create an image.
func (i *Images) Create(n interface{}) (p.Image, error) {
	image := n.(p.Image)
	file := fmt.Sprintf("%s/%s.ovf", i.path(), image.Name())
	c := exec.Command("VBoxManage", "export", image.Name(), "--output", file)
	output, err := c.CombinedOutput()
	if err != nil {
		return nil, err
	}
	if matched, err := regexp.MatchString("Success", string(output)); !matched {
		return nil, err
	}
	id, err := i.id(file)
	if err != nil {
		return nil, err
	}
	return &Image{id, image.Name(), file}, nil
}

// Destroy an image.
func (i *Images) Destroy(name string) (bool, error) {
	file := fmt.Sprintf("%s/%s.ovf", i.path(), name)
	err := os.Remove(file)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (i *Images) path(args ...string) string {
	if len(args) > 0 && args[0] != "" {
		i.Path = args[0]
		return i.Path
	}
	if i.Path == "" {
		var wd, _ = os.Getwd()
		i.Path = wd
	}
	return i.Path
}

func (i *Images) id(file string) (string, error) {
	hasher := sha1.New()
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := io.Copy(hasher, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
