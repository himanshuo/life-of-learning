package main
import (
    "fmt"
    "sort"
)
const (
    MaxInt32 = 1<<31 - 1
    MinInt32  = -1 << 31
    MaxInt64  = 1<<63 - 1

)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    testCases := [][]int{
        []int{},
        []int{3},
        []int{1,2},
        []int{3,5},
        []int{2,1},
        []int{5,3,6},
        []int{1, 1, 1},
        []int{1,3,6,8,9,0},
        []int{0,0,2,2,8,2},
        []int{4,6,9,4},
        []int{0,0,0,0},
        []int{1,4,7,7,6,6,5,5,5},
        []int{-1,1,0,3,-3},
        []int{-1},
        []int{MaxInt32, 3},
        []int{-20, -3916237, -357920, -3620601, 7374819, -7330761, 30, 6246457, -6461594, 266854},
        []int{-20, -3916237, -357920, -3620601, 7374819, -7330761, 30, 6246457, -6461594, 266854, -520, -470},
        []int{5, 4, 3, 2},

    }

    longCase := make([]int, 0)
    for i:=10000000; i>= -10000000; i--{
        longCase = append(longCase, i)
    }
    testCases = append(testCases, longCase)

    for _,t:=range testCases{
        fmt.Printf("%v -> %v\n", t, minAbsDiff(t))
    }




    /*var N int
    fmt.Scan(&N)
    arr := make([]int, N)
    for i:=0; i<N; i++ {
        var temp int
        fmt.Scan(&temp)
        arr[i] = temp
    }


    res := minAbsDiff(arr)

    for _,v := range res {
        fmt.Printf("%d %d ", v.left, v.right)
    }
*/





}
type Pair struct{
    left int
    right int
}

func minAbsDiff(arr []int) []Pair{
    //[]
    //[1]
    //[1 2]
    //[1 2 4 7]
    sort.IntSlice(arr).Sort()


    out := make([]Pair,len(arr))
    outIndex := 0
    minDiff := MaxInt32
    for i:=0;i<len(arr)-1;i++{
        if minDiff > absDiff(arr[i], arr[i+1]){
            minDiff = absDiff(arr[i], arr[i+1])
            outIndex = 0

        }

        if absDiff(arr[i], arr[i+1]) == minDiff{
            out[outIndex] = Pair{arr[i], arr[i+1]}
            outIndex++

        }
    }

    return out[0:outIndex]
}

func absDiff(a,b int) int{
    if a-b < 0 {
        return b-a
    } else {
        return a-b
    }
}


func insertionSort(arr []int){
    for i:=1; i<len(arr); i++ {
        cur := arr[i]
        for j:=i; j > 0; j-- {
            if cur < arr[j-1] {
                arr[j] = arr[j-1]
                arr[j-1] = cur
            } else {
                break
            }
        }
    }

    fmt.Println("sorted.")
}
