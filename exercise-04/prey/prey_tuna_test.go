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

func TestCreateTuna(t *testing.T) {
	t.Run("create tuna", func(t *testing.T) {
		// arrange
		tuna := prey.CreateTuna()

		// act
		speed := tuna.GetSpeed()
		//position := tuna.GetPosition()
		// assert
		expectedSpeed := speed >= 15.0

		require.True(t, expectedSpeed)
		//require.True(t, expectedPosition)
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
