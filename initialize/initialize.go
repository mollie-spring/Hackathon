package initialize

import (
	"github.com/fentec-project/gofe/abe" 
)



func Setup() (*abe.FAME, *abe.FAMEPubKey, 
	*abe.FAMEAttribKeys, *abe.FAMEAttribKeys, *abe.FAMEAttribKeys, *abe.FAMEAttribKeys, *abe.FAMEAttribKeys, *abe.FAMEAttribKeys, 
	*abe.MSP, *abe.MSP, *abe.MSP, *abe.MSP, *abe.MSP) {
	
	a := abe.NewFAME() // Create the scheme instance
	
	// Generate public and secret keys for instance
	pubKey, secKey, _ := a.GenerateMasterKeys()

	// Assign roles (or attributes) to each 'person' engaging in network
	manager := []string{"Manager"} 
	auditor := []string{"Auditor"} 
	supervisor := []string{"Manager", "Auditor"}
	participant1 := []string{"Part Type A", "Part Type B"}
	participant2 := []string{"Part Type A"}
	participant3 := []string{"Part Type B"}

	// Keys for each role assigned
	manager_keys, _ := a.GenerateAttribKeys(manager, secKey) // Generate keys with attributes of a manager
	auditor_keys, _ := a.GenerateAttribKeys(auditor, secKey) // Generate keys with attributes of an auditor
	supervisor_keys, _ := a.GenerateAttribKeys(supervisor, secKey) // Generate keys with attributes of a supervisor
	par1_keys, _ := a.GenerateAttribKeys(participant1, secKey) // Generate keys with attributes of Participant1
	par2_keys, _ := a.GenerateAttribKeys(participant2, secKey) // Generate keys with attributes of Participant2
	par3_keys, _ := a.GenerateAttribKeys(participant3, secKey) // Generate keys with attributes of Participant3
	
	mspAudit, mspBreakGlass, mspA, mspB, mspAll := init_policies()

	return a, pubKey, manager_keys, auditor_keys, supervisor_keys, par1_keys, par2_keys, par3_keys, mspAudit, mspBreakGlass, mspA, mspB, mspAll
}



func init_policies() (*abe.MSP, *abe.MSP, *abe.MSP, *abe.MSP, *abe.MSP) {
	// Policy List
	policy_audit := "Auditor OR Part Type A OR Part Type B"
	policy_breakGlass := "Auditor AND Manager"
	policy_a := "Part Type A"
	policy_b := "Part Type B"
	policy_all := "Part Type A OR Part Type B"


	// MSP structures for all policies
	mspAudit, _ := abe.BooleanToMSP(policy_audit, false) // Defining the policy allowing for an auditor or either type of Participant to see decrypt
	mspBreakGlass, _ := abe.BooleanToMSP(policy_breakGlass, false) // The MSP structure defining the policy for an extreme audit where a Manager/Auditor role is needed
	mspA, _ := abe.BooleanToMSP(policy_a, false) // The MSP structure defining the policy for Participant A
	mspB, _ := abe.BooleanToMSP(policy_b, false) // The MSP structure defining the policy for Participant B
	mspAll, _ := abe.BooleanToMSP(policy_all, false) // The MSP structure defining the policy for both Type A and B Participants

	return mspAudit, mspBreakGlass, mspA, mspB, mspAll
}

