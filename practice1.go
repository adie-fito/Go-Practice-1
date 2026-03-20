package main

import "fmt"

func main(){
// var revenue, expense, taxRate  float64 

revenue := getUserInput("Revenue: ")
expense := getUserInput("Expense: ")
taxRate := getUserInput("Tax Rate: ")

earningBeforeTax, earningAfterTax, ratio := calculateFinancials(revenue, expense, taxRate)

fmt.Print("Earning Before Tax/Profit: ")
fmt.Printf("%.1f\n", earningBeforeTax)
fmt.Print("Earning After Tax: ")
fmt.Printf("%.1f\n", earningAfterTax)
fmt.Print("Ratio: ")
fmt.Printf("%.3f\n", ratio)
}

// func userInputVariabelValues(variabel float64) (a string, b int){
// 	a = fmt.Sprint(`Future Value: %.0f`, variabel)
// 	b,_ = fmt.Scan(&variabel)
// return a, b
// }

func calculateFinancials(revenue, expense, taxRate float64)(float64, float64, float64){
 earningBeforeTax :=  revenue - expense
 earningAfterTax := earningBeforeTax *	(1 - taxRate / 100)
 ratio := earningBeforeTax / earningAfterTax
return earningBeforeTax, earningAfterTax, ratio
}

func getUserInput(infoText string)float64{
	var UserInput float64
	fmt.Print(infoText)
fmt.Scan(&UserInput)
return UserInput
}