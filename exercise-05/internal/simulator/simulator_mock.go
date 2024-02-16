package simulator

import "github.com/stretchr/testify/mock"

// NewCatchSimulatorMock creates a new CatchSimulatorMock
func NewCatchSimulatorMock() (simulator *CatchSimulatorMock) {
	simulator = &CatchSimulatorMock{}
	return
}

// CatchSimulatorMock is a mock for CatchSimulator
type CatchSimulatorMock struct {
	// mock.Mock is the mock
	mock.Mock
}

// CanCatch
func (m *CatchSimulatorMock) CanCatch(hunter, prey *Subject) (duration float64, ok bool) {
	// args is the arguments for the mock
	args := m.Called(hunter, prey)

	// return the values from the mock
	duration = args.Get(0).(float64)
	ok = args.Bool(1)
	return
}
