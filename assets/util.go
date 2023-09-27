package assets

import "image"

// CreateCells creates a slice of image.Rectangle pointers
// that can be used to draw sprites from a sprite sheet.
//
// the rectangles are from left to right, top to bottom
func CreateCells(columns, rows, width, height int) []image.Rectangle {
	var cells = make([]image.Rectangle, 0)
	var dimensions = image.Rect(0, 0, width, height)

	for y := 0; y < rows; y++ {
		for x := 0; x < columns; x++ {
			cell := dimensions.Add(image.Point{width * x, height * y})
			cells = append(cells, cell)
		}
	}

	return cells
}
