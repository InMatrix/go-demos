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

        // Question 1
        fmt.Println("\nWhat is the capital of France?")
        fmt.Println("a) Berlin")
        fmt.Println("b) Paris")
        fmt.Println("c) Rome")

        var answer1 string
        fmt.Print("Your answer: ")
        fmt.Scanln(&answer1)

        if answer1 == "b" || answer1 == "B" {
                fmt.Println("\033[32mCorrect!\033[0m")
        } else {
                fmt.Println("\033[31mIncorrect.\033[0m")
        }

        time.Sleep(1 * time.Second)

        // Question 2
        fmt.Println("\nWhat is the capital of Japan?")
        fmt.Println("a) Beijing")
        fmt.Println("b) Seoul")
        fmt.Println("c) Tokyo")

        var answer2 string
        fmt.Print("Your answer: ")
        fmt.Scanln(&answer2)

        if answer2 == "c" || answer2 == "C" {
                fmt.Println("\033[32mCorrect!\033[0m")
        } else {
                fmt.Println("\033[31mIncorrect.\033[0m")
        }
}