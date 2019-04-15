package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/coc1961/percentil/percentil"
)

func main() {
	var pointerPercentil = flag.Int("p", 0, "Percentil a Calcular debe ser un número > 0 y < 100")
	var pointerInputFile = flag.String("f", "", "Nombre de Archivo que contiene la tabla de datos, si se indica - se lee desde stdin")

	flag.Parse()

	if *pointerPercentil < 1 || *pointerPercentil > 99 {
		flag.PrintDefaults()
		os.Exit(1)
		return
	}
	if *pointerInputFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
		return
	}

	var fReader io.Reader = bufio.NewReader(os.Stdin)
	if *pointerInputFile != "-" {
		var err error
		fReader, err = os.Open(*pointerInputFile)
		if err != nil {
			fmt.Fprintf(flag.CommandLine.Output(), "%s\n\n", err.Error())
			os.Exit(1)
		}
	}
	reader := csv.NewReader(fReader)
	reader.Comma = ','
	reader.FieldsPerRecord = -1

	arr := make([][]string, 0)
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Fprintf(flag.CommandLine.Output(), "%s\n\n", error.Error())
			os.Exit(1)
		}
		arr = append(arr, line)
	}
	if len(arr) == 0 {
		return
	}

	per := percentil.New()

	if len(arr[0]) == 1 {
		//No Agrupado
		data := make([]float64, 0, len(arr))
		for _, x := range arr {
			if fl, err := strconv.ParseFloat(strings.Trim(x[0], " "), 64); err == nil {
				data = append(data, fl)
			} else {
				fmt.Fprintf(flag.CommandLine.Output(), "%s\n\n", err.Error())
				os.Exit(1)
			}
		}
		if per.SetTable(data).Error() == nil {
			if per.Calc(*pointerPercentil).Error() == nil {
				fmt.Fprintf(os.Stdout, "%.2f", per.Result())
				return
			}
		}
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n\n", per.Error().Error())
		os.Exit(1)
	} else if len(arr[0]) == 2 {
		//Agrupado
		data := make([][]float64, 0, len(arr))
		for _, x := range arr {
			if len(x) < 2 {
				continue
			}
			if fl, err := strconv.ParseFloat(strings.Trim(x[0], " "), 64); err == nil {
				if fl1, err := strconv.ParseFloat(strings.Trim(x[1], " "), 64); err == nil {
					data = append(data, []float64{fl, fl1})
				} else {
					fmt.Fprintf(flag.CommandLine.Output(), "%s\n\n", err.Error())
					os.Exit(1)
				}
			} else {
				fmt.Fprintf(flag.CommandLine.Output(), "%s\n\n", err.Error())
				os.Exit(1)
			}
		}
		if per.SetTable(data).Error() == nil {
			if per.Calc(*pointerPercentil).Error() == nil {
				fmt.Fprintf(os.Stdout, "%.2f", per.Result())
				return
			}
		}
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n\n", per.Error().Error())
		os.Exit(1)
	}

}
