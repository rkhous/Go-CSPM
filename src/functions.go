package functions

import (
	"./data"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func CheckIfAdmin(id string, listOfAdmins []string) bool {
	for _, n := range(listOfAdmins){
		if n == id{
			return true
		}
	}
	return false
}

func CheckIfPokemon(name string) bool {
	for _, n := range(pokeinfo.PokemonList){
		if strings.ToLower(name) == n{
			return true
		}
	}
	return false
}

func GrabStopInformation(name string) map[string]string {
	stopInformation := map[string]string{}
	f, err := os.Open("pokestops.csv")
	if err != nil{
		fmt.Println(err.Error())
	}
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if strings.Contains(strings.ToLower(record[0]), strings.ToLower(name)){
			stopInformation["name"] = record[0]
			stopInformation["lat,lon"] = record[1] + "," + record[2]
			stopInformation["img"] = record[3]
			return stopInformation
		}
	}
	return stopInformation
}

func SearchStops(name string) []string{
	stopsFound := []string{}
	f, err := os.Open("pokestops.csv")
	if err != nil{
		fmt.Println(err.Error())
	}
	r := csv.NewReader(bufio.NewReader(f))

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if strings.Contains(strings.ToLower(record[0]), strings.ToLower(name)){
			stopsFound = append(stopsFound, record[0] + "  |  " + record[1] + "," + record[2])
			}
		}

	return stopsFound
}