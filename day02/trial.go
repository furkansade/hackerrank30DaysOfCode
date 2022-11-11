package main

import (
    "fmt"
    "math"
)

func Calculate(meal_cost float64, tip_percent int32, tax_percent int32) {
    
    tip := (float64(tip_percent) / 100) * meal_cost
    tax := (float64(tax_percent) / 100) * meal_cost
    sum := math.Round(tip + tax + meal_cost)
    fmt.Println(sum)
    
}


func main() {
    var meal_cost float64
    var tip_percent int32
    var tax_percent int32

    fmt.Scanln(&meal_cost)
    fmt.Scanln(&tip_percent)
    fmt.Scanln(&tax_percent)

    Calculate(meal_cost, tip_percent, tax_percent )
}