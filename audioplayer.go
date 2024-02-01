package main

import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var f *os.File
var format beep.Format
var streamer beep.StreamSeekCloser
var pause bool = false
var audioContent *fyne.Container

func showAudioPlayer() {

	go func(msg string) {

		fmt.Println(msg)
		if streamer == nil {
		} else {
			//slider.Value = float64(streamer.Position())
			fmt.Println(fmt.Sprint(streamer.Len()))
		}

	}("going")

	time.Sleep(time.Second)

	w := myApp.NewWindow("Audio Player")
	w.Resize(fyne.NewSize(600, 600))

	img := canvas.NewImageFromFile("audioplayer.png")
	img.FillMode = canvas.ImageFillOriginal

	//logo := canvas.NewImageFromFile("audioplayer.png")
	//logo.FillMode = canvas.ImageFillOriginal

	toolbar := widget.NewToolbar(

		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {

			// f, _ = os.Open("hen.mp3")

			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			speaker.Play(streamer)

		}),

		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {

			if !pause {
				pause = true
				speaker.Lock()
			} else if pause {
				pause = false
				speaker.Unlock()
			}

		}),

		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			speaker.Clear()

			// speaker.Close()

		}),

		widget.NewToolbarSpacer(),
	)

	label := widget.NewLabel("^_^")
	label.Alignment = fyne.TextAlignCenter

	label2 := widget.NewLabel("Now Playing : ---")
	label2.Alignment = fyne.TextAlignCenter

	browse_files := widget.NewButton("Browse...", func() {

		fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, _ error) {
			streamer, format, _ = mp3.Decode(uc)
			label2.Text = uc.URI().Name()
			label2.Refresh()
		}, w)

		fd.Show()
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))

	})

	// slider := widget.NewSlider(0, 100)

	c := container.NewVBox(label, browse_files, label2, toolbar)
	w.SetContent(container.NewVBox(
		container.NewBorder(img, nil, nil, nil, c),
		container.NewVBox(),
	),
	)
	w.Show()

}
