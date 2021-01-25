package enc_dec

import (
	"github.com/fentec-project/gofe/abe"
	"fmt"
	"os"
	"encoding/json"
)



func PartipantEncrypts() (*abe.FAMECipher, *abe.FAMECipher, *abe.FAMECipher, *abe.FAMECipher, *abe.FAMECipher) {
	
	msg1 := "Here's lookin' at you, kid."
	msg2 := "Frankly, my dear, I don't give a damn."
	msg3 := "May the Force be with you."
	msg4 := "You're a wizard, Harry."
	msg5 := "I see dead people."

	var pubKey *abe.FAMEPubKey
	pubKeyBytes := os.Getenv("PUBKEY")
	json.Unmarshal([]byte(pubKeyBytes), &pubKey)

	fmt.Printf("A participant encrypts 5 different strings in 5 different policies.\n\n")

	var mspAudit *abe.MSP
	mspAuditBytes := os.Getenv("MSP_AUDIT")
	json.Unmarshal([]byte(mspAuditBytes), &mspAudit)

	var mspBreakGlass *abe.MSP
	mspBreakGlassBytes := os.Getenv("MSP_BREAKGLASS")
	json.Unmarshal([]byte(mspBreakGlassBytes), &mspBreakGlass)

	var mspA *abe.MSP
	mspABytes := os.Getenv("MSP_A")
	json.Unmarshal([]byte(mspABytes), &mspA)

	var mspB *abe.MSP
	mspBBytes := os.Getenv("MSP_B")
	json.Unmarshal([]byte(mspBBytes), &mspB)

	var mspAll *abe.MSP
	mspAllBytes := os.Getenv("MSP_ALL")
	json.Unmarshal([]byte(mspAllBytes), &mspAll)

	var scheme *abe.FAME
	schemeBytes := os.Getenv("SCHEME")
	json.Unmarshal([]byte(schemeBytes), &scheme)

	// encrypt function by a participant (auditor or any participants)
	cipher1, _ := scheme.Encrypt(msg1, mspAudit, pubKey)

	// // encrypt function by a participant (auditor and manager - supervisor only)
	cipher2, _ := scheme.Encrypt(msg2, mspBreakGlass, pubKey)

	// // encrypt function by a participant (Type A Par only! - Par 1 and 2)
	cipher3, _ := scheme.Encrypt(msg3, mspA, pubKey)

	// // encrypt function by a participant (Type B Par only - Par 1 and 3)
	cipher4, _ := scheme.Encrypt(msg4, mspB, pubKey)

	// // encrypt function by a participant (Either Participant Type - Par 1, 2, and 3)
	cipher5, _ := scheme.Encrypt(msg5, mspAll, pubKey)

	return cipher1, cipher2, cipher3, cipher4, cipher5
}


func Decryption(attribKey *abe.FAMEAttribKeys,
	cipher1 *abe.FAMECipher, cipher2 *abe.FAMECipher, cipher3 *abe.FAMECipher, cipher4 *abe.FAMECipher, cipher5 *abe.FAMECipher)	{

	var scheme *abe.FAME
	schemeBytes := os.Getenv("SCHEME")
	json.Unmarshal([]byte(schemeBytes), &scheme)

	var pubKey *abe.FAMEPubKey
	pubKeyBytes := os.Getenv("PUBKEY")
	json.Unmarshal([]byte(pubKeyBytes), &pubKey)
	
	// attempt to decrypt all ciphers using the given attribute-based key
	dec_cipher1, _ := scheme.Decrypt(cipher1, attribKey, pubKey)
	fmt.Printf("Decryption 1: %v \n", dec_cipher1)
	dec_cipher2, _ := scheme.Decrypt(cipher2, attribKey, pubKey)
	fmt.Printf("Decryption 2: %v \n", dec_cipher2)
	dec_cipher3, _ := scheme.Decrypt(cipher3, attribKey, pubKey)
	fmt.Printf("Decryption 3: %v \n", dec_cipher3)
	dec_cipher4, _ := scheme.Decrypt(cipher4, attribKey, pubKey)
	fmt.Printf("Decryption 4: %v \n", dec_cipher4)
	dec_cipher5, _ := scheme.Decrypt(cipher5, attribKey, pubKey)
	fmt.Printf("Decryption 5: %v \n", dec_cipher5)

	
}


func RoleDecrypts(role string, 	cipher1 *abe.FAMECipher, cipher2 *abe.FAMECipher, cipher3 *abe.FAMECipher, cipher4 *abe.FAMECipher, cipher5 *abe.FAMECipher)	{

	if role == "Manager" {
		var managerKeys	*abe.FAMEAttribKeys
		managerKeysBytes := os.Getenv("MANAGERKEYS")
		json.Unmarshal([]byte(managerKeysBytes), &managerKeys)

		Decryption(managerKeys, cipher1, cipher2, cipher3, cipher4, cipher5)
	
	} else if role == "Auditor" {
		var auditorKeys	*abe.FAMEAttribKeys
		auditorKeysBytes := os.Getenv("AUDITORKEYS")
		json.Unmarshal([]byte(auditorKeysBytes), &auditorKeys)

		Decryption(auditorKeys, cipher1, cipher2, cipher3, cipher4, cipher5)
	
	} else if role == "Supervisor" {
		var supervisorKeys	*abe.FAMEAttribKeys
		supervisorKeysBytes := os.Getenv("SUPERVISORKEYS")
		json.Unmarshal([]byte(supervisorKeysBytes), &supervisorKeys)

		Decryption(supervisorKeys, cipher1, cipher2, cipher3, cipher4, cipher5)
	
	} else if role == "Participant 1" {
		var par1Keys	*abe.FAMEAttribKeys
		par1KeysBytes := os.Getenv("PAR1KEYS")
		json.Unmarshal([]byte(par1KeysBytes), &par1Keys)

		Decryption(par1Keys, cipher1, cipher2, cipher3, cipher4, cipher5)
	
	} else if role == "Participant 2" {
		var par2Keys	*abe.FAMEAttribKeys
		par2KeysBytes := os.Getenv("PAR2KEYS")
		json.Unmarshal([]byte(par2KeysBytes), &par2Keys)

		Decryption(par2Keys, cipher1, cipher2, cipher3, cipher4, cipher5)
	
	} else if role == "Participant 3" {
		var par3Keys	*abe.FAMEAttribKeys
		par3KeysBytes := os.Getenv("PAR3KEYS")
		json.Unmarshal([]byte(par3KeysBytes), &par3Keys)
		
		Decryption(par3Keys, cipher1, cipher2, cipher3, cipher4, cipher5)
	
	} else {
		fmt.Printf("%s is not a valid role!\n", role)
	}
}
