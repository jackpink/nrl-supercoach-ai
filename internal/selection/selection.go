package selection


type Player struct {
	id int
	name string
	position1 string
	position2 string
	price int
	team string
}

type Squad struct {
	Fullbacks struct {
		players []Player
		size int
	}
}

func (squad Squad) addPlayer(player Player)

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