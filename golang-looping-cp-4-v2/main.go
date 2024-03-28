package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {

	emailSplit := strings.Split(email, "@")

	splitTitik := strings.Split(emailSplit[1], ".")
	// fmt.Println(splitTitik)

	var TLD string
	if len(splitTitik) > 2 {
		TLD = strings.Join(splitTitik[1:], ".")
	} else {
		TLD = splitTitik[1]
	}

	return "Domain: " + splitTitik[0] + " dan TLD: " + TLD // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.com"))
	fmt.Println(EmailInfo("ptmencaricintasejati@gmail.co.id"))
}
