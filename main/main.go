package main

import (
    "fmt"
	"initialize"
	"enc_dec"
	"os"
	"bufio"
)



func main() {
	
	// Scheme is established, roles assigned, policies made, public, secret and attribute-specific keys generated
	initialize.Setup()
	

	
	// A participant encrypts 1 message in various policies and submits to the manager
	cipher1, cipher2, cipher3, cipher4, cipher5 := enc_dec.PartipantEncrypts()
	

	fmt.Println("Enter Your First Name: ") 
  
    // var then variable name then variable type  
	
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
    role := scanner.Text()

	fmt.Printf("\n %s attempts to decrypt all 5 ciphers!\n", role)
	enc_dec.RoleDecrypts(role, cipher1, cipher2, cipher3, cipher4, cipher5)

}


