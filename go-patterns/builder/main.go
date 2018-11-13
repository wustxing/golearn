package main

import "github.com/0990/golearn/go-patterns/builder/car"

func main() {
	assembly := car.NewBuilder().Color(car.RedColor)

	familyCar := assembly.TopSpeed(50 * car.MPH).Build()
	familyCar.Drive()
	familyCar.Stop()
}
