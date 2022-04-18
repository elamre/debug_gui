package gui

import (
	"github.com/elamre/debug_gui/pkg/common"
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
)

type BasicElement struct {
	parent *Container
	BackupElement
}

type BackupElement interface {
	CanPress() bool
	SetSelected(bool)
	Selected() bool
	// Resize(screenWidth, screenHeight float64)
	Update(input *tentsuyu.InputController, stateChanger *common.StateChanger, gameState *tentsuyu.GameState)
	Draw(screen *ebiten.Image, camera *tentsuyu.Camera)
}

type ElementInterface interface {
	SetEnabled(enabled bool)
	Update(input *tentsuyu.InputController, positionX, positionY, width, height float64, stateChanger common.StateChanger, gameState tentsuyu.GameState)
	Draw(positionX, positionY, width, height float64, screen *ebiten.Image, camera *tentsuyu.Camera)
	GetMinWidth() float64
	GetMinHeight() float64
	IsEnabled() bool
}

type Element struct {
	element ElementInterface
	// Position could be absolute or relative. Depending on flags
	positionX, positionY float64
	width, height        float64

	flag ContainerFlag
}

func NewElement(elem ElementInterface, positionX, positionY float64, flag ContainerFlag) *Element {
	return &Element{element: elem, width: elem.GetMinWidth(), height: elem.GetMinHeight(), positionY: positionY, positionX: positionX, flag: flag}
}

func (e *Element) Update(input *tentsuyu.InputController, positionX, positionY, width, height float64, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
	e.element.Update(input, positionX, positionY, width, height, stateChanger, gameState)
}

func (e *Element) Draw(positionX, positionY, width, height float64, screen *ebiten.Image, camera *tentsuyu.Camera) {
	e.element.Draw(positionX, positionY, width, height, screen, camera)
}
