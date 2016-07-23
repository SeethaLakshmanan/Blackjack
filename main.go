package Blackjack

import (
"math/rand"
//"fmt"
)

type Card struct{
	number int //TODO: change it to string and assign appropriately
	value int
	suit string
}

type Game struct{
	players []Player
	deck []Card
}

type Player struct{
	score int
	cards []Card
	result string
}

func (g *Game) Deal(){
	//add 2 cards at random to both players
	players:=[]Player{}
	for _,p:=range g.players{
		score:=0
		for i:=0;i<2;i++{
			card:=Card{}
			card.number=rand.Intn(13)
			card.value=card.number
			score+=card.value
			p.cards = append(p.cards,card)			
		}	
		p.score=score
		players = append(players,p)	
	}
	g.players = players
}

func (g *Game) Init(){
	//add deck of cards
	for i:=1;i<=4;i++{
		for j:=1;j<=13;j++{
			card:=Card{}

			card.number=j
			card.value=j
			if i==1{
				card.suit="Hearts"
			}else if i==2{
				card.suit="Spades"
			}else if i==2{
				card.suit="Diamonds"
			}else if i==2{
				card.suit="Clubs"
			}

			g.deck = append(g.deck,card)
		}
	}

	//add players
	for i:=0;i<2;i++{
		player:=Player{score:0, cards:[]Card{}, result:""}
		g.players = append(g.players,player)
	}
}

