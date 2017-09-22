package main

import (
	//"fmt"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"sync"
	"time"
)

type Pixel struct {
	x, y  int
	color *RGBColor
}

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	da, _ := gtk.DrawingAreaNew()
	win.Add(da)

	world := NewWorld()

	var mux sync.Mutex
	var pixels []Pixel

	go func() {
		world.RenderScene(func(x, y int, color *RGBColor) {
			mux.Lock()
			pixels = append(pixels, Pixel{x, y, color})
			mux.Unlock()
		})
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 250)
			win.QueueDraw()
		}
	}()

	x, y := world.GetResolution()

	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {

		cr.SetSourceRGB(0, 0, 0)
		cr.Rectangle(0, 0, float64(x), float64(y))
		cr.Fill()

		mux.Lock()
		for _, pixel := range pixels {
			cr.SetSourceRGB(pixel.color.R, pixel.color.G, pixel.color.B)
			cr.Rectangle(float64(pixel.x), float64(pixel.y), 1, 1)
			cr.Fill()
		}
		mux.Unlock()
	})

	// Set the default window size.
	win.SetDefaultSize(x, y)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
