package main

import (
	"fmt"
	"github.com/elamre/debug_gui/pkg/composition"
	"github.com/elamre/debug_gui/pkg/gui"
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

const (
	WIDTH  = 1200
	HEIGHT = 1000
)

var chatBox *composition.ChatComposition
var inputController *tentsuyu.InputController
var camera *tentsuyu.Camera
var borderContainer *gui.Container

var exit bool

func main() {
	log.SetFlags(log.Flags() | log.Llongfile)
	inputController = tentsuyu.NewInputController()
	camera = tentsuyu.CreateCamera(WIDTH, HEIGHT)
	borderContainer = gui.NewBaseContainer(0, 0, WIDTH, HEIGHT)
	canvasContainer := gui.NewContainer(0, 0, 0.9, 0.9, gui.ANCHOR_HORIZONTAL_CENTER|gui.ANCHOR_VERTICAL_CENTER)
	canvasContainer.SetParent(borderContainer)
	canvasContainer.SetActive(true)

	chatBox = composition.NewChatComposition(0, 0, 0.3, 1, canvasContainer, gui.ANCHOR_RIGHT)
	userOne := gui.DebugPlayer{Name: "Elmar"}
	userOne1 := gui.DebugPlayer{Name: "Elmar2"}
	userOne2 := gui.DebugPlayer{Name: "Elmar45"}
	userOne3 := gui.DebugPlayer{Name: "Elmar1"}
	chatBox.OnUserJoined(&userOne)
	chatBox.OnUserJoined(&userOne1)
	chatBox.OnUserJoined(&userOne2)
	chatBox.OnUserJoined(&userOne3)
	chatBox.OnMessageReceived("TESTST", &userOne)
	chatBox.OnMessageReceived("sadrfsdf", &userOne1)
	chatBox.OnMessageReceived("sadrfsdfasdasdasdasdasdasd", &userOne1)
	ebiten.SetWindowSize(WIDTH, HEIGHT)

	lobby := composition.NewLobbyComposition(0, 0, 0.7, 1, canvasContainer, gui.ANCHOR_LEFT)
	lobby.CreateRoomButtonAction = func(roomName string) {
		lobby.RoomCreated(&composition.Room{Name: roomName})
	}

	var p *gui.PopupContainer
	p = composition.NewYesNoPopup("Exit Game?", canvasContainer, func(Action int) {
		log.Printf("Action: %d", Action)
		if Action == composition.NoAction {
			p.ClosePopupContainer()
		} else {
			exit = true
		}
	})

	p.SetMovable(true)
	//p.Cancellable(true)

	//canvasContainer.SetFocus(p.Container)

	if err := ebiten.RunGame(&TestGame{}); err != nil {
		log.Fatal(err)
	}
}

type TestGame struct{}

func (t *TestGame) Update() error {
	//TODO implement me
	inputController.Update()
	borderContainer.Update(inputController, nil, nil)
	//chatBox.Update(inputController, nil, nil)

	if exit {
		return fmt.Errorf("closing game")
	}
	return nil
}

func (t *TestGame) Draw(screen *ebiten.Image) {
	//chatBox.Draw(screen, camera)
	borderContainer.Draw(screen, camera)
	//TODO implement me
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Hello world: %f", ebiten.CurrentFPS()))

}

func (t *TestGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	//TODO implement me
	return outsideWidth, outsideHeight
}
