package world

import "gocraft/graphics"

type Renderer interface {
	CalculateMesh()
	Render(shader *graphics.ShaderProgram)
}
