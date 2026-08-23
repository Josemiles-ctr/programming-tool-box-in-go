package main
import (
	"fmt"
   "goProgramming/datastructures"

)
func main() {
	//arr:=[4]int{2, 5, 6,10}
	//out:=1
	//result:=datastructures.ProductExcept(arr, out)
	//fmt.Print("The product of array [")
	//for _, i:=range(arr){
	//	fmt.Printf(" %d ", i)		
	//}
	//fmt.Printf("] without element at index %d at is %d", out, result)
  first :=datastructures.LinkedList()
	for first.Next != nil {
 		fmt.Printf("%d -> ", first.Data)
		first = first.Next
	}
}
