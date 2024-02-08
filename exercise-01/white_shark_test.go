package hunt_test

import (
	hunt "testdoubles"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		//assert
		shark := hunt.NewWhiteShark(true, false, 10.0)
		tuna := hunt.NewTuna("tuna", 5.0)

		//act
		err := shark.Hunt(tuna)

		//arrange
		expectedHungry := false
		expectedTired := true
		require.NoError(t, err)
		require.Equal(t, expectedHungry, shark.Hungry)
		require.Equal(t, expectedTired, shark.Tired)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		//assert
		shark := hunt.NewWhiteShark(false, false, 10.0)
		tuna := hunt.NewTuna("tuna", 5.0)

		//act
		err := shark.Hunt(tuna)

		// arrange
		expectedError := hunt.ErrSharkIsNotHungry
		require.Error(t, err)
		require.ErrorIs(t, err, expectedError)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		// assert
		shark := hunt.NewWhiteShark(true, true, 10.0)
		tuna := hunt.NewTuna("tuna", 5.0)

		// act
		err := shark.Hunt(tuna)

		// arrange
		expectedError := hunt.ErrSharkIsTired
		require.Error(t, err)
		require.ErrorIs(t, err, expectedError)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		// assert
		shark := hunt.NewWhiteShark(true, false, 5.0)
		tuna := hunt.NewTuna("tuna", 10.0)

		// act
		err := shark.Hunt(tuna)

		// arrange
		expectedError := hunt.ErrSharkIsSlower
		require.Error(t, err)
		require.ErrorIs(t, err, expectedError)
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		// assert
		shark := hunt.NewWhiteShark(true, false, 10.0)
		var tuna *hunt.Tuna

		// act
		err := shark.Hunt(tuna)

		// arrange
		expectedError := hunt.ErrTunaIsNil
		require.Error(t, err)
		require.ErrorIs(t, err, expectedError)

	})
}

func TestWhiteSharkHuntWithTdd(t *testing.T) {
	type test struct {
		nameTest       string
		hungry         bool
		tired          bool
		speedShark     float64
		expectedHungry bool
		expectedTired  bool
		tuna           *hunt.Tuna
		expectedError  error
	}

	tests := []test{
		{"case 1: white shark hunts successfully", true, false, 10.0, false, true, hunt.NewTuna("tuna", 5.0), nil},
		{"case 2: white shark is not hungry", false, false, 10.0, false, false, hunt.NewTuna("tuna", 5.0), hunt.ErrSharkIsNotHungry},
		{"case 3: white shark is tired", true, true, 10.0, true, true, hunt.NewTuna("tuna", 5.0), hunt.ErrSharkIsTired},
		{"case 4: white shark is slower than the tuna", true, false, 5.0, true, false, hunt.NewTuna("tuna", 10.0), hunt.ErrSharkIsSlower},
		{"case 5: tuna is nil", true, false, 10.0, true, false, nil, hunt.ErrTunaIsNil},
	}

	for _, test1 := range tests {
		t.Run(test1.nameTest, func(t *testing.T) {
			//assert
			shark := hunt.NewWhiteShark(test1.hungry, test1.tired, test1.speedShark)
			tuna := test1.tuna

			//act
			err := shark.Hunt(tuna)

			//arrange
			if test1.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test1.expectedError)
			} else {
				require.NoError(t, err)
				require.Equal(t, test1.expectedHungry, shark.Hungry)
				require.Equal(t, test1.expectedTired, shark.Tired)
			}
		})
	}
}
