// intitialize go app into a package
// go mod init Booking-app

// All go must belong to a package
// the main function is the entrypoint of GO program
// Go programs are organized into packages
// Go standadrd library provides different core packages for us to use
// fmt is an example of a standard library
// go run 'filename' , compiles and run the code
// GO is a statiically data type you need to tell Go compiler, the data type when declaring variable
// Go is type inference, i.e can infer the data type when you assign the balue  for example var name = "Goddhi"
// Different functions for formatted input and output 1. print message 2. collect user input, 3. write into file
// A pointer is a variable that points to the memory address of another varible
// A pointer is called A special Variable in GO
// Arrays in GO has fixed size, size in this context mens a=how many elements the array can hold
// Slice is an abstraction of an array
// Slices are more flexible and powerful: variable-length or to get an sub-array of its own
// Slices are also index based and have a size, but is resized whhen needed
// append function for slice adds the element at the end of the slice
// the append function also grow the slice if a greappter capacity is needed and returns the update slice value
// Loop in GO is known as For loop. foreach
// infinite loop iterate multiple times
// for-each loop iterate over a list
// range iterates over elements for different structures (so not only arrays and slices)
// for arrays and slices, range provides the index and value for eah element
// Strings.field(), split the string with white space as seperator
// And returns a slice with the split elements
// Blank identifier with the symbol _  it ignores a variable  you dont want to use
// so with GO you need to make unused variable explicit
// conditions in loops such as for loop will keep on iterating as far the condition is true
// Encapsulate code  into its own container (=function) which logically belongs together
// like variable name, you should give a fucntion a descriptive name
// call the function by its name, whenever you want to execute this block of code
// Every program has at least one function, which is the main() function
// A function is used to reduce code duplication
// Parametes: are Information can be passed into  fucntion as parameters
// Parameters are also caleed arguments
// Return values: A fucntion can return data as a result
// so a function can take an input and return an output
// In Go you have to define to define the input and output( return variable) paramters incliding its type explicity
// when a function returns a value, the value is printed out in the main function
// A Go function can return multiple values
// Note: return multiple variables is possible in GO but not possible in most programming language
// Make code cleaner with package level variables
// Package Level Variables : they are defined at the top outside all functions
// They can be accessed inside any of the functions including the main function
// and in all files, which are in the same poackage
// Note: In Package level variables, we cant create a variable using this symbol :=   we create vatiable using the var keyword
// Local Variables meaning variables declared inside of a function : they are defined inside a function or a block
// They can be accessed only inside that function or block of code
// Best Practice: Create the variable when you need it
// More use cases of functions
// 1. They are group logic that belongs together
// 2. Reuse logic and so reducing duplication code
// Packages in Go:
// Applications in GO can be modularized: Applications consisting of multiples Go files
// Go programs are organized into packages
// A package is a collection of Go files
// Note for Package level : Again: variables and functions defined outside any function, can be accessed in all other files within the same package
// it is a best practice to decouple code into multiple files that all belong to the same package
//
//	there can be multiple packages in your application. this use case is used when each of the packages have different functionalities
//
// Best Practice, when you have multiple packages you should create a folder for them
// Note: The module in the file path go.mod named booking-app is an import path for all our packages in our booking-app applications
// Exporting a variable or function in a package means to make it available for all packages
// In Go, variables or functions are exported by capitalizing the first name of the variable and function
// you cant only export functions and variables, you can also export constant, types etc.
// There are 3 variable levels of Scope
// They are Local, Package, Global
// Local level variables: they are declared within the functions and can only be used only within that function and also Declaration within block(e.g for, if-else) : can be used only within that block
// Package level variables : They are declared outside all functions : can be used everywhere in the same package
// Global level variables : They are declared outside all functions and uppercase first letter : they can be used everywhere across all packages
// Variable Scope = is the region of a program, where a defined variable can accessed
// example of a Global level variables : var Myvar = "somevalues"
// Maps in Go are used extensively just like dictionary in Python..,,
// Maps is a unique key to values
// All keys have the same data type
// All values have the same data type
// Note in MAPS we cant, we cant have mix data types. i.e, we cant have different data types for keys and values. this also applies to string as well.. this is specific to GO language
// / Maps are mostly used to collect datas of the same data type. i.e Maps suports only 1 data type
// STRUCT, are mostly used to collect datas of different data types. e.g strings, integers, boolean, slice and maps
// Struct stands for structure and can hold mixed data types
// Struct is mostly used in tha package level
// The "type" keyword in Struct creates a new type, with the name you specify
// Create a type called UserData based on a struct of firstName, lastName
// Struct, it's like a lightweight class, but struct dosnt support inheritance
// Go Routines - Concurrency.. it is used to stimulate a long processimng task,  in this code logic we will use it to generate and send ticket
// / Generating and sending ticket is a longer task, it takes quite a long of process
// Handling this bloxking code with Goroutines
// By default: Go codes are executed Sequentially
// By default : In Go Tasks are processed 1 task or 1 code after another
// Concurrency is used to make our program more efficient
// Concurrency in GO is cheap and easy
// GO execute code in single thread execution i.e step by step execution of the code
// In this code logic the sendTicket function which takes longer to execute and blocks other functions or code from executing, to solve this issue we will execute the sendTIcket in a seperate thread i.e a seperate goroutine, so main thread doeesnt wait for sendTicket to finish, Next line of code is executed immediately
// once the seperate thread or goroutines is done executing it gets deleted with no interuption
// to implement goroutine in this code, we will make the sendTIcket function in a seperate thread using the "go" keyword
// The "go" keyword starts a new goroutine and a goroutine is a lightweiht thread managed by the Go runtime
// The sendTIcket function task now runs in the backgorund
// With Go goroutines the application now stays responsive and faster
// Synchronizing the Goroutines
// if the infinite loop is removed from the code, this means the code will be executed once.. which in turn leads to problem because by default, the main goroutine does NOT wait for other goroutines. in this code the "main" gorutines gets exited before the "sendTIcket" goroutine had time to start and executing the code
// To solve this problem, we need to tell the  "main" goroutine that it needs to wait until "sendTicket" goroutine is done
// we achieve this by creating a waitgroup
// The waitgroup makes the "main" goroutine to wait before the "sendTiket fun"
// The waitgroup utilizes a package called "sync", sync provides basic synchronization functionalitiy
// waitgroup has 3 fucntions named Add, Wait, Done
// The Add fucntion sets the numnber of goroutines to wait for (increases the counter by the provided number). in this code we will be passing 1 into the add function which means we only have 1 seperate(sendTicket) gorutine the "main" goroutines needs to wait for
// code WITHOUT the for loop
// The wait function blocks untill the waitGroup counter is 0
// The Done function decrement the waitGroup by counter  by 1, so this is called by the goroutine to indicate that its finished
package main

