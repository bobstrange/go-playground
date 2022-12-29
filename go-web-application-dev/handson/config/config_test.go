package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	expectedPort := 8888

	t.Setenv("PORT", fmt.Sprint(expectedPort))

	got, err := New()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}
	if got.Port != expectedPort {
		t.Errorf("expect port %d, got %d", expectedPort, got.Port)
	}

	expectedEnv := "development"
	if got.Env != expectedEnv {
		t.Errorf("expect env %s, got %s", expectedEnv, got.Env)
	}
}
