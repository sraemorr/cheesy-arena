// Copyright 2017 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)

package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotTriggered(t *testing.T) {
	touchpad := Touchpad{}
	touchpad.UpdateState(false, matchStartTime, timeAfterEnd(-10))

	touchpad.UpdateState(false, matchStartTime, timeAfterEnd(-1))
	assert.Equal(t, NotTriggered, touchpad.GetState(timeAfterEnd(-1)))
	assert.Equal(t, NotTriggered, touchpad.GetState(timeAfterEnd(2)))
}

func TestTriggeredReleasedEarly(t *testing.T) {
	touchpad := Touchpad{}
	touchpad.UpdateState(false, matchStartTime, timeAfterEnd(-10))

	touchpad.UpdateState(true, matchStartTime, timeAfterEnd(-5))
	assert.Equal(t, Triggered, touchpad.GetState(timeAfterEnd(-4.9)))
	assert.Equal(t, Held, touchpad.GetState(timeAfterEnd(-3)))
	touchpad.UpdateState(false, matchStartTime, timeAfterEnd(-1))
	assert.Equal(t, NotTriggered, touchpad.GetState(timeAfterEnd(-1.1)))
	assert.Equal(t, NotTriggered, touchpad.GetState(timeAfterEnd(2)))
}

func TestTriggeredTooShort(t *testing.T) {
	touchpad := Touchpad{}
	touchpad.UpdateState(false, matchStartTime, timeAfterEnd(-10))

	touchpad.UpdateState(true, matchStartTime, timeAfterEnd(-0.5))
	touchpad.UpdateState(true, matchStartTime, timeAfterEnd(0))
	assert.Equal(t, Triggered, touchpad.GetState(timeAfterEnd(0.2)))
	touchpad.UpdateState(false, matchStartTime, timeAfterEnd(0.4))
	assert.Equal(t, NotTriggered, touchpad.GetState(timeAfterEnd(0.5)))
	assert.Equal(t, NotTriggered, touchpad.GetState(timeAfterEnd(2)))

	touchpad.UpdateState(true, matchStartTime, timeAfterEnd(3))
	assert.Equal(t, NotTriggered, touchpad.GetState(timeAfterEnd(5)))
	touchpad.UpdateState(false, matchStartTime, timeAfterEnd(6))
	assert.Equal(t, NotTriggered, touchpad.GetState(timeAfterEnd(8)))
}

func TestTriggeredHeld(t *testing.T) {
	touchpad := Touchpad{}
	touchpad.UpdateState(false, matchStartTime, timeAfterEnd(-10))

	touchpad.UpdateState(true, matchStartTime, timeAfterEnd(-5))
	touchpad.UpdateState(true, matchStartTime, timeAfterEnd(-3))
	touchpad.UpdateState(true, matchStartTime, timeAfterEnd(1))
	assert.Equal(t, Held, touchpad.GetState(timeAfterEnd(2)))
}

func TestTriggeredReleased(t *testing.T) {
	touchpad := Touchpad{}
	touchpad.UpdateState(false, matchStartTime, timeAfterEnd(-10))

	touchpad.UpdateState(true, matchStartTime, timeAfterEnd(-5))
	touchpad.UpdateState(true, matchStartTime, timeAfterEnd(-3))
	touchpad.UpdateState(true, matchStartTime, timeAfterEnd(1))
	assert.Equal(t, Held, touchpad.GetState(timeAfterEnd(2)))
	touchpad.UpdateState(false, matchStartTime, timeAfterEnd(3))
	assert.Equal(t, Held, touchpad.GetState(timeAfterEnd(4)))
}

func TestReTriggered(t *testing.T) {
	touchpad := Touchpad{}
	touchpad.UpdateState(false, matchStartTime, timeAfterEnd(-10))

	touchpad.UpdateState(true, matchStartTime, timeAfterEnd(-5))
	assert.Equal(t, Held, touchpad.GetState(timeAfterEnd(-3)))
	touchpad.UpdateState(false, matchStartTime, timeAfterEnd(-1))
	assert.Equal(t, NotTriggered, touchpad.GetState(timeAfterEnd(-1.1)))
	touchpad.UpdateState(true, matchStartTime, timeAfterEnd(-0.1))
	assert.Equal(t, Triggered, touchpad.GetState(timeAfterEnd(0.1)))
	assert.Equal(t, Held, touchpad.GetState(timeAfterEnd(2)))
}

func TestTriggeredLate(t *testing.T) {
	touchpad := Touchpad{}
	touchpad.UpdateState(false, matchStartTime, timeAfterEnd(-10))

	touchpad.UpdateState(true, matchStartTime, timeAfterEnd(0.1))
	assert.Equal(t, NotTriggered, touchpad.GetState(timeAfterEnd(0.2)))
	assert.Equal(t, NotTriggered, touchpad.GetState(timeAfterEnd(2)))
}

func TestCountTouchpads(t *testing.T) {
	var touchpads [3]Touchpad
	touchpads[0].UpdateState(true, matchStartTime, timeAfterEnd(-5))
	touchpads[1].UpdateState(true, matchStartTime, timeAfterEnd(-2))
	touchpads[2].UpdateState(true, matchStartTime, timeAfterEnd(-0.1))

	assert.Equal(t, 0, CountTouchpads(&touchpads, timeAfterEnd(-6)))
	assert.Equal(t, 0, CountTouchpads(&touchpads, timeAfterEnd(-5.5)))
	assert.Equal(t, 1, CountTouchpads(&touchpads, timeAfterEnd(-3)))
	assert.Equal(t, 1, CountTouchpads(&touchpads, timeAfterEnd(-1.5)))
	assert.Equal(t, 2, CountTouchpads(&touchpads, timeAfterEnd(0)))
	assert.Equal(t, 3, CountTouchpads(&touchpads, timeAfterEnd(1)))
}
