package gui

//
//import (
//	"github.com/elamre/debug_gui/pkg/common"
//	"github.com/elamre/tentsuyu"
//	"github.com/hajimehoshi/ebiten/v2"
//
//	"log"
//	"sort"
//	"sync"
//)
//
//type LobbyListElement struct {
//	enabled bool
//	// TODO use different element than button
//
//	lobbies  map[string]*DebugButton
//	drawList []*DebugButton
//	selected *DebugButton
//	changed  bool
//
//	syn sync.Mutex
//}
//
//func (n *LobbyListElement) LobbyPressed(button *DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
//	n.selected = button
//}
//
//func NewLobbyListElement() *LobbyListElement {
//	return &LobbyListElement{enabled: true, lobbies: make(map[string]*DebugButton)}
//}
//
//func (l *LobbyListElement) AddLobby(lobbyName string /* TODO More information */) {
//	l.syn.Lock()
//	defer l.syn.Unlock()
//	b := NewDebugButton(lobbyName, l.LobbyPressed)
//	l.lobbies[lobbyName] = b
//	l.changed = true
//}
//
//func (l *LobbyListElement) SetEnabled(enabled bool) {
//	l.enabled = enabled
//}
//
//func (l *LobbyListElement) GetSelectedButton() *DebugButton {
//	return l.selected
//}
//
//func (l *LobbyListElement) Update(input *tentsuyu.InputController, positionX, positionY, width, height float64, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
//	l.syn.Lock()
//	if l.changed {
//		l.drawList = make([]*DebugButton, 0, len(l.lobbies))
//		for _, v := range l.lobbies {
//			l.drawList = append(l.drawList, v)
//		}
//		sort.SliceStable(l.drawList, func(i, j int) bool {
//			n := 0
//			for {
//				if num1, num2 := int(l.drawList[i].Text[n]), int(l.drawList[j].Text[n]); num1 != num2 {
//					return num1 < num2
//				}
//				n++
//				if n >= len(l.drawList[i].Text) || n >= len(l.drawList[j].Text) {
//					break
//				}
//			}
//			return false
//		})
//		log.Println(l.drawList)
//		l.changed = false
//	}
//	l.syn.Unlock()
//	extraY := float64(0)
//	for i := range l.drawList {
//		l.drawList[i].Update(input, positionX, positionY+extraY, width, l.drawList[i].GetMinHeight(), stateChanger, gameState)
//		extraY += l.drawList[i].GetMinHeight()
//	}
//}
//func (l *LobbyListElement) Draw(positionX, positionY, width, height float64, screen *ebiten.Image, camera *tentsuyu.Camera) {
//	extraY := float64(0)
//	for i := range l.drawList {
//		if l.drawList[i] == l.selected {
//			l.drawList[i].selected = true
//		}
//		l.drawList[i].Draw(positionX, positionY+extraY, width, l.drawList[i].GetMinHeight(), screen, camera)
//		extraY += l.drawList[i].GetMinHeight()
//	}
//}
//
//func (l *LobbyListElement) GetMinWidth() float64 {
//	return 64
//}
//
//func (l *LobbyListElement) GetMinHeight() float64 {
//	return 64
//}
//
//func (l *LobbyListElement) IsEnabled() bool {
//	return l.enabled
//}
//
//func (n *LobbyListElement) RemoveLobby(name string) {
//	n.syn.Lock()
//	defer n.syn.Unlock()
//	if b, ok := n.lobbies[name]; ok {
//		if b == n.selected {
//			n.selected = nil
//		}
//		delete(n.lobbies, name)
//	}
//	n.changed = true
//
//}
//
//type LobbyContainer struct {
//	Container
//	client      *client.NetInterface
//	joinButton  *DebugButton
//	createButon *DebugButton
//	lobbyName   *DebugInputElement
//	lobbyPass   *DebugInputElement
//	lobbyList   *LobbyListElement
//
//	lobbyMut          sync.Mutex
//	roomEventCallback client.TypeFunctionCallback
//}
//
//func NewLobbyContainer(cc *client.NetInterface, positionX, positionY, width, height float64, parent *Container, flag ContainerFlag) *LobbyContainer {
//	c := LobbyContainer{client: cc, lobbyList: NewLobbyListElement()}
//	c.Container = *NewContainer(positionX, positionY, width, height, flag)
//	botContainer := NewContainer(0, 0, 1, 0.1, ANCHOR_DOWN|ANCHOR_LEFT|STACK_HORIZONTAL|ALIGN_LEFT)
//	c.SetParent(parent)
//	botContainer.SetParent(&c.Container)
//
//	c.joinButton = NewDebugButton("Join", c.JoinLobbyPressed)
//	c.createButon = NewDebugButton("create", c.CreateLobbyPressed)
//	c.lobbyName = NewDebugInputElement(20)
//	botContainer.AddElement(NewElement(c.joinButton, 0, 0, 0))
//	botContainer.AddElement(NewElement(c.createButon, 0, 0, 0))
//	botContainer.AddElement(NewElement(c.lobbyName, 0, 0, 0))
//
//	c.AddElement(NewElement(c.lobbyList, 0, 0, ANCHOR_UP|EXTEND_WIDTH|EXTEND_HEIGHT))
//
//	c.roomEventCallback = client.TypeFunctionCallback{
//		Typ:          PackageType_ROOM_EVENT_MESSAGE,
//		CallBackFunc: c.roomActionHandler,
//	}
//	c.lobbyMut.Lock()
//	for k, _ := range cc.ClientLobby.Rooms {
//		c.lobbyList.AddLobby(k)
//	}
//	c.lobbyMut.Unlock()
//
//	return &c
//}
//
//func (n *LobbyContainer) CreateLobbyPressed(button *DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
//	n.client.CreateRoom(n.lobbyName.Text, nil)
//}
//func (n *LobbyContainer) JoinLobbyPressed(button *DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
//	if n.lobbyList.selected != nil {
//		n.client.JoinRoom(n.lobbyList.selected.Text, nil)
//	}
//}
//
//func (l *LobbyContainer) roomActionHandler(msg *NetMessage) {
//	l.lobbyMut.Lock()
//	defer l.lobbyMut.Unlock()
//	rm := msg.RoomEvent
//	switch rm.Action {
//	case RoomEventMessage_ROOM_CREATE:
//		l.lobbyList.AddLobby(rm.Room.Name)
//	case RoomEventMessage_ROOM_JOIN:
//	case RoomEventMessage_ROOM_UPDATE:
//	case RoomEventMessage_ROOM_DELETE:
//		l.lobbyList.RemoveLobby(rm.Room.Name)
//	}
//}
//
//func (l *LobbyContainer) Init(client *client.NetInterface) {
//	l.lobbyMut.Lock()
//	defer l.lobbyMut.Unlock()
//	client.AddAsyncListener(&l.roomEventCallback)
//}
//
//func (l *LobbyContainer) DeInit(client *client.NetInterface) {
//	client.RemoveAsyncListener(&l.roomEventCallback)
//}
