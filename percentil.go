package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/coc1961/percentil/internal/percentil"
)

func main() {
	reader := csv.NewReader(bufio.NewReader(os.Stdin))
	reader.Comma = ','
	reader.FieldsPerRecord = -1

	arr := make([][]string, 0)
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		arr = append(arr, line)
	}
	if len(arr) == 0 {
		return
	}

	per := percentil.New()
	_ = per

	if len(arr[0]) == 1 {
		//No Agrupado
		data := make([]float64, 0, len(arr))
		for _, x := range arr {
			if fl, err := strconv.ParseFloat(x[0], 64); err == nil {
				data = append(data, fl)
			}
		}
		per.SetTable(data)
		fmt.Println(per.Calc(25))
	} else if len(arr[0]) == 2 {
		//Agrupado
		data := make([][]float64, 0, len(arr))
		for _, x := range arr {
			if len(x) < 2 {
				continue
			}
			fl, err := strconv.ParseFloat(x[0], 64)
			fl1, err1 := strconv.ParseFloat(x[1], 64)
			if err == nil && err1 == nil {
				data = append(data, []float64{fl, fl1})
			}
			per.SetTable(data)
			fmt.Println(per.Calc(25))
		}
	}

}
