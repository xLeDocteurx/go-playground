package main

import (
	"fmt"
	"time"
	// "image"
	"image/color"
	// "os"

	// "gocv.io/x/gocv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)


const MinimumArea = 3000
// Settings
const windowWidth = 600
const windowHeight = 50

var fyneApp = app.New()
var fyneWindow = fyneApp.NewWindow("Up In Smoke")
var fyneWindowCanvas = fyneWindow.Canvas()
var fyneWindowContainer = container.New(layout.NewGridLayout(2))

func setUp() {


	rect := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 180, A: 255})
	// fyneWindowCanvas.SetContent(rect)

	clock := Clock{"", widget.NewLabel("")}
	clock.start()
	clock.update()

	button := widget.NewButton("Test alert", func() {
		alert("Alert", "Ceci est une alerte")
	})
	// fyneWindowCanvas.SetContent(button)
	fyneWindowContainer = container.New(layout.NewGridLayout(1), rect, clock.component, button);
}

func loop() {
	
}

func main() {
	fmt.Println("Starting...")

	setUp()

	fyneWindow.SetContent(fyneWindowContainer)
	fyneWindow.Resize(fyne.NewSize(windowWidth, windowHeight))
	fyneWindow.SetMaster()

	// Main Loop
	go func() {
		// fmt.Println("And the bass keeps running running")
		for {
			loop()
			// fmt.Println("And Running Running")
		}
	}()

	fyneWindow.Show()
	fyneApp.Run()
	// Executed after app exit
	fmt.Println("This app is Exiting...")
}

// components
type Clock struct {
	time string
	component *widget.Label
}

func (clock Clock) start() {
	clock.time = "Time: 00:00:00"
	clock.component = widget.NewLabel("")
}

func (clock Clock) update() {
	go func() {
		for range time.Tick(time.Second) {
			clock.time = time.Now().Format("Time: 03:04:05")
			clock.component.SetText(clock.time)
		}
	}()
}


// utils
func alert(errorCode string, message string) {
	fullMessage := errorCode + " : " + message
	alertWindow := fyneApp.NewWindow(errorCode)
	// alertWindowCanvas := alertWindow.Canvas()
	alertWindow.SetContent(widget.NewLabel(fullMessage))
	alertWindow.Show()
}

func updateTime(clock *widget.Label) {
	for range time.Tick(time.Second) {
		clock.SetText(time.Now().Format("Time: 03:04:05"))
	}
}
