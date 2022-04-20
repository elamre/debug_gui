package gui

import (
	"fmt"
	"github.com/elamre/debug_gui/pkg/common"
	"github.com/elamre/go_helpers/pkg/slice_helpers"
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"math/rand"
)

type PositionSizeParameters struct {
	positionX, positionY float64
	width, height        float64
}

type Container struct {
	Flag                                 ContainerFlag
	parent                               *Container
	children                             []*Container
	childrenToPosition                   map[*Container]PositionSizeParameters
	elements                             []*Element
	elementToPosition                    map[*Element]PositionSizeParameters
	originalPositionX, originalPositionY float64
	positionX, positionY                 float64
	pixelWidth, pixelHeight              float64
	percentageWidth, percentageHeight    float64
	r, g, b                              float64
	Name                                 string
	rectHelper                           *common.RectangleHelper
	DrawRect                             bool
	active                               bool
}

func (c *Container) SetActive(active bool) {
	c.DrawRect = active
	c.active = active
}

func (c Container) String() string {
	return fmt.Sprintf("x,y[%f, %f] orig[%f, %f] width,height perc[%f, %f] pixel[%f, %f]", c.positionX, c.positionY, c.originalPositionX, c.originalPositionY, c.percentageWidth, c.percentageHeight, c.pixelWidth, c.pixelHeight)
}

func newEmptyContainer(flag ContainerFlag) *Container {
	flag.CheckValid()
	return &Container{
		children:           make([]*Container, 0),
		elements:           make([]*Element, 0),
		childrenToPosition: make(map[*Container]PositionSizeParameters),
		elementToPosition:  make(map[*Element]PositionSizeParameters),
		r:                  rand.Float64(),
		b:                  rand.Float64(),
		g:                  rand.Float64(),
		Flag:               flag,
		rectHelper:         common.NewRectangleHelper(),
		DrawRect:           true,
		active:             true,
	}
}

func NewHorizontalSpacer() *Container {
	return NewSpacerContainer(1, 0, EXTEND_HEIGHT)
}

func NewVerticalSpacer() *Container {
	return NewSpacerContainer(0, 1, EXTEND_WIDTH)

}

func NewSpacerContainer(percentageWidth, percentageHeight float64, flag ContainerFlag) *Container {
	c := newEmptyContainer(flag)
	c.percentageHeight = percentageHeight
	c.percentageWidth = percentageWidth
	c.Name = fmt.Sprintf("Spacer:%d", flag)
	return c
}

func NewBaseContainer(positionX, positionY, width, height float64) *Container {
	c := newEmptyContainer(EXTEND_HEIGHT | EXTEND_WIDTH | ANCHOR_UP | ANCHOR_LEFT)
	c.originalPositionY = positionY
	c.originalPositionX = positionX
	c.pixelHeight = height
	c.pixelWidth = width
	return c
}

func NewContainer(positionX, positionY, percentageWidth, percentageHeight float64, flag ContainerFlag) *Container {
	c := newEmptyContainer(flag)
	c.originalPositionY = positionY
	c.originalPositionX = positionX
	c.percentageWidth = percentageWidth
	c.percentageHeight = percentageHeight
	return c
}

func (c *Container) SetParent(parent *Container) *Container {
	if parent == nil {
		panic("Parent == nil")
	}
	if parent == c {
		panic("Parent == self")
	}
	parent.AddChild(c)
	return c
}

func (c *Container) Draw(screen *ebiten.Image, camera *tentsuyu.Camera) {
	if !c.active {
		return
	}
	if c.DrawRect {
		c.rectHelper.Reset()
		c.rectHelper.SetColor(float32(c.r), float32(c.g), float32(c.b), 0.3)
		c.rectHelper.AddFilledRectangle(float32(c.positionX), float32(c.positionY), float32(c.pixelWidth), float32(c.pixelHeight))
		c.rectHelper.SetColor(float32(c.r), float32(c.g), float32(c.b), 1)
		c.rectHelper.AddRectangle(float32(c.positionX), float32(c.positionY), float32(c.pixelWidth), float32(c.pixelHeight))
		c.rectHelper.Draw(screen)
	}

	for i := range c.children {
		c.children[i].Draw(screen, camera)
	}

	for i := range c.elements {
		if c.elements[i].element.IsEnabled() {
			coordinates := c.elementToPosition[c.elements[i]]
			if coordinates.positionY+coordinates.height > (c.pixelHeight + c.positionY) {
				log.Printf("%f + %f = %f > %f", coordinates.positionY, coordinates.height, coordinates.positionY+coordinates.height, c.pixelHeight+c.positionY)
			} else {
				c.elements[i].Draw(coordinates.positionX, coordinates.positionY, coordinates.width, coordinates.height, screen, camera)
			}
		}
	}
}

