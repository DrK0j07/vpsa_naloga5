// Package vreme omogoča delo z vremenskimi postajami
// in njihovimi meritvami.
package vreme

import (
	"fmt"
	"sort"
)

// Postaja predstavlja vremensko postajo.
// Polja so zasebna in dostopna samo znotraj paketa vreme.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo ustvari novo vremensko postajo in jo doda v slovar postaj.
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{
		ime:  ime,
		kraj: kraj,
	}
}

// DodajMeritev doda temperaturno meritev izbrani postaji.
// Meritev sprejme samo, če je temperatura znotraj dovoljenih mej.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	p, obstaja := postaje[oznaka]

	if !obstaja {
		fmt.Printf("Napaka: postaja %q ne obstaja.\n", oznaka)
		return
	}

	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Printf("Napaka: temperatura %.1f °C ni veljavna.\n", temperatura)
		return
	}

	p.meritve = append(p.meritve, temperatura)
	postaje[oznaka] = p
}

// statistika izračuna najnižjo, najvišjo in povprečno temperaturo.
// Statistiko izračuna samo, če je na voljo dovolj meritev.
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (float64, float64, float64, bool) {
	p, obstaja := postaje[oznaka]

	if !obstaja || len(p.meritve) < stMeritev {
		return 0, 0, 0, false
	}

	min := p.meritve[0]
	max := p.meritve[0]
	vsota := 0.0

	for _, t := range p.meritve {
		if t < min {
			min = t
		}
		if t > max {
			max = t
		}
		vsota += t
	}

	povprecje := vsota / float64(len(p.meritve))

	return min, max, povprecje, true
}

// urejeneOznake vrne oznake postaj v abecednem vrstnem redu.
func urejeneOznake(postaje map[string]Postaja) []string {
	oznake := make([]string, 0, len(postaje))

	for oznaka := range postaje {
		oznake = append(oznake, oznaka)
	}

	sort.Strings(oznake)

	return oznake
}

// IzpisMeritev izpiše vse vremenske postaje in njihove meritve.
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")

	for _, oznaka := range urejeneOznake(postaje) {
		p := postaje[oznaka]
		fmt.Printf("%s - %s (%s): %v\n", oznaka, p.ime, p.kraj, p.meritve)
	}
}

// IzpisPorocila izpiše poročilo o temperaturah vseh vremenskih postaj.
// Za izračun statistike mora imeti postaja vsaj stMeritev meritev.
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
