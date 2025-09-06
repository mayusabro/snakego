package engine_test

import (
	"github.com/mayusabro/snakego/engine"
	"testing"
)

func TestEntity_Update(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		game *engine.Game
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var e engine.Entity
			e.Update(tt.game)
		})
	}
}
