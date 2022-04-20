package common

import "github.com/hajimehoshi/ebiten/v2"

var shaderProgram *ebiten.Shader

func init() {
	var err error
	shaderProgram, err = ebiten.NewShader([]byte(`
		package main
		
		func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
			size := imageDstTextureSize()
			aspect := float(size.y)/float(size.x)
			border_width := 0.005
			maxX := 1.0 - border_width;
   			minX := border_width;
   			maxY := (maxX / aspect) - (1 + border_width);
			minY := (minX / aspect) + border_width/2;
   			if (texCoord.x < maxX && texCoord.x > minX && texCoord.y < maxY && texCoord.y > minY) {
				return vec4(color.x/3,color.y/3,color.z/3,color.w/3)
			} else {
				return color
			}
		}
	`))
	if err != nil {
		panic(err)
	}
}

type RectangleHelper struct {
	red          float32
	green        float32
	blue         float32
	alpha        float32
	vertexAmount int
	vertices     []ebiten.Vertex
	indices      []uint16
	boxIndices   []uint16
}

func NewRectangleHelper() *RectangleHelper {
	r := RectangleHelper{
		red:          1,
		blue:         1,
		green:        1,
		vertexAmount: 0,
		vertices:     make([]ebiten.Vertex, 0, 8*2*3),
		indices:      make([]uint16, 0, 8*2*3),
		boxIndices:   []uint16{0, 1, 2, 1, 2, 3},
	}
	return &r
}

func (r *RectangleHelper) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawTrianglesShaderOptions{FillRule: ebiten.FillAll}
	screen.DrawTrianglesShader(r.vertices, r.indices, shaderProgram, op)
}

func (r *RectangleHelper) Reset() {
	r.vertexAmount = 0
	r.vertices = r.vertices[:0]
	r.indices = r.indices[:0]

}

func (r *RectangleHelper) SetColor(red, green, blue, alpha float32) {
	r.red = red
	r.green = green
	r.blue = blue
	r.alpha = alpha
}

func (r *RectangleHelper) AddFilledRectangle(x, y, width, height float32) {

}

func (r *RectangleHelper) AddRectangle(x, y, width, height float32) {
	r.vertices = append(r.vertices, []ebiten.Vertex{{
		DstX:   x,
		DstY:   y,
		SrcX:   0,
		SrcY:   0,
		ColorR: r.red,
		ColorG: r.green,
		ColorB: r.blue,
		ColorA: r.alpha,
	}, {
		DstX:   x + width,
		DstY:   y,
		SrcX:   1,
		SrcY:   0,
		ColorR: r.red,
		ColorG: r.green,
		ColorB: r.blue,
		ColorA: r.alpha,
	}, {
		DstX:   x,
		DstY:   y + height,
		SrcX:   0,
		SrcY:   1,
		ColorR: r.red,
		ColorG: r.green,
		ColorB: r.blue,
		ColorA: r.alpha,
	}, {
		DstX:   x + width,
		DstY:   y + height,
		SrcX:   1,
		SrcY:   1,
		ColorR: r.red,
		ColorG: r.green,
		ColorB: r.blue,
		ColorA: r.alpha,
	},
	}...)
	indiceCursor := uint16(r.vertexAmount * 4)
	r.indices = append(r.indices, []uint16{
		r.boxIndices[0] + indiceCursor,
		r.boxIndices[1] + indiceCursor,
		r.boxIndices[2] + indiceCursor,
		r.boxIndices[3] + indiceCursor,
		r.boxIndices[4] + indiceCursor,
		r.boxIndices[5] + indiceCursor,
	}...)
	r.vertexAmount++
}
