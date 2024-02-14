package simulator_test

import (
	"testdoubles/positioner"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCatch(t *testing.T) {
	t.Run("hunter can catch the prey", func(t *testing.T) {
		position := positioner.NewPositionStub()

		position.FuncGetLinearDistance = func(from, to *positioner.Position) (linerarDistance float64) {
			linerarDistance = 100
			return
		}

		sim := simulator.NewCatchSimulatorDefault(10, position)

		hunter := simulator.Subject{
			Position: &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Speed: 30,
		}

		prey := simulator.Subject{
			Position: &positioner.Position{
				X: 25,
				Y: 0,
				Z: 0,
			},
			Speed: 5,
		}

		ok := sim.CanCatch(&hunter, &prey)

		expectedCanCatch := true
		require.EqualValues(t, expectedCanCatch, ok)
	})

	t.Run("hunter cannot catch the prey because hunter is too slow", func(t *testing.T) {
		position := positioner.NewPositionStub()

		position.FuncGetLinearDistance = func(from, to *positioner.Position) (linerarDistance float64) {
			linerarDistance = 100
			return
		}

		sim := simulator.NewCatchSimulatorDefault(10, position)

		hunter := simulator.Subject{
			Position: &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Speed: 2,
		}

		prey := simulator.Subject{
			Position: &positioner.Position{
				X: 25,
				Y: 0,
				Z: 0,
			},
			Speed: 5,
		}

		ok := sim.CanCatch(&hunter, &prey)

		expectedCanCatch := false
		require.EqualValues(t, expectedCanCatch, ok)
	})

	t.Run("hunter cannot catch the prey because hunter is not has enough time", func(t *testing.T) {
		position := positioner.NewPositionStub()

		position.FuncGetLinearDistance = func(from, to *positioner.Position) (linerarDistance float64) {
			linerarDistance = 100
			return
		}

		sim := simulator.NewCatchSimulatorDefault(1000, position)

		hunter := simulator.Subject{
			Position: &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Speed: 30,
		}

		prey := simulator.Subject{
			Position: &positioner.Position{
				X: 25,
				Y: 0,
				Z: 0,
			},
			Speed: 5,
		}

		ok := sim.CanCatch(&hunter, &prey)

		expectedCanCatch := false
		require.EqualValues(t, expectedCanCatch, ok)
	})
}
