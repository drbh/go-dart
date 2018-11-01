package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func buildClasses() map[string][]string {
	classLookup := map[string][]string{}
	myClasses := []string{}
	prefixes := []string{
		"MKT", "CSI", "SPA", "CON",
		"ELC", "ARH", "MAT", "SOC",
		"PSY", "PHY", "BUS", "ENG",
	}
	for p := 0; p < len(prefixes); p++ {
		pre := prefixes[p]
		for i := 1; i < 10; i++ {
			final := pre + " " + strconv.Itoa(i) + "00"
			// fmt.Println(final)

			lku := classLookup[pre]
			lku = append(lku, final)
			classLookup[pre] = lku

			lku2 := classLookup[strconv.Itoa(i)+"00"]
			lku2 = append(lku2, final)
			classLookup[strconv.Itoa(i)+"00"] = lku2

			gsKey := "GS-A"
			if p%2 == 0 {
				gsKey = "GS-B"
			}
			if p%7 == 0 {
				gsKey = "GS-C"
			}

			lku3 := classLookup[gsKey]
			lku3 = append(lku3, final)
			classLookup[gsKey] = lku3

			myClasses = append(myClasses, final)
		}
	}
	// jsond, _ :=
	// json.MarshalIndent(myClasses, "", "\t")
	// fmt.Println(string(jsond))
	jsond, _ := json.Marshal(classLookup["GS-B"])
	fmt.Println(string(jsond))
	return classLookup
}

func main() {
	// lookup :=
	buildClasses()
	// fmt.Println(lookup["GS-A"])
	// fmt.Println("\n")
	// fmt.Println(lookup["GS-B"])
	// fmt.Println("\n")
	// fmt.Println(lookup["GS-C"])
	// fmt.Println(lookup)
}
