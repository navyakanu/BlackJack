package BlackJack

import "testing"

func TestPlayerScoreEquals21(t *testing.T) {
		expectedResult := 21
		if Score() != expectedResult {
			
			t.Errorf("Expected the result to be %d, but got %d", expectedResult, Score())
		}

	}

	


