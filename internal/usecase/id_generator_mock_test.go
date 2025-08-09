package usecase_test

type MockIDGenerator struct {
	ID string
}

func (m *MockIDGenerator) NewID() string {
	return m.ID
}
