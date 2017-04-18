	//Author: Timothy P. McClintock
	//Assignment: Extra Credit - Group Presentation: Go Programming Language
	//Professor Tinkham
	//Rowan University
	
	package main
	
	import (
		"fmt"
		"io/ioutil"
		"strings"
	)

var(
	name string
	num int
)

func Roles(file1 string, file2 string) []string {
	//Takes the employees  from a file and creates a byte slice
    empl,   _ := ioutil.ReadFile(file1)
    //Takes the assignment from a file and creates a byte slice
    assign, _ := ioutil.ReadFile(file2)
    
    //The byte slice of employee    data is converted into a string slice
    employees  := string(empl)
    //The byte slice of assignment data is converted into a string slice
    assignment := string(assign)
    
   
    //Simplify the string slices into ordered slices delineated by white space
    list := strings.Split(assignment, " ")
    
      //Phone Answering Schedulees
      fmt.Println("Answer Phones Ia: ", list[2], "Answer Phones Ib: ", list[3])
      //fmt.Println("Answer Phones I: ", list[2:4])
      fmt.Println("Answer Phones IIa: ", list[8], "Answer Phones IIb: ", list[9])
      //fmt.Println("Answer Phones II: ", list[8:10])
      fmt.Println("Answer Phones IIIa: ", list[14], "Answer Phones IIIb: ",  list[15])
      //fmt.Println("Answer Phones III: ", list[14:16])
      fmt.Println("Answer Phones IVa: ", list[20], "Answer Phones IVb: ", list[21])
      //fmt.Println("Answer Phones IV: ", list[20:22])
      //Computer Repair Schedulees
      fmt.Println("Computer Repair I:", list[4])
      //fmt.Println("Computer Repair I:", list[4:5])
      fmt.Println("Computer Repair II:", list[10])
      //fmt.Println("Computer Repair II:", list[10:11])
      fmt.Println("Computer Repair III:", list[16])
      //fmt.Println("Computer Repair III:", list[16:17])
      fmt.Println("Computer Repair IV:", list[22])
      //fmt.Println("Computer Repair IV:", list[22:23])
      //Network Schedules
      fmt.Println("Network I: ", list[5])
      //fmt.Println("Network I: ", list[5:6])
      fmt.Println("Network II: ", list[11])
      //fmt.Println("Network II: ", list[11:12])
      fmt.Println("Network III: ", list[17])
      //fmt.Println("Network III: ", list[17:18])
      fmt.Println("Network IV: ", list[23])
      //fmt.Println("Network IV: ", list[23:24])
    
    //Replaces the number with a valid skill
    repair := strings.Replace(employees, "0 1 0", "Computer Repair", -1)
    phones := strings.Replace(repair, "1 0 0", "Answer Phones", -1)
    network := strings.Replace(phones, "0 0 1", "Network", -1)
    all := strings.Replace(network,"1 1 1", "Answer Phones, Computer Repair, Network", -1)
    noPhones := strings.Replace(all, "0 1 1", "Computer Repair, Network", -1)
    noRepair := strings.Replace(noPhones, "1 0 1", "Answer Phones, Network", -1)
    noNet := strings.Replace(noRepair,"1 1 0", "Answer Phones, Computer Repair", -1)
    fmt.Println(noNet)
    
    num = strings.LastIndex(noNet, list[2]) + 1
    
    //Conditional evaluation of the scheduled individual as per their respective position
    if strings.Contains(noNet, list[2]) && strings.Contains(noNet, "Answer Phones") {
    	fmt.Println("Valid Schedule.")}
    if strings.Contains(noNet, list[8]) && strings.Contains(noNet, "Answer Phones") {
    	fmt.Println("Valid Schedule.")}
    if strings.Contains(noNet, list[14]) && strings.Contains(noNet, "Answer Phones") {
    	fmt.Println("Valid Schedule.")}
    if strings.Contains(noNet, list[20]) && strings.Contains(noNet, "Answer Phones") {
    	fmt.Println("Valid Schedule.")}
    
    //Conditional evaluation of the scheduled individual as per their respective position
    if strings.Contains(noNet, list[4]) && strings.Contains(noNet, "Computer Repair") {
    	fmt.Println("Valid Schedule.")}
    if strings.Contains(noNet, list[10]) && strings.Contains(noNet, "Computer Repair") {
    	fmt.Println("Valid Schedule.")}
    if strings.Contains(noNet, list[16]) && strings.Contains(noNet, "Computer Repair") {
    	fmt.Println("Valid Schedule.")}
    if strings.Contains(noNet, list[22]) && strings.Contains(noNet, "Computer Repair") {
    	fmt.Println("Valid Schedule.")}
    
    //Conditional evaluation of the scheduled individual as per their respective position
    if strings.Contains(noNet, list[5])  && strings.Contains(noNet, "Network") {
    	fmt.Println("Valid Schedule.")}
    if strings.Contains(noNet, list[11]) && strings.Contains(noNet, "Network") {
    	fmt.Println("Valid Schedule.")}
    if strings.Contains(noNet, list[17]) && strings.Contains(noNet, "Network") {
    	fmt.Println("Valid Schedule.")}
    if strings.Contains(noNet, list[23]) && strings.Contains(noNet, "Network") {
    	fmt.Println("Valid Schedule.")}
    
	    return list
}
	
	func main(){
			Roles("Comp.txt", "Sched.txt")
	}
