package design_parking_system

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	p := Constructor(1, 1, 0)

	assert.True(t, p.AddCar(1))
	assert.True(t, p.AddCar(2))
	assert.False(t, p.AddCar(3))
	assert.False(t, p.AddCar(1))
}

func TestCase2(t *testing.T) {
	p := Constructor(1, 1, 0)

	assert.True(t, p.AddCar(1))
	assert.False(t, p.AddCar(3))
	assert.True(t, p.AddCar(2))
	assert.False(t, p.AddCar(1))
}

func TestCase3(t *testing.T) {
	p := Constructor(1, 1, 0)

	assert.True(t, p.AddCar(2))
	assert.False(t, p.AddCar(3))
	assert.True(t, p.AddCar(1))
	assert.False(t, p.AddCar(1))
}

func TestCase4(t *testing.T) {
	p := Constructor(0, 0, 2)

	assert.False(t, p.AddCar(1))
	assert.False(t, p.AddCar(2))
	assert.True(t, p.AddCar(3))
}
