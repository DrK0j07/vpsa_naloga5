# vpsa_naloga5

Go modul za delo z vremenskimi postajami in njihovimi temperaturnimi meritvami.

Projekt je nastal v okviru 5. domače naloge pri predmetu Vzporedni in porazdeljeni sistemi in algoritmi.

## Namestitev

Modul lahko namestimo z ukazom:

```bash
go get github.com/DrK0j07/vpsa_naloga5@v1.0.0
```

## Paket vreme

Paket `vreme` omogoča:
- dodajanje vremenskih postaj,
- dodajanje temperaturnih meritev,
- izpis vseh meritev,
- izračun statistike in izpis temperaturnega poročila.

## Primer uporabe paketa

```go
package main

import "github.com/DrK0j07/vpsa_naloga5/vreme"

func main() {
    postaje := make(map[string]vreme.Postaja)

    vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
    vreme.DodajMeritev(postaje, "LJ01", 21.5, -50, 50)

    vreme.IzpisMeritev(postaje)
    vreme.IzpisPorocila(postaje, 1)
}
```

## Uporaba programa

Program omogoča nastavitev temperaturnih omejitev in najmanjšega števila meritev s pomočjo CLI zastavic.

| Zastavica | Privzeta vrednost | Opis |
|---|---:|---|
| `--minTemp` | -50 | Najnižja dovoljena temperatura |
| `--maxTemp` | 50 | Najvišja dovoljena temperatura |
| `--stMeritev` | 3 | Najmanjše število meritev za izračun statistike |

Primer zagona:

```bash
go run . --minTemp=-30 --maxTemp=40 --stMeritev=4
```

Pomoč programa:

```bash
go run . --help
```

## Licenca

Projekt je objavljen pod licenco MIT. Več informacij je v datoteki [LICENSE](LICENSE).
