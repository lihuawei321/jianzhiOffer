package main

//输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
//使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后半部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变

func reOrderArray(array []int) {
	if array == nil {
		return
	}
	i, j := 0, len(array)-1
	for i < j {
		for i < j && array[i]%2 == 1 {
			i++
		}
		for i < j && array[j]%2 == 0 {
			j--
		}
		array[i], array[j] = array[j], array[i]
	}
	return
}

func reOrderArray1(arr []int) {
	if arr == nil {
		return
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j]%2 == 0 && arr[j+1]%2 == 1 {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}
func main() {

}
