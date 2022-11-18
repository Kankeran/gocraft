package graphics

import (
	"embed"
	"fmt"
	"image"
	"image/draw"
	"image/png"

	"github.com/go-gl/gl/v4.5-core/gl"
)

//go:embed textures
var textureResources embed.FS

type Texture struct {
	handle        uint32
	width, height int
}

func NewTexture(name string) (*Texture, error) {
	imgFile, err := textureResources.Open("textures/" + name + ".png")
	if err != nil {
		return nil, fmt.Errorf("texture %q not found on disk: %w", name, err)
	}
	defer imgFile.Close()
	img, err := png.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return &Texture{texture, rgba.Rect.Size().X, rgba.Rect.Size().Y}, nil
}

func (t *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, t.handle)
}
