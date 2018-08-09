package main

import (
	"os"
	"time"

	"github.com/aequitas/tablewriter"
)

func main() {

	data := [][]string{
		[]string{"A", "The Good", "500"},
		[]string{"B", "The Very very Bad Man", "288"},
		[]string{"C", "The Ugly", "120"},
		[]string{"D", "The Gopher", "800"},
	}

	table := tablewriter.NewStreamWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sign", "Rating"})

	go table.Render() // Send output

	for _, v := range data {
		table.Append(v)
		time.Sleep(time.Second * 1)
	}

	table.Wg.Wait()
}
