package composition

import (
	"github.com/elamre/debug_gui/pkg/common"
	"github.com/elamre/debug_gui/pkg/gui"
	"github.com/elamre/tentsuyu"
	"log"
	"sync"
)

type SelectedRoomAction int

const (
	SelectedRoomUpdate  = SelectedRoomAction(iota)
	SelectedRoomDeleted = SelectedRoomAction(iota)
)

type CanEqual interface {
	Equal(other CanEqual) bool
}

type Room struct {
	Name         string
	Password     bool
	playerUpdate sync.Mutex
	Players      map[string]*gui.DebugPlayer
	Settings     CanEqual
}

func (r *Room) Update(other *Room) {
	r.Players = other.Players
	r.Password = other.Password
	r.Settings = other.Settings
}

func (r *Room) UserJoin(user *gui.DebugPlayer) {
	r.playerUpdate.Lock()
	defer r.playerUpdate.Unlock()
	r.Players[user.Name] = user
}

func (r *Room) UserLeft(user *gui.DebugPlayer) {
	r.playerUpdate.Lock()
	defer r.playerUpdate.Unlock()
	delete(r.Players, user.Name)
}

func (r *Room) Equal(other *Room) bool {
	if r.Password != other.Password {
		return false
	}
	if r.Password != other.Password {
		return false
	}
	if r.Name != other.Name {
		return false
	}
	if r.Settings != nil && other.Settings != nil {
		if !r.Settings.Equal(other.Settings) {
			return false
		}
	}
	if len(other.Players) != len(r.Players) {
		return false
	}
	for k, _ := range other.Players {
		if _, ok := r.Players[k]; !ok {
			return false
		}
	}
	return true
}

type LobbyComposition struct {
	rooms                  map[string]*Room
	roomsList              []*gui.DebugButton
	nameToElement          map[string]*gui.Element
	CreateRoomButtonAction func(roomName string)
	joinRoomButtonAction   func()
	selectedRoomAction     func(action SelectedRoomAction)
	selectedRoom           *Room
	// Room list
	// Room Select action
	// Create action
	// Join action
	Container         *gui.Container
	LobbyContainer    *gui.Container
	CreateLobbyButton *gui.DebugButton

	LobbyNameInput     *gui.DebugInputElement
	LobbyPasswordInput *gui.DebugInputElement
}

func NewLobbyComposition(positionX, positionY, width, height float64, parent *gui.Container, flag gui.ContainerFlag) *LobbyComposition {
	c := LobbyComposition{
		nameToElement: make(map[string]*gui.Element),
		rooms:         make(map[string]*Room),
	}
	c.Container = gui.NewContainer(positionX, positionY, width, height, flag)
	c.Container.SetParent(parent)
	InteractionContainer := gui.NewContainer(positionX, positionY, 1, 0.025, gui.ALIGN_RIGHT|gui.ANCHOR_DOWN|gui.STACK_HORIZONTAL)
	InteractionContainer.SetParent(c.Container)
	c.CreateLobbyButton = gui.NewDebugButton("Create", func(button *gui.DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
		if c.CreateRoomButtonAction != nil {
			c.CreateRoomButtonAction(c.LobbyNameInput.Text)
		}
	})
	InteractionContainer.AddElement(gui.NewElement(c.CreateLobbyButton, 0, 0, gui.ANCHOR_DOWN|gui.ANCHOR_RIGHT|gui.EXTEND_HEIGHT))
	c.LobbyNameInput = gui.NewDebugInputElement(12)
	InteractionContainer.AddElement(gui.NewElement(c.LobbyNameInput, 0, 0, gui.ANCHOR_DOWN|gui.ANCHOR_RIGHT|gui.EXTEND_HEIGHT))
	InteractionContainer.AddElement(gui.NewElement(gui.NewDebugTextElementWithText("Roomname:"), 0, 0, gui.ANCHOR_DOWN|gui.ANCHOR_RIGHT|gui.EXTEND_HEIGHT))

	c.LobbyContainer = gui.NewContainer(positionX, positionY, 1, 0.975, gui.ALIGN_UP|gui.ANCHOR_UP|gui.STACK_VERITCAL|gui.EXTEND_WIDTH)
	c.LobbyContainer.SetParent(c.Container)
	return &c
}

func (r *LobbyComposition) RoomCreated(room *Room) {
	r.rooms[room.Name] = room
	button := gui.NewDebugButton(room.Name, func(button *gui.DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
		r.selectedRoom = r.rooms[button.Text]
		log.Printf("Selected room: %s", r.selectedRoom.Name)
	})
	element := gui.NewElement(button, 0, 0, gui.EXTEND_WIDTH|gui.ANCHOR_HORIZONTAL_CENTER)
	r.LobbyContainer.AddElement(element)
	r.nameToElement[room.Name] = element
}

func (r *LobbyComposition) callSelectedIf(room *Room, action SelectedRoomAction) {
	if r.selectedRoom != nil {
		if r.selectedRoom.Equal(room) {
			if r.selectedRoomAction != nil {
				r.selectedRoomAction(action)
			}
		}
	}
}

func (r *LobbyComposition) RoomDeleted(room *Room) {
	delete(r.rooms, room.Name)
	r.callSelectedIf(room, SelectedRoomDeleted)
	r.LobbyContainer.RemoveElement(r.nameToElement[room.Name])
	if room == r.selectedRoom {
		r.selectedRoom = nil
	}
}

func (r *LobbyComposition) UpdateAllRooms(rooms []*Room) {
	for _, rr := range rooms {
		if rr == r.selectedRoom {
			if !rr.Equal(r.selectedRoom) {
				r.selectedRoom.Update(rr)
				if r.selectedRoomAction != nil {
					r.selectedRoomAction(SelectedRoomUpdate)
				}
			}
		}
		r.rooms[rr.Name] = rr
	}
}

func (r *LobbyComposition) UserJoined(room *Room, user *gui.DebugPlayer) {
	room.UserJoin(user)
	r.callSelectedIf(room, SelectedRoomUpdate)

}
func (r *LobbyComposition) UserLeft(room *Room, user *gui.DebugPlayer) {
	room.UserLeft(user)
	r.callSelectedIf(room, SelectedRoomUpdate)
}
