package composition

import (
	"github.com/elamre/debug_gui/pkg/common"
	"github.com/elamre/debug_gui/pkg/gui"
	"github.com/elamre/tentsuyu"
)

type YesNoCancelCallbackAction int

const (
	YesAction    = 1
	NoAction     = 2
	CancelAction = 3
)

type YesNoCancelCallback func(Action int)
type LogInCallback func(userName string, cancelled bool)

func NewLogInPopup(owner *gui.Container, callback LogInCallback) *gui.PopupContainer {
	var inputName gui.DebugInputElement
	contentContainer := gui.NewContainer(0, 0, 1, 1, gui.EXTEND_WIDTH|gui.EXTEND_HEIGHT|gui.STACK_VERITCAL|gui.ALIGN_DOWN|gui.ALIGN_RIGHT)
	contentContainer.AddElement(gui.NewElement(gui.NewDebugButton("cancel", func(button *gui.DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
		callback(inputName.Text, true)
	}), 0, 0, 0))
	contentContainer.AddElement(gui.NewElement(gui.NewDebugButton("connect", func(button *gui.DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
		callback(inputName.Text, false)
	}), 0, 0, 0))
	g := gui.NewPopupContainer(240, 240, contentContainer.GetMinWidth(), contentContainer.GetMinHeight())
	g.SetParent(owner)
	g.AddChild(contentContainer)
	return g
}

func NewYesNoCancelPopup(text string, owner *gui.Container, action YesNoCancelCallback) *gui.PopupContainer {

	contentContainer := gui.NewContainer(0, 0, 1, 1, gui.EXTEND_WIDTH|gui.EXTEND_HEIGHT|gui.STACK_VERITCAL|gui.ALIGN_DOWN|gui.ALIGN_RIGHT)
	text += "\n\n"

	contentContainer.AddElement(gui.NewElement(gui.NewDebugButton("cancel", func(button *gui.DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
		action(CancelAction)
	}), 0, 0, 0))
	contentContainer.AddElement(gui.NewElement(gui.NewDebugButton("no", func(button *gui.DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
		action(NoAction)
	}), 0, 0, 0))
	contentContainer.AddElement(gui.NewElement(gui.NewDebugButton("yes", func(button *gui.DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
		action(YesAction)
	}), 0, 0, 0))
	contentContainer.AddElement(gui.NewElement(gui.NewDebugTextElementWithText(text), 0, 0, gui.EXTEND_WIDTH|gui.ANCHOR_LEFT))
	g := gui.NewPopupContainer(240, 240, contentContainer.GetMinWidth(), contentContainer.GetMinHeight())
	g.SetParent(owner)
	g.AddChild(contentContainer)

	return g
}

func NewYesNoPopup(text string, owner *gui.Container, action YesNoCancelCallback) *gui.PopupContainer {

	contentContainer := gui.NewContainer(0, 0, 1, 1, gui.EXTEND_WIDTH|gui.EXTEND_HEIGHT|gui.STACK_VERITCAL|gui.ALIGN_DOWN|gui.ALIGN_RIGHT)
	text += "\n\n"

	contentContainer.AddElement(gui.NewElement(gui.NewDebugButton("no", func(button *gui.DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
		action(NoAction)
	}), 0, 0, 0))
	contentContainer.AddElement(gui.NewElement(gui.NewDebugButton("yes", func(button *gui.DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
		action(YesAction)
	}), 0, 0, 0))
	contentContainer.AddElement(gui.NewElement(gui.NewDebugTextElementWithText(text), 0, 0, gui.EXTEND_WIDTH|gui.ANCHOR_LEFT))
	g := gui.NewPopupContainer(240, 240, contentContainer.GetMinWidth(), contentContainer.GetMinHeight())
	g.SetParent(owner)
	g.AddChild(contentContainer)

	return g
}
