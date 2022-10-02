package world

import (
	"fmt"
	"gocraft/graphics"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

const renderChunks = 4

type World struct {
	chunks                                        []*Chunk
	playerPos                                     mgl32.Vec3
	grassBlockType, dirtBlockType, stoneBlockType *BlockType
	lines                                         [][2]mgl32.Vec3
}

func (w *World) PositionObserve(position mgl32.Vec3) {
	w.playerPos = position
}

func (w *World) Initialize() {
	for x := -16 * renderChunks; x <= 16*renderChunks; x += 16 {
		for y := 0; y <= 16*renderChunks; y += 16 {
			for z := -16 * renderChunks; z <= 16*renderChunks; z += 16 {
				w.generateChunk(mgl32.Vec3{float32(x), float32(y), float32(z)})
			}
		}
	}
}

func (w *World) generateChunk(position mgl32.Vec3) {
	var blocks []*Block
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			for z := 0; z < 16; z++ {
				if y == 15 && position.Y() == 16*renderChunks {
					blocks = append(blocks, w.grassBlockType.CreateBlock(mgl32.Vec3{float32(x), float32(y), float32(z)}))
					continue
				}
				if y > 10 && position.Y() == 16*renderChunks {
					blocks = append(blocks, w.dirtBlockType.CreateBlock(mgl32.Vec3{float32(x), float32(y), float32(z)}))
					continue
				}
				blocks = append(blocks, w.stoneBlockType.CreateBlock(mgl32.Vec3{float32(x), float32(y), float32(z)}))
			}
		}
	}
	w.chunks = append(w.chunks, NewChunk(position, blocks))
}

func (w *World) Render(shader *graphics.ShaderProgram) {
	for _, chunk := range w.chunks {
		chunk.Render(shader)
	}
}

func (w *World) CalculateMesh() {
	for _, chunk := range w.chunks {
		chunk.CalculateMesh()
	}
}

func (w *World) FindBlockByRaycast(from, direction mgl32.Vec3, length float32) *Block {
	var block *Block
	w.lines = w.lines[:0]
	var rayUnitStepSize = mgl32.Vec3{
		float32(math.Sqrt(float64(1 + (direction.Y()/direction.X())*(direction.Y()/direction.X()) + (direction.Z()/direction.X())*(direction.Z()/direction.X())))),
		float32(math.Sqrt(float64(1 + (direction.X()/direction.Y())*(direction.X()/direction.Y()) + (direction.Z()/direction.Y())*(direction.Z()/direction.Y())))),
		float32(math.Sqrt(float64(1 + (direction.Y()/direction.Z())*(direction.Y()/direction.Z()) + (direction.X()/direction.Z())*(direction.X()/direction.Z())))),
	}

	var rayLength1D = mgl32.Vec3{
		(float32(int(from.X())) - from.X()) * rayUnitStepSize.X(),
		(float32(int(from.Y())) - from.Y()) * rayUnitStepSize.Y(),
		(float32(int(from.Z())) - from.Z()) * rayUnitStepSize.Z(),
	}

	if direction.X() < 0 {
		rayLength1D[0] = -rayLength1D[0]
	}

	if direction.Y() < 0 {
		rayLength1D[1] = -rayLength1D[1]
	}

	if direction.Z() < 0 {
		rayLength1D[2] = -rayLength1D[2]
	}

	var shouldSearch = true
	var distance float32 = 0
	var lastPoint = from
	var currentPoint mgl32.Vec3
	var halfPoints []mgl32.Vec3
	for shouldSearch {
		if rayLength1D.X() < rayLength1D.Y() && rayLength1D.X() < rayLength1D.Z() {
			distance = rayLength1D[0]
			rayLength1D[0] += rayUnitStepSize[0]
		} else if rayLength1D.Y() < rayLength1D.X() && rayLength1D.Y() < rayLength1D.Z() {
			distance = rayLength1D[1]
			rayLength1D[1] += rayUnitStepSize[1]
		} else {
			distance = rayLength1D[2]
			rayLength1D[2] += rayUnitStepSize[2]
		}

		if distance < 0 {
			continue
		}

		if distance > length {
			shouldSearch = false
			distance = length
		}

		currentPoint = from.Add(direction.Mul(distance))
		halfPoints = append(halfPoints, lastPoint.Add(currentPoint.Sub(lastPoint).Mul(0.5)))
		lastPoint = currentPoint
	}

blockSearch:
	for i := 0; i < len(halfPoints); i++ {
		for _, chunk := range w.chunks {
			if chunk.IsDirIn(halfPoints[i]) {
				xIndex := int(halfPoints[i][0] - chunk.position[0])
				yIndex := int(halfPoints[i][1] - chunk.position[1])
				zIndex := int(halfPoints[i][2] - chunk.position[2])
				block = chunk.blocks[xIndex][yIndex][zIndex]
				if block != nil {
					if i > 0 {
						point := []int{
							SubFloat32ToInt(halfPoints[i-1][0], halfPoints[i][0]),
							SubFloat32ToInt(halfPoints[i-1][1], halfPoints[i][1]),
							SubFloat32ToInt(halfPoints[i-1][2], halfPoints[i][2]),
						}
						fmt.Println(point, halfPoints[i-1], halfPoints[i])
					}
					break blockSearch
				}
				break
			}
		}
	}

	return block
}

func SubFloat32ToInt(a, b float32) int {
	if int(a) == 0 && int(b) == 0 {
		if a < 0 && b > 0 {
			return -1
		}
		if b < 0 && a > 0 {
			return 1
		}
	}

	return int(a) - int(b)
}

func (w *World) Lines() [][2]mgl32.Vec3 {
	return w.lines
}
