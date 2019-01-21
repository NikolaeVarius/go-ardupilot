package main

import (
	"fmt"
	"image"
	"image/color"
	// "math"

	"gocv.io/x/gocv"
	// "gocv.io/x/gocv/contrib"
)

const MinimumArea = 3000

func main() {
	deviceID := 0

	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Hand Gestures")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	fmt.Printf("Start reading device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		status := fmt.Sprintf("Testing")

		white := color.RGBA{255, 255, 255, 0}

		gocv.PutText(&img, status, image.Pt(10, 20), gocv.FontHersheyPlain, 1.2, white, 2)

		window.IMShow(img)
		if window.WaitKey(1) == 27 {
			break
		}

	}
}
