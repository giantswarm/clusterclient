package config

type Customer struct {
	Name string
}

func DefaultCustomer() *Customer {
	return &Customer{
		Name: "",
	}
}
