package lib

import (
	"log"

	"github.com/awesome-gocui/gocui"
)

func Init() {
	gui, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		log.Fatalf("Failed to initialize GUI: %v", err)
	}
	defer gui.Close()

	gui.SetManagerFunc(layout)

	// legge til keybinds sånn at det er mulig å quitte appen
	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, Quit); err != nil {
		log.Fatalf("Failed to set quit key combination: %v", err)
	}
	// input field
	// if err := gui.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, sendMessageHandler(ws, gui)); err != nil {
	// 	log.Fatalf("Failed to set send message key combination: %v", err)
	// }

	// main loop til app gui
	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalf("Failed to start GUI main loop: %v", err)
	}
}

func Quit(*gocui.Gui, *gocui.View) error {
	return gocui.ErrQuit
}

// selve oppsettet til hvordan GUIet ser ut
func layout(gui *gocui.Gui) error {
	maxX, maxY := gui.Size()

	if v, err := gui.SetView("screen", 0, 0, maxX-1, maxY-3, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "ITDB"
		v.Wrap = true
		v.Autoscroll = true

		// kanskje ta bort ???
		v.FgColor = gocui.ColorWhite 
		v.BgColor = gocui.ColorDefault 
		v.TitleColor = gocui.ColorYellow
		v.FrameColor = gocui.ColorBlack
	}

	if v, err := gui.SetView("input", 0, maxY-3, maxX-1, maxY-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Input"
		v.TitleColor = gocui.ColorYellow
		v.FrameColor = gocui.ColorBlack

		v.Editable = true
		v.Wrap = true
		if _, err := gui.SetCurrentView("input"); err != nil {
			return err
		}
	}

	return nil
}