import (
	"Booking-app/helper"
	"fmt" // print messages

	//"strconv" //convert uint data type to string
	// "strings"
	"sync"
	"time" // this time package, which a sleep function contains, the sleep function stops or blocks the current "thread"(goroutine) execution for the defined duration
)

// Declared a Package level variable which is accessible by any functions
var conferenceName string = "Go Conference"

const conferenceTicket int = 50

var remainingTicket uint = 50

// var bookings = make([]map[string]string, 0) // initialize a list  of empty maps with the make function
var bookings = make([]UserData, 0) // initialize/create a list of empty UserData of struct type

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

// creating a waitgroup function
var wg = sync.WaitGroup{}

func main() {

	// greetUser(conferenceName, conferenceTicket, remainingTicket) // we dont need this line of code because we have declared a package level variable and all functions can now read the variable
	greetUser() // This functions prints out some package level variable

	firstName, lastName, email, userTickets := getUserInput()                                                                                         /// the return varibles were redeclaried here, so that it can be passed as an argument or parameter to the helper function can use it
	isValidFirstNameLastName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTicket) /// the return varibles were redeclaried here, so that the if statement can use it

	if isValidFirstNameLastName && isValidEmail && isValidTicketNumber { /// if all inputs are are true as defined in the helper functon execute the code below

		bookTicket(userTickets, firstName, lastName, email)    ///// the parameters passed here are gotten from the return variable of getUserInput variable
		go sendTIcket(userTickets, firstName, lastName, email) // with the "go" keyword, the sendTicket now runs in a seperate thread.. we just implemented concurency in our code

		wg.Add(1) // we only have one seperate goroutine
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTicket == 0 {
			fmt.Println("sorry, we have no tickets left")
		}
	} else {

		if !isValidFirstNameLastName {
			fmt.Println("input the valid firstname or lastname")
		}

		if !isValidEmail {
			fmt.Println("input the right email address")
		}

		if !isValidTicketNumber {
			fmt.Println("Input the valid ticket number")
		}

	}
	wg.Wait() // wait for all the sepeate threads to be execited before the main thread is executed
}

