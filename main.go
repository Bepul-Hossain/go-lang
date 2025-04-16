package main
import "fmt"
import (
    "example.com/mathlib"
)
func main(){
	mathlib.Add(4,5)
}

func init(){
	fmt.Println("hello world")
}