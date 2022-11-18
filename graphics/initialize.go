package graphics

var (
	diTextureAtlas              *Texture
	diShaderProgramBasicTexture *ShaderProgram
	diShaderProgramBasicColor   *ShaderProgram
	diShaderProgramBasic        *ShaderProgram
	diTexAtlas                  *TexAtlas
	diTextureCoordsGrass        *TextureCoords
	diTextureCoordsHalfGrass    *TextureCoords
	diTextureCoordsDirt         *TextureCoords
	diTextureCoordsStone        *TextureCoords
	diLine                      *Line
)

func ProvideTextureAtlas() *Texture {
	if diTextureAtlas != nil {
		return diTextureAtlas
	}
	tex, err := NewTexture("texAtlas")
	if err != nil {
		panic(err)
	}
	diTextureAtlas = tex

	return diTextureAtlas
}

func ProvideShaderProgramBasicTexture() *ShaderProgram {
	if diShaderProgramBasicTexture != nil {
		return diShaderProgramBasicTexture
	}
	shader, err := NewShader("basicTexture")
	if err != nil {
		panic(err)
	}
	diShaderProgramBasicTexture = shader

	return diShaderProgramBasicTexture
}

func ProvideShaderProgramBasicColor() *ShaderProgram {
	if diShaderProgramBasicColor != nil {
		return diShaderProgramBasicColor
	}
	shader, err := NewShader("basicColor")
	if err != nil {
		panic(err)
	}
	diShaderProgramBasicColor = shader

	return diShaderProgramBasicColor
}

func ProvideShaderProgramBasic() *ShaderProgram {
	if diShaderProgramBasic != nil {
		return diShaderProgramBasic
	}
	shader, err := NewShader("basic")
	if err != nil {
		panic(err)
	}
	diShaderProgramBasic = shader

	return diShaderProgramBasic
}

func ProvideTexAtlas() *TexAtlas {
	if diTexAtlas != nil {
		return diTexAtlas
	}
	diTexAtlas = NewTexAtlas(ProvideTextureAtlas().width, ProvideTextureAtlas().height, 16, 16)

	return diTexAtlas
}

func ProvideTextureCoordsGrass() *TextureCoords {
	if diTextureCoordsGrass != nil {
		return diTextureCoordsGrass
	}
	diTextureCoordsGrass = ProvideTexAtlas().GetCoordsByIDs(TextureIndexes{0, 0, 1, 1})

	return diTextureCoordsGrass
}

func ProvideTextureCoordsHalfGrass() *TextureCoords {
	if diTextureCoordsHalfGrass != nil {
		return diTextureCoordsHalfGrass
	}
	diTextureCoordsHalfGrass = ProvideTexAtlas().GetCoordsByIDs(TextureIndexes{1, 0, 2, 1})

	return diTextureCoordsHalfGrass
}

func ProvideTextureCoordsDirt() *TextureCoords {
	if diTextureCoordsDirt != nil {
		return diTextureCoordsDirt
	}
	diTextureCoordsDirt = ProvideTexAtlas().GetCoordsByIDs(TextureIndexes{2, 0, 3, 1})

	return diTextureCoordsDirt
}

func ProvideTextureCoordsStone() *TextureCoords {
	if diTextureCoordsStone != nil {
		return diTextureCoordsStone
	}
	diTextureCoordsStone = ProvideTexAtlas().GetCoordsByIDs(TextureIndexes{3, 0, 4, 1})

	return diTextureCoordsStone
}

func ProvideLine() *Line {
	if diLine != nil {
		return diLine
	}
	diLine = NewLine()

	return diLine
}
