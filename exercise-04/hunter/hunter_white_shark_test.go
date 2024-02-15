package hunter_test

import (
	"testdoubles/hunter"
	"testdoubles/positioner"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewWhiteShark(t *testing.T) {
	t.Run("NewWhiteShark should create a new WhiteShark", func(t *testing.T) {
		// arrange
		simulatorMock := simulator.NewCatchSimulatorMock()
		speed := 10.0
		position := &positioner.Position{
			X: 10.0,
			Y: 10.0,
			Z: 10.0,
		}

		// act
		whiteShark := hunter.NewWhiteShark(speed, position, simulatorMock)

		// assert
		require.NotNil(t, whiteShark)
	})
}

func TestHunter(t *testing.T) {
	t.Run("WhiteShark should catch the prey", func(t *testing.T) {
		// arrange
		simulatorMock := simulator.NewCatchSimulatorMock()
		simulatorMock.FuncCanCatch = func(hunter, prey *simulator.Subject) (ok bool) {
			return true
		}
		preyStub := prey.NewPreyStub()

		preyStub.FuncGetSpeed = func() float64 {
			return 10.0
		}

		preyStub.FuncGetPosition = func() *positioner.Position {
			return &positioner.Position{
				X: 10.0,
				Y: 10.0,
				Z: 10.0,
			}
		}

		whiteShark := hunter.CreateWhiteShark(simulatorMock)

		// act
		err := whiteShark.Hunt(preyStub)
		// assert
		require.NoError(t, err)
		require.Equal(t, simulatorMock.Calls.CanCatch, 1)
	})

	t.Run("WhiteShark should not catch the prey", func(t *testing.T) {
		// arrange
		simulatorMock := simulator.NewCatchSimulatorMock()
		simulatorMock.FuncCanCatch = func(hunter, prey *simulator.Subject) (ok bool) {
			return false
		}
		preyStub := prey.NewPreyStub()

		preyStub.FuncGetSpeed = func() float64 {
			return 10.0
		}

		preyStub.FuncGetPosition = func() *positioner.Position {
			return &positioner.Position{
				X: 10.0,
				Y: 10.0,
				Z: 10.0,
			}
		}

		whiteShark := hunter.CreateWhiteShark(simulatorMock)

		// act
		err := whiteShark.Hunt(preyStub)

		// assert
		expectedError := "can not hunt the prey: shark can not catch the prey"
		require.Error(t, err)
		require.EqualError(t, err, expectedError)
	})

}
