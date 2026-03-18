package main

import "fmt"

func main(){
var revenue, expense, taxRate  float64 

fmt.Print("Ravenue: ")
fmt.Scan(&revenue)
fmt.Print("Expanse: ")
fmt.Scan(&expense)
fmt.Print("Tax Rate: ")
fmt.Scan(&taxRate)

 earningBeforeTax :=  revenue - expense
 earningAfterTax := earningBeforeTax - taxRate
 ratio := earningBeforeTax / earningAfterTax

fmt.Print("Earning Before Tax/Profit: ")
fmt.Println(earningBeforeTax)

fmt.Print("Earning After Tax: ")
fmt.Println(earningAfterTax)

fmt.Print("Ratio: ")
fmt.Println(ratio)
}