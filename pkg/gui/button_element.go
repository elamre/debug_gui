package gui

import (
	"github.com/elamre/debug_gui/pkg/common"
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type DebugButtonPressedAction func(button *DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState)

type DebugButton struct {
	rectHelper *common.RectangleHelper

	Text           string
	Action         DebugButtonPressedAction
	selected       bool
	leftButtonDown bool
	enabled        bool
}

func NewDebugButton(text string, Action DebugButtonPressedAction) *DebugButton {
	return &DebugButton{Text: text, Action: Action, enabled: true, rectHelper: common.NewRectangleHelper()}
}

func (b *DebugButton) SetAction(d DebugButtonPressedAction) *DebugButton {
	b.Action = d
	return b
}

func (b *DebugButton) SetEnabled(enabled bool) { b.enabled = enabled }
func (b *DebugButton) IsEnabled() bool         { return b.enabled }

func (b *DebugButton) GetMinWidth() float64 {
	return 120
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
	if width < b.GetMinWidth() {
		width = b.GetMinWidth()
	}
	b.rectHelper.Reset()
	if b.selected {
		b.rectHelper.SetColor(float32(0), float32(1), float32(0), 1)
	} else {
		b.rectHelper.SetColor(float32(1), float32(0), float32(0), 1)
	}
	b.rectHelper.AddRectangle(float32(positionX), float32(positionY), float32(width), float32(height))
	b.rectHelper.Draw(screen)

	textLen := len(b.Text) * 6
	ebitenutil.DebugPrintAt(screen, b.Text, int(positionX+(width/2)-(float64(textLen)/2)), int(positionY))
}
