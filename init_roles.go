import (
	"os"
	"/Users/molliezechlin/go/src/github.com/fentec-project/gofe/abe"
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


	// Possible Attributes
	// 'Auditor' 'Part Type A' 'Part Type B' 'Manager'
	return manager_keys, auditor_keys, supervisor_keys, par1_keys, par2_keys, par3_keys
}


