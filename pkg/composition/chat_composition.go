package composition

import (
	"github.com/elamre/debug_gui/pkg/common"
	"github.com/elamre/debug_gui/pkg/gui"
	"github.com/elamre/tentsuyu"
	"log"
	"sync"
)

type OnSendMessageCb func(message string) bool

type ChatComposition struct {
	// Add message
	// Enable/Disable input
	// On Send pressed
	gui.Container
	sendButton       *gui.DebugButton
	input            *gui.DebugInputElement
	textElem         *gui.DebugTextElement
	userlist         *gui.DebugUserListElement
	chatToggleButton *gui.DebugButton

	sendCallback OnSendMessageCb

	enabled             bool
	previousChatEnabled bool

	messagesMutex sync.Mutex
}

func NewChatComposition(positionX, positionY, width, height float64, parent *gui.Container, flag gui.ContainerFlag) *ChatComposition {
	c := ChatComposition{textElem: gui.NewDebugTextElement(20)}
	c.Container = *gui.NewContainer(positionX, positionY, width, height, flag)
	c.sendButton = gui.NewDebugButton("Send", c.sendMessage)
	c.input = gui.NewDebugInputElement(22)
	c.userlist = gui.NewDebugUserListElement()
	c.SetParent(parent)
	v := gui.NewVerticalViewWithParent(1, 0.1, &c.Container, gui.ANCHOR_RIGHT|gui.ANCHOR_DOWN)
	v.AddEmptyContainers(2)

	v.GetContainer(1).AddElement(gui.NewElement(c.sendButton, 0, 0, gui.ANCHOR_LEFT|gui.ANCHOR_DOWN|gui.EXTEND_WIDTH|gui.EXTEND_HEIGHT))
	v.GetContainer(0).AddElement(gui.NewElement(c.input, 0, 0, gui.ANCHOR_LEFT|gui.ANCHOR_DOWN|gui.EXTEND_WIDTH|gui.EXTEND_HEIGHT))

	chatContainer := gui.NewContainer(0, 0, 1, 0.9, gui.ANCHOR_LEFT|gui.ANCHOR_UP)
	c.Container.AddChild(chatContainer)
	chatContainer.AddElement(gui.NewElement(c.userlist, 0, 0, gui.ANCHOR_LEFT|gui.ANCHOR_DOWN|gui.EXTEND_WIDTH|gui.EXTEND_HEIGHT))
	chatContainer.AddElement(gui.NewElement(c.textElem, 0, 0, gui.ANCHOR_LEFT|gui.ANCHOR_DOWN|gui.EXTEND_HEIGHT|gui.EXTEND_WIDTH))

	c.chatToggleButton = gui.NewDebugButton("Chat", c.toggleUsers)
	c.Container.AddElement(gui.NewElement(c.chatToggleButton, 0, 0, gui.ANCHOR_UP|gui.ANCHOR_RIGHT))
	c.SetEnabled(true)
	return &c
}

func (c *ChatComposition) SetEnabled(enabled bool) {
	log.Printf("%v + %v", enabled, c.enabled)
	if !enabled && c.enabled {
		c.previousChatEnabled = c.textElem.IsEnabled()

		c.sendButton.SetEnabled(false)
		c.input.SetEnabled(false)
		c.chatToggleButton.SetEnabled(false)
	} else if enabled && !c.enabled {
		c.userlist.SetEnabled(!c.previousChatEnabled)
		c.textElem.SetEnabled(c.previousChatEnabled)

		c.sendButton.SetEnabled(true)
		c.input.SetEnabled(true)
		c.chatToggleButton.SetEnabled(true)
	}
	c.enabled = enabled
}

func (c *ChatComposition) toggleUsers(button *gui.DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
	if c.userlist.IsEnabled() {
		button.Text = "Users "
	} else {
		button.Text = "Chat"
	}
	c.userlist.SetEnabled(!c.userlist.IsEnabled())
	c.textElem.SetEnabled(!c.textElem.IsEnabled())
}

func (c *ChatComposition) OnUserJoined(player *gui.DebugPlayer) {
	c.userlist.UserJoined(player)
}

func (c *ChatComposition) OnUserLeft(player *gui.DebugPlayer) {
	c.userlist.UserLeft(player)
}

func (c *ChatComposition) OnMessageReceived(message string, player *gui.DebugPlayer) {
	c.messagesMutex.Lock()
	defer c.messagesMutex.Unlock()
	c.textElem.Text = append(c.textElem.Text, player.Name+":"+message)
}

func (c *ChatComposition) sendMessage(button *gui.DebugButton, stateChanger common.StateChanger, gameState tentsuyu.GameState) {
	log.Println("Sending!")
	if c.sendCallback != nil {
		if c.sendCallback(c.input.Text) {
			c.input.Text = ""
		} else {
			// TODO
		}
	} else {
		// TODO
	}
}

func (c *ChatComposition) Resize(width, height int) {}
