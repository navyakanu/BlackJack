package BlackJack

import "fmt"
import "math/rand"


type Card struct{

	Name string
	Value int
}


type Deck []Card 


type Dealer struct {

	DealerDeck Deck


}



func GenerateRandomCard(Dd Deck) Card {
		key:=rand.Intn(52)
		return Dd[key]



}




func DeckToDealer() Deck {

		Dealer := new(Dealer)
		Dealer.DealerDeck=InitializeDeck()
		return Dealer.DealerDeck

}




func InitializeDeck() Deck {
 	cards := make([]Card,0,0)

	//Ace

	for i:=0;i<4;i++{
		cards = append(cards, Card{Name:"A",Value:11})
	}

	// for face value cards
	for i:=2;i<=10;i++{
		for j:=0;j<4;j++{
			cards = append(cards, Card{Name:fmt.Sprintf("%d",i),Value:i})
		}
	}

	aboveFaceValueCards := []string{"J","Q","K"}
	for _,item := range aboveFaceValueCards{
		for i:=0;i<4;i++{
			cards = append(cards, Card{Name:item,Value:10})
		}
	}

	return cards
}

