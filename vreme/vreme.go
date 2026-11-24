package vreme

import (
	"fmt"
)

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{ime: ime, kraj: kraj}
}

func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	p, obstaja := postaje[oznaka]
	if !obstaja {
		fmt.Printf("Napaka: postaja %q ne obstaja.\n", oznaka)
		return
	}
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Printf("Napaka: temperatura %.1f °C ni veljavna (dovoljeno od %.1f do %.1f °C).\n",
			temperatura, minTemp, maxTemp)
		return
	}
	p.meritve = append(p.meritve, temperatura)
	postaje[oznaka] = p
}
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for _, oznaka := range urejeneOznake(postaje) {
		p := postaje[oznaka]
		fmt.Printf("%s - %s (%s): %v\n", oznaka, p.ime, p.kraj, p.meritve)
	}
}

func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for _, oznaka := range urejeneOznake(postaje) {
		p := postaje[oznaka]
		min, max, povp, ok := statistika(postaje, oznaka, stMeritev)
		if !ok {
			fmt.Printf("%s: premalo meritev\n", p.ime)
			continue
		}

		var opis string
		switch {
		case povp >= 25:
			opis = "Vroče"
		case povp >= 15:
			opis = "Toplo"
		case povp >= 5:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}
		fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n",
			p.ime, min, max, povp, opis)
	}
}
