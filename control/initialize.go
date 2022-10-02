package control

import (
	"gocraft/services"

	"github.com/go-gl/mathgl/mgl32"
)

var (
	ParamCameraUp                  = new(services.Item)
	ParamStartPos                  = new(services.Item)
	ServiceCamera                  = new(services.Item)
	ServiceCameraMovementModelFree = new(services.Item)
	ServiceCameraMovementModelFps  = new(services.Item)
)

func init() {
	services.InjectParameter(func() (*services.Item, interface{}) {
		return ParamCameraUp, mgl32.Vec3{0, 1, 0}
	})
	services.InjectParameter(func() (*services.Item, interface{}) {
		return ParamStartPos, mgl32.Vec3{0, 17, 3}
	})
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceCamera, func() interface{} {
			return &Camera{
				position:      c.GetParameter(ParamStartPos).(mgl32.Vec3),
				frontPosition: mgl32.Vec3{0, 0, -1},
				upPosition:    c.GetParameter(ParamCameraUp).(mgl32.Vec3),
				firstMouse:    true,
				lock:          true,
				lastX:         0,
				lastY:         0,
				yaw:           -90,
				pitch:         0,
				movementModel: c.GetService(ServiceCameraMovementModelFree).(*CameraMovementModel),
			}
		}
	})
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceCameraMovementModelFree, func() interface{} {
			return &CameraMovementModel{
				cameraUp: c.GetParameter(ParamCameraUp).(mgl32.Vec3),
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
		}
	})
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceCameraMovementModelFps, func() interface{} {
			return &CameraMovementModel{
				cameraUp: c.GetParameter(ParamCameraUp).(mgl32.Vec3),
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
		}
	})
}
