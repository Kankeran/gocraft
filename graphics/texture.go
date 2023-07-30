package graphics

import (
	"embed"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"unsafe"

	"gocraft/gl"
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
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.CreateTextures(gl.TEXTURE_2D, 1, &texture)
	gl.TextureStorage2D(texture, 1, gl.RGBA8, int32(rgba.Rect.Dx()), int32(rgba.Rect.Dy()))

	gl.TextureParameteri(texture, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TextureParameteri(texture, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TextureParameteri(texture, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TextureParameteri(texture, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.TextureSubImage2D(
		texture,
		0,
		0,
		0,
		int32(rgba.Rect.Dx()),
		int32(rgba.Rect.Dy()),
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		unsafe.Pointer(&rgba.Pix[0]),
	)

	return &Texture{texture, rgba.Rect.Dx(), rgba.Rect.Dy()}, nil
}

func (t *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, t.handle)
}
