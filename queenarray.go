package golibrary

type Board [6][6] string
type Queen struct{
	Name string
	Current_xpos int
	Current_ypos int
	
}
func (c Queen) Occupied(xnew_position,ynew_position int,B Board) bool{
	var result=true
	
	if (B[xnew_position][ynew_position]!="em"){
		result = true
	}else{
		result = false
	}
	return result
}

func (c Queen) Canmove(xnew_position,ynew_position int,B Board) bool{
	
	var x,y,cx,cy int
	x=xnew_position
	y=ynew_position
	cx=c.Current_xpos 
	cy=c.Current_ypos 
	B[cx][cy]="em"
	
	var result bool

	result=false

	//8 cases for 8 directions
	//top right
	
	if(cx<6&&cy<6&&x<6&&y<6){
		for ((x<=5&&y<=5&&x>=0&&y>=0)&&(!c.Occupied(x,y,B))){
			
			if(!c.Occupied(x,y,B)&&(x==cx&&y==cy)){
				result=true
			}
			x=x+1
			y=y+1
		}
		//fmt.Println(result,x,y)
		//right
		x=xnew_position
		y=ynew_position
		for ((x<=5&&y<=5&&x>=0&&y>=0)&&(!c.Occupied(x,y,B))){
			
			if(!c.Occupied(x,y,B)&&(x==cx&&y==cy)){
				result=true
			}
			x=x+1
			y=y+0
		}
		//bottom right
		x=xnew_position
		y=ynew_position
		for ((x<=5&&y<=5&&x>=0&&y>=0)&&(!c.Occupied(x,y,B))){
			
			if(!c.Occupied(x,y,B)&&(x==cx&&y==cy)){
				result=true
			}
			x=x+1
			y=y-1
		}
		//bottom
		x=xnew_position
		y=ynew_position
		for ((x<=5&&y<=5&&x>=0&&y>=0)&&(!c.Occupied(x,y,B))){
			
			if(!c.Occupied(x,y,B)&&(x==cx&&y==cy)){
				result=true
			}
			x=x+0
			y=y-1
		}
		//bottom left
		x=xnew_position
		y=ynew_position
		for ((x<=5&&y<=5&&x>=0&&y>=0)&&(!c.Occupied(x,y,B))){
			
			if(!c.Occupied(x,y,B)&&(x==cx&&y==cy)){
				result=true
			}
			x=x-1
			y=y-1
		}
		//left
		x=xnew_position
		y=ynew_position
		for ((x<=5&&y<=5&&x>=0&&y>=0)&&(!c.Occupied(x,y,B))){
			
			if(!c.Occupied(x,y,B)&&(x==cx&&y==cy)){
				result=true
			}
			x=x-1
			y=y+0
		}
		//top left
		x=xnew_position
		y=ynew_position
		for ((x<=5&&y<=5&&x>=0&&y>=0)&&(!c.Occupied(x,y,B))){
			
			if(!c.Occupied(x,y,B)&&(x==cx&&y==cy)){
				result=true
			}
			x=x-1
			y=y+1
		}
		//up
		x=xnew_position
		y=ynew_position
		for ((x<6&&y<6&&x>=0&&y>=0)&&(!c.Occupied(x,y,B))){
			
			if(!c.Occupied(x,y,B)&&(x==cx&&y==cy)){
				result=true
			}
			x=x+0
			y=y+1
		}
	}
	if(c.Current_xpos==xnew_position&&c.Current_ypos==ynew_position){
		result = false
	}
	
	
	return result
}

func (c* Queen) Move(xnew_position,ynew_position int,B *Board){
	
	
	if(c.Canmove(xnew_position,ynew_position,*B)){
		
		B[xnew_position][ynew_position]=c.Name
		B[c.Current_xpos][c.Current_ypos]="em"
		c.Current_xpos=xnew_position
		c.Current_ypos=ynew_position
		
	}
}


func (c *Queen) Shoot(xnew_position int,ynew_position int,B *Board){
	var p int
	p=0
	if(c.Canmove(xnew_position,ynew_position,*B)){
		p=1
		
	}
	if(p==1){
		
		B[xnew_position][ynew_position]="xx"
		
	}
}
