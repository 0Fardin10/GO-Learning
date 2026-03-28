package main
import "fmt"
func main(){
 m:= map[string]int{
	"zabir": 90,
	"jack":80,
	"mj":85,
}
for key, value := range m{
	fmt.Println(key,value)
}

}
