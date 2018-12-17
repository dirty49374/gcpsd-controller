package verdb

import (
	"testing"
)

func TestCreate(t *testing.T) {
	_, err := NewDeploySpec("**")
	if err == nil {
		t.Errorf("** should not be created")
	}

	ds, err := NewDeploySpec("*")
	if err != nil {
		t.Errorf("* should be created")
	}

	if len(ds.Pods) != 0 {
		t.Errorf("length should be 0")
	}

	if ds.CurrentImage != "" {
		t.Errorf("CurrentImage should be empty")
	}
}
