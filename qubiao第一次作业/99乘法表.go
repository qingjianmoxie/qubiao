package main

import "fmt"

func main(){
	/*
	99乘法表
	*/
	for i:=1;i<10;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%d*%d=%d\t",j,i,i*j)
		}
		fmt.Println()
	}
}