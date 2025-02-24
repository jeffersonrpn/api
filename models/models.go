package models

import (
	"time"

	"github.com/dadosjusbr/proto/coleta"
	"github.com/dadosjusbr/storage"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// State - Struct cotains information of a state ans its agencies
type State struct {
	Name      string
	ShortName string
	FlagURL   string
	Agency    []AgencyBasic
}

// AgencyBasic - Basic information of a agency (name e category)
type AgencyBasic struct {
	Name           string
	FullName       string
	AgencyCategory string
}

// Employee - Represents an employee and his/her salary info
type Employee struct {
	Name   string
	Wage   float64
	Perks  float64
	Others float64
	Total  float64
	Type   string
	Active bool
}

// AgencySummary - Summary of an agency
type AgencySummary struct {
	FullName          string
	TotalEmployees    int
	TotalWage         float64
	TotalPerks        float64
	MaxWage           float64
	CrawlingTime      *timestamppb.Timestamp
	AgencyName        string
	TotalMembers      int
	TotalServants     int
	TotalInactives    int
	MaxPerk           float64
	TotalRemuneration float64
	HasNext           bool
	HasPrevious       bool
}

// AgencyTotalsYear - Represents the totals of an year
type AgencyTotalsYear struct {
	Year           int
	MonthTotals    []MonthTotals
	AgencyFullName string
	SummaryPackage *storage.Package `json:"SummaryPackage,omitempty"`
}

type ProcError struct {
	Stdout string `protobuf:"bytes,2,opt,name=stdout,proto3" json:"stdout,omitempty"` // String containing the standard output of the process.
	Stderr string `protobuf:"bytes,3,opt,name=stderr,proto3" json:"stderr,omitempty"` // String containing the standard error of the process.
}

// MonthTotals - Detailed info of a month (wage, perks, other)
type MonthTotals struct {
	Error              *ProcError
	Month              int
	BaseRemuneration   float64
	OtherRemunerations float64
}

// DataForChartAtAgencyScreen - contains all necessary data to load chart
type DataForChartAtAgencyScreen struct {
	Members     map[int]int
	Servers     map[int]int
	MaxSalary   float64
	PackageURL  string
	PackageHash string
}

// ProcInfoResult - contains information of the result of the process if something went wrong during parsing or crawling process
type ProcInfoResult struct {
	ProcInfo          *coleta.ProcInfo
	CrawlingTimestamp *timestamppb.Timestamp
}

// GeneralTotals - contains the summary from all DadosJusBr data
type GeneralTotals struct {
	AgencyAmount             int64
	MonthlyTotalsAmount      int64
	StartDate                time.Time
	EndDate                  time.Time
	RemunerationRecordsCount int
	GeneralRemunerationValue float64
}
