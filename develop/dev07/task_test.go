package main

import (
	"testing"
	"time"
)

func TestMergeDoneChannels(t *testing.T) {
	done1 := make(chan struct{})
	done2 := make(chan struct{})
	done3 := make(chan struct{})

	merged := mergeDoneChannels(done1, done2, done3)

	time.Sleep(1 * time.Second)
	close(done2)

	select {
	case _, ok := <-merged:
		if !ok {
			t.Log("Merged channel closed after one of the done channels was closed")
		} else {
			t.Error("Merged channel should be closed after one of the done channels is closed")
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for merged channel to close")
	}
}
