package model

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewParkingLot(t *testing.T) {
	testParkingLot := NewParkingLot(6)
	if testParkingLot == nil {
		t.Errorf("Failed to create parking lot")
	} else if testParkingLot.GetCapacity() != 6 {
		t.Errorf("Failed to create parking lot with want capacity %v, got %v", 6, testParkingLot.GetCapacity())
	}
}

func TestParkCar(t *testing.T) {
	type parklot struct {
		capacity int
		slots    []*ParkingSlot
	}
	type testCar struct {
		car Car
	}
	tests := []struct {
		name    string
		parklot parklot
		args    testCar
		want    *ParkingSlot
		wantErr bool
	}{
		{
			"TestCase 1: Parking slots is not full",
			parklot{
				capacity: 3,
				slots: []*ParkingSlot{
					{
						occupied: false,
						slotNo:   1,
						car:      nil,
					},
					{
						occupied: false,
						slotNo:   2,
						car:      nil,
					},
					{
						occupied: false,
						slotNo:   3,
						car:      nil,
					},
				},
			},
			testCar{car: Car{regNum: "KA01-1498", color: "Red"}},
			&ParkingSlot{
				occupied: true,
				slotNo:   1,
				car:      &Car{regNum: "KA01-1498", color: "Red"},
			},
			false,
		},
		{
			"TestCase 2: Parking slots is full",
			parklot{
				capacity: 3,
				slots: []*ParkingSlot{
					{
						occupied: true,
						slotNo:   1,
						car:      &Car{regNum: "BR01-1234", color: "White"},
					},
					{
						occupied: true,
						slotNo:   2,
						car:      &Car{regNum: "DL12-5608", color: "Blue"},
					},
					{
						occupied: true,
						slotNo:   3,
						car:      &Car{regNum: "KA01-1523", color: "Black"},
					},
				},
			},
			testCar{car: Car{regNum: "KA01-1498", color: "Red"}},
			nil,
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testParkingLot := &ParkingLot{
				capacity: test.parklot.capacity,
				slots:    test.parklot.slots,
			}
			got, err := testParkingLot.ParkCar(&test.args.car)
			if (err != nil) != test.wantErr {
				t.Errorf("\x1b[31;1mParkingLot.ParkCar() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("\x1b[31;1mParking.ParkCar() = %v, want %v\x1b[0m", got, test.want)
			}
		})
	}
}

func TestUnparkCar(t *testing.T) {
	type parklot struct {
		capacity int
		slots    []*ParkingSlot
	}

	tests := []struct {
		name    string
		parklot parklot
		args    int
		err     error
		wantErr bool
	}{
		{
			"TestCase 1: Unpark Car from parking slot where car is parked",
			parklot{
				capacity: 3,
				slots: []*ParkingSlot{
					{
						occupied: true,
						slotNo:   1,
						car:      &Car{regNum: "BR01-1234", color: "White"},
					},
					{
						occupied: false,
						slotNo:   2,
						car:      nil,
					},
					{
						occupied: false,
						slotNo:   3,
						car:      nil,
					},
				},
			},
			1,
			nil,
			false,
		},
		{
			"TestCase 2: Unpark Car from parking slot where no car is parked",
			parklot{
				capacity: 3,
				slots: []*ParkingSlot{
					{
						occupied: true,
						slotNo:   1,
						car:      &Car{regNum: "BR01-1234", color: "White"},
					},
					{
						occupied: false,
						slotNo:   2,
						car:      nil,
					},
					{
						occupied: true,
						slotNo:   3,
						car:      &Car{regNum: "KA01-1523", color: "Black"},
					},
				},
			},
			2,
			errors.New("Slot already empty"),
			true,
		},
		{
			"TestCase 3: Unpark Car from parking slot where slot no exceeds capacity",
			parklot{
				capacity: 3,
				slots: []*ParkingSlot{
					{
						occupied: true,
						slotNo:   1,
						car:      &Car{regNum: "BR01-1234", color: "White"},
					},
					{
						occupied: false,
						slotNo:   2,
						car:      nil,
					},
					{
						occupied: true,
						slotNo:   3,
						car:      &Car{regNum: "KA01-1523", color: "Black"},
					},
				},
			},
			4,
			errors.New("Wrong slot no. provided"),
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testParkingLot := &ParkingLot{
				capacity: test.parklot.capacity,
				slots:    test.parklot.slots,
			}
			err := testParkingLot.UnparkCar(test.args)
			if (err != nil) != test.wantErr {
				t.Errorf("\x1b[31;1mParkingLot.UnparkCar() error = %v, wantErr %v", err, test.wantErr)
				return
			}

		})
	}
}

func TestGetSlotsByColor(t *testing.T) {
	type parklot struct {
		capacity int
		slots    []*ParkingSlot
	}

	tests := []struct {
		name      string
		parklot   parklot
		args      string
		want      *ParkingSlot
		wantSlots int
	}{
		{
			"TestCase 1: Parking slots is not empty",
			parklot{
				capacity: 3,
				slots: []*ParkingSlot{
					{
						occupied: false,
						slotNo:   1,
						car:      nil,
					},
					{
						occupied: false,
						slotNo:   2,
						car:      nil,
					},
					{
						occupied: true,
						slotNo:   3,
						car:      &Car{regNum: "KA01-1498", color: "Red"},
					},
				},
			},
			"Red",
			&ParkingSlot{
				occupied: true,
				slotNo:   3,
				car:      &Car{regNum: "KA01-1498", color: "Red"},
			},
			1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testParkingLot := &ParkingLot{
				capacity: test.parklot.capacity,
				slots:    test.parklot.slots,
			}
			occupiedSlots := testParkingLot.GetSlotsByCarColor(test.args)
			if len(occupiedSlots) != test.wantSlots || !reflect.DeepEqual(occupiedSlots[0], test.want) {
				t.Errorf("\x1b[31;1mParkingLot.GetSlotByCarColor() error = , wantSlots %v", test.wantSlots)
			}
		})
	}
}
func TestGetSlotByCarregNum(t *testing.T) {
	type parklot struct {
		capacity int
		slots    []*ParkingSlot
	}

	tests := []struct {
		name    string
		parklot parklot
		args    string
		want    *ParkingSlot
		wantErr bool
	}{
		{
			"TestCase 1: Car present in parking lot",
			parklot{
				capacity: 3,
				slots: []*ParkingSlot{
					{
						occupied: false,
						slotNo:   1,
						car:      nil,
					},
					{
						occupied: false,
						slotNo:   2,
						car:      nil,
					},
					{
						occupied: true,
						slotNo:   3,
						car:      &Car{regNum: "KA01-1498", color: "Red"},
					},
				},
			},
			"KA01-1498",
			&ParkingSlot{
				occupied: true,
				slotNo:   3,
				car:      &Car{regNum: "KA01-1498", color: "Red"},
			},
			false,
		},
		{
			"TestCase 2: Car not present in parking lot",
			parklot{
				capacity: 3,
				slots: []*ParkingSlot{
					{
						occupied: false,
						slotNo:   1,
						car:      nil,
					},
					{
						occupied: false,
						slotNo:   2,
						car:      nil,
					},
					{
						occupied: true,
						slotNo:   3,
						car:      &Car{regNum: "KA01-1498", color: "Red"},
					},
				},
			},
			"KA01-12298",
			nil,
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testParkingLot := &ParkingLot{
				capacity: test.parklot.capacity,
				slots:    test.parklot.slots,
			}
			occupiedSlots, err := testParkingLot.GetSlotsByCarRegNum(test.args)
			if (err != nil) != test.wantErr {
				t.Errorf("\x1b[31;1mParkingLot.GetSlotByCarregNum() error = %v, wantSlots %v", err, test.wantErr)
			}
			if occupiedSlots != nil && !reflect.DeepEqual(occupiedSlots, test.want) {
				t.Errorf("\x1b[31;1mParkingLot.GetSlotByCarColor() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestGetFilledSlots(t *testing.T) {
	type parklot struct {
		capacity int
		slots    []*ParkingSlot
	}

	tests := []struct {
		name      string
		parklot   parklot
		want      *ParkingSlot
		wantSlots int
	}{
		{
			"TestCase 1: Parking slots is not empty",
			parklot{
				capacity: 3,
				slots: []*ParkingSlot{
					{
						occupied: false,
						slotNo:   1,
						car:      nil,
					},
					{
						occupied: false,
						slotNo:   2,
						car:      nil,
					},
					{
						occupied: true,
						slotNo:   3,
						car:      &Car{regNum: "KA01-1498", color: "Red"},
					},
				},
			},
			&ParkingSlot{
				occupied: true,
				slotNo:   3,
				car:      &Car{regNum: "KA01-1498", color: "Red"},
			},
			1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testParkingLot := &ParkingLot{
				capacity: test.parklot.capacity,
				slots:    test.parklot.slots,
			}
			occupiedSlots := testParkingLot.GetFilledSlots()
			if len(occupiedSlots) != test.wantSlots || !reflect.DeepEqual(occupiedSlots[0], test.want) {
				t.Errorf("\x1b[31;1mParkingLot.GetSlotByCarColor() error = , wantSlots %v", test.wantSlots)
			}
		})
	}
}
