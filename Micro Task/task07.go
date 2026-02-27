//string "25"  int  convert  5 add  & print ।
package main
import "fmt" 
import "strconv"

func main(){
	var str string = "25"
	var a int = 5
	convert,_ := strconv.Atoi(str)
	sum := convert + a
	fmt.Println(sum)
}