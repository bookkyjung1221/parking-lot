package model

type Car struct {
	regNum string
	color  string
}

func NewCar(regNum string, color string) *Car {
	c := new(Car)
	c.regNum = regNum
	c.color = color
	return c
}

func (c *Car) GetRegNum() string {
	return c.regNum
}

func (c *Car) GetColor() string {
	return c.color
}
