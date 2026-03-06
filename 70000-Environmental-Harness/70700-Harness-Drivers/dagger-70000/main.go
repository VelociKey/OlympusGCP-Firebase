package main

import (
	"context"
	"olympus.fleet/00SDLC/OlympusForge/70000-Environmental-Harness/dagger/olympusgcp-firebase/internal/dagger"
)

type OlympusGCPFirebase struct{}

func (m *OlympusGCPFirebase) HelloWorld(ctx context.Context) string {
	return "Hello from OlympusGCP-Firebase!"
}

func main() {
	dagger.Serve()
}
