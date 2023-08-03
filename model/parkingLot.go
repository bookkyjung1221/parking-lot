package model

import "errors"

type ParkingLot struct {
	capacity int
	slots    []*ParkingSlot
}

func NewParkingLot(capacity int) *ParkingLot {
	parkingLot := new(ParkingLot)
	parkingLot.capacity = capacity
	parkingLot.slots = make([]*ParkingSlot, capacity)
	return parkingLot
}

func (p *ParkingLot) GetCapacity() int {
	return p.capacity
}

func (p *ParkingLot) getNearestParkingSlot() *ParkingSlot {
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] == nil || p.slots[i].IsAvailable() {
			if p.slots[i] == nil {
				p.slots[i] = NewParkingSlot(i + 1)
			}
			return p.slots[i]
		}
	}
	return nil
}

func (p *ParkingLot) ParkCar(car *Car) (*ParkingSlot, error) {
	pAvailableSlot := p.getNearestParkingSlot()
	if pAvailableSlot == nil {
		err := errors.New("no empty parking slot available")
		return nil, err
	}
	err := pAvailableSlot.AllotCar(car)
	if err != nil {
		return nil, err
	}
	return pAvailableSlot, nil
}

func (p *ParkingLot) UnparkCar(slotNo int) error {
	if slotNo > p.capacity {
		err := errors.New("wrong slot no. provided")
		return err
	} else if p.slots[slotNo-1] == nil || p.slots[slotNo-1].IsAvailable() {
		err := errors.New("slot already empty")
		return err
	} else {
		p.slots[slotNo-1].FreeParkingSlot()
		return nil
	}
}

func (p *ParkingLot) GetFilledSlots() []*ParkingSlot {
	list := make([]*ParkingSlot, 0)
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] != nil && !p.slots[i].IsAvailable() {
			list = append(list, p.slots[i])
		}
	}
	return list
}

func (p *ParkingLot) GetSlotsByCarColor(color string) []*ParkingSlot {
	result := make([]*ParkingSlot, 0, p.capacity)
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] != nil && !p.slots[i].IsAvailable() {
			if p.slots[i].GetCar().GetColor() == color {
				result = append(result, p.slots[i])
			}
		}
	}
	return result
}

func (p *ParkingLot) GetSlotsByCarRegNum(regNum string) (*ParkingSlot, error) {
	var result *ParkingSlot
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] != nil && !p.slots[i].IsAvailable() {
			if p.slots[i].GetCar().GetRegNum() == regNum {
				result := p.slots[i]
				return result, nil
			}
		}
	}
	if result == nil {
		return result, errors.New("not found")
	}
	return result, nil
}
