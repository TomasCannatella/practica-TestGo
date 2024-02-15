package prey

import "testdoubles/positioner"

// PreyStub is a struct that represents a prey stub
type PreyStub struct {
	FuncGetSpeed    func() (speed float64)
	FuncGetPosition func() (position *positioner.Position)
	FuncCreateTuna  func() Prey
}

// NewPreyStub returns a new prey stub
func NewPreyStub() *PreyStub {
	return &PreyStub{}
}

// GetSpeed returns the hardcoded speed
func (s *PreyStub) GetSpeed() (speed float64) {
	return s.FuncGetSpeed()
}

// GetPosition returns the hardcoded position
func (s *PreyStub) GetPosition() (position *positioner.Position) {
	return s.FuncGetPosition()
}

// CreateTuna returns a new tuna
func (s *PreyStub) CreateTuna() Prey {
	return s.FuncCreateTuna()
}
