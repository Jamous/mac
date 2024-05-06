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
	//Get input
	input := parseInput()

	//Determin if a single input was passed, or a multiple imput. returns bool
	single := singleInputCheck(input)
	fmt.Println(single)

/*
	//If only one input, process as single input
	if len(inputs) == 1 {
		singleInput(inputs[0])
	} else {
	//Assume there are multipe inputs
		multipleInput(inputs)
	}
	*/
}

//Parse inputs
func parseInput() string {
	var input string
	
	//If onne value was wassed at runtime, assume this is the address
	if len(os.Args) >= 2 {
		input = os.Args[1]
	
	} else { 
	//If no values were passed at runtime, prompt the user for single address, or multi address lookup
		fmt.Println("\nMac address or addresses to convert: ")	
		reader := bufio.NewReader(os.Stdin)
		dinput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Could not get mac address from user.")
		}
		input = strings.TrimSpace(dinput)
	
	}

	return input
}

//Determin if a single input was passed, or a multiple imput.
func singleInputCheck(input string) bool {
	//Check how many new line characters are in the string
	lineCount := strings.Count(input, "\n")
	fmt.Println(lineCount)

	return false


}

func singleInput(input string) {
	//Convert to clean mac
	macMap, clean := convertMac(input)

	//Find the mac vendor
	vendor := findVend(clean)

	//Print results
	fmt.Println(vendor)
	fmt.Println(macMap["dashed"])
	fmt.Println(macMap["cisco"])
	fmt.Println(macMap["lowercase"])
	fmt.Println(macMap["uppercase"])
}

func multipleInput(inputs []string) {
	
	//Print results
	fmt.Println(inputs)

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
		"dashed":    fmt.Sprintf("%s-%s-%s-%s-%s-%s", input[:2], input[2:4], input[4:6], input[6:8], input[8:10], input[10:]),
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
