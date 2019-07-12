package test

import (
	"testing"

	"../handle"
)

func TestRegistryConflict(t *testing.T) {
	t.Log("test registry conflict")
	res, err := handle.RegistryConflict("cjxvdldng000l09380nyl6qw6", "")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
