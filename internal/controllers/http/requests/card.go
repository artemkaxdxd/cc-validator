package requests

import (
	"solidgate-test/config"
	"strconv"
	"time"
)

type Card struct {
	Number string `json:"number"`
	Month  string `json:"month"`
	Year   string `json:"year"`
}

func (c Card) ValidateNumber() error {
	total := 0
	isSecondDigit := false

	for i := len(c.Number) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(c.Number[i]))

		if isSecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		total += digit
		isSecondDigit = !isSecondDigit
	}

	if total%10 != 0 {
		return config.ErrInvalidCardNumber
	}
	return nil
}

func (c Card) ValidateDate() error {
	var year, month int
	var err error
	timeNow := time.Now()

	if len(c.Year) < 3 {
		fullYear := strconv.Itoa(timeNow.UTC().Year())[:2] + c.Year
		year, err = strconv.Atoi(fullYear)
		if err != nil {
			return config.ErrInvalidYear
		}
	} else {
		year, err = strconv.Atoi(c.Year)
		if err != nil {
			return config.ErrInvalidYear
		}
	}

	month, err = strconv.Atoi(c.Month)
	if err != nil {
		return config.ErrInvalidMonth
	}

	if month < 1 || month > 12 {
		return config.ErrInvalidMonth
	}

	if year < timeNow.UTC().Year() {
		return config.ErrExpiredCard
	}

	if year == timeNow.UTC().Year() && month < int(timeNow.UTC().Month()) {
		return config.ErrExpiredCard
	}
	return nil
}
