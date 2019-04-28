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
	deviceID := 2
	white := color.RGBA{255, 255, 255, 0}
	blue := color.RGBA{0, 0, 255, 0}
	webcam, err := gocv.VideoCaptureDevice(deviceID)
	var haarCascadeFile = "../data/haarcascade_frontalface_default.xml"

	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Drone Video Out")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	fmt.Printf("Start reading device: %v\n", deviceID)
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	if !classifier.Load(haarCascadeFile) {
		fmt.Printf("Error reading cascade file: %v\n", haarCascadeFile)
		return
	}

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		status := fmt.Sprintf("Scanning for Faces")

		// detect faces
		rects := classifier.DetectMultiScale(img)
		fmt.Printf("Found %d Faces\n", len(rects))

		gocv.PutText(&img, status, image.Pt(10, 20), gocv.FontHersheyPlain, 1.2, white, 2)

		// draw a rectangle around each face on the original image,
		// along with text identifying as "Human"
		for _, r := range rects {
			gocv.Rectangle(&img, r, blue, 3)

			size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
			pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
			gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
		}

		window.IMShow(img)
		if window.WaitKey(1) == 27 {
			break
		}

	}
}
