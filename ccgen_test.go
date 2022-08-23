package ccgensb_test

import (
	"testing"

	"github.com/brianredbeard/ccgensb"
)

var cardsToTests = []ccgensb.CardType{
	ccgensb.AmericanExpress,
	ccgensb.DinersClub,
	ccgensb.DinersClubUS,
	ccgensb.Discover,
	ccgensb.JCB,
	ccgensb.Laser,
	ccgensb.Maestro,
	ccgensb.Mastercard,
	ccgensb.Solo,
	ccgensb.Unionpay,
	ccgensb.Visa,
	ccgensb.Mir,
}

func TestGenerate(t *testing.T) {
	cnnumber := ccgensb.Generate()

	if len(cnnumber) == 0 {
		t.Error("Length is incorrect")
	}
}
func TestCardTypeGenerate(t *testing.T) {

	for _, cardType := range cardsToTests {
		cnumber := cardType.Generate()
		Suite(t, cnumber, cardType)
	}
}

func TestValidLength(t *testing.T) {
	if ccgensb.AmericanExpress.ValidLength(16) {
		t.Error("Length is incorrect")
	}

	if ccgensb.Visa.ValidLength(1) {
		t.Error("Length is incorrect")
	}
}

func TestGenerateOfLength(t *testing.T) {
	// test with valid card length
	cnumber := ccgensb.Visa.GenerateOfLength(13)
	Suite(t, cnumber, ccgensb.Visa)

	// test with invalid card length
	cnumber = ccgensb.Mastercard.GenerateOfLength(100)
	Suite(t, cnumber, ccgensb.Mastercard)
}

func TestValidNumber(t *testing.T) {
	cnumber := ccgensb.Visa.GenerateOfLength(13)

	// too short card length
	if ccgensb.Visa.ValidNumber(cnumber[:len(cnumber)-1]) {
		t.Error("Card number is not valid")
	}

}

func Suite(t *testing.T, cnumber string, card ccgensb.CardType) {
	// check length is valid
	if !card.ValidLength(len(cnumber)) {
		t.Error("Not valid card length", len(cnumber), card)
	}

	// Luhn check
	if !card.ValidNumber(cnumber) {
		t.Error("Luhn check failed: ", cnumber)
	}
}
