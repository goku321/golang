package controllers

// Controller passes state to our handlers
type Controller struct {
	storage Storage
}

// New is a controller constructor
func New(storage Storage) *Controller {
	return &Controller{
		storage: storage,
	}
}

type Payload struct {
	Value string `json:"value"`
}
