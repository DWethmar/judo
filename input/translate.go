package input

func WorldPosition(screenX, screenY int32, scale float64, screenWidth, screenHeight int) (x, y int32) {
	// Convert int values to float64 for calculations
	fScreenX := float64(screenX)
	fScreenY := float64(screenY)
	fScreenWidth := float64(screenWidth)
	fScreenHeight := float64(screenHeight)

	// Calculate the screen's center
	centerX, centerY := fScreenWidth/2.0, fScreenHeight/2.0

	// Convert screen position to world position
	worldX := (fScreenX-centerX)/scale + centerX
	worldY := (fScreenY-centerY)/scale + centerY

	return int32(worldX), int32(worldY)
}

func ScreenPosition(worldX, worldY int32, scale float64, screenWidth, screenHeight int32) (screenX, screenY int32) {
	// Convert int values to float64 for calculations
	fWorldX := float64(worldX)
	fWorldY := float64(worldY)
	fScreenWidth := float64(screenWidth)
	fScreenHeight := float64(screenHeight)

	// Calculate the screen's center
	centerX, centerY := fScreenWidth/2.0, fScreenHeight/2.0

	// Convert world position to screen position
	screenX = int32((fWorldX-centerX)*scale + centerX)
	screenY = int32((fWorldY-centerY)*scale + centerY)

	return screenX, screenY
}
