import (
	"policies"
)

func partipantEncrypts() {
	
	msg = "Here's lookin' at you kid."
	
	// encrypt function by a participant
	ciphertext1 = encrypt(msg, mspAudit) 

	// encrypt function by a participant
	ciphertext2 = encrypt(msg, mspBreakGlass) 

	// encrypt function by a participant
	ciphertext3 = encrypt(msg, mspA) 

	// encrypt function by a participant
	ciphertext4 = encrypt(msg, mspB) 

	// encrypt function by a participant
	ciphertext5 = encrypt(msg, mspAll) 

}
