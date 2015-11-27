// Practice Go. Practice Sorting
package main
import (
	"fmt"
)
func quicksort(arr []int, lo int, hi int) { 
	if lo < hi {
		if hi-lo < 4 {
			insertionSort(arr, lo, hi)
		} else { 
			pivot := partition(arr, lo, hi)
			quicksort(arr, lo, pivot-1)
			quicksort(arr, pivot+1, hi)
		}
	}
}

func insertionSort(arr []int, lo int, hi int){
	for i:=lo+1; i<= hi; i++ {
		j := i-1
		fmt.Println(arr, i, j)
		if arr[i] < arr[j] {
			
			for ;j>lo && arr[i]<arr[j]; j--{}
			swap(arr, i, j+1)
		}
	}

}

func partition(arr []int, lo int, hi int) int{
	pivotIndex, pivot := medianOfThree(arr, lo, hi)
	i := lo
	j := hi
	for { 
		for ; i<pivot; i++{}
		for ; j>pivot; j--{}
		if i < j {
			swap(arr, i, j)
		} else {
			swap(arr, i, pivotIndex )
			return pivotIndex
		}	
	}
}

func medianOfThree(arr []int, lo int, hi int) (int, int){
	a := arr[lo]
	b := arr[(hi-lo)/2]
	c := arr[hi]
	if (a > b && a < c) || (a<b && a>c) {
		return lo, a 
	} else if (b > a && b < c) || (b<a && b>c) {
		return (hi-lo)/2, b 
	} else {
		return hi, c
	}
}

func swap(arr []int, a int, b int) {
	fmt.Println(arr, "inside swap", a, b)
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
	fmt.Println(arr)
}


func main(){
	arrs := [][]int {
		//[]int{1,2,3},
		[]int{2,3,1},
	}
	for _, arr := range arrs {
		test(arr, quicksort)
	}
	fmt.Println("All tests passed")
}

func test(arr []int, f func([]int, int, int)) {
	if len(arr) < 2 {
		return
	}
	original := make([]int, len(arr))
	for i,val := range arr{
		original[i] = val
	}
	f(arr, 0, len(arr)-1)
	previous := arr[0]
	for _,cur := range arr{
		if previous > cur {
			panic(fmt.Sprint("Error. ORIGINAL:", original,"    SORTED:", arr ))
		}
		previous = cur
	}
	
}
