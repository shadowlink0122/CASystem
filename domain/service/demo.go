package service

// demoService is interface
type demoService interface {
	GetMessage() (string, error)
}

// Demo is object
type Demo struct{}

// NewDemo is Create new instanse
func NewDemo() demoService {
	return &Demo{}
}

// GetMessage is GET test message
func (d *Demo) GetMessage() (string, error) {
	return "Demo API", nil
}
