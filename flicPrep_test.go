package flicprep

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	fr "github.com/Ulbora/FileReader"
)

func TestFlicPrep_PrepRecords(t *testing.T) {
	var cr fr.CsvFileReader
	sourceFile, err := ioutil.ReadFile("../full_file.csv")
	fmt.Println("readFile err: ", err)
	rd := cr.GetNew()
	rec := rd.ReadCsvFile(sourceFile)
	fmt.Println("csv err: ", rec.CsvReadErr)
	fmt.Println("csv len: ", len(rec.CsvFileList))
	var fp FlicPrep
	recs := fp.PrepRecords(rec)
	if len(*recs) == 0 {
		t.Fail()
	}
}

func TestFlicPrep_buildDate2(t *testing.T) {
	var fp FlicPrep
	suc, m, y := fp.buildDate(2029, "9A")
	if !suc || y != 2029 || m != time.January {
		t.Fail()
	}

}

func TestFlicPrep_buildDate3(t *testing.T) {
	var fp FlicPrep
	suc, m, y := fp.buildDate(2021, "9B")
	if !suc || y != 2019 || m != time.February {
		t.Fail()
	}

}

func TestFlicPrep_buildDate4(t *testing.T) {
	var fp FlicPrep
	suc, m, y := fp.buildDate(2021, "6C")
	if suc || y != 0 || m != time.March {
		t.Fail()
	}

}

func TestFlicPrep_buildDate5(t *testing.T) {
	var fp FlicPrep
	suc, m, y := fp.buildDate(2121, "9D")
	if !suc || y != 2119 || m != time.April {
		t.Fail()
	}

}

func TestFlicPrep_buildDate6(t *testing.T) {
	var fp FlicPrep
	suc, m, y := fp.buildDate(2020, "7E")
	if !suc || y != 2017 || m != time.May {
		t.Fail()
	}

}

func TestFlicPrep_buildDate7(t *testing.T) {
	var fp FlicPrep
	suc, m, y := fp.buildDate(2129, "1F")
	if suc || y != 0 || m != time.June {
		t.Fail()
	}

}

func TestFlicPrep_buildDate8(t *testing.T) {
	var fp FlicPrep
	suc, m, y := fp.buildDate(2129, "1I")
	fmt.Println("month: ", m)
	if suc || y != 0 || m != 0 {
		t.Fail()
	}

}

func TestFlicPrep_buildDate(t *testing.T) {
	var fp FlicPrep
	suc, m, y := fp.buildDate(2020, "3G")
	if !suc || y != 2023 || m != time.July {
		t.Fail()
	}

}
