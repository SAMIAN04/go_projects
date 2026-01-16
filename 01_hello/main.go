package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Read input
    var threatLevel string
    var maxZonesStr string
    fmt.Scanln(&threatLevel)
    fmt.Scanln(&maxZonesStr)
    
    // Convert max zones to integer
    maxZones, _ := strconv.Atoi(maxZonesStr)
    
    // Security zones data
    zones := [][]string{
        {"low", "medium", "low"},
        {"medium", "high", "low"},
        {"critical", "medium", "high"},
        {"low", "critical", "medium"},
    }
    
    // TODO: Write your code below
    // Use nested loops with labeled break to search for the threat
    // Print threat location if found, or "Threat not found in searched zones" if not found
    // Then use switch with fallthrough for security alerts
    
    found := false
    
searchLoop:
    for i := 0; i < maxZones && i < len(zones); i++ {
        for j := 0; j < len(zones[i]); j++ {
            if zones[i][j] == threatLevel {
                fmt.Printf("Threat found at zone %d position %d\n", i, j)
                found = true
                break searchLoop
            }
        }
    }
    
    if !found {
        fmt.Println("Threat not found in searched zones")
    }
    
    switch threatLevel {
    case "critical":
        fmt.Println("CRITICAL: Lockdown initiated!")
        fallthrough
    case "high":
        fmt.Println("HIGH: Security breach detected!")
        fallthrough
    case "medium":
        fmt.Println("MEDIUM: Increased monitoring active")
    case "low":
        fmt.Println("LOW: Standard security protocols")
    }
}