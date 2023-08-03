package model

import "errors"

type ParkingSlot struct {
	occupied bool
	slotNo   int
	car      *Car
}

func NewParkingSlot(slotNo int) *ParkingSlot {
	p := new(ParkingSlot)
	p.occupied = false
	p.slotNo = slotNo
	p.car = nil
	return p
}

func (p *ParkingSlot) GetCar() *Car {
	return p.car
}

func (p *ParkingSlot) GetSlotNo() int {
	return p.slotNo
}

func (p *ParkingSlot) IsAvailable() bool {
	return !p.occupied
}

func (p *ParkingSlot) FreeParkingSlot() {
	p.occupied = false
	p.car = nil
}

func (p *ParkingSlot) AllotCar(c *Car) error {
	if p.IsAvailable() {
		p.occupied = true
		p.car = c
		return nil
	}
	return errors.New("slot is not empty")
}
