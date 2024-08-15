package gobook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello1(t *testing.T) {
	result := SayHello("irsyad")
	if result != "Hello irsyad" {
		t.Fatal("Result must be Hello irsyad")
	}
}

func TestSayHello2(t *testing.T) {
	result := SayHello("irsyad")
	assert.Equal(t, "Hello irsyad", result, "Result must be Hello irsyad")
}
