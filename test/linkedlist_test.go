package test

import (
	"testing"

	"github.com/neelkshah/pandora/pkg/linkedlist"
)

func TestFiver(t *testing.T) {
	var response = linkedlist.Fiver()

	if response != 5 {
		t.Errorf("Expected value 5, got %v", response)
	}
}
