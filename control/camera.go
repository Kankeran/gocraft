package control

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
)

type CameraMovementModel struct {
	cameraUp           mgl32.Vec3
	forward, right, up func(mgl32.Vec3, mgl32.Vec3, float32) mgl32.Vec3
}

func (c *CameraMovementModel) Forward(cameraFront mgl32.Vec3, speed float32) mgl32.Vec3 {
	return c.forward(cameraFront, c.cameraUp, speed)
}

func (c *CameraMovementModel) Right(cameraFront mgl32.Vec3, speed float32) mgl32.Vec3 {
	return c.right(cameraFront, c.cameraUp, speed)
}

func (c *CameraMovementModel) Up(cameraFront mgl32.Vec3, speed float32) mgl32.Vec3 {
	return c.up(cameraFront, c.cameraUp, speed)
}

type Camera struct {
	position, frontPosition, upPosition mgl32.Vec3
	firstMouse, lock                    bool
	lastX, lastY, yaw, pitch            float64
	movementModel                       *CameraMovementModel
	positionObservers                   []PositionObserver
}

func (c *Camera) AddPositionObserver(positionObserver PositionObserver) {
	c.positionObservers = append(c.positionObservers, positionObserver)
}

func (c *Camera) Lock() {
	c.lock = true
}

func (c *Camera) Unlock() {
	c.lock = false
	c.firstMouse = true
}

func (c *Camera) MouseMoved(xpos, ypos float64) {
	if c.lock {
		return
	}

	if c.firstMouse {
		c.lastX, c.lastY = xpos, ypos
		c.firstMouse = false
	}

	c.yaw += (xpos - c.lastX) * 0.1
	c.pitch += (c.lastY - ypos) * 0.1

	if c.pitch > 89.9 {
		c.pitch = 89.9
	}
	if c.pitch < -89.9 {
		c.pitch = -89.9
	}
	c.lastX, c.lastY = xpos, ypos

	c.frontPosition = mgl32.Vec3{
		float32(math.Cos(mgl64.DegToRad(c.yaw)) * math.Cos(mgl64.DegToRad(c.pitch))),
		float32(math.Sin(mgl64.DegToRad(c.pitch))),
		float32(math.Sin(mgl64.DegToRad(c.yaw)) * math.Cos(mgl64.DegToRad(c.pitch))),
	}.Normalize()
}

func (c *Camera) Mat4() mgl32.Mat4 {
	return mgl32.LookAtV(c.position, c.position.Add(c.frontPosition), c.upPosition)
}

func (c *Camera) MoveForward(speed float32) {
	c.position = c.position.Add(c.movementModel.Forward(c.frontPosition, speed))
	c.callObservers()
}

func (c *Camera) MoveBackward(speed float32) {
	c.position = c.position.Sub(c.movementModel.Forward(c.frontPosition, speed))
	c.callObservers()
}

func (c *Camera) MoveRight(speed float32) {
	c.position = c.position.Add(c.movementModel.Right(c.frontPosition, speed))
	c.callObservers()
}

func (c *Camera) MoveLeft(speed float32) {
	c.position = c.position.Sub(c.movementModel.Right(c.frontPosition, speed))
	c.callObservers()
}

func (c *Camera) MoveUp(speed float32) {
	c.position = c.position.Add(c.movementModel.Up(c.frontPosition, speed))
	c.callObservers()
}

func (c *Camera) MoveDown(speed float32) {
	c.position = c.position.Sub(c.movementModel.Up(c.frontPosition, speed))
	c.callObservers()
}

func (c *Camera) SetMovementModel(model *CameraMovementModel) {
	c.movementModel = model
}

func (c *Camera) callObservers() {
	for _, observer := range c.positionObservers {
		observer.PositionObserve(c.position)
	}
}

func (c *Camera) Direction() mgl32.Vec3 {
	return c.frontPosition
}

func (c *Camera) Position() mgl32.Vec3 {
	return c.position
}
