package car

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Builder interface {
	Color(Color) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

type CarBuilder struct {
	color    Color
	topSpeed Speed
}

func NewBuilder() Builder {
	return CarBuilder{}
}

type Car struct {
	carBuilder CarBuilder
}

func (p CarBuilder) Color(color Color) Builder {
	p.color = color
	return p
}

func (p CarBuilder) TopSpeed(speed Speed) Builder {
	p.topSpeed = speed
	return p
}

func (p CarBuilder) Build() Interface {
	return Car{
		carBuilder: p,
	}
}

func (p Car) Drive() error {
	fmt.Printf("car color:%v,topSpeed:%v drive\n", p.carBuilder.color, p.carBuilder.topSpeed)
	return nil
}

func (p Car) Stop() error {
	fmt.Printf("car color:%v,topSpeed:%v stop\n", p.carBuilder.color, p.carBuilder.topSpeed)
	return nil
}
