/*
Author: Ethan White
Version: 1.0.0
Date: 2025-11-21
Fileoverview: This program will calculate the cost of a car with certain add ons
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Constants
	const car = 25000
	const mats = 500
	const navigation = 1000
	const seats = 500
	const warranty = 2500
	const taxAmount = 0.13

	// Variables
	totalCost := car

	reader := bufio.NewReader(os.Stdin)

	// INPUTS
	fmt.Print("Do you want mats? (yes/no): ")
	matInput, _ := reader.ReadString('\n')
	matInput = strings.TrimSpace(strings.ToLower(matInput))

	if matInput == "yes" {
		totalCost += mats
	}

	fmt.Print("Do you want a navigation system? (yes/no): ")
	navInput, _ := reader.ReadString('\n')
	navInput = strings.TrimSpace(strings.ToLower(navInput))

	if navInput == "yes" {
		totalCost += navigation
	}

	fmt.Print("Do you want heated leather seats? (yes/no): ")
	seatsInput, _ := reader.ReadString('\n')
	seatsInput = strings.TrimSpace(strings.ToLower(seatsInput))

	if seatsInput == "yes" {
		totalCost += seats
	}

	fmt.Print("Do you want 5-year extended warranty? (yes/no): ")
	warrantyInput, _ := reader.ReadString('\n')
	warrantyInput = strings.TrimSpace(strings.ToLower(warrantyInput))

	if warrantyInput == "yes" {
		totalCost += warranty
	}

	// TAX + FINAL COST
	tax := float64(totalCost) * taxAmount
	finalCost := float64(totalCost) + tax

	// OUTPUT SECTION
	fmt.Println("\n----- Selected Add-Ons -----")

	if matInput == "yes" {
		fmt.Println("Mats        $500")
	}
	if navInput == "yes" {
		fmt.Println("Navigation  $1000")
	}
	if seatsInput == "yes" {
		fmt.Println("Seats       $500")
	}
	if warrantyInput == "yes" {
		fmt.Println("Warranty    $2500")
	}
  fmt.Println("Tax    " , tax)
	fmt.Printf("\nThe final cost is $%.2f\n", finalCost)
	fmt.Println("\nDone.")
}
