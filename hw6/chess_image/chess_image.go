package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	img := image.NewRGBA(image.Rect(0, 0, 400, 400))

	// Красим фон в белый.
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
	// Делаем черные квадраты 1
	for y1, y2 := 0, 50; y2 <= 400; y1, y2 = y1+100, y2+100 {
		for x1, x2 := 50, 100; x1 <= 400; x1, x2 = x1+100, x2+100 {
			img2 := image.Rect(x1, y1, x2, y2)
			draw.Draw(img, img2, &image.Uniform{black}, image.ZP, draw.Src)
		}
	}
	// Делаем черные квадраты 2
	for y1, y2 := 50, 100; y2 <= 400; y1, y2 = y1+100, y2+100 {
		for x1, x2 := 0, 50; x1 <= 400; x1, x2 = x1+100, x2+100 {
			img2 := image.Rect(x1, y1, x2, y2)
			draw.Draw(img, img2, &image.Uniform{black}, image.ZP, draw.Src)
		}
	}

	file, err := os.Create("chess_img.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()

	png.Encode(file, img)

}
