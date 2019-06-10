package main
import "fmt"
import "golibrary"

var board1 golibrary.Board
func listofmoves(p1 golibrary.Queen,B golibrary.Board,p int){
	if(p==1){
		fmt.Println("List of possible moves:")
	}
	if(p==2){
		fmt.Println("list of possible arrow shots")
	}
	for i:=0;i<6;i++{
		for j:=0;j<6;j++{			
			if(p1.Canmove(i,j,B)){
				fmt.Println(j,i)
			}
		}
	}
}
func noofmoves(p golibrary.Queen,B golibrary.Board) int{
	
	c := 0
	for i:=0;i<6;i++{
		for j:=0;j<6;j++{
			if(p.Canmove(i,j,board1)){
				c=c+1
			}
		}
	}
	return c
}
func game_finished(p11,p12,p21,p22 golibrary.Queen,B golibrary.Board) int{
	res := 0
	if(noofmoves(p11,B)==0&&noofmoves(p12,B)==0){
		res=1
	}
	if(noofmoves(p21,B)==0&&noofmoves(p22,B)==0){
		res=2
	}
	return res
}
func AI_choose(p,p1 *golibrary.Queen,B *golibrary.Board)int{
	if(noofmoves(*p,*B)==0){
		return 2
	}else{
		return 1
	}
}
func AI_moves(p *golibrary.Queen,B *golibrary.Board){
	xm:=0
	ym:=0
	
	for i:=0;i<6;i++{
		for j:=0;j<6;j++{
			if(p.Canmove(i,j,*B)){
				xm=i
				ym=j
				p.Move(xm,ym,B);
				return
			}
		}
	}
}
func AI_shoots(p *golibrary.Queen,B *golibrary.Board){
	var xm,ym int
	
	for i:=0;i<6;i++{
		for j:=0;j<6;j++{
			if(p.Canmove(i,j,*B)){
				xm=i
				ym=j
				p.Shoot(xm,ym,B);
				return
			}
		}
	}
}
func initialize(){
	for i:=0;i<6;i++{
		for j:=0;j<6;j++{
			board1[i][j]="em";
		}
	}
	
	board1[5][3]="f2"
	board1[3][0]="s1"
	board1[0][2]="f1"
	board1[2][5]="s2"
	
}
func show(){
	for i:=0;i<6;i++{
		for j:=0;j<6;j++{
			fmt.Printf("%s     ",board1[i][j]);
		}
		fmt.Printf("\n\n\n")
	}
}

