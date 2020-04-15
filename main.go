package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func rotate(arr [][]int, toLeft bool) [][]int {
	var res = make([][]int, len(arr))
	for i, row := range arr {
		res[i] = make([]int, len(row))
		for j := range row {
			if toLeft {
				res[i][j] = arr[j][len(arr)-i-1]
			} else {
				res[i][j] = arr[len(arr)-j-1][i]
			}
		}
	}
	return res
}

func toLeft(arr [][]int) [][]int {
	var res [][]int
	for _, row := range arr {
		var newRow []int
		newRow = rowSort(row, len(row))
		newRow = sum(newRow)
		newRow = rowSort(newRow, len(newRow))
		res = append(res, newRow)
	}
	return res

}

func reverse(matrix [][]int) [][]int {
	var res [][]int
	for _, arr := range matrix {
		for i := len(arr)/2 - 1; i >= 0; i-- {
			j := len(arr) - 1 - i
			arr[i], arr[j] = arr[j], arr[i]
		}
		res = append(res, arr)
	}
	return res
}

func rowSort(arr []int, pos int) []int {
	if pos == 1 {
		return arr
	}
	for i := 0; i < pos-1; i++ {
		if i+1 == len(arr) {
			continue
		}
		if arr[i] == 0 {
			arr[i] = arr[i+1]
			arr[i+1] = 0
		}

	}
	arr = rowSort(arr, pos-1)
	return arr
}

func sum(arr []int) []int {
	for i := range arr {
		if i+1 == len(arr) {
			continue
		}
		nextElem := arr[i+1]
		if arr[i] == nextElem {
			arr[i] = arr[i] + nextElem
			arr[i+1] = 0
		}
	}
	return arr
}

func fileDataToMatrix(fileData []byte) ([][]int, int) {
	var res [][]int
	var direction int
	for i, line := range strings.Split(string(fileData), "\n") {
		if i == 4 {
			direction, _ = strconv.Atoi(line)
			break
		}
		var arr []int
		for _, elem := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(elem)
			arr = append(arr, num)
		}

		res = append(res, arr)

	}
	return res, direction
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	matrix, direction := fileDataToMatrix(data)
	/* 	fmt.Println("input:")
	   	for _, res := range matrix {
	   		fmt.Println(res)
	   	} */
	start(matrix, direction)
}

func printResult(matrix [][]int) {
	for _, arr := range matrix {
		for _, elem := range arr {
			fmt.Print(elem, " ")
		}
		fmt.Println()
	}
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}

func start(matrix [][]int, direction int) [][]int {

	defer timeTrack(time.Now())
	var res [][]int

	switch direction {
	case 0: // left
		matrix = toLeft(matrix)
		res = matrix
		printResult(matrix)
	case 1: // up
		matrix = rotate(matrix, true)
		matrix = toLeft(matrix)
		matrix = rotate(matrix, false)
		res = matrix
		printResult(matrix)
	case 2: // right
		matrix = reverse(matrix)
		matrix = toLeft(matrix)
		matrix = reverse(matrix)
		res = matrix
		printResult(matrix)
	case 3: //down
		matrix = rotate(matrix, true)
		matrix = reverse(matrix)
		matrix = toLeft(matrix)
		matrix = reverse(matrix)
		matrix = rotate(matrix, false)
		res = matrix
		printResult(matrix)
	default:
		fmt.Println(direction, " direction is undefined")
	}
	return res
}
