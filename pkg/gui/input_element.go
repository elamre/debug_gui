package gui

import (
	"github.com/elamre/debug_gui/pkg/common"
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	rectHelper   *common.RectangleHelper
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
		rectHelper:   common.NewRectangleHelper(),
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
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || !d.enabled {
		if ebiten.IsKeyPressed(ebiten.KeyShiftLeft) || ebiten.IsKeyPressed(ebiten.KeyShiftRight) {
			d.Text += "\n"
		} else {
			d.selected = false
		}
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
	d.rectHelper.Reset()
	if d.selected {
		d.rectHelper.SetColor(0.3, 0.3, 0.3, 0.5)
	} else {
		d.rectHelper.SetColor(0.1, 0.1, 0.1, 0.5)
	}
	if len(d.description) > 0 {
		ebitenutil.DebugPrintAt(screen, d.description, int(positionX), int(positionY))
		descLen := float64(len(d.description) * PixelPerWidthCharacter)
		positionX += descLen
		width -= descLen
	}
	d.rectHelper.AddRectangle(float32(positionX), float32(positionY), float32(width), float32(height))
	d.rectHelper.Draw(screen)
	ebitenutil.DebugPrintAt(screen, d.Text, int(positionX)+2, int(positionY))
}

func (d *DebugInputElement) GetMinWidth() float64 {
	return float64(d.characters+len(d.description)) * PixelPerWidthCharacter
}

func (d *DebugInputElement) GetMinHeight() float64 {
	return PixelPerHeightCharacter
}
