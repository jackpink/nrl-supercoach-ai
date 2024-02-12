package selection

import _ "github.com/lib/pq"

import (
	"selection/players.sql.go"
)

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
	for i := 0; i < len(f.players); i++ {
		if players[i].id == 0{
			return true
		}
	}
	return false
}

func (f* fullbacks) addPlayer(p player) boolean {
	if f.isNotFull() {
		// add player
		for i := 0; i < len(f.players); i++ {
			if players[i].id == 0 {
				players[i].id = p.id
				players[i].name = p.name
				players[i].position1 = p.position1
				players[i].position2 = p.position2
				players[i].price = p.price
				players[i].team = p.team
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
	if !addedBySecondPosition {
		addedByPlayerSwap, err := 
	}
	
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
	// get starting rookies
	newStartingRookies := GetRookies(1, 2022, 175400)
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