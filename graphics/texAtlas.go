package graphics

const (
	LeftBottom int = iota
	LeftTop
	RightTop
	RightBottom
)

type TextureIndexes struct {
	minXID, minYID, maxXID, maxYID int
}

type TextureCoords struct {
	minX, minY, maxX, maxY float32
}

func (tc *TextureCoords) EdgeCoord(combine int) []float32 {
	switch combine {
	case LeftBottom:
		return []float32{tc.minX, tc.minY}
	case LeftTop:
		return []float32{tc.minX, tc.maxY}
	case RightTop:
		return []float32{tc.maxX, tc.maxY}
	case RightBottom:
		return []float32{tc.maxX, tc.minY}
	}
	panic("combine out of range")
}

type TexAtlas struct {
	texWidth, texHeight float32
}

func NewTexAtlas(atlasWidth, atlasHeight, texPixelsWidth, texPixelsHeight int) *TexAtlas {
	return &TexAtlas{
		float32(texPixelsWidth) / float32(atlasWidth),
		float32(texPixelsHeight) / float32(atlasHeight),
	}
}

func (ta *TexAtlas) GetCoordsByIDs(textureIndexes TextureIndexes) *TextureCoords {
	return &TextureCoords{
		ta.texWidth * float32(textureIndexes.minXID),
		ta.texHeight * float32(-textureIndexes.minYID-15),
		ta.texWidth * float32(textureIndexes.maxXID),
		ta.texHeight * float32(-textureIndexes.maxYID-15),
	}
}
