package main

import (
    "fmt"
    "strconv"
)

// TODO: Define your Sensor struct here

// TODO: Define your displayReading method with value receiver here

// TODO: Define your adjustTemperature method with value receiver here

func main() {
    // Read inputs
    var sensorID string
    var tempStr string
    var adjustStr string
    
    fmt.Scanln(&sensorID)
    fmt.Scanln(&tempStr)
    fmt.Scanln(&adjustStr)
    
    // Parse temperature and adjustment values
    temperature, _ := strconv.ParseFloat(tempStr, 64)
    adjustment, _ := strconv.ParseFloat(adjustStr, 64)
    
    // TODO: Create a Sensor instance and call the required methods
    
}