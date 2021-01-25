package initialize

import (
	"github.com/fentec-project/gofe/abe" 
	//"bytes"
	"encoding/json"
	//"fmt"
	"os"
)



func Setup() {
	
	scheme := abe.NewFAME() // Create the scheme instance
	schemeBytes, _ := json.Marshal(scheme)
	os.Setenv("SCHEME", string(schemeBytes))

	// Generate public and secret keys for instance
	pubKey, secKey, _ := scheme.GenerateMasterKeys()

	pubKeyBytes, _ := json.Marshal(pubKey)
	os.Setenv("PUBKEY", string(pubKeyBytes))

	secKeyBytes, _ := json.Marshal(secKey)
	os.Setenv("SECKEY", string(secKeyBytes))


	// Assign roles (or attributes) to each 'person' engaging in network
	manager := []string{"Manager"} 
	auditor := []string{"Auditor"} 
	supervisor := []string{"Manager", "Auditor"}
	participant1 := []string{"Part Type A", "Part Type B"}
	participant2 := []string{"Part Type A"}
	participant3 := []string{"Part Type B"}

	// Keys for each role assigned
	managerKeys, _ := scheme.GenerateAttribKeys(manager, secKey) // Generate keys with attributes of a manager
	auditorKeys, _ := scheme.GenerateAttribKeys(auditor, secKey) // Generate keys with attributes of an auditor
	supervisorKeys, _ := scheme.GenerateAttribKeys(supervisor, secKey) // Generate keys with attributes of a supervisor
	par1Keys, _ := scheme.GenerateAttribKeys(participant1, secKey) // Generate keys with attributes of Participant1
	par2Keys, _ := scheme.GenerateAttribKeys(participant2, secKey) // Generate keys with attributes of Participant2
	par3Keys, _ := scheme.GenerateAttribKeys(participant3, secKey) // Generate keys with attributes of Participant3
	
	managerKeysBytes, _ := json.Marshal(managerKeys)
	os.Setenv("MANAGERKEYS", string(managerKeysBytes))

	auditorKeysBytes, _ := json.Marshal(auditorKeys)
	os.Setenv("AUDITORKEYS", string(auditorKeysBytes))
	
	supervisorKeysBytes, _ := json.Marshal(supervisorKeys)
	os.Setenv("SUPERVISORKEYS", string(supervisorKeysBytes))
	
	par1KeysBytes, _ := json.Marshal(par1Keys)
	os.Setenv("PAR1KEYS", string(par1KeysBytes))
	
	par2KeysBytes, _ := json.Marshal(par2Keys)
	os.Setenv("PAR2KEYS", string(par2KeysBytes))

	par3KeysBytes, _ := json.Marshal(par3Keys)
	os.Setenv("PAR3KEYS", string(par3KeysBytes))

	init_policies()

}



func init_policies() {
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

	mspAuditBytes, _ := json.Marshal(mspAudit)
	os.Setenv("MSP_AUDIT", string(mspAuditBytes))

	mspBreakGlassBytes, _ := json.Marshal(mspBreakGlass)
	os.Setenv("MSP_BREAKGLASS", string(mspBreakGlassBytes))

	mspABytes, _ := json.Marshal(mspA)
	os.Setenv("MSP_A", string(mspABytes))

	mspBBytes, _ := json.Marshal(mspB)
	os.Setenv("MSP_B", string(mspBBytes))

	mspAllBytes, _ := json.Marshal(mspAll)
	os.Setenv("MSP_ALL", string(mspAllBytes))
}

