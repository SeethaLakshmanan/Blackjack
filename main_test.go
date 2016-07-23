package Blackjack

import "testing"

func TestDeck(t *testing.T){
	g:= Game{}

	g.Init()

	result:=len(g.deck)
	
	if result!=52{
		t.Errorf("Expected count of cards in deck - 52, got %d", result)
	}

	players:=len(g.players)
	if players!=2{
		t.Errorf("Expected count of players - 2, got %d", players)
	}
}

func TestDeal(t *testing.T){
	g:= Game{}

	g.Init()
	g.Deal()

	for _,p:=range g.players{
		result:=len(p.cards)
		if result!=2{
			t.Errorf("Expected count of cards in player - 2, got %d", result)
		}

		if p.score==0{
			t.Errorf("Expected score for player > 0, got %d", result)
		}
	}
}

