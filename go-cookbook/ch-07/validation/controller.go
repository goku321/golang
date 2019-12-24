package validation

// Controller ...
type Controller struct {
	ValidatePayload func(p *Payload) error
}

// New ...
func New() *Controller {
	return &Controller{
		ValidatePayload: ValidatePayload,
	}
}