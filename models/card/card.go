package card

import (
  "github.com/littlekbt/PorkerNotes-AppEngine/models/porker_notes"
)

type Number int

const (
  ONE Number = 1 << iota
  TWO
  THREE
  FOUR
  FIVE
  SIX
  SEVEN
  EIGHT
  NINE
  TEN
  ELEVEN
  TWELVE
  THIRTEEN
)

type Suit int
const (
  HEART Number = 8192 << iota
  DIAMOND
  SPADE
  CLUB
)

type Card struct {
  Suit porker_notes.Card_Suit `json:"suit"`
  Number porker_notes.Card_Number `json:"number"`
}


func (c Card) MkBinary() int {
  r := 0
  switch c.Suit {
  case porker_notes.Card_HEART:
    r = r | int(HEART)
  case porker_notes.Card_DIAMOND:
    r = r | int(DIAMOND)
  case porker_notes.Card_SPADE:
    r = r | int(SPADE)
  case porker_notes.Card_CLUB:
    r = r | int(CLUB)
  }

  switch c.Number {
  case porker_notes.Card_ONE:
    r = r | int(ONE)
  case porker_notes.Card_TWO:
    r = r | int(TWO)
  case porker_notes.Card_THREE:
    r = r | int(THREE)
  case porker_notes.Card_FOUR:
    r = r | int(FOUR)
  case porker_notes.Card_FIVE:
    r = r | int(FIVE)
  case porker_notes.Card_SIX:
    r = r | int(SIX)
  case porker_notes.Card_SEVEN:
    r = r | int(SEVEN)
  case porker_notes.Card_EIGHT:
    r = r | int(EIGHT)
  case porker_notes.Card_NINE:
    r = r | int(NINE)
  case porker_notes.Card_TEN:
    r = r | int(TEN)
  case porker_notes.Card_ELEVEN:
    r = r | int(ELEVEN)
  case porker_notes.Card_TWELVE:
    r = r | int(TWELVE)
  case porker_notes.Card_THIRTEEN:
    r = r | int(THIRTEEN)
  }

  return r
}