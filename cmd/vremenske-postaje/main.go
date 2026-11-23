// Vremenske-postaje je program, ki ustvari nabor vremenskih postaj, jim doda
// meritve temperature ter izpiše vse meritve in poročilo za vsako postajo.
//
// Uporaba:
//
//	vremenske-postaje [stikala]
//
// Stikala:
//
//	--minTemp float
//		najnižja dovoljena temperatura meritve v °C (privzeto -50)
//	--maxTemp float
//		najvišja dovoljena temperatura meritve v °C (privzeto 50)
//	--stMeritev int
//		najmanjše število meritev, potrebnih za izračun statistike (privzeto 3)
//	--help, -h
//		izpiše pomoč
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/NikZunic/vremenske-postaje/vreme"
	"github.com/urfave/cli/v3"
)

// zazeni ustvari testne postaje z meritvami ter izpiše meritve in poročilo.
// Meritve izven intervala [minTemp, maxTemp] so zavrnjene, statistika pa se
// izračuna le za postaje z vsaj stMeritev meritvami.
func zazeni(minTemp, maxTemp float64, stMeritev int) {

	postaje := make(map[string]vreme.Postaja)

	vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Slovenija")
	for _, t := range []float64{20.5, 21.0, 19.8, 22.1, 23.0} {
		vreme.DodajMeritev(postaje, "LJ01", t, minTemp, maxTemp)
	}

	vreme.DodajPostajo(postaje, "LJ02", "Ljubljana Vič", "Slovenija")
	for _, t := range []float64{18.0, 19.5, 17.7, 20.1, 21.0, 22.3} {
		vreme.DodajMeritev(postaje, "LJ02", t, minTemp, maxTemp)
	}

	vreme.DodajPostajo(postaje, "LJ03", "Ljubljana Šiška", "Slovenija")
	for _, t := range []float64{15.0, 16.5, 14.7, 17.1, 18.0} {
		vreme.DodajMeritev(postaje, "LJ03", t, minTemp, maxTemp)
	}

	vreme.DodajPostajo(postaje, "LJ04", "Ljubljana Jarše", "Slovenija")
	for _, t := range []float64{10.0, 11.5, 9.7, 12.1, 13.0, 14.3} {
		vreme.DodajMeritev(postaje, "LJ04", t, minTemp, maxTemp)
	}

	vreme.DodajPostajo(postaje, "KP01", "Koper Luka", "Koper")
	for _, t := range []float64{22.0, 23.5, 21.7, 24.1, 25.0, 26.3} {
		vreme.DodajMeritev(postaje, "KP01", t, minTemp, maxTemp)
	}

	vreme.DodajPostajo(postaje, "KP02", "Koper Center", "Koper")
	for _, t := range []float64{29.0, 31.5, 30.7, 27.1, 29.0, 28.3} {
		vreme.DodajMeritev(postaje, "KP02", t, minTemp, maxTemp)
	}

	vreme.DodajPostajo(postaje, "KR01", "Kredarica", "Triglav")
	for _, t := range []float64{5.0, 4.5, 6.2, 3.8, 2.1, 1.0, 0.5} {
		vreme.DodajMeritev(postaje, "KR01", t, minTemp, maxTemp)
	}

	vreme.DodajPostajo(postaje, "MB01", "Maribor Center", "Slovenija")
	for _, t := range []float64{15.0, 16.5, 14.7, 17.1, 18.0, 19.3} {
		vreme.DodajMeritev(postaje, "MB01", t, minTemp, maxTemp)
	}

	vreme.DodajPostajo(postaje, "MB02", "Maribor Tabor", "Slovenija")
	for _, t := range []float64{12.0, 13.5, 11.7, 14.1, 15.0, 16.3} {
		vreme.DodajMeritev(postaje, "MB02", t, minTemp, maxTemp)
	}

	vreme.DodajPostajo(postaje, "MB03", "Maribor Studenci", "Slovenija")
	for _, t := range []float64{10.0, 11.5, 9.7, 12.1, 13.0, 14.3} {
		vreme.DodajMeritev(postaje, "MB03", t, minTemp, maxTemp)
	}

	vreme.DodajPostajo(postaje, "CE01", "Celje Center", "Celje")
	for _, t := range []float64{14.2, 15.1} {
		vreme.DodajMeritev(postaje, "CE01", t, minTemp, maxTemp)
	}

	vreme.DodajMeritev(postaje, "LJ01", 57.0, minTemp, maxTemp) // Napaka! temperatura je izven mej
	vreme.DodajMeritev(postaje, "GO01", 30.0, minTemp, maxTemp) // Napaka! postaja s to oznako ne obstaja

	vreme.IzpisMeritev(postaje)

	fmt.Println("----------------------------------------------------")

	vreme.IzpisPorocila(postaje, stMeritev)

}

func main() {
	cmd := &cli.Command{
		Name:  "vremenske-postaje",
		Usage: "izpis meritev in poročila vremenskih postaj",
		Flags: []cli.Flag{
			&cli.Float64Flag{Name: "minTemp", Value: -50, Usage: "najnižja dovoljena temperatura meritve (°C)"},
			&cli.Float64Flag{Name: "maxTemp", Value: 50, Usage: "najvišja dovoljena temperatura meritve (°C)"},
			&cli.IntFlag{Name: "stMeritev", Value: 3, Usage: "najmanjše število meritev za izračun statistike"},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			zazeni(cmd.Float64("minTemp"), cmd.Float64("maxTemp"), cmd.Int("stMeritev"))
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
