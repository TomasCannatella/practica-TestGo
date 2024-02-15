package positioner_test

import (
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStubGetLinearDistance(t *testing.T) {
	t.Run("success - Get linear distance stub", func(t *testing.T) {
		// arrange
		positionerStub := positioner.NewPositionStub()

		positionerStub.FuncGetLinearDistance = func(from, to *positioner.Position) (linearDistance float64) {
			linearDistance = 0.0
			return
		}

		// act
		linearDistance := positionerStub.GetLinearDistance(&positioner.Position{}, &positioner.Position{})

		// assert
		expectedLinearDistance := 0.0
		require.Equal(t, expectedLinearDistance, linearDistance)
	})
}
