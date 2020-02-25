package main

func minTimeToVisitAllPoints(points [][]int) int {
	moves := 0
	for i := 0; i < len(points)-1; i++ {
		currentPoint := points[i]
		nextPoint := points[i+1]
		moves += calculateMoves(currentPoint, nextPoint)
	}
	return moves
}

func calculateMoves(currentPoint []int, nextPoint []int) int {
	dx := abs(nextPoint[0] - currentPoint[0])
	dy := abs(nextPoint[1] - currentPoint[1])
	if dx < dy {
		return dx + (dy - dx)
	} else if dx > dy {
		return dy + (dx - dy)
	} else {
		return dx
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
