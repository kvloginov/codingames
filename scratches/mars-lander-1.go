package main

import (
	"fmt"
	"os"
)

// https://www.codingame.com/ide/puzzle/mars-lander-episode-1
func main() {
    // surfaceN: the number of points used to draw the surface of Mars.
    var surfaceN int
    fmt.Scan(&surfaceN)
	surface := NewSurfaceFromStdIn(surfaceN)

    
    for {
        // hSpeed: the horizontal speed (in m/s), can be negative.
        // vSpeed: the vertical speed (in m/s), can be negative.
        // fuel: the quantity of remaining fuel in liters.
        // rotate: the rotation angle in degrees (-90 to 90).
        // power: the thrust power (0 to 4).
        var X, Y, hSpeed, vSpeed, fuel, rotate, power int
        fmt.Scan(&X, &Y, &hSpeed, &vSpeed, &fuel, &rotate, &power)
    

		hToSurface := surface.GetHeightToSurface(point{X: X, Y: Y})
		fmt.Fprintf(os.Stderr, "height: %v \n", hToSurface)
        
		thrustPower := 1;

		if (hToSurface < 500 || vSpeed < -20 ) {
			thrustPower = 4
		}
        
        // 2 integers: rotate power. rotate is the desired rotation angle (should be 0 for level 1), power is the desired thrust power (0 to 4).
        fmt.Printf("0 %v\n", thrustPower)
    }
}


type point struct{
	X int
	Y int
}


type surface struct{
	coords []point
}

func NewSurfaceFromStdIn(pointsCount int) *surface{
	coords := make([]point, 0, pointsCount)

	for i := 0; i < pointsCount; i++ {
        // landX: X coordinate of a surface point. (0 to 6999)
        // landY: Y coordinate of a surface point. By linking all the points together in a sequential fashion, you form the surface of Mars.
        var landX, landY int
        fmt.Scan(&landX, &landY)
		coords = append(coords, point{X: landX, Y: landY})
    }

	return &surface{
		coords: coords,
	}
}

// it works a little incorrectly :)
func (s *surface) GetHeightToSurface(p point) int {
	start, end := s.getSegmentBelowPoint(p)
	surfaceHeight := getXProjectionToSegment(start, end, p)
	fmt.Fprintf(os.Stderr, "HeightOfSurfaceThere: %v\n", surfaceHeight)
	
	return p.Y - int(surfaceHeight)
}

// not log(N), but ok for < 30 points
func (s *surface) getSegmentBelowPoint(p point) (point, point) {
    prev := s.coords[0]
	for i:= 1; i<len(s.coords); i++ {
		next := s.coords[i]
		if (next.X > p.X) {
			return prev, next
		}
		prev = next
	}
	return prev, prev
}

func getXProjectionToSegment(A point, B point, p point) float32 {
	if A.X == B.X {
		return float32(A.Y)
	}

	yc := float32((p.X - A.X) * (B.Y - A.Y)) / float32(B.X - A.X) + float32(A.Y)
	return yc
}
