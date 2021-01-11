import (
	"os"
	"/Users/molliezechlin/go/src/github.com/fentec-project/gofe/abe"
	"init_roles"
	"policies"
)


func setup()  {
	a := abe.NewFAME() // Create the scheme instance
	pubKey, secKey, _ := a.GenerateMasterKeys()
	// policies should be activated here
	manager_keys, auditor_keys, supervisor_keys, par1_keys, par2_keys, par3_keys := init_roles(secKey)
	
	return pubKey, secKey, manager_keys, auditor_keys, supervisor_keys, par1_keys, par2_keys, par3_keys

}


