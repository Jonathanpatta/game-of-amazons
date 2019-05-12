package main
import "fmt"
var board[6][6] string
type game struct{
	//board[6][6] int
}
func (c game) initialize(){
	for i:=0;i<6;i++{
		for j:=0;j<6;j++{
			board[i][j]="a";
		}
	}
	board[5][3]="p12"
	board[3][0]="p21"
	board[0][2]="p11"
	board[2][5]="p22"
}
func (c game) show(){
	for i:=0;i<6;i++{
		for j:=0;j<6;j++{
			fmt.Printf("%s    ",board[i][j]);
		}
		fmt.Printf("\n\n")
	}
}
func main(){
	var l game
	var p queen
	l.initialize();
	l.show();
}