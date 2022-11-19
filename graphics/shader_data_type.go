package graphics

import "gocraft/gl"

type ShaderDataType uint32

const (
	Float ShaderDataType = iota
	Float2
	Float3
	Float4
	Mat3
	Mat4
	Int
	Int2
	Int3
	Int4
	Bool
)

func (sType *ShaderDataType) Size() int32 {
	switch *sType {
	case Float, Float2, Float3, Float4:
		return 4 * sType.ComponentCount()
	case Mat3, Mat4:
		return 4 * sType.ComponentCount()
	case Int, Int2, Int3, Int4:
		return 4 * sType.ComponentCount()
	case Bool:
		return 1
	}

	panic("Unknown ShaderDataType!")
}

func (sType *ShaderDataType) ComponentCount() int32 {
	switch *sType {
	case Float:
		return 1
	case Float2:
		return 2
	case Float3:
		return 3
	case Float4:
		return 4
	case Mat3:
		return 3 * 3
	case Mat4:
		return 4 * 4
	case Int:
		return 1
	case Int2:
		return 2
	case Int3:
		return 3
	case Int4:
		return 4
	case Bool:
		return 1
	}

	panic("Unknown ShaderDataType!")
}

func (sType *ShaderDataType) OpenGLType() uint32 {
	switch *sType {
	case Float, Float2, Float3, Float4, Mat3, Mat4:
		return gl.FLOAT
	case Int, Int2, Int3, Int4:
		return gl.INT
	case Bool:
		return gl.BOOL
	}

	panic("Unknown ShaderDataType!")
}
