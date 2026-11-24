// Program naloga5 prikazuje uporabo paketa vreme.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/DrK0j07/naloga5/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "naloga5",
		Usage: "Primer uporabe modula za vremenske postaje",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Value: -50,
				Usage: "najnižja dovoljena temperatura meritve (°C)",
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Value: 50,
				Usage: "najvišja dovoljena temperatura meritve (°C)",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "najmanjše število meritev za izračun statistike",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float64("minTemp")
			maxTemp := cmd.Float64("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			if minTemp >= maxTemp {
				return fmt.Errorf("minTemp mora biti manjši od maxTemp")
			}
			if stMeritev < 1 {
				return fmt.Errorf("stMeritev mora biti vsaj 1")
			}

			postaje := make(map[string]vreme.Postaja)
			vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
			vreme.DodajPostajo(postaje, "MB01", "Maribor Tabor", "Maribor")
			vreme.DodajPostajo(postaje, "KP01", "Koper Luka", "Koper")
			vreme.DodajPostajo(postaje, "KR01", "Kredarica", "Triglav")

			// Začetne meritve iz 1. domače naloge.
			for _, m := range []struct {
				oznaka string
				temp   float64
			}{
				{"LJ01", 18.2}, {"LJ01", 21.5}, {"LJ01", 19.8}, {"LJ01", 16.5},
				{"MB01", 12.4},
				{"KP01", 24.5}, {"KP01", 27}, {"KP01", 26.1}, {"KP01", 28.4},
				{"KR01", 2}, {"KR01", -1.5}, {"KR01", 3.5},
			} {
				vreme.DodajMeritev(postaje, m.oznaka, m.temp, minTemp, maxTemp)
			}

			fmt.Println("--- Meritve ---")
			vreme.IzpisMeritev(postaje)
			fmt.Println("\n--- Poročilo ---")
			vreme.IzpisPorocila(postaje, stMeritev)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
