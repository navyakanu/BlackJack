package BlackJack

import "testing"

func TestPlayerScoreEquals21(t *testing.T) {
		expectedResult := 21
		if Score(10,11) == expectedResult {
			if Winner() != 1 {
			t.Errorf("Expected the result to be %d, but got %d", expectedResult, Score(10,11))
		}
	}

	}


func TestPlayersTwoCardsGeneration(t *testing.T) {
	var card1,card2 = Generate()
	if card1 == 1 || card2 == 0 {
		t.Errorf("Expected the result to be greater than 0 ")
		}
	
}



	


