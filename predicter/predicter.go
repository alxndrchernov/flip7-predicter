package predicter

import (
	"strconv"
	"strings"
)

const cardNum = 94
const actionCardsNum = 9
const modifierCardsNum = 6

type Player struct {
	currentCards map[uint8]uint8
}

func NewPlayer(cards []string) *Player {
	return &Player{
		currentCards: getCardPoints(cards),
	}
}

func (p *Player) Predict(removedCards []string) float64 {
	removedCardPoints := getCardPoints(removedCards)
	var removedCardsCount uint8
	for _, c := range removedCardPoints {
		removedCardsCount += c
	}
	allCardsCount := cardNum - removedCardsCount - actionCardsNum - modifierCardsNum
	var res float64
	for card := range p.currentCards {
		inPack := card - 1
		if countRemovedCards, ok := removedCardPoints[card]; ok && countRemovedCards > 0 {
			inPack -= countRemovedCards
		}
		if inPack <= 0 {
			continue
		}
		res += float64(inPack) / float64(allCardsCount)
	}
	return res
}

func getCardPoints(cards []string) map[uint8]uint8 {
	res := make(map[uint8]uint8, len(cards))
	for _, card := range cards {
		card = strings.TrimSpace(card)
		cardPoint, err := strconv.Atoi(card)
		if err == nil {
			res[uint8(cardPoint)]++
		}
	}
	return res
}
