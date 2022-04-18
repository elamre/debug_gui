package gui

import (
	"github.com/elamre/debug_gui/pkg/common"
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type HorizontalView struct {
	Container
	view []*Container
}

func NewHorizontalViewWithParent(width, height float64, parent *Container, flag ContainerFlag) *HorizontalView {
	n := HorizontalView{
		view: make([]*Container, 0),
	}
	n.Container = *NewContainer(parent.positionX, parent.positionY, width, height, flag)
	parent.AddChild(&n.Container)
	return &n
}

func NewBaseHorizontalView(width, height, positionX, positionY float64, flag ContainerFlag) *HorizontalView {
	n := HorizontalView{
		view: make([]*Container, 0),
	}
	n.Container = *NewBaseContainer(positionX, positionY, width, height)
	return &n
}

func (v *HorizontalView) AddEmptyContainers(count int) *HorizontalView {
	for i := 0; i < count; i++ {
		v.AddContainer(NewHorizontalSpacer())
	}
	return v
}

func (v *HorizontalView) GetContainer(idx int) *Container {
	return v.view[idx]
}

func (v *HorizontalView) AddContainer(container *Container) *HorizontalView {
	container.Flag = EXTEND_HEIGHT
	v.AddChild(container)
	v.view = append(v.view, container)
	v.ReOrder()
	return v
}

func (v *HorizontalView) ReOrder() {
	if len(v.view) > 0 {
		widthPerPiece := float64(1) / float64(len(v.view)) * v.pixelWidth
		for i := 0; i < len(v.view); i++ {
			v.view[i].percentageWidth = float64(1) / float64(len(v.view))
			v.view[i].percentageHeight = 1
			// TODO add extend width to the last one
			v.view[i].positionX = v.positionX + (widthPerPiece * float64(i))
			v.view[i].calculateContainerPositioning()
		}
	}
}

func (v *HorizontalView) Draw(screen *ebiten.Image, camera *tentsuyu.Camera) {
	log.Println("DRAW")
	v.Container.Draw(screen, camera)
	for i := range v.view {
		if v.view[i] != nil {
			v.view[i].Draw(screen, camera)
		}
	}
}

func (v *HorizontalView) Update(input *tentsuyu.InputController, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
	log.Println("UPDATE")
	v.Container.Update(input, stateChanger, gameState)
	for i := range v.view {
		if v.view[i] != nil {
			v.view[i].Update(input, stateChanger, gameState)
		}
	}
}
