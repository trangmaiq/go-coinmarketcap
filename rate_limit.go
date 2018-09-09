package coinmarketcap

const (
	NumberOfSecoundsInDay = 86400
	NumberOfMinutesInDay  = 1440
	NumberHoursInDay      = 24

	DailyCreditLimitStarter       = 200
	MonthlyCreditLimitStarter     = 6000
	ApiCallRateLimitStarterPerMin = 10

	DailyCreditLimitHobbyist       = 2000
	MonthlyCreditLimitHobbyist     = 60000
	ApiCallRateLimitHobbyistPerMin = 30

	DailyCreditLimitStandard       = 10000
	MonthlyCreditLimitStandard     = 300000
	ApiCallRateLimitStandardPerMin = 30

	DailyCreditLimitProfessional       = 100000
	MonthlyCreditLimitProfessional     = 3000000
	ApiCallRateLimitProfessionalPerMin = 60
)

const (
	Starter CmcPlan = iota
	Hobbyist
	Standard
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
	case int(Hobbyist):
		return &CreditLimit{
			DailyCreditLimit:   numberKeys * DailyCreditLimitHobbyist,
			MonthlyCreditLimit: numberKeys * MonthlyCreditLimitHobbyist,
			APICallRateLimit:   numberKeys * ApiCallRateLimitHobbyistPerMin,
		}
	case int(Standard):
		return &CreditLimit{
			DailyCreditLimit:   numberKeys * DailyCreditLimitStandard,
			MonthlyCreditLimit: numberKeys * MonthlyCreditLimitStandard,
			APICallRateLimit:   numberKeys * ApiCallRateLimitStandardPerMin,
		}
	case int(Professional):
		return &CreditLimit{
			DailyCreditLimit:   numberKeys * DailyCreditLimitProfessional,
			MonthlyCreditLimit: numberKeys * MonthlyCreditLimitProfessional,
			APICallRateLimit:   numberKeys * ApiCallRateLimitProfessionalPerMin,
		}
	default:
		return nil
	}
}
