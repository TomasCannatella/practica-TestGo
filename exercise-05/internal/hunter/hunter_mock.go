package hunter

import (
	"testdoubles/internal/positioner"
	"testdoubles/internal/prey"

	"github.com/stretchr/testify/mock"
)

// NewHunter return a mock implementation of Hunter
func NewHunterMock() *HunterMock {
	return &HunterMock{}
}

// Hunter is a mock implementation of Hunter
type HunterMock struct {
	mock.Mock
}

func (ht *HunterMock) Hunt(pr prey.Prey) (duration float64, err error) {
	// observers
	args := ht.Called(pr)

	duration, err = args.Get(0).(float64), args.Error(1)
	return
}

func (ht *HunterMock) Configure(speed float64, position *positioner.Position) {
	// observers
	ht.Called(speed, position)
}
