package customer

type Customer struct {
	name     string
	surname  string
	lastname string
	age      int
}

func (c Customer) Getcustomer() string {
	return c.name
}

func (c *Customer) Setcustomer(name string) {
	c.name = name
}