func greetUser() {
	fmt.Printf("welcome to our %v Booking Application\n", conferenceName)
	fmt.Printf("we have total of %v ticket and remaining %v  Booking Application \n", conferenceName, remainingTicket)
	fmt.Println("Get your tickets here to attend")
}
func getFirstNames() []string { // data type of output parameters
	firstNames := []string{}
	for _, value := range bookings { /// looping through the list of struct
		// var names = strings.Fields(value) /// this function converts each values into a slice and with a whitespace sperating the values

		// firstNames = append(firstNames, value["firstName"])  // getting the firstName key value from a map list
		firstNames = append(firstNames, value.firstName)
	}
	return firstNames // output paramters
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName) // Scan function can now assign the user's value to the userName variable, because it has a pointer to  its memory address
	// Note: without the pointer above we wouldnt have gotten the username value

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")

	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTicket = remainingTicket - userTickets
	//bookings[0] = firstName  + " " + lastName  // implementation of getting the first value of list  using array

	// create a map for user using map
	// var userData = make(map[string]string) // creating a map with keys and values
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberofTickets"] = strconv.FormatUint(uint64(userTickets), 10) // converts unit to string

	// create a map for user using map
	var userData = UserData{ // utilizing UserData struct data type to create keys and values
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberofTickets: userTickets,
	}
	bookings = append(bookings, userData)          // appending userData maps to the empty list of maps(bookings)
	fmt.Printf("List of bookings is %v", bookings) // printing the bookings maps
	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaing for the %v\n", remainingTicket, conferenceName)

}

// creating a function that generate a ticket and send it to a user
func sendTIcket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) // the sleep function stops the thread for 10 seconds and then resume the program after 10 seconds
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("Sending ticket:\n %v into email address %v\n", ticket, email)
	fmt.Println("##############")
	wg.Done() // this removes the  sendTIcket thread from the waiting list, basically it tells the main thread not to wait for it anymore because its done executin

}

///  code WITH the for loop iteration
// package main

// import (
// 	"Booking-app/helper"
// 	"fmt" // print messages
// 	//"strconv" //convert uint data type to string
// 	// "strings"
// 	"time" // this time package, which a sleep function contains, the sleep function stops or blocks the current "thread"(goroutine) execution for the defined duration
// )

// // Declared a Package level variable which is accessible by any functions
// var conferenceName string = "Go Conference"
// const conferenceTicket int = 50
// var remainingTicket uint = 50
// // var bookings = make([]map[string]string, 0) // initialize a list  of empty maps with the make function
// var bookings = make([]UserData, 0) // initialize/create a list of empty UserData of struct type

// type UserData struct {
// 	firstName string
//     lastName string
// 	email string
// 	numberofTickets uint
// }

// func main() {

// 	// greetUser(conferenceName, conferenceTicket, remainingTicket) // we dont need this line of code because we have declared a package level variable and all functions can now read the variable
//     greetUser() // This functions prints out some package level variable

// 	for { /// infinite loop

// 		firstName, lastName, email, userTickets := getUserInput()  /// the return varibles were redeclaried here, so that it can be passed as an argument or parameter to the helper function can use it
// 		isValidFirstNameLastName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTicket) /// the return varibles were redeclaried here, so that the if statement can use it

