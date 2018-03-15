package main

import(
"fmt"
"net/http"
"net/http/httputil"
"regexp"
"strconv"
"math"
"time"
)

/*
TODO:
- Regex: allow no minutes or no hours
- Markdown help: fixer souci echappement backtips/chevron
*/

func RequestDebug(r *http.Request) {
	fmt.Printf("\n\n--------------------------\n\n")	
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
	fmt.Printf("\n\n--------------------------\n\n")	
}

func ConvertDuration(s string) float64 {

	var re = regexp.MustCompile(`(?i)(?P<hour>[0-9]{1,2})h(?P<minute>[0-9]{1,2})m`)
	found := re.FindStringSubmatch(s)

	//clean leading zeros
	for i, f := range found {
		if i == 0 {
			continue
		}

		if f[:1] == "0"{
			found[i] = f[1:]
		}
	}

	//set parts
	parts := map[string]int{}
	for k, v := range re.SubexpNames(){
		parts[v] = k
	}

	hour, _ := strconv.ParseFloat(found[parts["hour"]], 64)
	minute, _ := strconv.ParseFloat(found[parts["minute"]], 64)
	var converted float64
	converted = ((hour * 60) + minute) / 60

	return converted
}

func ConvertFloat(f float64) string {

	intpart, floatpart := math.Modf(f) //float parts (int and divider)
	hours := strconv.FormatFloat(intpart, 'f', 0, 64)
	tmp := math.Ceil(floatpart * 60) //minus one if 60 to avoid 1:60 e.g
	if tmp == 60{
		tmp--
	}	
	minutes:= strconv.FormatFloat(tmp, 'f', 0, 64)
	converted := hours + "h" + minutes + "m"

	return converted
}

func GetHelp() string{
	var text string
	text = "## Un peu d'aide sur les commandes possibles :\n"
	text += "|Je veux...|Commande|Arguments requis|Arguments optionnels|Exemple|\n"
	text += "|:-|:-|:-|:-|:-|:-|:-|\n"
	text += "|**Lister mes temps**|`ls`|-|`[yyyy-mm-dd]`|`/mtm ls` ou `/mtm ls 2018-03-15`|\n"
	text += "|**Ajouter une tâche**|`add`|`<duree>` et `<task>`|`[yyyy-mm-dd]`|`/mtm add 3h10m \"Refonte du front\"` ou `/mtm add 01h5m \"Tests unitaires\" 2018-03-23`|\n"
	text += "|**Supprimer un temps saisi**|`rm`|`<id>`|-|`/mtm rm 5aaa23831d6ef46644aec592`|\n"
	text += "|**Supprimer tous mes temps du jour**|`clear`|-|-|`/mtm clear`|\n"
	text += "|**Obtenir de l'aide**|`help`|-|-|`/mtm help`|\n"

	return text
}

func GetTodayDate() string{
	date := time.Now().Format("2006-01-02") // today default

	return date
}