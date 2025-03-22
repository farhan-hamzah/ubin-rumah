package main
import "fmt"
var ukuranUbin int = 30*30

func conversionToCm(length, width, hasil, a *int){
	*length = *length*100
	*width = *width*100
	*hasil = (*width)*(*length)
	*a = countTiles(*hasil)
	fmt.Print("You need: ", *a," tiles")
}
func countTiles(roomArea int)int{
	if roomArea%ukuranUbin != 0{
		return (roomArea/ukuranUbin+1)
	}
	return roomArea/ukuranUbin
}
func main(){
	var l, w, result, b int
	fmt.Scan(&l, &w)
	conversionToCm(&l, &w, &result, &b)
}
