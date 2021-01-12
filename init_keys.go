import (
	"os"
	"/Users/molliezechlin/go/src/github.com/fentec-project/gofe/abe"
	"init_roles"
	"policies"
)


func setup()  {
	a := abe.NewFAME() // Create the scheme instance
	pubKey, secKey, _ := a.GenerateMasterKeys()

	f_pub, __ := os.Create("/keys/pubkey")
    defer f_pub.Close()
	n1, __ := f_pub.WriteString(pubkey)
	
	f_sec, __ := os.Create("/keys/seckey")
    defer f_sec.Close()
    n2, __ := f_sec.WriteString(seckey)

	// policies should be activated here
	init_roles(secKey)

}


