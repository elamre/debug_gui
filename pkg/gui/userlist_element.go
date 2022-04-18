package gui

import (
	"github.com/elamre/debug_gui/pkg/common"
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"strings"
	"sync"
)

type DebugPlayer struct {
	Name      string
	OtherData interface{}
}

type DebugUserListElement struct {
	enabled    bool
	users      []*DebugPlayer
	ownName    string
	usersMutex sync.Mutex
}

func NewDebugUserListElement() *DebugUserListElement {
	d := &DebugUserListElement{
		enabled:    true,
		users:      make([]*DebugPlayer, 0),
		usersMutex: sync.Mutex{},
	}

	return d
}

func (d *DebugUserListElement) SetEnabled(enabled bool) {
	d.enabled = enabled
}

func (n *DebugUserListElement) UserJoined(netPlayer *DebugPlayer) {
	n.usersMutex.Lock()
	defer n.usersMutex.Unlock()
	n.users = append(n.users, netPlayer)
}

func (n *DebugUserListElement) UserLeft(netPlayer *DebugPlayer) {
	n.usersMutex.Lock()
	defer n.usersMutex.Unlock()
	i := 0
	for range n.users {
		if strings.Compare(netPlayer.Name, n.users[i].Name) == 0 {
			if len(n.users) == 1 {
				n.users = []*DebugPlayer{}
			} else if i == 0 {
				n.users = n.users[1:]
			} else if i == len(n.users)-1 {
				n.users = n.users[:len(n.users)-1]
			} else {
				n.users = append(n.users[0:i-1], n.users[:i+1]...)
			}
			break
		}
		i++
	}
}

func (d *DebugUserListElement) Update(input *tentsuyu.InputController, positionX, positionY, width, height float64, stateChanger common.StateChanger, gameState tentsuyu.GameState) {

}

func (b *DebugUserListElement) IsEnabled() bool { return b.enabled }

func (d *DebugUserListElement) Draw(positionX, positionY, width, height float64, screen *ebiten.Image, camera *tentsuyu.Camera) {
	drawClr := color.RGBA{R: 30, G: 30, B: 30, A: 255}
	ebitenutil.DrawRect(screen, positionX, positionY, width, height, drawClr)
	ebitenutil.DebugPrintAt(screen, d.ownName, int(positionX), int(positionY))
	for i, u := range d.users {
		ebitenutil.DebugPrintAt(screen, u.Name, int(positionX), int(positionY)+((i+1)*16))
	}
}

func (d *DebugUserListElement) GetMinWidth() float64 {
	return float64(10) * PixelPerWidthCharacter
}

func (d *DebugUserListElement) GetMinHeight() float64 {
	return PixelPerHeightCharacter
}
