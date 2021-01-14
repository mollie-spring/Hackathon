package main

import (
    "fmt"

	"initialize"
	//"github.com/fentec-project/gofe/abe"
	"enc_dec"
)




func main() {
	
	// Scheme is established, roles assigned, policies made, public, secret and attribute-specific keys generated
	scheme, pubKey, manager_keys, auditor_keys, supervisor_keys, par1_keys, par2_keys, par3_keys, mspAudit, mspBreakGlass, mspA, mspB, mspAll := initialize.Setup()
	
	
	// A participant encrypts 1 message in various policies and submits to the manager
	cipher1, cipher2, cipher3, cipher4, cipher5 := enc_dec.PartipantEncrypts(scheme, pubKey, mspAudit, mspBreakGlass, mspA, mspB, mspAll)
	


	// attempt to decrypt all ciphers by the Manager who shouldn't be able to
	fmt.Printf("\nThe manager attempts to decrypt all 5 ciphers! (None should be returned)\n")
	enc_dec.Decryption(scheme, manager_keys, pubKey, cipher1, cipher2, cipher3, cipher4, cipher5)

	// attempt to decrypt all ciphers by the Auditor who shouldn't be able to
	fmt.Printf("\nThe auditor attempts to decrypt all 5 ciphers! (should be able to decrypt 1)\n")
	enc_dec.Decryption(scheme, auditor_keys, pubKey, cipher1, cipher2, cipher3, cipher4, cipher5)

	// attempt to decrypt all ciphers by the Supervisor who shouldn't be able to
	fmt.Printf("\nThe supervisor (who possesses both auditor and manager attributes) attempts to decrypt all 5 ciphers! (should be able to decrypt 1, and 2)\n")
	enc_dec.Decryption(scheme, supervisor_keys, pubKey, cipher1, cipher2, cipher3, cipher4, cipher5)

	// attempt to decrypt all ciphers by the 1st Participant who shouldn't be able to
	fmt.Printf("\nThe 1st participant attempts to decrypt all 5 ciphers! (should be able to decrypt 1, 3, 4, and 5)\n")
	enc_dec.Decryption(scheme, par1_keys, pubKey, cipher1, cipher2, cipher3, cipher4, cipher5)

	// attempt to decrypt all ciphers by the 2nd Participant who shouldn't be able to
	fmt.Printf("\nThe 2nd participant attempts to decrypt all 5 ciphers! (should be able to decrypt 1, 3, and 5)\n")
	enc_dec.Decryption(scheme, par2_keys, pubKey, cipher1, cipher2, cipher3, cipher4, cipher5)

	// attempt to decrypt all ciphers by the 3rd Participant who shouldn't be able to
	fmt.Printf("\nThe 3rd Participant attempts to decrypt all 5 ciphers! (should be able to decrypt 1, 4, and 5)\n")
	enc_dec.Decryption(scheme, par3_keys, pubKey, cipher1, cipher2, cipher3, cipher4, cipher5)


}


