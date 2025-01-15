package main

import (
        "fmt"
        "os"
        "time"
)

func main() {
        // Clear the terminal screen
        os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")

        // Display welcome message
        fmt.Println("\033[32mWelcome to the Go Quiz Game!\033[0m")
        time.Sleep(2 * time.Second)

        // Ask a question
        fmt.Println("\nWhat is the capital of France?")
        fmt.Println("a) Berlin")
        fmt.Println("b) Paris")
        fmt.Println("c) Rome")

        // Get user input
        var answer string
        fmt.Print("Your answer: ")
        fmt.Scanln(&answer)

        // Check answer and provide feedback
        if answer == "b" || answer == "B" {
                fmt.Println("\033[32mCorrect!\033[0m")
        } else {
                fmt.Println("\033[31mIncorrect.\033[0m")
        }
}
