package initialize

import (
	"github.com/fentec-project/gofe/abe" 
	"encoding/json"
	"os"
)



func Setup() {
	
	scheme := abe.NewFAME() // Create the scheme instance
	schemeBytes, _ := json.Marshal(scheme)
	f1, _ := os.Create("../variables/scheme")
    defer f1.Close()
    f1.Write(schemeBytes)

	// Generate public and secret keys for instance
	pubKey, secKey, _ := scheme.GenerateMasterKeys()

	pubKeyBytes, _ := json.Marshal(pubKey)
	f2, _ := os.Create("../variables/pubkey")
    defer f2.Close()
	f2.Write(pubKeyBytes)

	secKeyBytes, _ := json.Marshal(secKey)
	f3, _ := os.Create("../variables/seckey")
    defer f3.Close()
    f3.Write(secKeyBytes)


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
	f4, _ := os.Create("../variables/managerkeys")
    defer f4.Close()
    f4.Write(managerKeysBytes)

	auditorKeysBytes, _ := json.Marshal(auditorKeys)
	f5, _ := os.Create("../variables/auditorkeys")
    defer f5.Close()
    f5.Write(auditorKeysBytes)
	
	supervisorKeysBytes, _ := json.Marshal(supervisorKeys)
	f6, _ := os.Create("../variables/supervisorkeys")
    defer f6.Close()
    f6.Write(supervisorKeysBytes)
	
	par1KeysBytes, _ := json.Marshal(par1Keys)
	f7, _ := os.Create("../variables/par1keys")
    defer f7.Close()
    f7.Write(par1KeysBytes)
	
	par2KeysBytes, _ := json.Marshal(par2Keys)
	f8, _ := os.Create("../variables/par2keys")
    defer f8.Close()
    f8.Write(par2KeysBytes)

	par3KeysBytes, _ := json.Marshal(par3Keys)
	f9, _ := os.Create("../variables/par3keys")
    defer f9.Close()
    f9.Write(par3KeysBytes)

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
	f1, _ := os.Create("../policies/mspaudit")
    defer f1.Close()
    f1.Write(mspAuditBytes)

	mspBreakGlassBytes, _ := json.Marshal(mspBreakGlass)
	f2, _ := os.Create("../policies/mspbreakglass")
    defer f2.Close()
    f2.Write(mspBreakGlassBytes)

	mspABytes, _ := json.Marshal(mspA)
	f3, _ := os.Create("../policies/msp_a")
    defer f3.Close()
    f3.Write(mspABytes)

	mspBBytes, _ := json.Marshal(mspB)
	f4, _ := os.Create("../policies/msp_b")
    defer f4.Close()
    f4.Write(mspBBytes)

	mspAllBytes, _ := json.Marshal(mspAll)
	f5, _ := os.Create("../policies/msp_all")
    defer f5.Close()
    f5.Write(mspAllBytes)
}

