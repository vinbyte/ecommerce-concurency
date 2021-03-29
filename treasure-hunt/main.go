package main

import (
	"fmt"
	"time"
)

var arena [][]string

func main() {
	startingPoint := []int{3, 0}
	treasurePosition := []int{}
	buildTheArea()
	fmt.Println("The arena : ")
	prettyPrint(arena, startingPoint, treasurePosition)
	fmt.Print("\n\n")
	result := findTreasure(startingPoint)
	fmt.Print("\n\n")
	fmt.Println("The possible location of treasure is : ")
	fmt.Println("the index location : ", result)
	for _, t := range result {
		prettyPrint(arena, startingPoint, t)
		fmt.Print("\n")
	}
}

func buildTheArea() [][]string {
	//we will use nested map instead of slice, to avoid the complex nested loop
	//assign the obstacle as "#"
	var blockedIndex = make(map[int]map[int]bool)
	blockedIndex[1] = make(map[int]bool)
	blockedIndex[1][1] = true
	blockedIndex[1][2] = true
	blockedIndex[1][3] = true
	blockedIndex[2] = make(map[int]bool)
	blockedIndex[2][3] = true
	blockedIndex[2][5] = true
	blockedIndex[3] = make(map[int]bool)
	blockedIndex[3][1] = true
	//assign the player index position
	// playerIndex := []int{3, 0}
	//begin loop row
	for i := 0; i < 4; i++ {
		rows := []string{}
		//begin loop column
		for j := 0; j < 6; j++ {
			//check if a row exist in blockedIndex map. if does not exist, print "."
			if blockedIndex[i] == nil {
				rows = append(rows, ".")
			} else {
				//if any obstacle in a row, check the column in the next map
				blocked := blockedIndex[i]
				if blocked[j] {
					//if any obstable in blockedIndex[i][j], print "#"
					rows = append(rows, "#")
				} else {
					// else, print ".". but we need also check the playerIndex array.
					// if playerIndex[0] == i && playerIndex[1] == j {
					// 	rows = append(rows, "X")
					// } else {

					// }
					rows = append(rows, ".")
				}
			}
		}
		arena = append(arena, rows)
	}
	return arena
}

func prettyPrint(arena [][]string, playerPosition []int, treasurePosition []int) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 8; j++ {
			if j == 0 || j == 7 {
				fmt.Print("#")
			} else {
				if i == 0 {
					fmt.Print("#")
				} else if i == 5 {
					fmt.Print("#")
				} else {
					objIndexX := i - 1
					objIndexY := j - 1
					objArena := arena[objIndexX][objIndexY]
					if len(playerPosition) > 0 {
						if objIndexX == playerPosition[0] && objIndexY == playerPosition[1] {
							fmt.Print("X")
						} else {
							//check the treasure location
							if len(treasurePosition) > 0 {
								if objIndexX == treasurePosition[0] && objIndexY == treasurePosition[1] {
									fmt.Print("$")
								} else {
									fmt.Print(objArena)
								}
							} else {
								fmt.Print(objArena)
							}
						}
					} else {
						//check the treasure location
						if len(treasurePosition) > 0 {
							if objIndexX == treasurePosition[0] && objIndexY == treasurePosition[1] {
								fmt.Print("$")
							} else {
								fmt.Print(objArena)
							}
						} else {
							fmt.Print(objArena)
						}
					}

				}
			}

			if j == 7 {
				fmt.Print("\n")
			}
		}
	}
}

func findTreasure(position []int) [][]int {
	// index := 0
	maxTopIndex := 0
	maxRightIndex := 5
	var isBlocked bool
	var newPosition []int
	var movementCounter int
	var moveUpStep = 1
	var moveRightStep = 1
	var moveDownStep = 1
	var result [][]int
	treasurePosition := []int{}
	for {
		movementCounter++
		fmt.Printf("\n Trial %d \n", movementCounter)
		prettyPrint(arena, position, treasurePosition)
		time.Sleep(time.Second * 1)
		newPosition, isBlocked = moveUp(position, moveUpStep)
		//if cannot go up, then we end
		if isBlocked {
			break
		}
		prettyPrint(arena, newPosition, treasurePosition)
		time.Sleep(time.Second * 1)
		newPosition, isBlocked = moveRight(newPosition, moveRightStep)
		if isBlocked {
			//if we cannot move right again, move up. reset moveRightStep and moveDownStep
			moveUpStep++
			moveRightStep = 1
			moveDownStep = 1
			continue
		}
		prettyPrint(arena, newPosition, treasurePosition)
		time.Sleep(time.Second * 1)
		newPosition, isBlocked = moveDown(newPosition, moveDownStep)
		if isBlocked {
			//if we cannot move down again, try move right. reset moveDownStep
			moveRightStep++
			moveDownStep = 1
			continue
		}
		prettyPrint(arena, newPosition, treasurePosition)
		time.Sleep(time.Second * 1)
		//if all clear, try move down.
		moveDownStep++
		//add the new index to result
		result = append(result, newPosition)
		if newPosition[0] == maxTopIndex && newPosition[1] == maxRightIndex {
			break
		}
	}
	return result
}

func moveUp(position []int, step int) (newPosition []int, isBlocked bool) {
	newPosition = make([]int, 2)
	newPosition[0] = position[0] - step
	newPosition[1] = position[1]
	if newPosition[0] == -1 {
		//if we reach the max top index, return isBlocked = true
		return newPosition, true
	}
	//check to arena
	obj := arena[newPosition[0]][newPosition[1]]
	if obj == "#" {
		return newPosition, true
	}
	return newPosition, false
}

func moveRight(position []int, step int) (newPosition []int, isBlocked bool) {
	newPosition = make([]int, 2)
	newPosition[0] = position[0]
	newPosition[1] = position[1] + step
	if newPosition[1] == 6 {
		//if we reach the max right index, return isBlocked = true
		return newPosition, true
	}
	//check to arena
	obj := arena[newPosition[0]][newPosition[1]]
	if obj == "#" {
		return newPosition, true
	}
	return newPosition, false
}

func moveDown(position []int, step int) (newPosition []int, isBlocked bool) {
	newPosition = make([]int, 2)
	newPosition[0] = position[0] + step
	newPosition[1] = position[1]
	if newPosition[0] == 4 {
		//if we reach the max down index, return isBlocked = true
		return newPosition, true
	}
	//check to arena
	obj := arena[newPosition[0]][newPosition[1]]
	if obj == "#" {
		return newPosition, true
	}
	return newPosition, false
}
