package main 
import "fmt"
func main(){
//2d array 
array :=[2][3]string{
{"apple","banana","orange"} ,
{"onion","garlic","carrot"}
}

for i:=0;i<2;i++ { // outer loop
  for j:=0;j<3;j++{//innet loop
  fmt.Println(array[i][j])
  }
  }
  
}
