package prey_test

import (
	"testdoubles/positioner"
	"testdoubles/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStubGetSpeed(t *testing.T) {
	t.Run("get speed", func(t *testing.T) {
		// arrange
		preyStub := prey.NewPreyStub()

		preyStub.FuncGetSpeed = func() (speed float64) {
			speed = 252.0
			return
		}

		// act
		speed := preyStub.GetSpeed()

		// assert
		expectedSpeed := 252.0
		require.Equal(t, expectedSpeed, speed)
	})
}

func TestStubGetPosition(t *testing.T) {
	t.Run("get position", func(t *testing.T) {
		// arrange
		preyStub := prey.NewPreyStub()

		preyStub.FuncGetPosition = func() (position *positioner.Position) {
			position = &positioner.Position{}
			return
		}

		// act
		position := preyStub.GetPosition()

		// assert
		expectedPosition := &positioner.Position{}
		require.Equal(t, expectedPosition, position)
	})
}

func TestStubCreateTuna(t *testing.T) {

	t.Run("create tuna", func(t *testing.T) {
		// arrange
		preyStub := prey.NewPreyStub()

		preyStub.FuncCreateTuna = func() prey.Prey {
			tuna := prey.NewTuna(0, &positioner.Position{})
			return tuna
		}

		// act
		tuna := preyStub.CreateTuna()

		// assert
		speed := tuna.GetSpeed()
		position := tuna.GetPosition()

		expectedSpeed := 0.0
		expectedPosition := &positioner.Position{
			X: 0,
			Y: 0,
			Z: 0,
		}

		require.Equal(t, expectedSpeed, speed)
		require.Equal(t, expectedPosition, position)
	})
}
