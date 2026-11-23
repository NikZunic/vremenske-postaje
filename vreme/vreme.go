// Package vreme omogoča delo z vremenskimi postajami in njihovimi meritvami temperature.
//
// Postaje so shranjene v slovarju tipa map[string]Postaja, kjer je ključ
// oznaka postaje (na primer "LJ01"). Postajo ustvarimo s funkcijo [DodajPostajo],
// meritve pa ji dodajamo s funkcijo [DodajMeritev]. Funkcija [IzpisMeritev]
// izpiše vse zbrane meritve, funkcija [IzpisPorocila] pa za vsako postajo
// izpiše najnižjo, najvišjo in povprečno temperaturo.
//
// Primer uporabe:
//
//	postaje := make(map[string]vreme.Postaja)
//	vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Slovenija")
//	vreme.DodajMeritev(postaje, "LJ01", 20.5, -50, 50)
//	vreme.DodajMeritev(postaje, "LJ01", 21.0, -50, 50)
//	vreme.DodajMeritev(postaje, "LJ01", 19.8, -50, 50)
//	vreme.IzpisMeritev(postaje)
//	vreme.IzpisPorocila(postaje, 3)
package vreme

import (
	"fmt"
	"maps"
	"slices"
)

// Postaja predstavlja vremensko postajo z imenom, krajem in seznamom
// izmerjenih temperatur v °C.
//
// Polja strukture niso izvožena, zato postajo ustvarimo s funkcijo
// [DodajPostajo], meritve pa ji dodajamo s funkcijo [DodajMeritev].
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// statistika vrne najnižjo, najvišjo in povprečno temperaturo postaje z dano oznako.
// Vrednost ok je false, če postaja ne obstaja ali ima manj kot stMeritev meritev.
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

// DodajMeritev doda temperaturo v °C k meritvam postaje z dano oznako.
//
// Meritev je dodana le, če je temperatura med minTemp in maxTemp (vključno z
// mejama). Če je temperatura izven mej ali postaja z dano oznako ne obstaja,
// funkcija izpiše napako in meritve ne doda.
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

// IzpisMeritev izpiše oznako, ime, kraj in vse meritve vsake postaje.
// Postaje so izpisane urejene po oznaki.
func IzpisMeritev(postaje map[string]Postaja) {

	fmt.Println("MERITVE:")
	kljuci := slices.Sorted(maps.Keys(postaje))

	for _, oznaka := range kljuci {

		postaja := postaje[oznaka]
		fmt.Printf("%s - %s (%s): ", oznaka, postaja.ime, postaja.kraj)
		fmt.Println(postaja.meritve)
	}

}

// IzpisPorocila za vsako postajo izpiše najnižjo, najvišjo in povprečno
// temperaturo ter oceno glede na povprečje:
//
//   - Vroče, če je povprečje vsaj 25 °C,
//   - Toplo, če je povprečje vsaj 15 °C in manj kot 25 °C,
//   - Hladno, če je povprečje vsaj 5 °C in manj kot 15 °C,
//   - Mrzlo, če je povprečje manj kot 5 °C.
//
// Če ima postaja manj kot stMeritev meritev, namesto statistike izpiše
// "premalo meritev". Postaje so izpisane urejene po oznaki.
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

// DodajPostajo v slovar postaje doda novo postajo z dano oznako, imenom in
// krajem. Nova postaja nima meritev.
//
// Če postaja z dano oznako že obstaja, funkcija izpiše napako in obstoječe
// postaje ne spremeni.
func DodajPostajo(postaje map[string]Postaja, oznaka string, ime string, kraj string) {

	if _, obstaja := postaje[oznaka]; obstaja {

		fmt.Println("Napaka! Postaja s to oznako že obstaja")
	} else {
		postaje[oznaka] = Postaja{ime: ime, kraj: kraj}
	}

}
