package main

import (
    "golang.org/x/tour/tree"
    "fmt"
)

func Walk(t *tree.Tree, ch chan int){
    defer close(ch)
    var f func(tr *tree.Tree){
      if tr == nil {
	return
      }
      f(tr.Left)
      ch <- tr.Value
      f(tr.Right)
    }
    f(t)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool{
    ch1 := make(chan int)
    ch2 := make(chan int)
    result := make(map[int]bool)
    i, j := 0, 0
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    for{
        select{
        case v := <-ch1:
            _, ok := result[v]
            if !ok {
                result[v]=true
            }else{
                delete(result, v)
            }
            i++
        case v := <-ch2:
            _, ok := result[v]
            if !ok {
                result[v]=true
            }else{
                delete(result, v)
            }
            j++
        }
        if i == 10 && j == 10 {
            break
        }
    }
    
    if len(result) == 0 {
        return true
    }else {
        return false
    }
}

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    t11 := tree.New(1)
    t12 := tree.New(1)
    go Walk(t11, ch1)
    for v := range ch1 {
        fmt.Println(v)
    }
    fmt.Println("xxxxxxxxxx")
    go Walk(t12, ch2)
    for v := range ch2 {
        fmt.Println(v)
    }
    //t1 :=tree.New(1)
    //t2 :=tree.New(2)
    //fmt.Println(Same(t1, t2))
}
