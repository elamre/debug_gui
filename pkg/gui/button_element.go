package gui

import (
	"github.com/elamre/debug_gui/pkg/common"
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"image/color"
)

type DebugButtonPressedAction func(button *DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState)

type DebugButton struct {
	Text           string
	Action         DebugButtonPressedAction
	selected       bool
	leftButtonDown bool
	enabled        bool
}

func NewDebugButton(text string, Action DebugButtonPressedAction) *DebugButton {
	return &DebugButton{Text: text, Action: Action, enabled: true}
}

func (b *DebugButton) SetAction(d DebugButtonPressedAction) *DebugButton {
	b.Action = d
	return b
}

func (b *DebugButton) SetEnabled(enabled bool) { b.enabled = enabled }
func (b *DebugButton) IsEnabled() bool         { return b.enabled }

func (b *DebugButton) GetMinWidth() float64 {
	return 100
}
func (b *DebugButton) GetMinHeight() float64 {
	return 16
}

func (b *DebugButton) Update(input *tentsuyu.InputController, positionX, positionY, width, height float64, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
	if !b.enabled {
		return
	}
	mX, mY := input.GetMouseCoords()
	if nn, mm := mX-positionX, mY-positionY; nn > 0 && nn < width && mm > 0 && mm < height {
		b.selected = true
	} else {
		b.selected = false
	}
	if input.LeftClick().Down() {
		if b.selected && !b.leftButtonDown {
			b.leftButtonDown = true
			if b.Action != nil {
				b.Action(b, stateChanger, gameState)
			}
		}
	} else {
		b.leftButtonDown = false
	}
}
func (b *DebugButton) Draw(positionX, positionY, width, height float64, screen *ebiten.Image, camera *tentsuyu.Camera) {
	var drawClr color.Color
	if b.selected {
		drawClr = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	} else {
		drawClr = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	}
	tentsuyu.DrawLine(screen, positionX, positionY, positionX, positionY+height, drawClr, camera)
	tentsuyu.DrawLine(screen, positionX+width, positionY, positionX+width, positionY+height, drawClr, camera)
	tentsuyu.DrawLine(screen, positionX, positionY, positionX+width, positionY, drawClr, camera)
	tentsuyu.DrawLine(screen, positionX, positionY+height, positionX+width, positionY+height, drawClr, camera)
	textLen := len(b.Text) * 6
	ebitenutil.DebugPrintAt(screen, b.Text, int(positionX+(width/2)-(float64(textLen)/2)), int(positionY))
}
