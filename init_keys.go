package main
import (
	"os"
	"github.com/fentec-project/gofe/abe" 
)

func setup()  {
	a := abe.NewFAME() // Create the scheme instance
	var pubKey, secKey, _ = a.GenerateMasterKeys()
	
	f_pub, __ := os.Create("keys/pubkey")
    defer f_pub.Close()
	n1, __ := f_pub.Write(pubKey)
	
	f_sec, __ := os.Create("keys/seckey")
    defer f_sec.Close()
    n2, __ := f_sec.WriteString(secKey)

	// policies should be activated here
	init_roles(secKey)

}


