package world

import (
	"gocraft/graphics"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	Front int = iota
	Back
	Left
	Right
	Top
	Bottom
)

type BlockType struct {
	textureCoordsGetter func(faceID int) *graphics.TextureCoords
}

func (b *BlockType) CreateBlock(position mgl32.Vec3) *Block {
	return &Block{position: position, textureCoordsGetter: b.textureCoordsGetter}
}
