package audio

import "log"

// This audio package is a small stub for procedural audio.
// For portability and simplicity we provide lightweight stubs.
// You can replace these with proper Ebiten audio code later.

// Init initializes audio subsystem (no-op for now).
func Init() error {
	// Placeholder: initialize audio context with ebiten/audio if desired.
	// Returning nil so the game continues without audio.
	return nil
}

// PlayShot plays a short shooting sound (stub).
func PlayShot() {
	go func() {
		// TODO: implement a procedural beep using ebiten/audio.
		// This stub intentionally does nothing so builds remain simple.
		_ = 0
	}()
}

// PlayPop plays a small pop sound (stub).
func PlayPop() {
	go func() {
		// TODO: implement a pop sound here.
		_ = 0
	}()
}

func warn(msg string) { log.Printf("audio: %s", msg) }
