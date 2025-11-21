/**
 * @author Ethan White
 * @version 1.0.0
 * @date 2025-11-21
 * @fileoverview This program will allow you to calculate the cost of a car with certain add-ons
 */

//create constants
const car: number= 25000;
const mats: number = 500;
const navigation: number = 1000;
const seats: number = 500;
const warranty: number = 2500;
const yes: string = "yes";
const no: string = "no";
const taxAmount: number = 0.13;

//Variables
let totalCost: number=25000

//Input
let matInput = prompt("Do you want mats?")
if(matInput=yes) {
  totalCost: totalCost + 500;
}
else{

}
let navigationInput = prompt("Do you want a navigation system?")
if (navigationInput = yes) {
  totalCost: totalCost + 1000;
}
else {

}
let seatsInput = prompt("Do you want heated leather seats?")
if (seatsInput = yes) {
totalCost: totalCost + 500;

}
else {

}
let warrantyInput = prompt("Do you want 5 year extended warranty?")
if (warrantyInput = yes) {
totalCost: totalCost + 2500;

}
else {

}
let tax = totalCost * taxAmount;
let finalCost: number = totalCost + tax;
if (matInput = yes) {
  console.log(`${"Mats".padStart(10)} ${"$ 500".padEnd(8)} `)
}
else {

}
if (navigationInput = yes) {
  console.log(`${"Navigation".padStart(10)} ${"$ 1000".padEnd(8)} `)
}
else {

}
if (seatsInput = yes) {
  console.log(`${"Seats".padStart(10)} ${"$ 500".padEnd(8)} `)

}
else {

}
if (warrantyInput = yes) {
  console.log(`${"Warranty".padStart(10)} ${"$ 2500".padEnd(8)} `)

}
else {

}
console.log(`${"Tax  $".padStart(10)} ${tax} ${"".padEnd(8)} `)
console.log("The final cost is $" + finalCost)

console.log("\nDone.")