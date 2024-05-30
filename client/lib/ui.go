package lib

import (
	"fmt"
	"log"
	"strings"

	"github.com/awesome-gocui/gocui"
	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	
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
	if err := gui.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, inputHandler(gui)); err != nil {
		log.Fatalf("Failed to set send message key combination: %v", err)
	}

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

func inputHandler(gui *gocui.Gui) func(*gocui.Gui, *gocui.View) error {
	return func(gui *gocui.Gui, v *gocui.View) error {
		input := strings.TrimSpace(v.Buffer())
		if input == "" {
			return nil
		}

		command, args := ParseInput(input)

		switch command {
		case "search":
			go func() {
				response, err := HandleSearch(args)
				if err != nil {
					log.Printf("Error performing search: %v", err)
					return
				}
				displayResponse(gui, response)
			}()
		case "review":
			go func() {
				response, err := HandleReview(args)
				if err != nil {
					log.Printf("Error fetching review: %v", err)
					return
				}
				displayResponse(gui, response)
			}()
		case "register":
			go func() {
				response, err := HandleRegister(args)
				if err != nil {
					log.Printf("Error registering user: %v", err)
				}
				displayResponse(gui, response)
			}()
		case "login":
			go func() {
				response, err := HandleLogin(args)
				if err != nil {
					log.Printf("Error logging in user: %v", err)
				}
				displayResponse(gui, response)
			}()
		case "user":
			go func() {
				response := HandleUser()
				displayResponse(gui, response)
			}()
		default:
			displayResponse(gui, "Unknown command: "+command)
		}

		v.Clear()
		v.SetCursor(0, 0)
		return nil
	}
}


func displayResponse(gui *gocui.Gui, response string) {
	gui.Update(func(g *gocui.Gui) error {
		v, err := g.View("screen")
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, response)
		return nil
	})
}
