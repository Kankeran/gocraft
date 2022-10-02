package world

import (
	"gocraft/control"
	"gocraft/graphics"
	"gocraft/services"

	"github.com/go-gl/mathgl/mgl32"
)

var (
	ServiceBlockTypeGrass = new(services.Item)
	ServiceBlockTypeDirt  = new(services.Item)
	ServiceBlockTypeStone = new(services.Item)
	ServiceWorld          = new(services.Item)
)

func init() {
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceBlockTypeGrass, func() interface{} {
			return &BlockType{func(faceID int) *graphics.TextureCoords {
				switch faceID {
				case Front, Back, Left, Right:
					return c.GetService(graphics.ServiceTextureCoordsHalfGrass).(*graphics.TextureCoords)
				case Top:
					return c.GetService(graphics.ServiceTextureCoordsGrass).(*graphics.TextureCoords)
				case Bottom:
					return c.GetService(graphics.ServiceTextureCoordsDirt).(*graphics.TextureCoords)
				}
				panic("incorrect face ID")
			}}
		}
	})
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceBlockTypeDirt, func() interface{} {
			return &BlockType{func(faceID int) *graphics.TextureCoords {
				switch faceID {
				case Bottom, Top, Front, Back, Left, Right:
					return c.GetService(graphics.ServiceTextureCoordsDirt).(*graphics.TextureCoords)
				}
				panic("incorrect face ID")
			}}
		}
	})
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceBlockTypeStone, func() interface{} {
			return &BlockType{func(faceID int) *graphics.TextureCoords {
				switch faceID {
				case Bottom, Top, Front, Back, Left, Right:
					return c.GetService(graphics.ServiceTextureCoordsStone).(*graphics.TextureCoords)
				}
				panic("incorrect face ID")
			}}
		}
	})
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceWorld, func() interface{} {
			return &World{
				playerPos:      c.GetParameter(control.ParamStartPos).(mgl32.Vec3),
				grassBlockType: c.GetService(ServiceBlockTypeGrass).(*BlockType),
				dirtBlockType:  c.GetService(ServiceBlockTypeDirt).(*BlockType),
				stoneBlockType: c.GetService(ServiceBlockTypeStone).(*BlockType),
			}
		}
	})
	services.AddInitializer(func(c *services.Container) {
		c.GetService(ServiceWorld).(*World).Initialize()
		c.GetService(control.ServiceCamera).(*control.Camera).AddPositionObserver(c.GetService(ServiceWorld).(*World))
	})
}
