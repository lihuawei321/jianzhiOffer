package main

import "fmt"

/*
在一个二维数组中（每个一维数组的长度相同），
每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	low, high := 0, row*col-1

	for low <= high {
		mid := (low + high) / 2
		currentrow := mid / col
		currentcol := mid % col
		fmt.Println(low, mid, high)
		if matrix[currentrow][currentcol]<target{
			low=mid+1
		}
		if matrix[currentrow][currentcol]>target{
			high=mid-1
		}
		if matrix[currentrow][currentcol]==target{
			return true
		}
	}

	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	row := len(matrix)
	if row <= 0 {
		return false
	}
	col := len(matrix[0])
	if col <= 0 {
		return false
	}

	i, j := 0, col - 1
	for ;i < row && j >= 0 ;{
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
func main() {
	
}
