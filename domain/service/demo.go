package service

// DemoService is interface
type DemoService interface {
	GetMessage() (string, error)
}

// Demo is object
type Demo struct{}

// NewDemo is Create new instanse
func NewDemo() DemoService {
	return Demo{}
}

// GetMessage is GET test message
func (d Demo) GetMessage() (string, error) {
	return "Hello World", nil
}
