package graphics

type BufferLayout struct {
	elements []*BufferElement
	stride   int32
}

func NewBufferLayout(elements ...*BufferElement) *BufferLayout {
	var bufferLayout = &BufferLayout{
		elements: elements,
	}
	var offset uintptr
	for _, element := range elements {
		element.offset = offset
		offset += uintptr(element.size)
		bufferLayout.stride += element.size
	}

	return bufferLayout
}

func (layout *BufferLayout) Elements() []*BufferElement {
	return layout.elements
}

func (layout *BufferLayout) Stride() int32 {
	return layout.stride
}
