package domain

type Payer struct {
	ID    string
	Name  string
	Email string
}

func NewPayer(id, name, email string) *Payer {
	return &Payer{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
