package action

import (
	"io/ioutil"
	"testing"

	"github.com/andrewlunde/glide/cfg"
	"github.com/andrewlunde/glide/msg"
)

func TestAddPkgsToConfig(t *testing.T) {
	// Route output to discard so it's not displayed with the test output.
	o := msg.Default.Stderr
	msg.Default.Stderr = ioutil.Discard

	conf := new(cfg.Config)
	dep := new(cfg.Dependency)
	dep.Name = "github.com/andrewlunde/cookoo"
	dep.Subpackages = append(dep.Subpackages, "convert")
	conf.Imports = append(conf.Imports, dep)

	names := []string{
		"github.com/andrewlunde/cookoo/fmt",
		"github.com/andrewlunde/semver",
	}

	addPkgsToConfig(conf, names, false, true, false)

	if !conf.HasDependency("github.com/andrewlunde/semver") {
		t.Error("addPkgsToConfig failed to add github.com/andrewlunde/semver")
	}

	d := conf.Imports.Get("github.com/andrewlunde/cookoo")
	found := false
	for _, s := range d.Subpackages {
		if s == "fmt" {
			found = true
		}
	}
	if !found {
		t.Error("addPkgsToConfig failed to add subpackage to existing import")
	}

	// Restore messaging to original location
	msg.Default.Stderr = o
}
