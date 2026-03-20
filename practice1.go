package main

import (
	"errors"
	"fmt"
	"os"
)

func main(){
// var revenue, expense, taxRate  float64 

revenue, err := getUserInput("Revenue: ")
if err != nil{
	fmt.Println(err)
	return 
}
expense, err := getUserInput("Expense: ")
if err != nil{
	fmt.Println(err)
	return
}
taxRate, err := getUserInput("Tax Rate: ")
if err != nil{
	fmt.Println(err)
	return
}


earningBeforeTax, earningAfterTax, ratio := calculateFinancials(revenue, expense, taxRate)

StoreCalculatedResult(earningBeforeTax, earningAfterTax, ratio)

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

func getUserInput(infoText string)(float64, error){
	var UserInput float64
	fmt.Print(infoText)
fmt.Scan(&UserInput)
if(UserInput <= 0){	
	return 0, errors.New("Value Must be positif Numbers")
}
return UserInput, nil
}

func StoreCalculatedResult(beforeResult, afterResult, ratio float64){
	resultText := fmt.Sprintf("Before Tax: %.1f\nAfter Tax: %.1f\nRatio: %.3f",beforeResult,afterResult,ratio )
os.WriteFile("Result.txt",[]byte(resultText), 0644)
}


