package prey_test

import (
	"testdoubles/positioner"
	"testdoubles/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSpeed(t *testing.T) {
	t.Run("speed has default values", func(t *testing.T) {
		// arrange
		tuna := prey.NewTuna(252.0, &positioner.Position{})

		// act
		speed := tuna.GetSpeed()

		// assert
		expectedSpeed := 252.0
		require.Equal(t, expectedSpeed, speed)
	})
}

func TestGetPosition(t *testing.T) {
	t.Run("position has default values", func(t *testing.T) {
		// arrange
		tuna := prey.NewTuna(0, &positioner.Position{})

		// act
		position := tuna.GetPosition()

		// assert

		expectedPosition := &positioner.Position{}
		require.Equal(t, expectedPosition, position)
	})
}
