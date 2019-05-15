package main
import "fmt"
type board [][] string
type queen struct{
	current_xpos int
	current_ypos int
}
func (c queen) occupied(xnew_position,ynew_position int,B board) bool{
	fmt.Printf("")
	if B[xnew_position][ynew_position]!="a"{//.TrimRight(B[xnew_position][ynew_position], "\n") != "a"{
		return true
	}else{
		return false
	}
}

func (c queen) canmove(xnew_position,ynew_position int,B board) bool{
	
	var x,y,cx,cy int
	x=xnew_position
	y=ynew_position
	cx=c.current_xpos 
	cy=c.current_ypos 

	var result bool

	result=false

	//8 cases for 8 directions
	//top right
	for (x!=cx&&y!=cy)&&(x<6&&y<6){
		x=x+1
		y=y+1
		if(!c.occupied(x,y,B)){
			result=true
		}else{
			result=false
		}
	}
	//right
	for (x!=cx&&y!=cy)&&(x<6&&y<6){
		x=x+1
		y=y+0
		if(!c.occupied(x,y,B)){
			result=true
		}else{
			result=false
		}
	}
	//bottom right
	for (x!=cx&&y!=cy)&&(x<6&&y<6){
		x=x+1
		y=y-1
		if(!c.occupied(x,y,B)){
			result=true
		}else{
			result=false
		}
	}
	//bottom
	for (x!=cx&&y!=cy)&&(x<6&&y<6){
		x=x+0
		y=y-1
		if(!c.occupied(x,y,B)){
			result=true
		}else{
			result=false
		}
	}
	//bottom left
	for (x!=cx&&y!=cy)&&(x<6&&y<6){
		x=x-1
		y=y-1
		if(!c.occupied(x,y,B)){
			result=true
		}else{
			result=false
		}
	}
	//left
	for (x!=cx&&y!=cy)&&(x<6&&y<6){
		x=x-1
		y=y+0
		if(!c.occupied(x,y,B)){
			result=true
		}else{
			result=false
		}
	}
	//top left
	for (x!=cx&&y!=cy)&&(x<6&&y<6){
		x=x-1
		y=y+1
		if(!c.occupied(x,y,B)){
			result=true
		}else{
			result=false
		}
	}
	//up
	for (x!=cx&&y!=cy)&&(x<6&&y<6){
		x=x+0
		y=y+1
		if(!c.occupied(x,y,B)){
			result=true
		}else{
			result=false
		}
	}
	return result
}

func (c queen) move(xnew_position,ynew_position int,B board){
	for !c.occupied(xnew_position,ynew_position , B) && c.canmove(xnew_position,ynew_position , B){

	}
}


func (c queen) shoot(xposition int,yposition int,B board){

}
