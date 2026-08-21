//datastructures
package datastructures

func ProductExcept(arr[4] int, out int) int{
product:=1
	for index, i:=range(arr){
		
		if index!=out{
			product*=i
		}
	}
	return product
}
