package src

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func validNumber(cardNumber string) bool {

	if len(cardNumber) < 13 || len(cardNumber) >= 19 {
		return false
	}
	sum := 0
	fmt.Println("card number ", cardNumber)
	for i := len(cardNumber) - 1; i >= 0; i = i - 2 {
		sum += int(cardNumber[i]) - int('0')
		value, _ := (strconv.Atoi(string(cardNumber[i-1])))
		value *= 2
		if value > 9 {
			value = (value - 10) + 1
		}
		sum += value
		fmt.Println("card number and ", value, sum)
	}
	fmt.Println("card number", sum)
	if sum%10 == 0 {
		return true
	}
	return false
}

func ValidateExpiration(month, year string) error {
	var yearInt, monthInt int
	var err error
	timeNow := time.Now()

	if len(year) < 3 {
		yearInt, err = strconv.Atoi(strconv.Itoa(timeNow.UTC().Year())[:2] + year)
		if err != nil {
			return errors.New("Invalid year")
		}
	} else {
		yearInt, err = strconv.Atoi(year)
		if err != nil {
			return errors.New("Invalid year")
		}
	}

	monthInt, err = strconv.Atoi(month)
	if err != nil {
		return errors.New("Invalid month")
	}

	if monthInt < 1 || 12 < monthInt {
		return errors.New("Invalid month")
	}

	if yearInt < timeNow.UTC().Year() {
		return errors.New("Credit card has expired")
	}

	if yearInt == timeNow.UTC().Year() && monthInt < int(timeNow.UTC().Month()) {
		return errors.New("Credit card has expired")
	}

	return nil
}
