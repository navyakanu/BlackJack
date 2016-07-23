package BlackJack

import "fmt"


import "math/rand"

type Player struct {
	choice bool
}


func Generate() (int,int) {

	card1 := rand.Intn(11)
	card2 := rand.Intn(10)
	fmt.Println("Player's Card 1")
	fmt.Println(card1)
	fmt.Println("Player's Card 2")
	fmt.Println(card2)



	return card1,card2 


}



func Score(card1 int,card2 int) int {

	
	fmt.Println(rand.Intn(52))
	


	return card1+card2
}





func Winner() int {

	if Score(10,11) == 21 {
		fmt.Println("Player is the winner!!")
		return 1
	}
	return 0
}
