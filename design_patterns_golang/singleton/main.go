package main

import "fmt"

/*

--------------- DEFINITION: singleton pattern

Singleton is a creational design pattern, which ensures that only
one object of its kind exists and provides a single point of access
to it for any other code.

--------------- PROBLEM + APPLICATION

1. Ensure that a class has just a single instance.
to control access to some shared resource—for example, a database or a file.

2. Provide a global access point to that instance.
those global variables that you used to store some essential objects? 
While they’re very handy, they’re also very unsafe since any code can 
potentially overwrite the contents of those variables and crash the app.

--------------- IMPLEMENT
1. Add a private static field to the class for storing the singleton instance.

2. Declare a public static creation method for getting the singleton instance.

3. Implement “lazy initialization” inside the static method. It should create 
a new object on its first call and put it into the static field. The method 
should always return that instance on all subsequent calls.

4. Make the constructor of the class private. The static method of the class 
will still be able to call the constructor, but not the other objects.

5. Go over the client code and replace all direct calls to the singleton’s 
constructor with calls to its static creation method.

*/

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
