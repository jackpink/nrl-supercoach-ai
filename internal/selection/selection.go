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

type fullbacks struct {
	players [2]player
}

func (f fullbacks) isNotFull() boolean {
	for player in f.players {
		if player.id == 0{
			return true
		}
	}
	return false
}

func (f* fullbacks) addPlayer(p player) boolean {
	if f.isNotFull() {
		// add player
		for player in f.players {
			if player.id == 0 {
				player.id = .id
				player.name = p.name
				player.position1 = p.position1
				player.position2 = p.position2
				player.price = p.price
				player.team = p.team
			}
		}
		return true
	}
	return false
}

// Map for positions needed (ie FRF to FrontRowForward)

func (squad* squad) addPlayer(player player) {
	// try to add first position
	position := positions(player.position1)
	added, err := squad.position.addPlayer(player)
	// if unsuccessfull try to add second position
	if !added {
		addedBySecondPosition, err := squad.position.addPlayer(player)
	}
	// if unsuccessful try to add via player swap
	if !ad {
		addedByPlayerSwap
	}
	if 
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
	for playersAdded < numberOfRookies {
		// increment price by $50,000
		squad, playersAdded = selectCheapiesInRange()
	}
}

func selectStartingRookies(playersAdded int, numberOfRoookies int, squad Squad) Squad {
	var startingRookies []Player
	for i := 0; i < len(startingRookies); i++ {
		playerAdded := squad.addPlayer(startingRookies[i])
		if playerAdded {
			playersAdded++
		}
		if playersAdded == numberOfRoookies {
			return squad
		}
	}
	return squad

}

func selectBenchRookies() {

}

func selectCheapiesInRange() {

}