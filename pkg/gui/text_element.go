package gui

import (
	"github.com/elamre/debug_gui/pkg/common"
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"strings"
)

type DebugTextElement struct {
	rectHelper *common.RectangleHelper

	characters   int
	Text         []string
	backspaceWas bool
	selected     bool
	enabled      bool
	description  string
}

func NewDebugTextElement(characters int) *DebugTextElement {
	return &DebugTextElement{
		characters:   characters,
		Text:         make([]string, 0),
		backspaceWas: false,
		enabled:      true,
		rectHelper:   common.NewRectangleHelper(),
	}
}

func NewDebugTextElementWithText(text string) *DebugTextElement {
	d := &DebugTextElement{
		characters:   len(text),
		Text:         strings.Split(text, "\n"),
		backspaceWas: false,
		enabled:      true,
		rectHelper:   common.NewRectangleHelper(),
	}
	return d
}

func (b *DebugTextElement) IsEnabled() bool { return b.enabled }

func (d *DebugTextElement) SetDescription(newDescription string) *DebugTextElement {
	d.description = newDescription
	return d
}

func (d *DebugTextElement) SetEnabled(enabled bool) {
	d.enabled = enabled
}

func (d *DebugTextElement) Update(input *tentsuyu.InputController, positionX, positionY, width, height float64, stateChanger common.StateChanger, gameState tentsuyu.GameState) {

}

func (d *DebugTextElement) Draw(positionX, positionY, width, height float64, screen *ebiten.Image, camera *tentsuyu.Camera) {
	d.rectHelper.Reset()
	d.rectHelper.SetColor(0.1, 0.1, 0.1, 1)
	d.rectHelper.AddRectangle(float32(positionX), float32(positionY), float32(width), float32(height))
	d.rectHelper.Draw(screen)

	if float64(len(d.Text))*16 > height {
		d.Text = d.Text[1:]
	}
	//TODO make align an option
	for i := range d.Text {
		ebitenutil.DebugPrintAt(screen, d.Text[len(d.Text)-1-i], int(positionX), int(height+positionY-(float64(i+1)*16)))
	}
	//for i, u := range d.Text {
	//	ebitenutil.DebugPrintAt(screen, u, int(positionX), int(positionY)+((i)*16))
	//}
}

func (d *DebugTextElement) GetMinWidth() float64 {
	return float64(d.characters+len(d.description)) * PixelPerWidthCharacter
}

func (d *DebugTextElement) GetMinHeight() float64 {
	return PixelPerHeightCharacter
}
