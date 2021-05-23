package main

import (
	"fmt"
	"os"
)

type localeSettings struct {
	language string
	locale   string
}

func initLocales() []localeSettings {
	loc1 := localeSettings{"EN", "en_US"}
	loc2 := localeSettings{"EN", "en_CN"}
	loc3 := localeSettings{"EN", "en_AU"}
	loc4 := localeSettings{"FR", "fr_CN"}
	loc5 := localeSettings{"FR", "fr_FR"}
	loc6 := localeSettings{"FR", "fr_BE"}
	loc7 := localeSettings{"FR", "fr_CH"}
	loc8 := localeSettings{"RU", "ru_RU"}
	loc9 := localeSettings{"RU", "ru_BR"}

	localities := []localeSettings{
		loc1,
		loc2,
		loc3,
		loc4,
		loc5,
		loc6,
		loc7,
		loc8,
		loc9,
	}
	return localities
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Locale code not passed")
		os.Exit(1)
	}

	selected := os.Args[1]

	locales := initLocales()
	supported := false
	for i := 0; i < len(locales); i++ {
		if selected == locales[i].locale {
			supported = true
			break
		}
	}

	if supported == true {
		fmt.Println("Locale passed is supported:", selected)
	} else {
		fmt.Println("Locale not supported:", selected)
	}
}
