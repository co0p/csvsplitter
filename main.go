package main

import "os"
import "flag"
import "fmt"
import "log"
import "encoding/csv"
import "github.com/fatih/color"

var fileName string

func init() {
	flag.StringVar(&fileName, "input", "", "The path to the csv file")
    flag.StringVar(&fileName, "i", "", "The path to the csv file (shorthand)")
}

func main() {
    flag.Parse()

    // make it pretty
    blueColor := color.New(color.FgCyan).Add(color.Underline)
    redColor := color.New(color.FgRed)

    blueColor.Println(".o( R3d!nG CSV l!k3 a Pr0 )o.")
	color.Unset();

    // make sure we have a file
    if len(fileName) < 1 {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

    // read the file already
    f, err := os.Open(fileName)
	redColor.Println("Done: OPENING FILE")


    // REAAAADDDDD
    if err != nil {
        redColor.Println(err)
        os.Exit(1)
    } else {
        r := csv.NewReader(f)
        defer f.Close()
        r.Comma = ';'
		r.LazyQuotes = true

        records, err := r.ReadAll()
    	if err != nil {
    		log.Fatal(err)
    	}
		redColor.Println("Done: READING")

    	fmt.Println(records)
        os.Exit(0)

    }
}
