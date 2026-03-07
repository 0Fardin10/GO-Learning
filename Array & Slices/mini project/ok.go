package main

import ("fmt")
func main(){
	a:=[5] int{}
	fmt.Println("enter 5 number:") 
	    for i := 0; i <5; i++ {
			fmt.Scan(&a[i])
		}
		fmt.Println("Now revars time")
		for i:=4;i>=0;i--{
			fmt.Printf("%d\n",a[i])
		}
		fmt.Println()
}