// 		if isValidFirstNameLastName && isValidEmail && isValidTicketNumber { /// if all inputs are are true as defined in the helper functon execute the code below

// 			bookTicket(userTickets, firstName, lastName, email) ///// the parameters passed here are gotten from the return variable of getUserInput variable
// 			go sendTIcket(userTickets, firstName, lastName, email) // with the "go" keyword, the sendTicket now runs in a seperate thread.. we just implemented concurency in our code

// 			firstNames := getFirstNames()
// 			fmt.Printf("The first names of bookings are: %v\n", firstNames)

// 			if remainingTicket == 0 {
// 				fmt.Println("sorry, we have no tickets left")
// 				break
// 			}
// 		} else {

// 			if !isValidFirstNameLastName {
// 				fmt.Println("input the valid firstname or lastname")
// 			}

// 			if !isValidEmail {
// 				fmt.Println("input the right email address")
// 			}

// 			if !isValidTicketNumber {
// 				fmt.Println("Input the valid ticket number")
// 			}

// 		}

// 	}

// }

// func greetUser() {
// 	fmt.Printf("welcome to our %v Booking Application\n", conferenceName)
// 	fmt.Printf("we have total of %v ticket and remaining %v  Booking Application \n", conferenceName, remainingTicket)
// 	fmt.Println("Get your tickets here to attend")
// }
// func getFirstNames() []string { // data type of output parameters
// 	firstNames := []string{}
// 	for _, value := range bookings { /// looping through the list of struct
// 		// var names = strings.Fields(value) /// this function converts each values into a slice and with a whitespace sperating the values

// 		// firstNames = append(firstNames, value["firstName"])  // getting the firstName key value from a map list
// 		firstNames = append(firstNames, value.firstName)
// 	}
// 	return firstNames // output paramters
// }

// func getUserInput () (string, string, string, uint) {
// 	var firstName string
// 	var lastName string
// 	var email string
// 	var userTickets uint

// 	fmt.Println("Enter your first name")
// 	fmt.Scan(&firstName) // Scan function can now assign the user's value to the userName variable, because it has a pointer to  its memory address
// 	// Note: without the pointer above we wouldnt have gotten the username value

// 	fmt.Println("Enter your last name: ")
// 	fmt.Scan(&lastName)

// 	fmt.Println("Enter your email address: ")
// 	fmt.Scan(&email)

// 	fmt.Println("Enter number of tickets: ")

// 	fmt.Scan(&userTickets)

// 	return firstName, lastName, email, userTickets
// }

// func bookTicket(userTickets uint, firstName string, lastName string, email string)  {
// 	remainingTicket = remainingTicket - userTickets
// 	//bookings[0] = firstName  + " " + lastName  // implementation of getting the first value of list  using array

// 	// create a map for user using map
// 	// var userData = make(map[string]string) // creating a map with keys and values
// 	// userData["firstName"] = firstName
// 	// userData["lastName"] = lastName
// 	// userData["email"] = email
// 	// userData["numberofTickets"] = strconv.FormatUint(uint64(userTickets), 10) // converts unit to string

// 	// create a map for user using map
// 	var userData = UserData { // utilizing UserData struct data type to create keys and values
// 		firstName: firstName,
// 		lastName: lastName,
// 		email: email,
// 		numberofTickets: userTickets,
// 	}
// 	bookings = append(bookings, userData) // appending userData maps to the empty list of maps(bookings)
// 	fmt.Printf("List of bookings is %v", bookings) // printing the bookings maps
// 	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
// 	fmt.Printf("%v tickets remaing for the %v\n", remainingTicket, conferenceName)

// }

// // creating a function that generate a ticket and send it to a user
// func sendTIcket(userTickets uint, firstName string, lastName string, email string) {
// 	time.Sleep(10 * time.Second) // the sleep function stops the thread for 10 seconds and then resume the program after 10 seconds
// 	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
// 	fmt.Println("##############")
// 	fmt.Printf("Sending ticket:\n %v into email address %v\n", ticket, email )
// 	fmt.Println("##############")

// }
