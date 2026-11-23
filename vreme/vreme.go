package vreme

import (
	"fmt"
	"maps"
	"slices"
)

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {

	postaja, obstaja := postaje[oznaka]

	if obstaja == false {
		return 0, 0, 0, false
	}

	if len(postaja.meritve) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja = slices.Min(postaja.meritve)
	najvisja = slices.Max(postaja.meritve)

	sum := 0.0
	for _, num := range postaja.meritve {
		sum += num
	}

	povprecna = sum / float64(len(postaja.meritve))

	return najnizja, najvisja, povprecna, true
}

func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, minTemp float64, maxTemp float64) {

	if temperatura <= maxTemp && temperatura >= minTemp {

		if postaja, obstaja := postaje[oznaka]; obstaja {

			postaja.meritve = append(postaja.meritve, temperatura)

			postaje[oznaka] = postaja

		} else {
			fmt.Println("Napaka! postaja s to oznako ne obstaja")
		}

	} else {

		fmt.Println("Napaka! temperatura je izven mej")

	}

}

func IzpisMeritev(postaje map[string]Postaja) {

	fmt.Println("MERITVE:")
	kljuci := slices.Sorted(maps.Keys(postaje))

	for _, oznaka := range kljuci {

		postaja := postaje[oznaka]
		fmt.Printf("%s - %s (%s): ", oznaka, postaja.ime, postaja.kraj)
		fmt.Println(postaja.meritve)
	}

}

func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {

	kljuci := slices.Sorted(maps.Keys(postaje))

	for _, oznaka := range kljuci {

		postaja := postaje[oznaka]

		min, max, avg, ok := statistika(postaje, oznaka, stMeritev)

		fmt.Printf("%s: ", postaja.ime)

		if ok {
			fmt.Printf("min %.1f °C, max %.1f °C, povprečje %.1f °C ", min, max, avg)

			switch {

			case avg >= 25:
				fmt.Println("-> Vroče")
			case avg >= 15 && avg < 25:
				fmt.Println("-> Toplo")
			case avg >= 5 && avg < 15:
				fmt.Println("-> Hladno")
			case avg < 5:
				fmt.Println("-> Mrzlo")

			}

		} else {
			fmt.Println("premalo meritev")
		}

	}

}

func DodajPostajo(postaje map[string]Postaja, oznaka string, ime string, kraj string) {

	if _, obstaja := postaje[oznaka]; obstaja {

		fmt.Println("Napaka! Postaja s to oznako že obstaja")
	} else {
		postaje[oznaka] = Postaja{ime: ime, kraj: kraj}
	}

}
