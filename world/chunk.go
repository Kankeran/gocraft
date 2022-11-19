package world

import (
	"errors"
	"gocraft/graphics"

	"gocraft/gl"

	"github.com/go-gl/mathgl/mgl32"
)

type Indices []uint32

func (i Indices) Move(value uint32) Indices {
	if len(i) != 6 {
		panic(errors.New("Indices must have 6 elements"))
	}
	return Indices{i[0] + value, i[1] + value, i[2] + value, i[3] + value, i[4] + value, i[5] + value}
}

var faceInd = Indices{0, 1, 3, 1, 2, 3}

type Chunk struct {
	position      mgl32.Vec3
	blocks        [16][16][16]*Block //x, y, z top=-y bottom=+y left=-x right=+x back=-z front=+z
	vao, vbo, ebo uint32
	indices       []uint32
	// vertices      []float32
}

func NewChunk(position mgl32.Vec3, blocks []*Block) *Chunk {
	var c = &Chunk{position: position}
	for _, block := range blocks {
		c.blocks[int(block.position.X())][int(block.position.Y())][int(block.position.Z())] = block
		block.chunk = c
	}
	gl.GenVertexArrays(1, &c.vao)
	gl.GenBuffers(1, &c.vbo)
	gl.GenBuffers(1, &c.ebo)

	return c
}

func (c *Chunk) CalculateMesh() {
	gl.BindVertexArray(c.vao)

	var vertices []float32
	var indsCount uint32
	var blockPosition mgl32.Vec3
	for xIndex, xyblocks := range c.blocks {
		for yIndex, xblocks := range xyblocks {
			for zIndex, block := range xblocks {
				if block == nil {
					continue
				}
				for faceID, vertexIndexes := range inds {
					if c.HasNeighbor(faceID, xIndex, yIndex, zIndex) {
						continue
					}
					for j, index := range vertexIndexes {
						blockPosition = block.position
						vertices = append(vertices, ver[index][0]+blockPosition.X(), ver[index][1]+blockPosition.Y(), ver[index][2]+blockPosition.Z())
						vertices = append(vertices, block.TextureCoords(faceID).EdgeCoord(j)...)
					}
					c.indices = append(c.indices, faceInd.Move(indsCount)...)
					indsCount += 4
				}
			}
		}
	}

	gl.BindBuffer(gl.ARRAY_BUFFER, c.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, c.ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(c.indices)*4, gl.Ptr(c.indices), gl.STATIC_DRAW)

	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 5*4, 0)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)
	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)
}

func (c *Chunk) Render(shader *graphics.ShaderProgram) {
	shader.SetUniformMat4("model\x00", mgl32.Translate3D(c.position.X(), c.position.Y(), c.position.Z()))
	gl.BindVertexArray(c.vao)
	gl.DrawElements(gl.TRIANGLES, int32(len(c.indices)), gl.UNSIGNED_INT, nil)
}

func (c *Chunk) HasNeighbor(faceID, x, y, z int) bool {
	switch faceID {
	case Front:
		z += 1
	case Back:
		z -= 1
	case Left:
		x -= 1
	case Right:
		x += 1
	case Top:
		y += 1
	case Bottom:
		y -= 1
	}

	return z >= 0 && z < 16 && x >= 0 && x < 16 && y >= 0 && y < 16 && c.blocks[x][y][z] != nil
}

func (c *Chunk) IsDirIn(dir mgl32.Vec3) bool {
	return dir[0] > c.position[0] &&
		dir[0] < c.position[0]+16 &&
		dir[1] > c.position[1] &&
		dir[1] < c.position[1]+16 &&
		dir[2] > c.position[2] &&
		dir[2] < c.position[2]+16
}
