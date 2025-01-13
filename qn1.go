/*1. Write a program to calculate the simple interest
First-line has the comma-separated values of the Principal, rate and time (in years) respective
*constraints: Round simple interest float value to 2 decimal places*/

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main(){
var input string

if _,err:=fmt.Scan(&input);err!=nil{
	log.Fatal("Error accepting input")
}

result:=strings.Split(input, ",")

var principal,e1 =strconv.ParseFloat(result[0],64)
if e1!=nil{
	log.Fatal("error invalid principal amount")
}

var rate,e2 =strconv.ParseFloat(result[1],64)
if e2!=nil{
	log.Fatal("error invalid rate")
}

var time,e3 =strconv.ParseFloat(result[2],64)
if e3!=nil{
	log.Fatal("error invalid time ")
}

var simple_interest float64 
simple_interest=(principal*rate*time)/100

fmt.Printf("simple interest %.2f",simple_interest)

}