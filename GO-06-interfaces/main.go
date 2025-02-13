package main

import "fmt"

type auto struct {
	distance, time float64
}

type motorbike struct {
	distance, time float64
}

type bike struct {
	distance, time float64
}

type vehicles interface {
	getSpeed() float64
}

func retrieveVelocity(v vehicles) float64 {
	return v.getSpeed()
}

func (a *auto) getSpeed() float64 {
	return a.distance / a.time
}

func (m *motorbike) getSpeed() float64 {
	return m.distance / m.time
}

func (b *bike) getSpeed() float64 {
	return b.distance / b.time
}

func main() {
	car := auto{distance: 50, time: 1}
	moto := motorbike{distance: 77, time: 1}
	bike := bike{distance: 10, time: 1}

	listOfVehicles := make(map[string]vehicles)
	listOfVehicles["car"] = &car
	listOfVehicles["moto"] = &moto
	listOfVehicles["bike"] = &bike

	for k, v := range listOfVehicles {
		fmt.Printf("Velocity of the %v  is: %v km/h\n", k, retrieveVelocity(v)*3.6)
	}
}
