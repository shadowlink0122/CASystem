package service

// DemoService is interface
type DemoService interface {
	GetMessage() (string, error)
}

// Demo is object
type Demo struct{}

// New is Create new instanse
func New() DemoService {
	return &Demo{}
}

// GetMessage is GET test message
func (d *Demo) GetMessage() (string, error) {
	return "Demo API", nil
}
