package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"https://github.com/sebastiannesheim/Minyr.git"
)

const inputFile = "kjevik-temp-celsius-20220318-20230318.csv"
const outputFile = "kjevik-temp-fahr-20220318-20230318.csv"

func main() {
	choice := presentOptions()

	switch choice {
	case "convert":
		if err := handleConvertOption(); err != nil {
			log.Fatal(err)
		}
	case "average":
		if err := handleAverageOption(); err != nil {
			log.Fatal(err)
		}
	case "exit", "quit", "q":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice.")
	}
}

func presentOptions() string {
	fmt.Println("***********************************************")
	fmt.Println("*                                             *")
	fmt.Println("*  Welcome to the temperature converter!      *")
	fmt.Println("*                                             *")
	fmt.Println("***********************************************")
	fmt.Println("Please select an option:")
	fmt.Println("Type 'convert' to convert create a new file with temperatures in Fahrenheit.")
	fmt.Println("Type 'average' to calculate the average temperature from the files.")
	fmt.Println("Type 'exit' to exit the program.")

	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Trim newline character and convert to lowercase
	choice = strings.TrimSpace(strings.ToLower(choice))

	return choice
}

func handleConvertOption() error {
	// Check if output file already exists
	if _, err := os.Stat(outputFile); err == nil {
		// Output file already exists, ask user if they want to generate it again
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Output file already exists. Generate again? (y/n): ")
		confirm, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		confirm = strings.TrimSpace(strings.ToLower(confirm))

		switch confirm {
		case "y", "yes":
			// Generate output file again
			if err := generateOutputFile(); err != nil {
				return err
			}
			fmt.Println("Output file generated successfully.")
		default:
			// Do not generate output file again
			fmt.Println("Exiting program.")
		}
	} else {
		// Output file does not exist, generate it
		if err := generateOutputFile(); err != nil {
			return err
		}
		fmt.Println("Output file generated successfully.")
	}
	return nil
}

func handleAverageOption() error {

	// Prompt user for unit of measurement
	fmt.Print("What unit of measurement do you want the average temperature in? (c/f): ")

	// Read user input
	reader := bufio.NewReader(os.Stdin)
	unit, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	unit = strings.TrimSpace(strings.ToLower(unit))

	switch unit {
	case "c":
		// Calculate average temperature in Celsius from input file and print
		average, err := calculateAverageTemperature(inputFile, "c")
		if err != nil {
			return err
		}
		fmt.Printf("Average temperature: %.2f °C\n", average)
	case "f":
		// Calculate average temperature in Fahrenheit from output file and print
		average, err := calculateAverageTemperature(outputFile, "f")
		if err != nil {
			return err
		}
		fmt.Printf("Average temperature: %.2f °F\n", average)
	default:
		fmt.Println("Invalid unit of measurement.")
	}

	return nil
}

func calculateAverageTemperature(filepath, unit string) (float64, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("File does not exist. You must first convert the file from celsius to fahrenheit.")
		fmt.Println("Exiting program.")
		return 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var sum float64
	var count int

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		}

		// Split the line into parts
		parts := strings.Split(strings.TrimSpace(line), ";")
		if len(parts) != 4 {
			// Skip line if it does not contain four parts
			fmt.Printf("Skipping line: %s", line)
			continue
		}

		// Parse the temperature
		temperature, err := strconv.ParseFloat(parts[3], 64)
		if err != nil {
			// Skip line if conversion fails
			fmt.Printf("Skipping line: %s", line)
			continue
		}

		// Add temperature to sum
		sum += temperature
		count++
	}

	// Calculate average temperature
	var average float64
	if count > 0 {
		average = sum / float64(count)
	}

	return average, nil
}

func generateOutputFile() error {
	inputFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	// get number of lines in file
	totalLines, err := yr.GetNumberOfLines(inputFile.Name())
	if err != nil {
		fmt.Println("Error counting lines:", err)
	} else {
		fmt.Println("Number of lines in inputfile:", totalLines)
	}

	lineCount := 0
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		if lineCount == 1 {
			// Write the first line to the output file as is
			_, err = writer.WriteString(line + "\n")
			if err != nil {
				return err
			}
			continue
		}

		// Process the line (convert temperature and format output)
		processedLine, err := yr.CelsiusToFahrenheitLine(line)
		if err != nil {
			return err
		}

		if lineCount < totalLines {
			// Write processed line to output file
			_, err = writer.WriteString(processedLine + "\n")
			if err != nil {
				return err
			}
		} else {
			// Write test string for the last line
			_, err = writer.WriteString("Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Kjell Lindberg")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}
