package simulator

// CatchSimulatorMock is a mock of CatchSimulator
type CatchSimulatorMock struct {
	// FuncCanCatch is the function that will be called when CanCatch is called
	FuncCanCatch func(hunter, prey *Subject) (ok bool)

	Calls struct {
		CanCatch      int
		CurrentHunter *Subject
		CurrentPrey   *Subject
	}
}

// NewCatchSimulatorMock returns a new CatchSimulatorMock
func NewCatchSimulatorMock() *CatchSimulatorMock {
	return &CatchSimulatorMock{}
}

// CanCatch returns the hardcoded value
func (m *CatchSimulatorMock) CanCatch(hunter, prey *Subject) (ok bool) {
	// increment the number of times CanCatch was called
	m.Calls.CanCatch++

	// save the current hunter and prey
	m.Calls.CurrentHunter = hunter
	m.Calls.CurrentPrey = prey

	return m.FuncCanCatch(hunter, prey)
}
