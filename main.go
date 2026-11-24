package main

import (
	"fmt"
	"sort"
)

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func dodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64) {
	p, obstaja := postaje[oznaka]
	if !obstaja {
		fmt.Printf("Napaka: postaja %q ne obstaja.\n", oznaka)
		return
	}
	if temperatura < -50 || temperatura > 50 {
		fmt.Printf("Napaka: temperatura %.1f °C ni veljavna (dovoljeno od -50 do 50 °C).\n", temperatura)
		return
	}
	p.meritve = append(p.meritve, temperatura)
	postaje[oznaka] = p
}

func statistika(postaje map[string]Postaja, oznaka string) (najnizja, najvisja, povprecna float64, ok bool) {
	p, obstaja := postaje[oznaka]
	if !obstaja || len(p.meritve) < 3 {
		return 0, 0, 0, false
	}
	najnizja = p.meritve[0]
	najvisja = p.meritve[0]
	vsota := 0.0
	for _, t := range p.meritve {
		if t < najnizja {
			najnizja = t
		}
		if t > najvisja {
			najvisja = t
		}
		vsota += t
	}
	povprecna = vsota / float64(len(p.meritve))
	return najnizja, najvisja, povprecna, true
}

func urejeneOznake(postaje map[string]Postaja) []string {
	oznake := make([]string, 0, len(postaje))
	for oznaka := range postaje {
		oznake = append(oznake, oznaka)
	}
	sort.Strings(oznake)
	return oznake
}

func izpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for _, oznaka := range urejeneOznake(postaje) {
		p := postaje[oznaka]
		fmt.Printf("%s - %s (%s): %v\n", oznaka, p.ime, p.kraj, p.meritve)
	}
}

func izpisPorocila(postaje map[string]Postaja) {
	for _, oznaka := range urejeneOznake(postaje) {
		p := postaje[oznaka]
		min, max, povp, ok := statistika(postaje, oznaka)
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

		fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", p.ime, min, max, povp, opis)
	}
}

func main() {
	postaje := map[string]Postaja{
		"LJ01": {ime: "Ljubljana Bežigrad", kraj: "Ljubljana", meritve: []float64{18.2, 21.5, 19.8, 16.5}},
		"MB01": {ime: "Maribor Tabor", kraj: "Maribor", meritve: []float64{12.4}},
		"KP01": {ime: "Koper Luka", kraj: "Koper", meritve: []float64{24.5, 27, 26.1, 28.4}},
		"KR01": {ime: "Kredarica", kraj: "Triglav", meritve: []float64{2, -1.5, 3.5}},
	}

	fmt.Println("--- Začetno stanje ---")
	izpisMeritev(postaje)
	fmt.Println()
	izpisPorocila(postaje)

	fmt.Println("\n--- Dodajanje meritev ---")
	dodajMeritev(postaje, "MB01", 14.1)
	dodajMeritev(postaje, "MB01", 15.3)
	dodajMeritev(postaje, "LJ01", 75.0)
	dodajMeritev(postaje, "KR01", -60.2)
	dodajMeritev(postaje, "XX99", 20.0)

	fmt.Println("\n--- Stanje po dodajanju ---")
	izpisMeritev(postaje)
	fmt.Println()
	izpisPorocila(postaje)

	fmt.Println("\n--- Neposreden klic statistike ---")
	if min, max, povp, ok := statistika(postaje, "LJ01"); ok {
		fmt.Printf("LJ01: min %.1f, max %.1f, povprečje %.1f\n", min, max, povp)
	}
	if _, _, _, ok := statistika(postaje, "XX99"); !ok {
		fmt.Println("XX99: statistike ni mogoče izračunati")
	}
}
