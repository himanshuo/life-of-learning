package main

import "fmt"

func main() {
	board := make([][]int, 0)
	for r:=0; r<9; r++ {
		val := make([]int, 0)
		for c:=0; c<9; c++{
			val = append(val, -1)
		}
		board = append(board, val)
	}
	
	board[0][0] = 1
	board[1][1] = 1
	
	for _,row := range board{
		fmt.Println(row)
	}
	
	
	if(validSudoku(board)){
	    fmt.Println("valid sudoku")    
	} else {
	    fmt.Println("invalid sudoku")    
	}
}


func validSudoku(board [][]int) bool{
	resp := make(chan bool, 3)
	go checkRows(board, resp)
	go checkCols(board, resp)
	go checkBoxes(board, resp) 
	
	rows := <- resp
	cols := <- resp
	boxes := <- resp
	panic("adsf")
	return rows && cols && boxes
}


func checkRows(board [][]int, resp chan bool ){
	
	for r := 0; r<9; r++ {
        set := make(map[int]bool)
        for c:=0; c<9; c++ {
        	if board[r][c] != -1 {
            	if _, ok := set[board[r][c]]; ok {
            	    resp <- false  
            	      
            	} else {
            	    set[board[r][c]] = true   
            	}
        	}
        }    
     }
     
     resp <- true
     	
}

func checkCols(board [][]int, resp chan bool){
	
	for r := 0; r<9; r++ {
        set := make(map[int]bool)
        for c:=0; c<9; c++ {
        	if board[c][r] != -1 {
            	if _, ok := set[board[c][r]]; ok {
            	    resp <- false    
            	} else {
            	    set[board[c][r]] = true   
            	}
        	}
        }    
     }
     resp <- true
     	
}

func alreadyExistsOrAdd(board [][]int,  set map[int]bool,  r int,  c int) bool{
    if _,ok:= set[board[r][c]]; ok {
        return true
    } else {
        set[board[r][c]] = true
        return false
    }
}

func checkBox(board [][]int, r int,c int) bool{
    toCheck := [][]int{
        []int{r,c},
        []int{r,c+1},
        []int{r,c+2},
        []int{r+1,c},
        []int{r+1,c+1},
        []int{r+1,c+2},
        []int{r+2,c},
        []int{r+2,c+1},
        []int{r+2,c+2},
    }
    set := map[int]bool{}
    alreadyExists := false
    for _, val := range toCheck {
        if alreadyExists = alreadyExistsOrAdd(board, set, val[0], val[1]); alreadyExists {
            return false
        }    
    }
    return true
}

func checkBoxes( board [][]int, resp chan bool){
    for r:=0; r<3;r++{
        for c:=0; c<3; c++{
            
            if res := checkBox(board,r, c); !res{
                resp <- false
            } 
        }    
    }
    resp <- true
}
