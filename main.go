package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/endobit/oui"
)

type addrStruct struct {
	IP        string
	Vendor    string
	Dashed    string
	Cisco     string
	Lowercase string
	Uppercase string
	netIP     net.IP
}

func main() {
	//Print hello
	fmt.Println("Welcome to mac. For more details visit https://github.com/Jamous/mac")
	//Get input
	inputs := parseInput()

	//Determin if a single input was passed.
	if len(inputs) == 1 {
		singleInput(inputs[0])

	} else {
		//Multipe inputs
		multipleInput(inputs)
	}
}

// Parse inputs
func parseInput() []string {
	var input []string

	//If onne value was wassed at runtime, assume this is the address
	if len(os.Args) >= 2 {
		input = append(input, os.Args[1])

	} else {
		//If no values were passed at runtime, prompt the user for single address, or multi address lookup
		fmt.Println("\nMac address or addresses to convert (press enter twice): ")
		reader := bufio.NewReader(os.Stdin)

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
			}

			// Trim whitespace from the line
			line = strings.TrimSpace(line)

			// If the line is empty, terminate input
			if line == "" {
				break
			}

			input = append(input, line)
		}
	}

	return input
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
	fmt.Println(macMap["uppercase_clean"])
}

func multipleInput(inputs []string) {
	var addrList []addrStruct

	//Iterate through inputs, ignore blank inputs
	for _, input := range inputs {
		if input != "" {
			//Find the mac address
			mac := findMac(input)

			//Find the IP address
			ip := findIp(input)

			//Stop processing if mac was not found found
			if mac != "" {
				//Convert to clean mac and find vendor
				macMap, clean := convertMac(mac)
				vendor := findVend(clean)

				//Add to addrList
				addrList = append(addrList, addrStruct{IP: ip, Vendor: vendor, Dashed: macMap["dashed"], Cisco: macMap["cisco"], Lowercase: macMap["lowercase"], Uppercase: macMap["uppercase"]})
			}
		}
	}

	//Sort results
	sortedList := sortResults(addrList)

	//Write results
	writeResults(sortedList)
}

func findMac(input string) string {
	//Search for standard mac, if found return results
	standardRegex := regexp.MustCompile(`\b([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})\b`)
	standard := standardRegex.FindString(input)
	if standard != "" {
		return standard
	}

	//Search for cisco mac, if found return results
	ciscoRegex := regexp.MustCompile(`\b([0-9A-Fa-f]{4}\.){2}[0-9A-Fa-f]{4}\b`)
	cisco := ciscoRegex.FindString(input)
	if cisco != "" {
		return cisco
	}

	return ""
}

func findIp(input string) string {
	//Search for IPv4 address
	v4Regex := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)
	v4 := v4Regex.FindString(input)
	if v4 != "" {
		return v4
	}

	//Search for IPv6 address
	v6Regex := regexp.MustCompile(`\b(?:[0-9a-fA-F]{1,4}::?){1,7}(?:[0-9a-fA-F]{1,4})?\b`)
	v6 := v6Regex.FindString(input)
	if v6 != "" {
		return v6
	}

	return ""
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
		"dashed":          fmt.Sprintf("%s-%s-%s-%s-%s-%s", input[:2], input[2:4], input[4:6], input[6:8], input[8:10], input[10:]),
		"cisco":           fmt.Sprintf("%s.%s.%s", input[:4], input[4:8], input[8:]),
		"lowercase":       fmt.Sprintf("%s:%s:%s:%s:%s:%s", input[:2], input[2:4], input[4:6], input[6:8], input[8:10], input[10:]),
		"uppercase":       strings.ToUpper(fmt.Sprintf("%s:%s:%s:%s:%s:%s", input[:2], input[2:4], input[4:6], input[6:8], input[8:10], input[10:])),
		"uppercase_clean": strings.ToUpper(input),
	}

	return macMap, input
}

func findVend(mac string) string {
	vendor := oui.Vendor(mac)

	return vendor
}

// Sorts results by IP address
func sortResults(addrList []addrStruct) []addrStruct {
	//We know the size of the slice, so lets go ahead and preallocate it.
	sortedList := make([]addrStruct, 0, len(addrList))

	//We will add NetIP to the slice
	for _, addr := range addrList {
		sortedList = append(sortedList, addrStruct{IP: addr.IP, Vendor: addr.Vendor, Dashed: addr.Dashed, Cisco: addr.Cisco, Lowercase: addr.Lowercase, Uppercase: addr.Uppercase, netIP: net.ParseIP(addr.IP)})
	}

	//Sort slice by NetIP
	sort.Slice(sortedList, func(i, j int) bool {
		return bytes.Compare(sortedList[i].netIP, sortedList[j].netIP) < 0
	})

	return sortedList
}

func writeResults(addrList []addrStruct) {
	//Print results
	for _, addr := range addrList {
		fmt.Printf("%s   %s   %s\n", addr.IP, addr.Lowercase, addr.Vendor)
	}

	//Convert to dataframe and output to csv
	//data := dataframe.LoadStructs(addrList)
	//out := makeCSV("mac.csv")
	//defer out.Close()
	//data.WriteCSV(out)
}

func makeCSV(csvName string) *os.File {
	//Check if file exists, if it does increment by 1 and try again
	//variabels
	var tempName, finalName string
	ncount := 0

	//Seperate file extension and name
	extension := csvName[strings.Index(csvName, "."):]
	name := csvName[:strings.Index(csvName, ".")]

	//Start a for loop to find an appropiate file name
	for {
		//Set tempName to filename on the first iteration
		if ncount == 0 {
			tempName = csvName
		} else {
			tempName = fmt.Sprintf("%s_%d%s", name, ncount, extension)
		}

		//Check if file exists, an error will indicate it does not, which is what we ant.
		_, err := os.Stat(tempName)
		if err != nil {
			//Set finalName and exit loop
			finalName = tempName
			break
		}

		//Increment ncount and try again
		ncount++
	}

	out, err := os.Create(finalName)
	if err != nil {
		errm := fmt.Sprintf("Failed to create csv file for %s, error: %s", csvName, err)
		log.Fatal(errm)
	}

	return out
}
