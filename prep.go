package flicprep

import (
	"time"

	fr "github.com/Ulbora/FileReader"
)

//Flic Flic
type Flic struct {
	Key            string
	Lic            string
	Type           int
	ExpDate        time.Time
	LicName        string
	BusName        string
	PremiseAddress string
	PremiseZip     string
	MailingAddress string
	Phone          string
}

//RecordPrep RecordPrep
type RecordPrep interface {
	PrepRecords(files *fr.CsvFiles) *[]Flic
}

//build map by address store lic type in object with flic in map
//FoundFlic{
// 	licType int
// 	flic Flic
//}
//before insert search map if lic type num is less, swap the map value
//this eliminates duplicates

//go mod init github.com/Ulbora/FlicPrep
