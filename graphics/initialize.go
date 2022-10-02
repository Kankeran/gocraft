package graphics

import (
	"gocraft/services"
)

var (
	ServiceTextureAtlas              = new(services.Item)
	ServiceShaderProgramBasicTexture = new(services.Item)
	ServiceShaderProgramBasicColor   = new(services.Item)
	ServiceShaderProgramBasic        = new(services.Item)
	ServiceTexAtlasTexAtlas          = new(services.Item)
	ServiceTextureCoordsGrass        = new(services.Item)
	ServiceTextureCoordsHalfGrass    = new(services.Item)
	ServiceTextureCoordsDirt         = new(services.Item)
	ServiceTextureCoordsStone        = new(services.Item)
	ServiceLine                      = new(services.Item)
)

func init() {
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceTextureAtlas, func() interface{} {
			tex, err := NewTexture("texAtlas")
			if err != nil {
				panic(err)
			}
			return tex
		}
	})
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceShaderProgramBasicTexture, func() interface{} {
			shader, err := NewShader("basicTexture")
			if err != nil {
				panic(err)
			}
			return shader
		}
	})
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceShaderProgramBasicColor, func() interface{} {
			shader, err := NewShader("basicColor")
			if err != nil {
				panic(err)
			}
			return shader
		}
	})
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceShaderProgramBasic, func() interface{} {
			shader, err := NewShader("basic")
			if err != nil {
				panic(err)
			}
			return shader
		}
	})
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceTexAtlasTexAtlas, func() interface{} {
			texture := c.GetService(ServiceTextureAtlas).(*Texture)
			return NewTexAtlas(texture.width, texture.height, 16, 16)
		}
	})

	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceTextureCoordsGrass, func() interface{} {
			return c.GetService(ServiceTexAtlasTexAtlas).(*TexAtlas).GetCoordsByIDs(TextureIndexes{0, 0, 1, 1})
		}
	})
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceTextureCoordsHalfGrass, func() interface{} {
			return c.GetService(ServiceTexAtlasTexAtlas).(*TexAtlas).GetCoordsByIDs(TextureIndexes{1, 0, 2, 1})
		}
	})
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceTextureCoordsDirt, func() interface{} {
			return c.GetService(ServiceTexAtlasTexAtlas).(*TexAtlas).GetCoordsByIDs(TextureIndexes{2, 0, 3, 1})
		}
	})
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceTextureCoordsStone, func() interface{} {
			return c.GetService(ServiceTexAtlasTexAtlas).(*TexAtlas).GetCoordsByIDs(TextureIndexes{3, 0, 4, 1})
		}
	})
	services.InjectInvoker(func(c *services.Container) (*services.Item, services.Invoker) {
		return ServiceLine, func() interface{} {
			return NewLine()
		}
	})
}
