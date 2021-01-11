import (
	"os"
	"/Users/molliezechlin/go/src/github.com/fentec-project/gofe/abe"
	"init_roles"
	"policies"
)


func main()  {
	
	pubKey, secKey, manager_keys, auditor_keys, supervisor_keys, par1_keys, par2_keys, par3_keys = setup()
	// eventually make it so all the keys are written to files so that the main program doesn't 'see' them

	// encrypt function by a participant

	// attempt to decrypt by a different participant

	// attempt to decrypt by an auditor
	


}


