package main

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
	"math"
	"strconv"
)

func main() {

	//mainGrid := mat64.NewDense(300, 300, []float64{})

	gridWithPowers := GetGridWithPowers(9005)

	resultGrids := getAll3x3gridsandPowers(gridWithPowers)

	result := getMaxPowerGrid(resultGrids)

	fmt.Printf("X = %d, Y = %d, Power = %f", result.X, result.Y, result.Power)
}

type gridWithPowers struct {
	X     int
	Y     int
	Power float64
}

func getMaxPowerGrid(input []gridWithPowers) gridWithPowers {

	var maxpowerGrid gridWithPowers
	var maxPower float64 = 0.0

	for _, i := range input {
		if i.Power > maxPower {
			maxPower = i.Power
			maxpowerGrid = i
		}
	}

	return maxpowerGrid
}

func getAll3x3gridsandPowers(inputGrid mat64.Dense) []gridWithPowers {

	var result []gridWithPowers
	r, c := inputGrid.Dims()

	for i := 0; i+3 <= r; i++ {
		for j := 0; j+3 <= c; j++ {
			gridSlice := inputGrid.Slice(i, i+3, j, j+3)
			power := mat64.Sum(gridSlice)
			result = append(result, gridWithPowers{X: i, Y: j, Power: power})
		}
	}

	return result
}

func GetGridWithPowers(input float64) mat64.Dense {


	data :=  make([]float64, 90000)

	grid := mat64.NewDense(300, 300, data)

	rows, cols := grid.Dims()
	for i := 0; i < rows; i++ {
		var power float64
		for j := 0; j < cols; j++ {
			rackID := float64(i + 10)
			power = rackID * float64(j)
			power = power + input
			power = power * rackID
			resultPower := findHundreadsDigit(power)
			resultPower = resultPower - 5

			grid.Set(i, j, float64(resultPower))
		}
	}

	return *grid
}

func findHundreadsDigit(input float64) int {
	// get zero precision of float
	i := fmt.Sprintf("%.0f", input)
	j, _ := strconv.Atoi(i)
	r := j % int(math.Pow(10, float64(3)))
	return r / int(math.Pow(10, float64(3-1)))
}

//func getTotalPower(grid mat64.Dense) float64 {
//	var totalPower float64 = 0.0
//	r, c := grid.Dims()
//	for i := 0; i < r; i++ {
//		for j := 0; j < c; j++ {
//			totalPower = totalPower + grid.At(i, j)
//		}
//	}
//	return totalPower
//}
