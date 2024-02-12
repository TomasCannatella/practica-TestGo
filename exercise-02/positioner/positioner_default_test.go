package positioner_test

import (
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPositionerGetLinearDistance(t *testing.T) {
	t.Run("success - the coordinates are negative", func(t *testing.T) {
		// arrange
		from := &positioner.Position{
			X: -1,
			Y: -1,
			Z: -1,
		}

		to := &positioner.Position{
			X: -1,
			Y: -1,
			Z: -1,
		}
		positioner := positioner.NewPositionerDefault()

		// act
		linearDistance := positioner.GetLinearDistance(from, to)

		// assert
		expectedLinearDistance := 0.0
		require.Equal(t, expectedLinearDistance, linearDistance)
	})

	t.Run("the coordinates are positive", func(t *testing.T) {
		// arrange
		from := &positioner.Position{
			X: 1,
			Y: 1,
			Z: 1,
		}

		to := &positioner.Position{
			X: 2,
			Y: 2,
			Z: 2,
		}
		positioner := positioner.NewPositionerDefault()

		// act
		linearDistance := positioner.GetLinearDistance(from, to)

		// assert
		expectedLinearDistance := 1.7320508075688772
		require.Equal(t, expectedLinearDistance, linearDistance)
	})

	t.Run("the coordinates return a linear distance without decimals", func(t *testing.T) {
		// arrange
		from := &positioner.Position{
			X: 7,
			Y: 0,
			Z: 0,
		}

		to := &positioner.Position{
			X: 0,
			Y: 0,
			Z: 0,
		}

		positioner := positioner.NewPositionerDefault()

		// act
		linearDistance := positioner.GetLinearDistance(from, to)

		// assert
		expectedLinearDistance := 7.0
		require.Equal(t, expectedLinearDistance, linearDistance)
	})
}
