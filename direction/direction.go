package direction

type Direction int

const (
	None Direction = iota
	Top
	Bottom
	Left
	Right
	TopLeft
	TopRight
	BottomLeft
	BottomRight
)

func GetDirection(sX, sY, dX, dY int32) Direction {
	deltaX := dX - sX
	deltaY := dY - sY

	// Diagonal checks
	if deltaX > 0 && deltaY > 0 {
		return BottomRight
	} else if deltaX < 0 && deltaY > 0 {
		return BottomLeft
	} else if deltaX > 0 && deltaY < 0 {
		return TopRight
	} else if deltaX < 0 && deltaY < 0 {
		return TopLeft
	}

	// Straight direction checks
	if sX == dX {
		if dY > sY {
			return Bottom
		}
		if dY < sY {
			return Top
		}
	} else if sY == dY {
		if dX > sX {
			return Right
		}
		if dX < sX {
			return Left
		}
	}

	// If we reach here, there's no clear direction or they're the same coordinates
	return None
}
