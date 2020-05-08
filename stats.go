package lspartitioninglib

const (
	AmountOfItterations    = "amountofitteration"
	AmountOfTrueMark       = "amountoftruemark"
	AmountOfMarkCount      = "amountofmarkcount"
	AmountOfFalseMark      = "amountoffalsemark"
	AmountOfDisbalanceFail = "amountofdisbfail"
	AmountOfParamFail      = "amountofparamfail"
	OverallMarkDerivative  = "overallmarkderivative"
	MarkOneDerivative      = "markonederivative"
)

type stats struct {
	m map[string]int64
}

func (s *stats) GetStats() *map[string]int64 {
	return &(s.m)
}

func (s *stats) Zero() {
	for k, _ := range s.m {
		s.m[k] = 0
	}
}

var Statistic = stats{m: make(map[string]int64)}
