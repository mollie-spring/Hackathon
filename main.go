package main
import (
)



func main()  {
	
	setup()
	// eventually make it so all the keys are written to files so that the main program doesn't 'see' them

	// A participant encrypts 1 message in various policies and submits to the manager
	cipher1, cipher2, cipher3, cipher4, cipher5 = partipantEncrypts()


	// attempt to decrypt by participant1
	 //potential_msg1 = decrypt()


}

