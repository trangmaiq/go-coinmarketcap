package coinmarketcap

const (
	NumberOfSecoundsInDay = 86400
	NumberOfMinutesInDay  = 1440
	NumberHoursInDay      = 24

	DailyCreditLimitStarter       = 200
	MonthlyCreditLimitStarter     = 6000
	ApiCallRateLimitStarterPerMin = 10
)

const (
	Starter CmcPlan = iota
	Hobbyist
	Professional
	Enterprise
)

type CmcPlan int

func (p *CmcPlan) String(i int) string {
	return [...]string{"Starter Plan", "Hobbyist Plan", "Professional Plan", "Enterprise Plan"}[i]
}

type PlanOptions struct {
	PlanID int
	Keys   []string
}

type CreditLimit struct {
	DailyCreditLimit   int
	MonthlyCreditLimit int
	APICallRateLimit   int
}

func CalculateCreditLimitWithPlan(opt *PlanOptions) *CreditLimit {
	numberKeys := len(opt.Keys)
	if numberKeys < 1 {
		return nil
	}

	switch opt.PlanID {
	case int(Starter):
		return &CreditLimit{
			DailyCreditLimit:   numberKeys * DailyCreditLimitStarter,
			MonthlyCreditLimit: numberKeys * MonthlyCreditLimitStarter,
			APICallRateLimit:   numberKeys * ApiCallRateLimitStarterPerMin,
		}
	// Todo: Add another plan
	default:
		return nil
	}
}
