// This doc is to establish main policies used within the system

package main

import (
	"github.com/fentec-project/gofe/abe"
)

func init_policies()  {
	// Policy List
	policy_audit := "Auditor"
	policy_breakGlass := "Auditor AND Manager"
	policy_a := "Part Type A"
	policy_b := "Part Type B"
	policy_all := "Part Type A OR Part Type B"


	// MSP structures for all policies
	mspAudit, _ := abe.BooleanToMSP(policy_audit, false) // The MSP structure defining the policy for an Audit case
	mspBreakGlass, _ := abe.BooleanToMSP(policy_breakGlass, false) // The MSP structure defining the policy for an extreme audit where a Manager/Auditor role is needed
	mspA, _ := abe.BooleanToMSP(policy_a, false) // The MSP structure defining the policy for Participant B
	mspB, _ := abe.BooleanToMSP(policy_b, false) // The MSP structure defining the policy for Participant A
	mspAll, _ := abe.BooleanToMSP(policy_all, false) // The MSP structure defining the policy for both Type A and B Participants
}

// Possible Attributes
// 'Auditor' 'Part Type A' 'Part Type B' 'Manager'