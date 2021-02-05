package main

import (
    "fmt"
	"initialize"
	"enc_dec"
	"os"
	"bufio"
)



func main() {
	
	
	fmt.Println("Enter Your Intent (Setup, Encrypt, Decrypt): ") 
  
	
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
    intent := scanner.Text()

	
	if intent == "Setup" {

		// Scheme is established, roles assigned, policies made, public, secret and attribute-specific keys generated
		initialize.Setup()
		fmt.Printf("Attribute Based Encryption initialized\n")

	} else if intent == "Encrypt" {

		fmt.Printf("What would you like to encrypt?\n")		
		
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		plaintext := scanner.Text()

		fmt.Println("Enter the Policy (Audit, Break glass, Type A, Type B, All Types): ") 
  
    	// var then variable name then variable type  
	
    	scanner.Scan() // use `for scanner.Scan()` to keep reading
    	policy := scanner.Text()

		// A participant encrypts 1 message in various policies and submits to the manager
		enc_dec.PartipantEncrypts(plaintext, policy)

	} else if intent == "Decrypt" {

		fmt.Println("Enter Your Role: ") 
  
    	// var then variable name then variable type  

    	scanner.Scan() // use `for scanner.Scan()` to keep reading
    	role := scanner.Text()

		fmt.Printf("\n %s attempts to decrypt the ciphertext!\n", role)
		enc_dec.RoleDecrypts(role)
	
	}

	
}


