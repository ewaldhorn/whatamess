package main

import (
	"fmt"
	"math"
)

const (
	width, height int = 160, 44
)

var (
	A, B, C             float64
	cubeWidth, K1       float64 = 20.0, 40.0
	zBuffer                     = make([]float64, width*height*4)
	buffer                      = make([]rune, width*height)
	backgroundASCIIcode         = ' '
	incrementSpeed      float64 = 0.6
	distanceFromCamera          = 100
	xp, yp, idx         int
	x, y, z, ooz        float64
)

func calculateX(i, j, k float64) float64 {
	return j*math.Sin(A)*math.Sin(B)*math.Cos(C) -
		k*math.Cos(A)*math.Sin(B)*math.Cos(C) +
		j*math.Cos(A)*math.Sin(C) +
		k*math.Sin(A)*math.Sin(C) +
		i*math.Cos(B)*math.Cos(C)
}

func calculateY(i, j, k float64) float64 {
	return j*math.Cos(A)*math.Cos(C) +
		k*math.Sin(A)*math.Cos(C) -
		j*math.Sin(A)*math.Sin(B)*math.Sin(C) +
		k*math.Cos(A)*math.Sin(B)*math.Sin(C) -
		i*math.Cos(B)*math.Sin(C)
}

func calculateZ(i, j, k float64) float64 {
	return k*math.Cos(A)*math.Cos(B) -
		j*math.Sin(A)*math.Cos(B) +
		i*math.Sin(B)
}

func calculateForSurface(cubeX, cubeY, cubeZ float64, ch rune) {
	x = calculateX(cubeX, cubeY, cubeZ)
	y = calculateY(cubeX, cubeY, cubeZ)
	z = calculateZ(cubeX, cubeY, cubeZ) + float64(distanceFromCamera)

	ooz = 1 / z
	xp = (int)(float64(width)/2.0 + K1*ooz*x*2.0)
	yp = (int)(float64(height)/2.0 + K1*ooz*y)

	idx = xp + yp*width

	if idx >= 0 && idx < width*height {
		if ooz > zBuffer[idx] {
			zBuffer[idx] = ooz
			buffer[idx] = ch
		}
	}
}

func main() {
	fmt.Print("\x1b[2J")

	for {
		// init the buffer with spaces
		for tmp := range buffer {
			buffer[tmp] = backgroundASCIIcode
		}

		// now clear the zBuffer
		for tmp := range zBuffer {
			zBuffer[tmp] = 0.0
		}

		var cubeX float64
		var cubeY float64

		for cubeX = -cubeWidth; cubeX < cubeWidth; cubeX += incrementSpeed {
			for cubeY = -cubeWidth; cubeY < cubeWidth; cubeY += incrementSpeed {
				calculateForSurface(cubeX, cubeY, -cubeWidth, '.')
				calculateForSurface(cubeWidth, cubeY, cubeX, '$')
				calculateForSurface(-cubeWidth, cubeY, -cubeX, '~')
				calculateForSurface(-cubeX, cubeY, cubeWidth, '#')
				calculateForSurface(cubeX, -cubeWidth, -cubeY, ';')
				calculateForSurface(cubeX, cubeWidth, cubeY, '+')
			}
		}

		fmt.Print("\x1b[H")

		for k := 0; k < width*height; k++ {
			if k%width != 0 {
				fmt.Printf("%c", buffer[k])
			} else {
				fmt.Println()
			}
		}

		A += 0.005
		B += 0.005

	}
}