func main(){
	
	var p[4] golibrary.Queen
	
	p[0] = golibrary.Queen{Name:"f1",Current_xpos:0,Current_ypos:2}
	p[1] = golibrary.Queen{Name:"f2",Current_xpos:5,Current_ypos:3}
	p[2] = golibrary.Queen{Name:"s1",Current_xpos:3,Current_ypos:0}
	p[3] = golibrary.Queen{Name:"s2",Current_xpos:2,Current_ypos:5}
	initialize();
	show();
	turn := 1
	for game_finished(p[0],p[1],p[2],p[3],board1)==0{
		if(turn==1){
			var xm,ym,xs,ys int
			fmt.Println("player 1's turn")
			fmt.Println("enter which piece")
			fmt.Println("1 to play with f1")
			fmt.Println("2 to play with f2")
			ch := 0
			fmt.Scan(&ch)
			if(noofmoves(p[0],board1)==0){
				ch=2
				fmt.Println("f1 cannot move hence playing with f2")
			}
			if(noofmoves(p[1],board1)==0){
				ch=1
				fmt.Println("f2 cannot move hence playing with f1")
			}
			if(ch==1){
				listofmoves(p[0],board1,1)
			}else{
				listofmoves(p[1],board1,1)
			}
			fmt.Println("enter x coordinates to move your piece")
			fmt.Scan(&ym)
			fmt.Println("enter y coordinates to move your piece")
			fmt.Scan(&xm)
			
			for !(xm<6&&ym<6&&xm>=0&&ym>=0){
				fmt.Println("please enter valid coordinates")
				fmt.Println("enter x coordinates to move your piece")
				fmt.Scan(&ym)
				fmt.Println("enter y coordinates to move your piece")
				fmt.Scan(&xm)
				
			}
			if(ch==1){
				for !(p[0].Canmove(xm,ym,board1)){
					fmt.Println("your piece cannot move to this position please enter coordinates again")
					fmt.Println("enter x coordinates to move your piece")
					fmt.Scan(&ym)
					fmt.Println("enter y coordinates to move your piece")
					fmt.Scan(&xm)
					
				}
				p[0].Move(xm,ym,&board1);
				show();
				if(ch==1){
					listofmoves(p[0],board1,2)
				}else{
					listofmoves(p[1],board1,2)
				}
				fmt.Println("enter coordinates to shoot arrow")
				fmt.Println("enter x coordinates to shoot arrow")
				fmt.Scan(&ys)
				fmt.Println("enter y coordinates to shoot arrow")
				fmt.Scan(&xs)
				
				for !(xs<6&&ys<6&&xs>=0&&ys>=0){
					fmt.Println("please enter valid coordinates")
					fmt.Println("enter x coordinates to shoot arrow")
					fmt.Scan(&ys)
					fmt.Println("enter y coordinates to shoot arrow")
					fmt.Scan(&xs)
					
				}
				for !(p[0].Canmove(xs,ys,board1)){
					fmt.Println("your piece cannot shoot to this position please enter coordinates again")
					fmt.Println("enter x coordinates to shoot arrow")
					fmt.Scan(&ys)
					fmt.Println("enter y coordinates to shoot arrow")
					fmt.Scan(&xs)
					
				}
				p[0].Shoot(xs,ys,&board1);
				show();
			}else{
				for !(p[1].Canmove(xm,ym,board1)){
					fmt.Println("your piece cannot move to this position please enter coordinates again")
					fmt.Println("enter x coordinates to move your piece")
					fmt.Scan(&ym)
					fmt.Println("enter y coordinates to move your piece")
					fmt.Scan(&xm)
					
				}
				p[1].Move(xm,ym,&board1);
				show();
				if(ch==1){
					listofmoves(p[0],board1,2)
				}else{
					listofmoves(p[1],board1,2)
				}
				fmt.Println("enter coordinates to shoot arrow")
				fmt.Println("enter x coordinates to shoot arrow")
				fmt.Scan(&ys)
				fmt.Println("enter y coordinates to shoot arrow")
				fmt.Scan(&xs)
				
				for !(xs<6&&ys<6&&xs>=0&&ys>=0){
					fmt.Println("please enter valid coordinates")
					fmt.Println("enter x coordinates to shoot arrow")
					fmt.Scan(&ys)
					fmt.Println("enter y coordinates to shoot arrow")
					fmt.Scan(&xs)
					
				}
				for !(p[1].Canmove(xs,ys,board1)){
					fmt.Println("your piece cannot shoot to this position please enter coordinates again")
					fmt.Println("enter x coordinates to shoot arrow")
					fmt.Scan(&ys)
					fmt.Println("enter y coordinates to shoot arrow")
					fmt.Scan(&xs)
					
				}
				p[1].Shoot(xs,ys,&board1);
				show();
			}
			turn=2
		}else{
			//var xm,ym,xs,ys int
			fmt.Println("AI's turn")
			pawn := AI_choose(&p[2],&p[3],&board1);
			if(pawn==1){
				AI_moves(&p[2],&board1);
			}else{
				AI_moves(&p[3],&board1);
			}
			if(pawn==1){
				AI_shoots(&p[2],&board1);
			}else{
				AI_shoots(&p[3],&board1);
			}
			
			show();
			
			/*fmt.Println("enter which piece")
			fmt.Println("1 to play with s1")
			fmt.Println("2 to play with s2")
			ch := 0
			fmt.Scan(&ch)
			if(noofmoves(p[2],board1)==0){
				ch=2
				fmt.Println("s1 cannot move hence playing with s2")
			}
			if(noofmoves(p[3],board1)==0){
				ch=3
				fmt.Println("s2 cannot move hence playing with s1")
			}
			if(ch==1){
				listofmoves(p[2],board1,1)
			}else{
				listofmoves(p[3],board1,1)
			}
			fmt.Println("enter coordinates to move your piece")
			fmt.Println("enter x coordinates to move your piece")
			fmt.Scan(&ym)
			fmt.Println("enter y coordinates to move your piece")
			fmt.Scan(&xm)
			
			for !(xm<6&&ym<6&&xm>=0&&ym>=0){
				fmt.Println("please enter valid coordinates")
				fmt.Println("enter x coordinates to move your piece")
				fmt.Scan(&ym)
				fmt.Println("enter y coordinates to move your piece")
				fmt.Scan(&xm)
				
			}
			if(ch==1){
				for !(p[2].Canmove(xm,ym,board1)){
					fmt.Println("your piece cannot move to this position please enter coordinates again")
					fmt.Println("enter x coordinates to move your piece")
					fmt.Scan(&ym)
					fmt.Println("enter y coordinates to move your piece")
					fmt.Scan(&xm)
					
				}
				p[2].Move(xm,ym,&board1);
				show();
				if(ch==1){
					listofmoves(p[2],board1,2)
				}else{
					listofmoves(p[3],board1,2)
				}
				fmt.Println("enter coordinates to shoot arrow")
				fmt.Println("please enter valid coordinates")
				fmt.Println("enter x coordinates to shoot arrow")
				fmt.Scan(&ys)
				fmt.Println("enter y coordinates to shoot arrow")
				fmt.Scan(&xs)
				
				for !(xs<6&&ys<6&&xs>=0&&ys>=0){
					fmt.Println("please enter valid coordinates")
					fmt.Println("enter x coordinates to shoot arrow")
					fmt.Scan(&ys)
					fmt.Println("enter y coordinates to shoot arrow")
					fmt.Scan(&xs)
					
				}
				for !(p[2].Canmove(xs,ys,board1)){
					fmt.Println("your piece cannot shoot to this position please enter coordinates again")
					fmt.Println("enter x coordinates to shoot arrow")
					fmt.Scan(&ys)
					fmt.Println("enter y coordinates to shoot arrow")
					fmt.Scan(&xs)
					
				}
				p[2].Shoot(xs,ys,&board1);
				show();
			}else{
				for !(p[3].Canmove(xm,ym,board1)){
					fmt.Println("your piece cannot move to this position please enter coordinates again")
					fmt.Println("enter x coordinates to move your piece")
					fmt.Scan(&ym)
					fmt.Println("enter y coordinates to move your piece")
					fmt.Scan(&xm)
					
				}
				p[3].Move(xm,ym,&board1);
				show();
				if(ch==1){
					listofmoves(p[2],board1,2)
				}else{
					listofmoves(p[3],board1,2)
				}
				fmt.Println("enter coordinates to shoot arrow")
				fmt.Println("please enter valid coordinates")
				fmt.Println("enter x coordinates to shoot arrow")
				fmt.Scan(&ys)
				fmt.Println("enter y coordinates to shoot arrow")
				fmt.Scan(&xs)
				
				for !(xs<6&&ys<6&&xs>=0&&ys>=0){
					fmt.Println("please enter valid coordinates")
					fmt.Println("enter x coordinates to shoot arrow")
					fmt.Scan(&ys)
					fmt.Println("enter y coordinates to shoot arrow")
					fmt.Scan(&xs)
					
				}
				for !(p[3].Canmove(xs,ys,board1)){
					fmt.Println("your piece cannot shoot to this position please enter coordinates again")
					fmt.Println("enter x coordinates to shoot arrow")
					fmt.Scan(&ys)
					fmt.Println("enter y coordinates to shoot arrow")
					fmt.Scan(&xs)
					
				}
				p[3].Shoot(xs,ys,&board1);
				show();
			}*/
			turn=1

		}
	}
	if(game_finished(p[0],p[1],p[2],p[3],board1)==1){
		fmt.Println("player 2 has won!!!")
	}else{
		fmt.Println("player 1 has won!!")
	}
}
