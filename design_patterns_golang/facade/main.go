package main

/*

----------------------- DEFINITION: FACADE PATTERNS
Facade is a structural design pattern that provides a simplified 
(but limited) interface to a complex system of classes, library or 
framework.

---------------- PROBLEM
Imagine that you must make your code work with a broad set of objects 
that belong to a sophisticated library or framework. Ordinarily, you’d 
need to initialize all of those objects, keep track of dependencies, 
execute methods in the correct order, and so on.

---------------- SOLUTION
A facade is a class that provides a simple interface to a complex 
subsystem which contains lots of moving parts. A facade might provide 
limited functionality in comparison to working with the subsystem 
directly. 

----------------- APPLICATION
Use the Facade pattern when you need to have a limited but 
straightforward interface to a complex subsystem.

Use the Facade when you want to structure a subsystem into layers.

------------------ IMPLEMENT
Check whether it’s possible to provide a simpler interface than what 
an existing subsystem already provides. You’re on the right track if 
this interface makes the client code independent from many of the 
subsystem’s classes.

Declare and implement this interface in a new facade class. The 
facade should redirect the calls from the client code to appropriate 
objects of the subsystem. The facade should be responsible for 
initializing the subsystem and managing its further life cycle unless 
the client code already does this.

To get the full benefit from the pattern, make all the client code 
communicate with the subsystem only via the facade. Now the client 
code is protected from any changes in the subsystem code. For 
example, when a subsystem gets upgraded to a new version, 
you will only need to modify the code in the facade.

*/ 

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}

/*

Starting create account
Account created

Starting add money to wallet
Account Verified
SecurityCode Verified
Wallet balance added successfully
Sending wallet credit notification
Make ledger entry for accountId abc with txnType credit for amount 10

Starting debit money from wallet
Account Verified
SecurityCode Verified
Wallet balance is Sufficient
Sending wallet debit notification
Make ledger entry for accountId abc with txnType debit for amount 5

*/
