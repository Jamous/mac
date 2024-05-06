package main

import (
	"bufio"
	"fmt"
	"github.com/endobit/oui"
	"log"
	"os"
	"strings"
)

func main() {
	input := parseInput()

	//Convert to clean mac
	macMap, clean := convertMac(input)
	fmt.Println(macMap)

	//Find the mac vendor
	vendor := findVend(clean)

	//Print results
	fmt.Println(vendor)
	fmt.Println(macMap["cisco"])
	fmt.Println(macMap["lowercase"])
	fmt.Println(macMap["uppercase"])
}

func parseInput() string {
	var input string

	//Prompt if impot was not passed at runtime
	if len(os.Args) >= 2 {
		input = os.Args[1]
	} else {
		fmt.Println("\nMac address to convert: ")
		reader := bufio.NewReader(os.Stdin)
		dinput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Could not get mac address from user.")
		}
		input = strings.TrimSpace(dinput)
	}

	return input
}

func convertMac(input string) (map[string]string, string) {
	//Remove characters from string and convert to lowercase
	removeChars := []string{".", ":", "-", " "}
	for _, char := range removeChars {
		input = strings.Replace(input, char, "", -1)
	}
	input = strings.ToLower(input)

	//validate mac is a the correct lenght, otherwise throw error
	if len(input) != 12 {
		log.Fatalf("Input is not of the correct lenght. Expected 12 clean characters, received %d. %s", len(input), input)
	}

	//Convert to useable formats
	macMap := map[string]string{
		"cisco":     fmt.Sprintf("%s.%s.%s", input[:4], input[4:8], input[8:]),
		"lowercase": fmt.Sprintf("%s:%s:%s:%s:%s:%s", input[:2], input[2:4], input[4:6], input[6:8], input[8:10], input[10:]),
		"uppercase": strings.ToUpper(fmt.Sprintf("%s:%s:%s:%s:%s:%s", input[:2], input[2:4], input[4:6], input[6:8], input[8:10], input[10:])),
	}

	return macMap, input
}

func findVend(mac string) string {
	vendor := oui.Vendor(mac)

	return vendor
}
