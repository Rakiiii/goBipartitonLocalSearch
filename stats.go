package lspartitioninglib

const (
	AmountOfItterations    = "amountofitteration"
	AmountOfTrueMark       = "amountoftruemark"
	AmountOfMarkCount      = "amountofmarkcount"
	AmountOfFalseMark      = "amountoffalsemark"
	AmountOfDisbalanceFail = "amountofdisbfail"
	AmountOfParamFail      = "amountofparamfail"
)

type stats struct {
	m map[string]int64
}

func (s *stats) GetStats() *map[string]int64 {
	return &(s.m)
}

var Statistic = stats{m: make(map[string]int64)}
