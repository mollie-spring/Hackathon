package enc_dec

import (
	"github.com/fentec-project/gofe/abe"
)



// func Encryption(a *abe.FAME, msg string, msp *abe.MSP, pubKey *abe.FAMEPubKey) (*abe.FAMECipher) {

// 	cipher, _ := a.Encrypt(msg, msp, pubKey)
// 	return cipher
// }

// func Decryption(a *abe.FAME, cipher *abe.FAMECipher, keys *abe.FAMEPubKey) (*abe.FAME) {

// 	dec, _ := a.Decrypt(cipher, keys, pubKey) // Decrypt the message
// 	return dec
// }


func PartipantEncrypts(a *abe.FAME, pubKey *abe.FAMEPubKey, mspAudit *abe.MSP, mspBreakGlass *abe.MSP, mspA *abe.MSP, mspB *abe.MSP, mspAll *abe.MSP) (*abe.FAMECipher, *abe.FAMECipher, *abe.FAMECipher, *abe.FAMECipher, *abe.FAMECipher) {
	
	msg := "Here's lookin' at you kid."

	// encrypt function by a participant
	cipher1, _ := a.Encrypt( msg, mspAudit, pubKey)

	// // encrypt function by a participant
	cipher2, _ := a.Encrypt( msg, mspBreakGlass, pubKey)

	// // encrypt function by a participant
	cipher3, _ := a.Encrypt(msg, mspA, pubKey)

	// // encrypt function by a participant
	cipher4, _ := a.Encrypt(msg, mspB, pubKey)

	// // encrypt function by a participant
	cipher5, _ := a.Encrypt( msg, mspAll, pubKey)

	return cipher1, cipher2, cipher3, cipher4, cipher5
}


