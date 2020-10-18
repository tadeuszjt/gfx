package main

import (
    "math"
	"github.com/tadeuszjt/geom/32"
)

const (
	playerSpeed           = 0.1
    playerLookSensitivity = 0.015
)

var (
	player = struct {
		position geom.Vec3
		bearing  geom.Angle
		pitch    geom.Angle
	}{
		position: geom.Vec3{0, 0, 3},
	}

	keys struct{ w, a, s, d bool }
)

func playerUpdate() {
	if keys.w {
		player.position.Z -= playerSpeed * float32(math.Cos(float64(player.bearing)))
		player.position.X += playerSpeed * float32(math.Sin(float64(player.bearing)))
		player.position.Y -= playerSpeed * float32(math.Sin(float64(player.pitch)))
	}
	if keys.a {
		player.position.X -= playerSpeed * float32(math.Cos(float64(player.bearing)))
		player.position.Z -= playerSpeed * float32(math.Sin(float64(player.bearing)))
	}
	if keys.s {
		player.position.Z += playerSpeed * float32(math.Cos(float64(player.bearing)))
		player.position.X -= playerSpeed * float32(math.Sin(float64(player.bearing)))
		player.position.Y += playerSpeed * float32(math.Sin(float64(player.pitch)))
	}
	if keys.d {
		player.position.X += playerSpeed * float32(math.Cos(float64(player.bearing)))
		player.position.Z += playerSpeed * float32(math.Sin(float64(player.bearing)))
	}
}

func playerLook(dx, dy float32) {
    player.bearing.PlusEquals(geom.MakeAngle(dx * playerLookSensitivity))
    player.pitch.PlusEquals(geom.MakeAngle(dy * playerLookSensitivity))
    player.pitch.Clamp(-geom.Angle90Deg, geom.Angle90Deg)
}
