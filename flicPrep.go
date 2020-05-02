package flicprep

import (
	"strconv"
	"strings"
	"time"

	fr "github.com/Ulbora/FileReader"
)

//FlicPrep FlicPrep
type FlicPrep struct {
	flicMap map[string]Flic
}

//PrepRecords PrepRecords
func (p *FlicPrep) PrepRecords(files *fr.CsvFiles) *[]Flic {
	var rtn []Flic
	p.flicMap = make(map[string]Flic)
	if files != nil {
		//fmt.Println("csv len: ", len(files.CsvFileList))
		for i, cf := range files.CsvFileList {
			if i == 0 {
				continue
			}
			// fmt.Println("csv rec: ", cf)
			if cf[3] == "01" || cf[3] == "02" || cf[3] == "07" {
				var key = cf[8] + cf[9] + cf[10]
				key = strings.Replace(key, " ", "", -1)
				//fmt.Println("key: ", key)
				fflic := p.flicMap[key]
				// fmt.Println("fflic: ", fflic)
				if fflic.Type != 0 {
					fty, ierr := strconv.Atoi(cf[3])
					if ierr == nil {
						if fflic.Type > fty {
							suc, newflic := p.parseRecord(cf)
							//fmt.Println("newflic: ", newflic)
							if suc {
								p.flicMap[key] = newflic
							}

						}
					}
				} else {
					suc, newflic := p.parseRecord(cf)
					//fmt.Println("newflic: ", newflic)
					if suc {
						p.flicMap[key] = newflic
					}
				}
			}
		}

		//fmt.Println("map len: ", len(p.flicMap))
		for _, v := range p.flicMap {
			rtn = append(rtn, v)
		}
		//fmt.Println("rtn len: ", len(rtn))
		// fmt.Println("rtn 1: ", rtn[0])
		// fmt.Println("rtn 2: ", rtn[20])
		//fmt.Println("rtn: ", rtn)
	}

	return &rtn
}

func (p *FlicPrep) parseRecord(cf []string) (bool, Flic) {
	var rtn Flic
	var key = cf[0] + cf[1] + cf[2] + cf[3] + cf[4] + cf[5]
	key = strings.Replace(key, " ", "", -1)
	rtn.Key = key
	//fmt.Println("key: ", key)
	lic := p.buildLic(cf)
	//fmt.Println("license no:", lic)
	rtn.Lic = lic
	fty, _ := strconv.Atoi(cf[3])
	rtn.Type = fty
	rtn.LicName = cf[6]
	rtn.BusName = cf[7]
	rtn.PremiseAddress = cf[8] + "\n" + cf[9] + ", " + cf[10] + " " + cf[11]
	rtn.PremiseZip = cf[11]
	// fmt.Println("prim address:")
	// fmt.Println(rtn.PremiseAddress)
	rtn.MailingAddress = cf[12] + "\n" + cf[13] + ", " + cf[14] + " " + cf[15]
	//fmt.Println("mail address:")
	//fmt.Println(rtn.MailingAddress)
	//fmt.Println("lic: ", lic)
	///fmt.Println("cf[4]: ", string(cf[4][0]))
	tphone := cf[16]
	tphone = strings.Replace(tphone, " ", "", -1)
	//fmt.Println("tphone:", tphone)
	if len(tphone) == 10 {
		rtn.Phone = string(tphone[0:3]) + "-" + string(tphone[3:6]) + "-" + string(tphone[6:])
	} else {
		rtn.Phone = tphone
	}

	//fmt.Println("tphone format:", rtn.Phone)

	//cf4, _ := strconv.Atoi(string(cf[4][0]))
	y, _, _ := time.Now().Date()
	suc, mon, year := p.buildDate(y, cf[4])
	if suc {
		expDate := time.Date(year, mon, 1, 0, 0, 0, 0, time.UTC)
		rtn.ExpDate = expDate
	}

	//fmt.Println("expDate: ", expDate)
	//////rtn.ExpDate = expd

	return suc, rtn
}

func (p *FlicPrep) buildLic(cf []string) string {
	var lic string
	lic = (cf[0] + "-")

	if len(cf[1]) == 1 {
		lic += ("0" + cf[1] + "-")
	} else {
		lic += (cf[1] + "-")
	}

	if len(cf[2]) == 1 {
		lic += ("00" + cf[2] + "-")
	} else if len(cf[2]) == 2 {
		lic += ("0" + cf[2] + "-")
	} else {
		lic += (cf[2] + "-")
	}

	if len(cf[3]) == 1 {
		lic += ("0" + cf[3] + "-")
	} else {
		lic += (cf[3] + "-")
	}

	lic += (cf[4] + "-")

	if len(cf[5]) == 1 {
		lic += ("0000" + cf[5])
	} else if len(cf[5]) == 2 {
		lic += ("000" + cf[5])
	} else if len(cf[5]) == 3 {
		lic += ("00" + cf[5])
	} else if len(cf[5]) == 4 {
		lic += ("0" + cf[5])
	} else {
		lic += (cf[5])
	}
	return lic
}

func (p *FlicPrep) buildDate(currentYear int, cf4 string) (suc bool, month time.Month, year int) {
	suc = true
	var licYear int
	currentYearStr := strconv.Itoa(currentYear)

	currentYearMin3Str := strconv.Itoa(currentYear - 3)

	testYearStr := string(currentYearStr[0:3]) + string(cf4[0])
	testYear, _ := strconv.Atoi(testYearStr)

	if (testYear - currentYear) > 3 {
		ly := string(currentYearMin3Str[0:3]) + string(cf4[0])
		lyint, _ := strconv.Atoi(ly)
		licYear = lyint
	} else {
		licYear = testYear
	}

	if licYear < (currentYear-3) || licYear > (currentYear+3) {
		licYear = 0
		suc = false
	}
	var monthVals = map[string]time.Month{
		"A": time.January,
		"B": time.February,
		"C": time.March,
		"D": time.April,
		"E": time.May,
		"F": time.June,
		"G": time.July,
		"H": time.August,
		"J": time.September,
		"K": time.October,
		"L": time.November,
		"M": time.December,
	}
	month = monthVals[string(cf4[1])]

	if month == 0 {
		suc = false
	}
	year = licYear

	return suc, month, year
}

//GetNew GetNew
func (p *FlicPrep) GetNew() RecordPrep {
	return p
}
