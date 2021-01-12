import (
	"os"
	"/Users/molliezechlin/go/src/github.com/fentec-project/gofe/abe"
	"init_roles"
	"policies"
)


func main()  {
	
	setup()
	// eventually make it so all the keys are written to files so that the main program doesn't 'see' them


	// attempt to decrypt by an auditor
	


}


mspAudit, _ := abe.BooleanToMSP(policy_audit, false) // The MSP structure defining the policy for an Audit case
mspBreakGlass, _ := abe.BooleanToMSP(policy_breakGlass, false) // The MSP structure defining the policy for an extreme audit where a Manager/Auditor role is needed
mspA, _ := abe.BooleanToMSP(policy_a, false) // The MSP structure defining the policy for Participant B
mspB, _ := abe.BooleanToMSP(policy_b, false) // The MSP structure defining the policy for Participant A
mspAll, _ := abe.BooleanToMSP(policy_all, false) // The MSP structure defining the policy for both Type A and B Participants
