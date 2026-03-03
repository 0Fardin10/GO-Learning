package main 
import "fmt"
func main(){

	 a:= [] int{1,2,3,4,5}//declar a array
	 fmt.Println(a)
//slice1
b:=append(a,6,7)	 
	fmt.Println(b,len(b),len(a))
 
}
