import (
	"os"
	"/Users/molliezechlin/go/src/github.com/fentec-project/gofe/abe"
	"seckey"
)


func init_roles(secKey)  {
	
	// Attributes assigned to roles
	manager := []string{"Manager"} 
	auditor := []string{"Auditor"} 
	supervisor := []string{"Manager", "Auditor"}
	participant1 := []string{"Part Type A", "Part Type B"}
	participant2 := []string{"Part Type A"}
	participant3 := []string{"Part Type B"}

	// Keys for each role assigned
	
	manager_keys, _ := a.GenerateAttribKeys(manager, secKey) // Generate keys for the entity with attributes of a manager
	auditor_keys, _ := a.GenerateAttribKeys(auditor, secKey) // Generate keys for the entity with attributes of an auditor
	supervisor_keys, _ := a.GenerateAttribKeys(supervisor, secKey) // Generate keys for the entity with attributes of a supervisor
	par1_keys, _ := a.GenerateAttribKeys(participant1, secKey) // Generate keys for the entity with attributes of Participant1
	par2_keys, _ := a.GenerateAttribKeys(participant2, secKey) // Generate keys for the entity with attributes of Participant2
	par3_keys, _ := a.GenerateAttribKeys(participant3, secKey) // Generate keys for the entity with attributes of Participant1
	
	// Write keys to files
	f_man, __ := os.Create("/keys/mankey")
    defer f_man.Close()
	n1, __ := f_man.WriteString(mankey)
	
	f_aud, __ := os.Create("/keys/audkey")
    defer f_aud.Close()
    n2, __ := f_aud.WriteString(audkey)

	f_sup, __ := os.Create("/keys/supkey")
    defer f_sup.Close()
	n3, __ := f_sup.WriteString(supkey)
	
	f_par1, __ := os.Create("/keys/par1key")
    defer f_par1.Close()
	n4, __ := f_par1.WriteString(par1key)
	
	f_par2, __ := os.Create("/keys/par2key")
    defer f_par2.Close()
	n5, __ := f_par2.WriteString(par2key)
	
	f_par3, __ := os.Create("/keys/par3key")
    defer f_par3.Close()
    n6, __ := f_par3.WriteString(par3key)


	// Possible Attributes
	// 'Auditor' 'Part Type A' 'Part Type B' 'Manager'
}


