package main

import (
    "fmt"

	"initialize"
	//"github.com/fentec-project/gofe/abe"
	"enc_dec"
)

/* func main() {
    // Get a greeting message and print it.
    message := initialize.Hello("Gladys")
	fmt.Println(message)
	


}
 */


func main() {
	
	scheme, pubKey, manager_keys, auditor_keys, supervisor_keys, mspAudit, mspBreakGlass, mspA, mspB, mspAll := initialize.Setup()
	// eventually make it so all the keys are written to files so that the main program doesn't 'see' them
	
	// A participant encrypts 1 message in various policies and submits to the manager
	cipher1, cipher2, cipher3, cipher4, cipher5 := enc_dec.PartipantEncrypts(scheme, pubKey, mspAudit, mspBreakGlass, mspA, mspB, mspAll)


	// Keys for each role assigned
	
	
	// attempt to decrypt by participant1
	
	dec1_cipher1, _ := scheme.Decrypt(cipher1, manager_keys, pubKey)
	dec2_cipher1, _ := scheme.Decrypt(cipher1, auditor_keys, pubKey)
	dec3_cipher1, _ := scheme.Decrypt(cipher1, supervisor_keys, pubKey)
	fmt.Printf("%v This should not decrypt for the manager\n", dec1_cipher1)
	fmt.Printf("%v This should decrypt for the auditor\n", dec2_cipher1)
	fmt.Printf("%v This should decrypt for the supervisor\n", dec3_cipher1)

	dec1_cipher2, _ := scheme.Decrypt(cipher2, manager_keys, pubKey)
	dec2_cipher2, _ := scheme.Decrypt(cipher2, auditor_keys, pubKey)
	dec3_cipher2, _ := scheme.Decrypt(cipher2, supervisor_keys, pubKey)
	fmt.Printf("%v This should not decrypt for the manager\n", dec1_cipher2)
	fmt.Printf("%v This should not decrypt for the auditor\n", dec2_cipher2)
	fmt.Printf("%v This should decrypt for the supervisor\n", dec3_cipher2)

	dec1_cipher3, _ := scheme.Decrypt(cipher3, manager_keys, pubKey)
	dec2_cipher3, _ := scheme.Decrypt(cipher3, auditor_keys, pubKey)
	dec3_cipher3, _ := scheme.Decrypt(cipher3, supervisor_keys, pubKey)
	fmt.Printf("%v This should not decrypt for the manager\n", dec1_cipher3)
	fmt.Printf("%v This should not decrypt for the auditor\n", dec2_cipher3)
	fmt.Printf("%v This should not decrypt for the supervisor\n", dec3_cipher3)

	dec1_cipher4, _ := scheme.Decrypt(cipher4, manager_keys, pubKey)
	dec2_cipher4, _ := scheme.Decrypt(cipher4, auditor_keys, pubKey)
	dec3_cipher4, _ := scheme.Decrypt(cipher4, supervisor_keys, pubKey)
	fmt.Printf("%v This should not decrypt for the manager\n", dec1_cipher4)
	fmt.Printf("%v This should not decrypt for the auditor\n", dec2_cipher4)
	fmt.Printf("%v This should not decrypt for the supervisor\n", dec3_cipher4)

	dec1_cipher5, _ := scheme.Decrypt(cipher5, manager_keys, pubKey)
	dec2_cipher5, _ := scheme.Decrypt(cipher5, auditor_keys, pubKey)
	dec3_cipher5, _ := scheme.Decrypt(cipher5, supervisor_keys, pubKey)
	fmt.Printf("%v This should not decrypt for the manager\n", dec1_cipher5)
	fmt.Printf("%v This should not decrypt for the auditor\n", dec2_cipher5)
	fmt.Printf("%v This should not decrypt for the supervisor\n", dec3_cipher5)
	
}


