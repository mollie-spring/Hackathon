package enc_dec

import (
	"github.com/fentec-project/gofe/abe"
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)



func PartipantEncrypts(msg string, policy string) {
	
	var pubKey *abe.FAMEPubKey
	pubKeyBytes, _ := ioutil.ReadFile("../variables/pubkey")
	json.Unmarshal([]byte(pubKeyBytes), &pubKey)

	var scheme *abe.FAME
	schemeBytes, _ := ioutil.ReadFile("../variables/scheme")
	json.Unmarshal([]byte(schemeBytes), &scheme)

	var msp *abe.MSP

	if policy == "Audit" {
		mspBytes, _ := ioutil.ReadFile("../policies/mspaudit")
		json.Unmarshal([]byte(mspBytes), &msp)
	} else if policy == "Break glass" {
		mspBytes, _ := ioutil.ReadFile("../policies/mspbreakglass")
		json.Unmarshal([]byte(mspBytes), &msp)
	} else if policy == "Type A" {
		mspBytes, _ := ioutil.ReadFile("../policies/msp_a")
		json.Unmarshal([]byte(mspBytes), &msp)
	} else if policy == "Type B" {
		mspBytes, _ := ioutil.ReadFile("../policies/msp_b")
		json.Unmarshal([]byte(mspBytes), &msp)
	} else if policy == "All Types" {
		mspBytes, _ := ioutil.ReadFile("../policies/msp_all")
		json.Unmarshal([]byte(mspBytes), &msp)
	} else {
		fmt.Printf("%s is not a valid policy!\n", policy)
	}

	// encrypt function by a participant (auditor or any participants)
	ciphertext, _ := scheme.Encrypt(msg, msp, pubKey)

	cipherBytes, _ := json.Marshal(ciphertext)

	f, _ := os.Create("ciphertext")
    defer f.Close()
    f.Write(cipherBytes)

}


func Decryption(attribKey *abe.FAMEAttribKeys)	{

	var scheme *abe.FAME
	schemeBytes, _ := ioutil.ReadFile("../variables/scheme")
	json.Unmarshal([]byte(schemeBytes), &scheme)

	var pubKey *abe.FAMEPubKey
	pubKeyBytes, _ := ioutil.ReadFile("../variables/pubkey")
	json.Unmarshal([]byte(pubKeyBytes), &pubKey)
	
	var ciphertext *abe.FAMECipher
	cipherBytes, _ := ioutil.ReadFile("ciphertext")
	json.Unmarshal([]byte(cipherBytes), &ciphertext)

	// attempt to decrypt all ciphers using the given attribute-based key
	dec_cipher, _ := scheme.Decrypt(ciphertext, attribKey, pubKey)
	fmt.Printf("Decryption: %v \n", dec_cipher)
	
}


func RoleDecrypts(role string)	{	

	if role == "Manager 1" {
		var managerKeys	*abe.FAMEAttribKeys
		managerKeysBytes, _ := ioutil.ReadFile("../variables/managerkeys")
		json.Unmarshal([]byte(managerKeysBytes), &managerKeys)

		Decryption(managerKeys)
	
	} else if role == "Auditor 1" {
		var auditorKeys	*abe.FAMEAttribKeys
		auditorKeysBytes, _ := ioutil.ReadFile("../variables/auditorkeys")
		json.Unmarshal([]byte(auditorKeysBytes), &auditorKeys)

		Decryption(auditorKeys)
	
	} else if role == "Supervisor 1" {
		var supervisorKeys	*abe.FAMEAttribKeys
		supervisorKeysBytes, _ := ioutil.ReadFile("../variables/supervisorkeys")
		json.Unmarshal([]byte(supervisorKeysBytes), &supervisorKeys)

		Decryption(supervisorKeys)
	
	} else if role == "Participant 1" {
		var par1Keys	*abe.FAMEAttribKeys
		par1KeysBytes, _ := ioutil.ReadFile("../variables/par1keys")
		json.Unmarshal([]byte(par1KeysBytes), &par1Keys)

		Decryption(par1Keys)
	
	} else if role == "Participant 2" {
		var par2Keys	*abe.FAMEAttribKeys
		par2KeysBytes, _ := ioutil.ReadFile("../variables/par2keys")
		json.Unmarshal([]byte(par2KeysBytes), &par2Keys)

		Decryption(par2Keys)
	
	} else if role == "Participant 3" {
		var par3Keys	*abe.FAMEAttribKeys
		par3KeysBytes, _ := ioutil.ReadFile("../variables/par3keys")
		json.Unmarshal([]byte(par3KeysBytes), &par3Keys)
		
		Decryption(par3Keys)
	
	} else {
		fmt.Printf("%s is not a valid role!\n", role)
	}
}
