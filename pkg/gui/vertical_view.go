package gui

type VerticalView struct {
	Container
}

func NewVerticalViewWithParent(width, height float64, parent *Container, flag ContainerFlag) *VerticalView {
	n := VerticalView{
		Container: *NewContainer(parent.positionX, parent.positionY, width, height, flag),
	}
	n.Name = "            vertical"
	parent.AddChild(&n.Container)
	return &n
}

func (v *VerticalView) AddEmptyContainers(count int) *VerticalView {
	for i := 0; i < count; i++ {
		v.Container.AddChild(NewVerticalSpacer())
	}
	v.ReOrder()
	return v
}

func (v *VerticalView) GetContainer(idx int) *Container {
	return v.children[idx]
}

func (v *VerticalView) ReOrder() {
	if len(v.children) > 0 {
		percentage := float64(1) / float64(len(v.children))
		heightPerPiece := percentage * v.pixelHeight
		for i := 0; i < len(v.children); i++ {
			v.children[i].percentageWidth = 1
			v.children[i].percentageHeight = percentage
			v.children[i].originalPositionY = heightPerPiece*float64(i)
			v.children[i].originalPositionX = 0
			v.calculateContainerPositioning()
		}
	}
}
