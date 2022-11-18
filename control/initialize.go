package control

import (
	"github.com/go-gl/mathgl/mgl32"
)

var (
	ParamCameraUp             = mgl32.Vec3{0, 1, 0}
	ParamStartPos             = mgl32.Vec3{0, 17, 3}
	diCamera                  *Camera
	diCameraMovementModelFree *CameraMovementModel
	diCameraMovementModelFps  *CameraMovementModel
)

func ProvideCamera() *Camera {
	if diCamera != nil {
		return diCamera
	}
	diCamera = &Camera{
		position:      ParamStartPos,
		frontPosition: mgl32.Vec3{0, 0, -1},
		upPosition:    ParamCameraUp,
		firstMouse:    true,
		lock:          true,
		lastX:         0,
		lastY:         0,
		yaw:           -90,
		pitch:         0,
		movementModel: ProvideCameraMovementModelFree(),
	}

	return diCamera
}

func ProvideCameraMovementModelFree() *CameraMovementModel {
	if diCameraMovementModelFree != nil {
		return diCameraMovementModelFree
	}
	diCameraMovementModelFree = &CameraMovementModel{
		cameraUp: ParamCameraUp,
		forward: func(cameraFront mgl32.Vec3, cameraUp mgl32.Vec3, speed float32) mgl32.Vec3 {
			return cameraFront.Mul(speed)
		},
		right: func(cameraFront mgl32.Vec3, cameraUp mgl32.Vec3, speed float32) mgl32.Vec3 {
			return cameraFront.Cross(cameraUp).Normalize().Mul(speed)
		},
		up: func(cameraFront mgl32.Vec3, cameraUp mgl32.Vec3, speed float32) mgl32.Vec3 {
			return cameraUp.Mul(speed)
		},
	}

	return diCameraMovementModelFree
}

func ProvideCameraMovementModelFps() *CameraMovementModel {
	if diCameraMovementModelFps != nil {
		return diCameraMovementModelFps
	}
	diCameraMovementModelFps = &CameraMovementModel{
		cameraUp: ParamCameraUp,
		forward: func(cameraFront mgl32.Vec3, cameraUp mgl32.Vec3, speed float32) mgl32.Vec3 {
			return mgl32.Vec3{cameraFront.X(), 0, cameraFront.Z()}.Normalize().Mul(speed)
		},
		right: func(cameraFront mgl32.Vec3, cameraUp mgl32.Vec3, speed float32) mgl32.Vec3 {
			return mgl32.Vec3{cameraFront.X(), 0, cameraFront.Z()}.Normalize().Cross(cameraUp).Normalize().Mul(speed)
		},
		up: func(cameraFront mgl32.Vec3, cameraUp mgl32.Vec3, speed float32) mgl32.Vec3 {
			return cameraUp.Mul(speed)
		},
	}

	return diCameraMovementModelFps
}
