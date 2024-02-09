package selection


type player struct {
	id int
	name string
	position1 string
	position2 string
	price int
	team string
}

func (p player) isEmpty() boolean {
	if p.id == 0 {
		return true
	}
	return false
} 

type squad struct {
	fullbacks struct {
		players [2]Player
	}
}

// Map for positions needed (ie FRF to FrontRowForward)

func (squad* squad) addPlayer(player player) {
	// try to add first position

	// if unsuccessfull try to add second position
	// if unsuccessful try to add via player swap
}

func selection(year int, numberOfRookies int) {
	// initialise empty squad
	// run rookies selection
	// run squad selection
	// return squad
}

func selectRookies(squad Squad, year int, numberOfRookies int) Squad {
	// get starting rookies
	playersAdded := 0
	squad, playersAdded = selectStartingRookies()
	squad, playersAdded = selectBenchRookies()
	while playersAdded < numberOfRookies {
		// increment price by $50,000
		squad, playersAdded = selectCheapiesInRange()
	}
}

func selectStartingRookies(playersAdded int, numberOfRoookies int, squad Squad) Squad {
	var startingRookies []Player
	while playersAdded < numberOfRookies {
		squad.addPlayer(startingRookies.pop())
	}
	return squad

}

func selectBenchRookies() {

}

func selectCheapiesInRange() {

}