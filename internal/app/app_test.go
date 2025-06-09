package app

import (
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/benidevo/website/internal/config"
)

func TestNew(t *testing.T) {
	cfg := &config.Config{
		Settings: &config.Settings{
			Port: "8080",
		},
	}

	app := New(cfg)

	assert.NotNil(t, app)
	assert.Equal(t, cfg, app.cfg)
	assert.NotNil(t, app.done)
}

func TestApp_WaitForShutdown(t *testing.T) {
	cfg := &config.Config{
		Settings: &config.Settings{
			Port: "0",
		},
	}

	app := New(cfg)
	app.server = &http.Server{
		Addr: ":0",
	}

	// Simulate shutdown signal
	go func() {
		time.Sleep(100 * time.Millisecond)
		app.done <- syscall.SIGTERM
	}()

	// This should complete without hanging
	done := make(chan bool)
	go func() {
		app.WaitForShutdown()
		done <- true
	}()

	select {
	case <-done:
		// Success
	case <-time.After(2 * time.Second):
		t.Fatal("WaitForShutdown did not complete in time")
	}
}

func TestApp_SignalChannel(t *testing.T) {
	cfg := &config.Config{
		Settings: &config.Settings{
			Port: "8080",
		},
	}

	app := New(cfg)

	select {
	case app.done <- os.Interrupt:
	default:
		t.Fatal("Signal channel should be buffered")
	}
}
