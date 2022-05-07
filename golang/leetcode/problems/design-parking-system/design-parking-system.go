package design_parking_system

// https://leetcode.com/problems/design-parking-system/

type ParkingSystem struct {
	big, medium, small int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 3:
		if this.small > 0 {
			this.small--
			return true
		}
	case 2:
		if this.medium > 0 {
			this.medium--
			return true
		}
	case 1:
		if this.big > 0 {
			this.big--
			return true
		}
	}

	return false
}
