## CryptoHackathon Project using ABE FE in order to address audit capabilities

### Summary of Project
My project focuses on a use case of Attribute-based Encryption in which a network of participants send encryptions meant for other participants to a managing server. The participants are able to use attributes to manage who is able to decrypt their ciphertext by implementing policies. While creating this project I focused on an auditing case in which a participant may enable auditing by encrypting it with the policy “Participant OR Auditor” and therefore either another participant or an auditor (as designated by the managing server) would be able to decrypt the ciphertext. This method also enables the participants who distinguish between types of participants who would be able to decrypt their ciphertext. Although it is not displayed in the demo, this project also focused on layering this encryption in a fashion that would allow for a transmission of one ciphertext that had another ciphertext revealed once decrypted allowing only a further subset of attribute-owners to see the second layer ciphertext. The prototype written during this hackathon uses the FAME scheme within the GoFE library.

WARNING ---- this file is not up to date with the latest commits!
## The Demo files
The demo shown by running `main.go` has one participant encrypt 5 different plaintexts with 5 different policies and then have each of the given roles attempt to decrypt all 5 ciphertexts. This accents how the effect ABE has on the exchange of data and restricting that exchange within a network of entities.

The file `enc_dec.go` sets up functions to enable the demo. The function `ParticipantEncrypts` takes 5 different sample messages and encrypts them each using a different policy that was initiated in setup. If policies are added, they need to be added here as well to enable the demo in the way intended.

The function `Decryption` attempts to decrypt each given ciphertext with the given attribute-based key. The results shown describe what you should be seeing based on what permissions each of the roles have. 


## Initialize
`initialize.go` contains the necessary functions to setup the the ABE FAME scheme given specific roles (or people within the network) and policies. To create this instance and enable encryption and decryptions, begin by calling `initialize.Setup()` from the main file.

### Adding a role
To add a role to the list (in other words, add a person to the network given specific attributes that apply to them), a string containing their attributes must be defined. For example, adding ` manager2 := []string{"Manager"}` would add a second manager with the same permissions as the first. 

After adding the role, you must also generate the attribute-based key, for example, 
    
    manager2_keys, _ := a.GenerateAttribKeys(manager2, secKey). 
    
Note: It is important to generate a separate attribute-based key for each person using the network even if they have the same attributes, otherwise recovering from attack and compromise is made more difficult.

### Adding a policy
The function `init_policies` lists the policies available and creates the necessary MSPs to use them. To add a policy, you must first add a string that describes the policy. For example, to add a policy that restricts access so that only a role with the attributes  `Participant Type A` and   `Auditor`, you would add the following code to `init_policies`:

    policy_new := "Auditor AND Part Type A"
    mspNew, _ := abe.BooleanToMSP(policy_new, false)
    

## Future Developments
In a further iteration of this project I would like to structure encryptions to allow for a layered approach to decryption access. This would be most useful in cases in which the policy used to encrypt used an "OR" statement and the information available upon decrypting should be available only in part to one of the roles allowed in the policy. For example, say a participant in the network wants to make their data available to other participants and to an auditor (not a participant that is also an auditor), however the auditor shouldn't be able to see the cleartext data, just proof of another nature, while the participant requesting data within the network should be able to see all the available data. In this case, the encrypting participant would first encrypt what it would only want the other participant to see using a policy allowing only participants to see it, then encrypt this ciphertext as well as the other data that is meant for the auditor as well using a policy in which either the participant "OR" the auditor is able to decrypt it. 
