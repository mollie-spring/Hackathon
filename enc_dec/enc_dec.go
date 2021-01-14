package enc_dec

import (
	"github.com/fentec-project/gofe/abe"
	"fmt"
)




func PartipantEncrypts(a *abe.FAME, pubKey *abe.FAMEPubKey, mspAudit *abe.MSP, mspBreakGlass *abe.MSP, mspA *abe.MSP, mspB *abe.MSP, mspAll *abe.MSP) (*abe.FAMECipher, *abe.FAMECipher, *abe.FAMECipher, *abe.FAMECipher, *abe.FAMECipher) {
	
	msg1 := "Here's lookin' at you, kid."
	msg2 := "Frankly, my dear, I don't give a damn."
	msg3 := "May the Force be with you."
	msg4 := "You're a wizard, Harry."
	msg5 := "I see dead people."


	fmt.Printf("A participant encrypts 5 different strings in 5 different policies.\n\n")

	// encrypt function by a participant (auditor or any participants)
	cipher1, _ := a.Encrypt( msg1, mspAudit, pubKey)

	// // encrypt function by a participant (auditor and manager - supervisor only)
	cipher2, _ := a.Encrypt( msg2, mspBreakGlass, pubKey)

	// // encrypt function by a participant (Type A Par only! - Par 1 and 2)
	cipher3, _ := a.Encrypt(msg3, mspA, pubKey)

	// // encrypt function by a participant (Type B Par only - Par 1 and 3)
	cipher4, _ := a.Encrypt(msg4, mspB, pubKey)

	// // encrypt function by a participant (Either Participant Type - Par 1, 2, and 3)
	cipher5, _ := a.Encrypt( msg5, mspAll, pubKey)

	return cipher1, cipher2, cipher3, cipher4, cipher5
}


func Decryption(a *abe.FAME, attribKey *abe.FAMEAttribKeys, pubKey *abe.FAMEPubKey,
	cipher1 *abe.FAMECipher, cipher2 *abe.FAMECipher, cipher3 *abe.FAMECipher, cipher4 *abe.FAMECipher, cipher5 *abe.FAMECipher)	{

	// attempt to decrypt all ciphers using the given attribute-based key
	dec_cipher1, _ := a.Decrypt(cipher1, attribKey, pubKey)
	fmt.Printf("Decryption 1: %v \n", dec_cipher1)
	dec_cipher2, _ := a.Decrypt(cipher2, attribKey, pubKey)
	fmt.Printf("Decryption 2: %v \n", dec_cipher2)
	dec_cipher3, _ := a.Decrypt(cipher3, attribKey, pubKey)
	fmt.Printf("Decryption 3: %v \n", dec_cipher3)
	dec_cipher4, _ := a.Decrypt(cipher4, attribKey, pubKey)
	fmt.Printf("Decryption 4: %v \n", dec_cipher4)
	dec_cipher5, _ := a.Decrypt(cipher5, attribKey, pubKey)
	fmt.Printf("Decryption 5: %v \n", dec_cipher5)
}



