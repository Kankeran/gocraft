package world

import (
	"gocraft/control"
	"gocraft/graphics"
)

var (
	diBlockTypeGrass *BlockType
	diBlockTypeDirt  *BlockType
	diBlockTypeStone *BlockType
	diWorld          *World
)

func ProvideBlockTypeGrass() *BlockType {
	if diBlockTypeGrass != nil {
		return diBlockTypeGrass
	}
	diBlockTypeGrass = &BlockType{func(faceID int) *graphics.TextureCoords {
		switch faceID {
		case Front, Back, Left, Right:
			return graphics.ProvideTextureCoordsHalfGrass()
		case Top:
			return graphics.ProvideTextureCoordsGrass()
		case Bottom:
			return graphics.ProvideTextureCoordsDirt()
		}
		panic("incorrect face ID")
	}}

	return diBlockTypeGrass
}

func ProvideTypeDirt() *BlockType {
	if diBlockTypeDirt != nil {
		return diBlockTypeDirt
	}
	diBlockTypeDirt = &BlockType{func(faceID int) *graphics.TextureCoords {
		switch faceID {
		case Bottom, Top, Front, Back, Left, Right:
			return graphics.ProvideTextureCoordsDirt()
		}
		panic("incorrect face ID")
	}}

	return diBlockTypeDirt
}

func ProvideBlockTypeStone() *BlockType {
	if diBlockTypeStone != nil {
		return diBlockTypeStone
	}
	diBlockTypeStone = &BlockType{func(faceID int) *graphics.TextureCoords {
		switch faceID {
		case Bottom, Top, Front, Back, Left, Right:
			return graphics.ProvideTextureCoordsStone()
		}
		panic("incorrect face ID")
	}}

	return diBlockTypeStone
}

func ProvideWorld() *World {
	if diWorld != nil {
		return diWorld
	}
	diWorld = &World{
		playerPos:      control.ParamStartPos,
		grassBlockType: ProvideBlockTypeGrass(),
		dirtBlockType:  ProvideTypeDirt(),
		stoneBlockType: ProvideBlockTypeStone(),
	}

	return diWorld
}

func init() {
	control.ProvideCamera().AddPositionObserver(ProvideWorld())
}
