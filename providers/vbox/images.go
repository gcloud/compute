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

// The images available from the compute service.
type Images struct {
	Path string
}

// List images available to the account.
func (i *Images) List() ([]byte, error) {
	results, err := filepath.Glob(fmt.Sprintf("%s/*.ovf", i.path()))
	if err != nil {
		return nil, err
	}
	responses := make([]interface{}, 0)
	for _, r := range results {
		id, err := i.id(r)
		if err != nil {
			return nil, err
		}
		responses = append(responses, map[string]interface{}{
			"Id":     id,
			"Name":   strings.Replace(path.Base(r), path.Ext(r), "", 1),
			"Source": r,
		})
	}
	b, err := json.Marshal(responses)
	return b, err
}

// Show image information for a given id.
func (i *Images) Show(name string) ([]byte, error) {
	file := fmt.Sprintf("%s/%s.ovf", i.path(), name)
	id, err := i.id(file)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(map[string]interface{}{
		"Id":     id,
		"Name":   name,
		"Source": file,
	})
	return b, err
}

// Create an image.
func (i *Images) Create(n *p.Image) ([]byte, error) {
	file := fmt.Sprintf("%s/%s.ovf", i.path(), n.Name)
	c := exec.Command("VBoxManage", "export", n.Name, "--output", file)
	output, err := c.CombinedOutput()
	if err != nil {
		return output, err
	}
	if matched, err := regexp.MatchString("Success", string(output)); !matched {
		return output, err
	}
	id, _ := i.id(file)
	n.Id = id
	n.Source = file
	b, err := json.Marshal(n)
	return b, err
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
