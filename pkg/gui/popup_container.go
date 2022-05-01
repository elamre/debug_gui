package gui

import (
	"github.com/elamre/debug_gui/pkg/common"
	"github.com/elamre/tentsuyu"
)

type PopupContainer struct {
	*Container
	movable   bool
	mX, mY    float64
	canMove   bool
	canCancel bool
}

func NewPopupContainer(positionX, positionY, width, height float64) *PopupContainer {
	p := PopupContainer{Container: newEmptyContainer(POS_FLOATING)}

	p.positionX = positionX
	p.positionY = positionY
	p.pixelHeight = height
	p.pixelWidth = width
	p.updateSpecifics = p.Update
	//p.Container.Update = p.Update
	return &p
}

func (c *PopupContainer) Cancellable(canCancel bool) {
	c.canCancel = canCancel
}

func (c *PopupContainer) SetMovable(canMove bool) {
	c.canMove = canMove
}

func (c *PopupContainer) ClosePopupContainer() {
	c.parent.RemoveContainer(c.Container)
}

func (c *PopupContainer) Update(input *tentsuyu.InputController, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
	mX, mY := input.GetMouseCoords()

	if input.LeftClick().JustPressed() && c.canMove {
		if nn, mm := mX-c.positionX, mY-c.positionY; nn > 0 && nn < c.pixelWidth && mm > 0 && mm < c.pixelHeight {
			c.mX = mX - c.positionX
			c.mY = mY - c.positionY
			c.movable = true
		} else if c.canCancel {
			c.ClosePopupContainer()
			return
		}
	}
	if input.LeftClick().Up() {
		c.movable = false
	}

	if c.movable && input.LeftClick().Down() {
		c.positionX = mX - c.mX
		c.positionY = mY - c.mY
	}
}
