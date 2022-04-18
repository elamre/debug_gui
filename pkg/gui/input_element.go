package gui

import (
	"github.com/elamre/debug_gui/pkg/common"
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"strings"
)

const PixelPerWidthCharacter = 7
const PixelPerHeightCharacter = 16

func validateInput(input string) string {
	input = strings.ToLower(input)
	retval := ""
	for c := range input {
		if (int(input[c]) >= int('a') && int(input[c]) <= int('z')) || (int(input[c]) >= int('0') && int(input[c]) <= int('9')) {
			retval += string(input[c])
		}
	}
	return retval
}

type DebugInputElement struct {
	characters   int
	Text         string
	backspaceWas bool
	selected     bool
	enabled      bool
	description  string
}

func (b *DebugInputElement) IsEnabled() bool { return b.enabled }

func NewDebugInputElement(characters int) *DebugInputElement {
	return &DebugInputElement{
		characters:   characters,
		Text:         "",
		backspaceWas: false,
		enabled:      true,
	}
}

func (d *DebugInputElement) SetDescription(newDescription string) *DebugInputElement {
	d.description = newDescription
	return d
}

func (d *DebugInputElement) SetEnabled(enabled bool) {
	d.enabled = enabled
}

func (d *DebugInputElement) Update(input *tentsuyu.InputController, positionX, positionY, width, height float64, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
	if input.LeftClick().Down() {
		mX, mY := input.GetMouseCoords()
		if nn, mm := mX-positionX, mY-positionY; nn > 0 && nn < width && mm > 0 && mm < height {
			d.selected = true
		} else {
			d.selected = false
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) || !d.enabled {
		d.selected = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) && d.enabled {
		if !d.backspaceWas {
			d.backspaceWas = true
			if len(d.Text) > 0 {
				d.Text = d.Text[:len(d.Text)-1]
			}
		}
	} else {
		d.backspaceWas = false
	}
	if d.selected {
		if len(d.Text) <= int(d.characters)+1 {
			d.Text += validateInput(string(ebiten.InputChars()))
		}
	}
}

func (d *DebugInputElement) Draw(positionX, positionY, width, height float64, screen *ebiten.Image, camera *tentsuyu.Camera) {
	var drawClr color.Color
	if d.selected {
		drawClr = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	} else {
		drawClr = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	}
	if len(d.description) > 0 {
		ebitenutil.DebugPrintAt(screen, d.description, int(positionX), int(positionY))
		descLen := float64(len(d.description) * PixelPerWidthCharacter)
		positionX += descLen
		width -= descLen
	}
	tentsuyu.DrawLine(screen, positionX, positionY, positionX, positionY+height, drawClr, camera)
	tentsuyu.DrawLine(screen, positionX+width, positionY, positionX+width, positionY+height, drawClr, camera)
	tentsuyu.DrawLine(screen, positionX, positionY, positionX+width, positionY, drawClr, camera)
	tentsuyu.DrawLine(screen, positionX, positionY+height, positionX+width, positionY+height, drawClr, camera)
	ebitenutil.DebugPrintAt(screen, d.Text, int(positionX)+2, int(positionY))
}

func (d *DebugInputElement) GetMinWidth() float64 {
	return float64(d.characters+len(d.description)) * PixelPerWidthCharacter
}

func (d *DebugInputElement) GetMinHeight() float64 {
	return PixelPerHeightCharacter
}