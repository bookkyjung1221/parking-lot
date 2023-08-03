package model

import (
	"errors"
	"testing"
)

func TestCreateParkingSlot(t *testing.T) {
	testParkingSlot := NewParkingSlot(1)
	if testParkingSlot == nil {
		t.Errorf("Failed to create parking slot")
	} else if testParkingSlot.GetSlotNo() != 1 {
		t.Errorf("Failed to create parking lot with slotNo %v, got %v", 1, testParkingSlot.GetSlotNo())
	}
}

func TestAllotCar(t *testing.T) {
	type parkslot struct {
		occupied bool
		slotNo   int
		car      *Car
	}
	type testCar struct {
		car Car
	}
	tests := []struct {
		name     string
		parkslot parkslot
		args     testCar
		want     error
		wantErr  bool
	}{
		{
			"TestCase 1: Parking slot is empty where Car is to park",
			parkslot{
				occupied: false,
				slotNo:   1,
				car:      nil,
			},
			testCar{car: Car{regNum: "KA01-1498", color: "Red"}},
			nil,
			false,
		},
		{
			"TestCase 2: Parking slot is full where Car is to park",
			parkslot{
				occupied: true,
				slotNo:   1,
				car:      &Car{regNum: "KA01-1491", color: "Red"},
			},
			testCar{car: Car{regNum: "KA01-1498", color: "Red"}},
			errors.New("Slot is not empty"),
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testParkingSlot := &ParkingSlot{
				occupied: test.parkslot.occupied,
				slotNo:   test.parkslot.slotNo,
				car:      test.parkslot.car,
			}
			err := testParkingSlot.AllotCar(&test.args.car)
			if (err != nil) != test.wantErr {
				t.Errorf("\x1b[31;1mParkingSlot.AllotCar() error = %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}
