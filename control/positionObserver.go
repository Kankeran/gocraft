package control

import "github.com/go-gl/mathgl/mgl32"

type PositionObserver interface {
	PositionObserve(position mgl32.Vec3)
}
