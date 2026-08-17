package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if currency != "USD" && currency != "EUR" {
		return "", errors.New("invalid currency")
	}
	if locale != "en-US" && locale != "nl-NL" {
		return "", errors.New("invalid locale")
	}

	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	sort.Slice(entriesCopy, func(i, j int) bool {
		a := entriesCopy[i]
		b := entriesCopy[j]
		if a.Date != b.Date {
			return a.Date < b.Date
		}
		if a.Description != b.Description {
			return a.Description < b.Description
		}
		return a.Change < b.Change
	})

	var header string
	switch locale {
	case "en-US":
		header = "Date       | Description               | Change       \n"
	case "nl-NL":
		header = "Datum      | Omschrijving              | Verandering  \n"
	}

	var sb strings.Builder
	sb.WriteString(header)

	for _, e := range entriesCopy {
		if len(e.Date) != 10 || e.Date[4] != '-' || e.Date[7] != '-' {
			return "", errors.New("invalid date")
		}

		year := e.Date[:4]
		month := e.Date[5:7]
		day := e.Date[8:10]

		var displayDate string
		switch locale {
		case "en-US":
			displayDate = fmt.Sprintf("%s/%s/%s", month, day, year)
		case "nl-NL":
			displayDate = fmt.Sprintf("%s-%s-%s", day, month, year)
		}

		desc := e.Description
		if len(desc) > 25 {
			desc = desc[:22] + "..."
		}
		desc = fmt.Sprintf("%-25s", desc)

		amountStr, err := formatMoney(e.Change, currency, locale)
		if err != nil {
			return "", err
		}

		line := fmt.Sprintf("%-10s | %s | %13s\n", displayDate, desc, amountStr)
		sb.WriteString(line)
	}

	return sb.String(), nil
}

func formatMoney(cents int, currency string, locale string) (string, error) {
	isNegative := false
	if cents < 0 {
		isNegative = true
		cents = -cents
	}

	dollars := cents / 100
	frac := cents % 100

	dStr := strconv.Itoa(dollars)
	var groups []string
	for i := len(dStr); i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		groups = append([]string{dStr[start:i]}, groups...)
	}

	var symbol string
	switch currency {
	case "USD":
		symbol = "$"
	case "EUR":
		symbol = "€"
	default:
		return "", errors.New("bad currency")
	}

	var val string
	switch locale {
	case "en-US":
		intPart := strings.Join(groups, ",")
		val = fmt.Sprintf("%s%s.%02d", symbol, intPart, frac)
		if isNegative {
			val = "(" + val + ")"
		} else {
			val += " "
		}
		return val, nil
	case "nl-NL":
		intPart := strings.Join(groups, ".")
		if isNegative {
			val = fmt.Sprintf("%s -%s,%02d", symbol, intPart, frac)
		} else {
			val = fmt.Sprintf("%s %s,%02d", symbol, intPart, frac)
		}
		val += " "
		return val, nil
	default:
		return "", errors.New("bad locale")
	}
}