package main
import "fmt"
type board [][]int
type queen struct{
	current_xpos int
	current_ypos int
}
func (c queen) canmove(xnew_position,ynew_position int,B board) bool{
	
	var x,y int
	xnew_position=x
	ynew_position=y
	for (x+y==c.current_xpos+c.current_ypos){
		fmt.Printf("")
	}
	return true
}
func (c queen) occupied(xnew_position,ynew_position int,B board) bool{
	return true
}
func (c queen) move(xnew_position,ynew_position int,B board){
	for !c.occupied(xnew_position,ynew_position , B) && c.canmove(xnew_position,ynew_position , B){

	}
}
func (c queen) shoot(xposition int,yposition int,B board){

}