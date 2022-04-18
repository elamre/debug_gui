package gui

type ContainerFlag uint32

const (
	EXTEND_WIDTH             = ContainerFlag(0b0000_0000_0000_0001) // Extend element to screen width
	EXTEND_HEIGHT            = ContainerFlag(0b0000_0000_0000_0010) // Extend element to screen height
	FILL_CELL                = ContainerFlag(0b0000_0000_0000_0011) // Fill up a complete cell if in a setup
	INV_EXTEND               = ContainerFlag(0b1111_1111_1111_1100)
	STACK_HORIZONTAL         = ContainerFlag(0b0000_0000_0000_0100)
	STACK_VERITCAL           = ContainerFlag(0b0000_0000_0000_1000)
	DISABLE_UPDATE           = ContainerFlag(0b0000_0000_0001_0000) // Don't call update method on element
	DISABLE_DRAW             = ContainerFlag(0b0000_0000_0010_0000) // Don't call draw method on element
	POS_FLOATING             = ContainerFlag(0b0000_0000_0100_0000)
	ALIGN_LEFT               = ContainerFlag(0b0000_0001_0000_0000)
	ALIGN_RIGHT              = ContainerFlag(0b0000_0010_0000_0000)
	ALIGN_UP                 = ContainerFlag(0b0000_0100_0000_0000)
	ALIGN_DOWN               = ContainerFlag(0b0000_1000_0000_0000)
	ANCHOR_HORIZONTAL_CENTER = ContainerFlag(0b0011_0000_0000_0000)
	ANCHOR_VERTICAL_CENTER   = ContainerFlag(0b1100_0000_0000_0000)
	ANCHOR_LEFT              = ContainerFlag(0b0001_0000_0000_0000)
	ANCHOR_RIGHT             = ContainerFlag(0b0010_0000_0000_0000)
	ANCHOR_UP                = ContainerFlag(0b0100_0000_0000_0000)
	ANCHOR_DOWN              = ContainerFlag(0b1000_0000_0000_0000)
)

func (c ContainerFlag) CheckValid() {
	anchorSum := (c & ContainerFlag(0xFF000000)) >> 24
	if anchorSum == (ANCHOR_UP|ANCHOR_DOWN) || (anchorSum == (ANCHOR_RIGHT | ANCHOR_LEFT)) {
		panic("Invalid anchor")
	}
	if c&(ANCHOR_LEFT  | ANCHOR_UP | ANCHOR_DOWN| ANCHOR_RIGHT) !=0  && c&POS_FLOATING == POS_FLOATING {
		panic("Cant anchor and float")
	}
	if (c&STACK_HORIZONTAL == STACK_HORIZONTAL) && (c&ALIGN_LEFT != c&ALIGN_LEFT && c&ALIGN_RIGHT != ALIGN_RIGHT) {
		panic("Need align strategy when stacking horizontally")
	}
	if (c&STACK_VERITCAL == STACK_VERITCAL) && (c&ALIGN_UP != c&ALIGN_UP && c&ALIGN_DOWN != ALIGN_DOWN) {
		panic("Need align strategy when stacking vertically")
	}
}
