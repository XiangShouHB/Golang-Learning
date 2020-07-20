package main

import (
	"fmt"
	"math"
)

func CalcTriangle(a,b int) int  {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func main()  {
	fmt.Println("3 ,4 对应的斜边为：",CalcTriangle(3,4))
}