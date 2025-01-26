package main

import (
	"fmt"
	"math"
	"os"
)

type Unit struct {
	unitId, faction                                                 int
	unitType                                                        string
	health, positionX, positionY, velocityX, velocityY, gunCooldown float64
}

func findEnemyUnit(units []Unit) (float64, float64) {
	x := -1.0
	y := -1.0

	for i := 0; i < len(units); i++ {
		if x == -1.0 {
			x = units[i].positionX
			y = units[i].positionY

			// fmt.Fprintf(os.Stderr, "Enemy ship at %f %f\n", x, y)

			break
		}
	}

	return x, y
}

func main() {

	// fmt.Fprintf(os.Stderr, "Top left at %d,%d and bottom right at %d,%d\n", minX, minY, maxX, maxY)

	missileCache := 8
	counter := 0
	lastLaunch := 0

	accx := 5.0
	accy := 10.0

	for {
		counter += 1
		var unitCount int
		fmt.Scan(&unitCount)

		units := make([]Unit, unitCount)
		myUnits := make([]Unit, 0)
		myMissiles := make([]Unit, 0)
		enemyUnits := make([]Unit, 0)

		for i := 0; i < unitCount; i++ {
			// unitId: unit's unique ID
			// faction: 1 if the unit belongs to the player, -1 if to opponent
			// unitType: 'S' for ship, 'B' for bullet, 'M' for missile
			// health: remaining unit's health points
			// positionX: X coordinate of the unit's position
			// positionY: Y coordinate of the unit's position
			// velocityX: X coordinate of the unit's velocity
			// velocityY: Y coordinate of the unit's velocity
			// gunCooldown: number of rounds till the next bullet can be fired if this is a ship, -1 otherwise
			var unitId, faction int
			var unitType string
			var health, positionX, positionY, velocityX, velocityY, gunCooldown float64
			fmt.Scan(&unitId, &faction, &unitType, &health, &positionX, &positionY, &velocityX, &velocityY, &gunCooldown)
			newUnit := Unit{unitId: unitId, faction: faction, unitType: unitType, health: health, positionX: positionX, positionY: positionY, velocityX: velocityX, velocityY: velocityY, gunCooldown: gunCooldown}
			units = append(units, newUnit)

			if faction == 1 && unitType == "S" {
				myUnits = append(myUnits, newUnit)
			} else if faction == -1 && unitType == "S" {
				enemyUnits = append(enemyUnits, newUnit)
			} else if faction == 1 && unitType == "M" {
				myMissiles = append(myMissiles, newUnit)
			}
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")

		// One line for each of the (actively) controlled unit with at least one action specified
		// unitId [| A x y] [| F x y] [| M x y] [| W]

		x, y := findEnemyUnit(enemyUnits)

		if x > 700 {
			accx = -5.0
		} else {
			accx = 5.0
		}

		if y > 500 {
			accy = -10.0
		} else {
			accy = 10.0
		}

		for i := 0; i < len(myUnits); i++ {
			// fmt.Fprintf(os.Stderr, "Enemy at %f %f\n", x, y)
			// fmt.Fprintf(os.Stderr, "I'm doing %f %f\n", myUnits[i].velocityX, myUnits[i].velocityY)

			if x > -1 {
				mx := myUnits[i].positionX
				my := myUnits[i].positionY

				dist := math.Sqrt(math.Pow(mx-x, 2.0) + math.Pow(my-y, 2))

				aimy := 0.0
				aimx := 0.0
				// fmt.Fprintf(os.Stderr, "Distance to enemy is %f\n", dist)

				if dist < 1500 || (missileCache == 0 && dist < 800) {

					if mx > x {
						aimx = x - mx + (15.0 * myUnits[i].velocityX)
					} else if mx < x {
						aimx = -1.0*(mx-x) - (15.0 * myUnits[i].velocityX)
					}

					aimx = aimx / 3.0

					if my > y {
						aimy = -1.0*(my-y) + (15.0 * myUnits[i].velocityY)
					} else if my < y {
						aimy = y - my - (15.0 * myUnits[i].velocityY)
					}

					fmt.Printf("%d ", myUnits[i].unitId)

					// start moving
					if myUnits[i].positionX > 1500 {
						accx = -700 - -1.0 * math.Abs(myUnits[i].velocityX)
					} else if myUnits[i].positionX < 350 {
						accx = 700 + math.Abs(myUnits[i].velocityX)
					}

					if myUnits[i].positionY > 750 {
						accy = -700 - -1.0 * math.Abs(myUnits[i].velocityY)
					} else if myUnits[i].positionY < 250 {
						accy = 700 + math.Abs(myUnits[i].velocityY)
					}

					fmt.Printf("| A %f %f", accx, accy)

					// end moving

					if missileCache == 0 && (counter - lastLaunch > 10){
						fmt.Printf(" | F %f %f", aimx, aimy)
					}

					if missileCache > 0 && counter%1 == 0 {
						fmt.Fprintln(os.Stderr, "Missile to ", aimx, aimy)
						fmt.Printf(" | M %f %f\n", aimx, aimy)
						missileCache -= 1
						lastLaunch = counter
					} else {
						fmt.Printf(" | W\n")
					}
				} else {
					// we need to move
					mvx := 0
					mvy := 0

					if mx > x {
						mvx = 1
					} else if mx < x {
						mvx = -1
					}

					if my > y {
						mvy = -2
					} else if my < y {
						mvy = 2
					}

					if myUnits[i].velocityX == 0 || myUnits[i].velocityY == 0 {
						fmt.Printf("%d | A %d %d\n", myUnits[i].unitId, mvx, mvy)
					} else {
						fmt.Printf("S | W\n")
					}
				}

			} else {
				fmt.Printf("S | W\n")
			}
		}

		for i := 0; i < len(myMissiles); i++ {
			mx := myMissiles[i].positionX
			my := myMissiles[i].positionY

			dist := math.Sqrt(math.Pow(mx-x, 2.0) + math.Pow(my-y, 2))

			fmt.Fprintf(os.Stderr, "Enemy at %f %f\n", x, y)
			fmt.Fprintf(os.Stderr, "I'm %d at %f %f %f\n", myMissiles[i].unitId, mx, my, dist)

			var xDiff float64
			var yDiff float64

			if mx > x {
				xDiff = 1.5 * (x - mx)
			} else if mx < x {
				xDiff = -1.5 * (mx - x)
			}

			if my > y {
				yDiff = -1.5 * (my - y)
			} else if my < y {
				yDiff = 1.5 * (y - my)
			}

			fmt.Fprintf(os.Stderr, "XDiff %f  YDiff %f\n\n", xDiff, yDiff)

			if dist < 190 {
				fmt.Printf("%d | D\n", myMissiles[i].unitId)
			} else {
				if xDiff != 0 || yDiff != 0 {
					fmt.Printf("%d | A %f %f\n", myMissiles[i].unitId, xDiff, yDiff)
				} else {
					fmt.Printf("%d | W\n", myMissiles[i].unitId)
				}
			}
		}
	}
}
