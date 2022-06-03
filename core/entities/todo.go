package entities

// my domain entitie, for example User
// in that case, the file name must be user.go
type ToDo struct {
	ID string `json:"id,omitempty"`
	To string `json:"to,omitempty" form:"to"`
	Do string `json:"do,omitempty" form:"do"`
}
