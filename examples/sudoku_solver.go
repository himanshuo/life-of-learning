/*PLAN:
 * Input: 9x9 matrix (from input.txt)
 * output: solved 9x9 matrix (to output.txt)
 * 
 * Brute force, however, going to use goroutines.
 * 
 * SUDO CODE:
 * 
 * solution := make([3][3]int)
 * solved := false
 * func solve(graph [3][3]int)
 * 		numSolved := 0 
 * 		for r in graph:
 * 			for c in graph[r]:
 * 				if graph[r][c] is blank && !solved
 * 					possible = determinePossible(graph,r,c)
 * 					for p in possible:
 * 						temp_graph = copy(graph)
 * 						temp_graph[r][c] = p
 * 						go solve(temp_graph)
 * 				else:
 * 					numSolved += 1
 * 		if numSolved == 81:
 * 			solution = graph
 * 			solved = true
 * 
 * func determinePossible(graph, r, c)
 * 	possible = 0:9	
 * 	//check row,col
 * 	for i in 0:9:
 * 		val = graph[i][c]
 * 		val2 = graph[r][i]
 * 		possible.remove(val)
 * 		possible.remove(val2)
 *  
 * 	//check box
 * 	rowsToCheck = []
 *  if r < 3
 * 		rowsToCheck = [1:3]
 * 	else if r > 6
 * 		rowsToCheck = [4:6]
 * 	else
 * 		rowsToCheck = [7:9]
 *  colsToCheck = []
 *  if c < 3
 * 		colsToCheck = [1:3]
 * 	else if r > 6
 * 		colsToCheck = [4:6]
 * 	else
 * 		colsToCheck = [7:9]
 *  for row in rowsToCheck:
 * 		for col in colsToCheck:
 * 			possible.remove(graph[row][col])
 * 		
 */ 


package main
import (
	"fmt"
)

var solved bool
var solution [][]int

func solve(graph [][]int){
  	numSolved := 0 
  	for r,_ := range graph {
  		for c,_ := range graph[r] {
  			if graph[r][c] == 0 && !solved {
				fmt.Println("checking ", graph[r][c])
  				possible := determinePossible(graph,r,c)
 				for _, p := range possible {
					tempGraph := make([][]int, len(graph))
					for i,_ := range tempGraph{
						tempGraph[i] = make([]int, len(graph[i]))
					}
  					copy(graph, tempGraph)
  					fmt.Println(r,len(tempGraph), c, len(tempGraph[0]))
  					tempGraph[r][c] = p
  					go solve(tempGraph)
				}
  			} else {
  				numSolved += 1
			}
  		}
	}
  	if numSolved == 81 {
		solution = graph
  		solved = true
  	}
}






func buildList(start int, end int) []int {
	var out []int
	for i:=start; i<=end; i++{
		out = append(out, i)
	}
	return out
}

func remove(arr []int, valToRemove int) []int {
	if valToRemove == 0{
		return arr
	}
	for i,v := range arr{
		if v == valToRemove {
			
			arr = append(arr[:i], arr[i+1:]...)
			
		}
	}
	return arr
}


func determinePossible(graph [][]int, r int, c int) []int{
	possible := buildList(1,9)
	
	//check row,col
  	for i,_ := range make([]int, 9){
 		val := graph[i][c]
  		val2 := graph[r][i]

  		possible = remove(possible, val) 
  		possible = remove(possible, val2) 
   }
 	//check box
 	var rowsToCheck []int
	if r < 3 {
		rowsToCheck = buildList(0,2)
  	} else if r > 6 {
  		rowsToCheck = buildList(6,8)
  	} else {
  		rowsToCheck = buildList(3,5)
	}
   
	var colsToCheck []int
	if r < 3 {
		colsToCheck = buildList(0,2)
  	} else if r > 6 {
  		colsToCheck = buildList(6,8)
  	} else {
  		colsToCheck = buildList(3,5)
  	}
  	
	for _,row :=range rowsToCheck{
		for _,col :=range colsToCheck{
			
			possible = remove(possible, graph[row][col])
		}
	}
	return possible

}

//func indexToGrid(n int) int {
	//if n < 0 || n > 8 {
		//panic("invalid index")
	//}
	//return n + 1
//}

//func gridToIndex(n int) int {
	//if n < 1 || n > 9 {
		//panic("invalid grid")
	//}
	//return n - 1
//}



func main(){
	t1 := [][]int {
		[]int{0,0,0,0,0,7,1,0,9},
		[]int{0,0,7,0,0,1,8,6,2},
		[]int{2,0,0,3,0,0,0,5,4},
		[]int{8,0,9,6,3,0,0,0,0},
		[]int{0,0,6,0,0,0,2,0,0},
		[]int{0,0,0,0,1,9,4,0,6},
		[]int{9,7,0,0,0,8,0,0,5},
		[]int{4,8,1,9,0,0,6,0,0},
		[]int{6,0,5,1,0,0,0,0,0},
	}
	solve(t1)
	fmt.Println(solution)

}
