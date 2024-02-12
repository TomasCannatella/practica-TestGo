package positioner

// PositionStub is a struct that represents a position stub
type PositionStub struct {
	FuncGetLinearDistance func(from, to *Position) (linearDistance float64)
}

// NewPositionStub returns a new position stub
func NewPositionStub() *PositionStub {
	return &PositionStub{}
}

// GetLinearDistance returns the hardcoded linear distance
func (s *PositionStub) GetLinearDistance(from, to *Position) (linearDistance float64) {
	return s.FuncGetLinearDistance(from, to)
}
