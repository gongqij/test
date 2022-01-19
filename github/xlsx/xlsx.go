package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"time"
)

type structTest struct {
	Name       string    `xlsx:"0"`
	Date       time.Time `xlsx:"1"`
	BelongTo   string    `xlsx:"2"`
	IgnoredVal string    `xlsx:"3"`
	Boss       string    `xlsx:"4"`
	Type       string    `xlsx:"5"`
	Price      int       `xlsx:"6"`
	Num        int       `xlsx:"7"`
	RMB        int       `xlsx:"8"`
}

func main() {
	/*	structVal := structTest{
				IntVal:     16,
				StringVal:  "heyheyhey :)!",
				FloatVal:   3.14159216,
				IgnoredVal: 7,
				BoolVal:    true,
			}
			//create a new xlsx file and write a struct
			//in a new row
			f := xlsx.NewFile()
			sheet, _ := f.AddSheet("TestRead")
			row := sheet.AddRow()
			row.WriteStruct(&structVal, -1)

			//read the struct from the same row
			readStruct := &stru
		ctTest{}
			err := row.ReadStruct(readStruct)
			if err != nil {
				panic(err)
			}
			fmt.Println(readStruct)
			f.Save("./hahah.xlsx")*/
	readStruct := &structTest{}
	f, _ := xlsx.OpenFile("/home/gongqi/Desktop/test.xlsx")
	for i := 1; i < f.Sheets[2].MaxRow; i++ {
		r := f.Sheets[2].Row(i)
		_ = r.ReadStruct(readStruct)
		if readStruct.Name == "" {
			continue
		}
		fmt.Println(*readStruct)
	}
}
