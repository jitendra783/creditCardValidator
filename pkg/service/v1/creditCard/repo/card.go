package repo

import (
	"errors"
	"strconv"
)

type Company struct {
	Name string
}
type digits [6]int

func (s srvSt) CompanyDetails(cardNumber string) (interface{}, error) {
	ccLen := len(cardNumber)
	ccDigits := digits{}
	var err error
	for i := 0; i < 6; i++ {
		if i < ccLen {
			ccDigits[i], err = strconv.Atoi(cardNumber[:i+1])
			if err != nil {
				return nil, err
			}
		}
	}
	switch {
	case isAmex(ccDigits):
		return Company{"American Express"}, nil
	case isBankCard(ccDigits):
		return Company{" Bank Card"}, nil
	case isCabal(ccDigits):
		return Company{"Cabal"}, nil
	case isElo(ccDigits):
		return Company{"Elo"}, nil
	case isDiscover(ccDigits):
		return Company{"Discover"}, nil
	case isDinersClubCarteBlanche(ccDigits, ccLen):
		return Company{"Diners Club Carte Blanche"}, nil
	case isDinersClubEnroute(ccDigits):
		return Company{"Diners Club Enroute"}, nil
	case isDinersClubInternational(ccDigits, ccLen):
		return Company{"Diners Club International"}, nil
	case isHipercard(ccDigits):
		return Company{"Hiper Card"}, nil
	case isInstapayment(ccDigits, ccLen):
		return Company{"Insta Payment"}, nil
	case isJCB(ccDigits):
		return Company{"JCB"}, nil
	case isNaranja(ccDigits):
		return Company{"Naranja"}, nil
	case isMaestro(cardNumber, ccDigits):
		return Company{"Maestro"}, nil
	case isDankort(ccDigits):
		return Company{"Dankort"}, nil
	case isMasterCard(ccDigits):
		return Company{"MasterCard"}, nil
	case isVisaElectron(ccDigits):
		return Company{"Visa Electron"}, nil
	case isVisa(ccDigits):
		return Company{"Visa"}, nil
	case isAura(ccDigits):
		return Company{"Aura"}, nil

	default:
		return Company{""}, errors.New("Unknown credit card method")
	}
}
