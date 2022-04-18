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

func main() {
	inputController = tentsuyu.NewInputController()
	camera = tentsuyu.CreateCamera(WIDTH, HEIGHT)
	borderContainer := gui.NewBaseContainer(0, 0, WIDTH, HEIGHT)
	canvasContainer := gui.NewContainer(0, 0, 0.9, 0.9, gui.ANCHOR_HORIZONTAL_CENTER|gui.ANCHOR_VERTICAL_CENTER)
	canvasContainer.SetParent(borderContainer)
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
	if err := ebiten.RunGame(&TestGame{}); err != nil {
		log.Fatal(err)
	}
}

type TestGame struct{}

func (t *TestGame) Update() error {
	//TODO implement me
	inputController.Update()
	chatBox.Update(inputController, nil, nil)
	return nil
}

func (t *TestGame) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Hello world: %f", ebiten.CurrentFPS()))
	chatBox.Draw(screen, camera)
	//TODO implement me
}

func (t *TestGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	//TODO implement me
	return outsideWidth, outsideHeight
}
