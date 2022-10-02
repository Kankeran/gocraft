package world

import (
	"fmt"
	"gocraft/graphics"

	"github.com/go-gl/mathgl/mgl32"
)

var ver = [][]float32{
	{0, 0, 1}, // 0 bottom left front
	{0, 1, 1}, // 1 top left front
	{1, 1, 1}, // 2 top right front
	{1, 0, 1}, // 3 bottom right front

	{1, 0, 0}, // 4 bottom right back
	{1, 1, 0}, // 5 top right back
	{0, 1, 0}, // 6 top left back
	{0, 0, 0}, // 7 bottom left back
}

var inds = [][]uint32{
	Front:  {0, 1, 2, 3},
	Back:   {4, 5, 6, 7},
	Left:   {7, 6, 1, 0},
	Right:  {3, 2, 5, 4},
	Top:    {1, 6, 5, 2},
	Bottom: {7, 0, 3, 4},
}

type Block struct {
	position            mgl32.Vec3
	textureCoordsGetter func(face int) *graphics.TextureCoords
	chunk               *Chunk
}

func (b *Block) Position() mgl32.Vec3 {
	return b.position
}

func (b *Block) TextureCoords(faceID int) *graphics.TextureCoords {
	return b.textureCoordsGetter(faceID)
}

func (b *Block) DeleteFromWorld() {
	b.chunk.blocks[int(b.position[0])][int(b.position[1])][int(b.position[2])] = nil
	b.chunk.CalculateMesh()
}

func (b *Block) String() string {
	return fmt.Sprint("block", b.position)
}
