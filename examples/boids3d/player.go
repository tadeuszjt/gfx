package main

import (
	"github.com/tadeuszjt/geom/generic"
	"math"
	//"fmt"
)

const (
	playerSpeed           = 2
	playerLookSensitivity = 0.015
)

var (
	player = struct {
		position geom.Vec3[float32]
		bearing  float32
		pitch    float32
	}{
		position: geom.Vec3[float32]{0, 0, 3},
	}

	keys struct{ w, a, s, d bool }
)

func playerUpdate() {
	forward := geom.Vec3NormPitchYaw(player.pitch, player.bearing)
	right := geom.Vec3NormPitchYaw(0, player.bearing+float32(math.Pi/2))

	if keys.w {
		player.position.PlusEquals(forward.ScaledBy(playerSpeed))
	}
	if keys.a {
		player.position.MinusEquals(right.ScaledBy(playerSpeed))
	}
	if keys.s {
		player.position.MinusEquals(forward.ScaledBy(playerSpeed))
	}
	if keys.d {
		player.position.PlusEquals(right.ScaledBy(playerSpeed))
	}
}

func playerLook(dx, dy float32) {
	player.bearing += dx * playerLookSensitivity
	player.pitch += dy * playerLookSensitivity

	if player.pitch > (math.Pi / 2) {
		player.pitch = math.Pi / 2
	} else if player.pitch < (-math.Pi / 2) {
		player.pitch = -math.Pi / 2
	}
}
