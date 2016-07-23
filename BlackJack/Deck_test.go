package BlackJack

import "testing"


func TestDeckHas52Cards(t *testing.T){

		if len(InitializeDeck()) != 52 {
			t.Errorf("Expected the result is not 52")
		}
							
}



func TestDealerGetsTheDeck(t* testing.T){

	if len(DeckToDealer())!=52 {
			t.Errorf("Expected the result is not 52")
}

}
/*

func TestRandomCardFromDeck(t* testing.T){


	cards := new(Card)
	cards=GenerateRandomCard(DeckToDealer())
	if (cards == nil) {
		t.Errorf("Card is not generated and assigned to the dealer")
	}
}
*/