func (c *Container) calculateContainerPositioning() {
	var previousPositionX = c.positionX
	var previousPositionY = c.positionY
	for i := range c.children {
		positionX := c.positionX + c.children[i].originalPositionX
		positionY := c.positionY + c.children[i].originalPositionY

		width := c.children[i].percentageWidth * c.pixelWidth
		height := c.children[i].percentageHeight * c.pixelHeight
		flag := c.children[i].Flag

		if c.Flag&STACK_HORIZONTAL != STACK_HORIZONTAL {
			if flag&EXTEND_WIDTH == EXTEND_WIDTH {
				width = c.pixelWidth
			}
			if flag&ANCHOR_HORIZONTAL_CENTER == ANCHOR_HORIZONTAL_CENTER {
				positionX = (c.pixelWidth / 2) + c.positionX - (width / 2)
			} else if flag&ANCHOR_LEFT == ANCHOR_LEFT {
				positionX = c.positionX
			} else if flag&ANCHOR_RIGHT == ANCHOR_RIGHT {
				positionX = c.positionX + c.pixelWidth - width
			}
		} else {
			if c.Flag&ALIGN_LEFT == ALIGN_LEFT {
				positionX = previousPositionX
				previousPositionX += width
			} else {
				positionX = previousPositionX - width
				previousPositionX = positionX
			}
		}

		if c.Flag&STACK_VERITCAL != STACK_VERITCAL {
			if flag&EXTEND_HEIGHT == EXTEND_HEIGHT {
				width = c.pixelHeight
			}
			if flag&ANCHOR_VERTICAL_CENTER == ANCHOR_VERTICAL_CENTER {
				positionY = (c.pixelHeight / 2) + c.positionY - (height / 2)
			} else if flag&ANCHOR_UP == ANCHOR_UP {
				positionY = c.positionY
			} else if flag&ANCHOR_DOWN == ANCHOR_DOWN {
				positionY = c.positionY + c.pixelHeight - height
			}
		} else {
			if c.Flag&ALIGN_UP == ALIGN_UP {
				positionY = previousPositionY
				previousPositionY += height
			} else {
				positionY = previousPositionY - height
				previousPositionY = positionY
			}
		}

		c.children[i].positionX = positionX
		c.children[i].positionY = positionY
		c.children[i].pixelWidth = width
		c.children[i].pixelHeight = height

		c.childrenToPosition[c.children[i]] = PositionSizeParameters{
			positionX: positionX,
			positionY: positionY,
			width:     width,
			height:    height,
		}
	}
}

func (c *Container) calculateElementPositioning() {
	var previousPositionX float64
	var previousPositionY float64

	if c.Flag&ALIGN_LEFT == ALIGN_LEFT {
		previousPositionX = c.positionX
	} else if c.Flag&ALIGN_RIGHT == ALIGN_RIGHT {
		previousPositionX = c.positionX + c.pixelWidth
	}

	if c.Flag&ALIGN_UP == ALIGN_UP {
		previousPositionY = c.positionY
	} else if c.Flag&ALIGN_DOWN == ALIGN_DOWN {
		previousPositionY = c.positionY + c.pixelHeight
	}

	for i := range c.elements {
		positionX := c.positionX + c.elements[i].positionX
		positionY := c.positionY + c.elements[i].positionY
		width := c.elements[i].width
		height := c.elements[i].height
		flag := c.elements[i].flag

		if c.Flag&STACK_HORIZONTAL != STACK_HORIZONTAL {
			if flag&ANCHOR_LEFT == ANCHOR_LEFT {
				positionX = c.positionX
				if flag&EXTEND_WIDTH == EXTEND_WIDTH {
					width = c.pixelWidth
				}
			} else if flag&ANCHOR_RIGHT == ANCHOR_RIGHT {
				positionX = c.positionX + c.pixelWidth - width
				if flag&EXTEND_WIDTH == EXTEND_WIDTH {
					positionX = c.positionX
					width = c.pixelWidth
				}
			} else {
				if flag&EXTEND_WIDTH == EXTEND_WIDTH {
					width = c.pixelWidth - positionX
				} else {
					//Centering
					positionX = (c.pixelWidth / 2) + c.positionX - (width / 2)
				}
			}
		} else {
			if c.Flag&ALIGN_LEFT == ALIGN_LEFT {
				positionX = previousPositionX
				previousPositionX += width
			} else {
				positionX = previousPositionX - width
				previousPositionX = positionX
			}
			// Stacking
		}

		if c.Flag&STACK_VERITCAL != STACK_VERITCAL {
			if flag&ANCHOR_UP == ANCHOR_UP {
				positionY = c.positionY
				if flag&EXTEND_HEIGHT == EXTEND_HEIGHT {
					height = c.pixelWidth
				}
			} else if flag&ANCHOR_DOWN == ANCHOR_DOWN {
				positionY = c.positionY + c.pixelHeight - height
				if flag&EXTEND_HEIGHT == EXTEND_HEIGHT {
					positionY = c.positionY
					height = c.pixelHeight
				}
			} else {
				if flag&EXTEND_HEIGHT == EXTEND_HEIGHT {
					height = c.pixelHeight - positionY
				} else {
					//Centering
					positionY = (c.pixelHeight / 2) + c.positionY - (height / 2)
				}
			}
		} else {
			if c.Flag&ALIGN_UP == ALIGN_UP {
				positionY = previousPositionY
				previousPositionY += height
			} else {
				positionY = previousPositionY - height
				previousPositionY = positionY
			}
		}
		c.elementToPosition[c.elements[i]] = PositionSizeParameters{
			positionX: positionX,
			positionY: positionY,
			width:     width,
			height:    height,
		}
	}
}

func (c *Container) Update(input *tentsuyu.InputController, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
	if !c.active {
		return
	}
	for i := range c.children {
		c.children[i].Update(input, stateChanger, gameState)
	}
	for i := range c.elements {
		coordinates := c.elementToPosition[c.elements[i]]
		c.elements[i].Update(input, coordinates.positionX, coordinates.positionY, coordinates.width, coordinates.height, stateChanger, gameState)
	}
}

func (c *Container) AddElement(element *Element) {
	c.elements = append(c.elements, element)
	c.calculateElementPositioning()
}

func (c *Container) RemoveElement(element *Element) {
	c.elements = slice_helpers.RemoveFromList(element, c.elements)
	c.calculateElementPositioning()
}

func (c *Container) AddChild(child *Container) {
	child.parent = c
	c.children = append(c.children, child)
	c.calculateContainerPositioning()
}
