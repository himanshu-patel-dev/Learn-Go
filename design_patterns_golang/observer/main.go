package main

/*

Observer is a behavioral design pattern that allows some objects
to notify other objects about changes in their state.

*/

func main() {
	shirtItem := newItem("Nike Shirt")

	firstObserver := newCustomer("abc@gmail.com")
	secondObserver := newCustomer("xyz@gmail.com")

	shirtItem.register(firstObserver)
	shirtItem.register(secondObserver)

	shirtItem.updateAvailability()
}

//Item Nike Shirt is now in stock
//Sending email to customer abc@gmail.com for item Nike Shirt
//Sending email to customer xyz@gmail.com for item Nike Shirt
