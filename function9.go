package main

import "fmt"

func main() {
	doDBOperations()
}

func connectToDB() {
	fmt.Println("ok,connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok,disconnected form db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB()
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash pr network error ...")
	fmt.Println("Returning form function here!")
	return
}
