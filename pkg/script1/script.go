package main

import (
    "encoding/csv"
    "fmt"
    "os"
	"io"
)

func main() {
    // Open the file
    file, err := os.Open("./script1/file.csv")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }

    // Create a CSV reader
    reader := csv.NewReader(file)

    // Read the rows
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }

        // Print the row
        for _, value := range record {
            fmt.Println(value)
        }
    }
}