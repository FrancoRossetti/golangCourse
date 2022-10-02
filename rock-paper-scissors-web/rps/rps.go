package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0
	PAPER        = 1
	SCISSORS     = 2
	PLAYERWINS   = 0
	COMPUTERWINS = 2
	DRAW         = 1
)

var RoundResponses = [][]string{
	{"Ganó el jugador!!!", "Le ganaste a la máquina, crack!", "Se la lleva el player!"},
	{"Empataste!!", "No ganó ninguno de los dos, probá de nuevo!", "Termina en empate."},
	{"Perdiste!!", "Ganó la computadora...", "Probá de nuevo, te ganó la PC!"},
}
var RoundResults = []string{
	"Gana Jugador",
	"Empate",
	"Gana Computadora",
}

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
	Winner         int    `json:"winner"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())

	computerValue := rand.Intn(3)
	computerChoice := ""

	//Switch to know what did the computer chose
	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}

	winner := 0
	//Check who won
	if playerValue == computerValue {
		winner = DRAW
	} else if playerValue == (computerValue+1)%3 {
		winner = PLAYERWINS
	} else {
		winner = COMPUTERWINS
	}

	// Get random message
	// get a random number to use to pick a random response
	randomIndex := rand.Intn(len(RoundResponses[winner]))
	// assign randomly selected response to output
	var output = RoundResponses[winner][randomIndex]

	var result Round
	result.Message = output
	result.ComputerChoice = computerChoice
	result.RoundResult = RoundResults[winner]
	result.Winner = winner
	return result

}
