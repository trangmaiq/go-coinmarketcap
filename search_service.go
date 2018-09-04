package coinmarketcap

type SearchService service

type Search struct {
	ID   string `json:"_id,omitempty" bson:"_id"`
	Name struct {
		First  string `json:"first" bson:"first"`
		Middle string `json:"middle" bson:"middle"`
		Last   string `json:"last" bson:"last"`
	} `json:"name" bson:"name"`
	SignupPlanID int         `json:"signup_plan_id" bson:"signup_plan_id"`
	Notes        interface{} `json:"notes" bson:"notes"` // return []null - Todo: Check
	Status       struct {
		Log interface{} `json:"log" bson:"log"` // return []null - Todo: Check
	} `json:"status" bson:"status"`
	TimeCreated string  `json:"timeCreated" bson:"timeCreated"`
	User        User    `json:"user" bson:"user"`
	KeyPlan     KeyPlan `json:"keyPlan" bson:"keyPlan"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type KeyPlan struct {
	Key       string `json:"key" bson:"key"`
	Enabled   bool   `json:"enabled" bson:"enabled"`
	Timestamp struct {
		Created  string      `json:"created" bson:"created"`
		Updated  string      `json:"updated" bson:"updated"`
		LastUsed interface{} // sanbox return null - Todo: check
	} `json:"timestamp" bson:"timestamp"`
	Plan Plan `json:"plan" bson:"plan"`
}

type Plan struct {
	PlanID                 int         `json:"plan_id" bson:"plan_id"`
	Label                  string      `json:"label" bson:"label"`
	Description            string      `json:"description" bson:"description"`
	LimitDaily             int         `json:"limit_daily" bson:"limit_daily"`
	LimitMonthly           int         `json:"limit_monthly" bson:"limit_monthly"`
	HistoricalAccess       int         `json:"historical_access" bson:"historical_access"`
	PaymentRequired        bool        `json:"payment_required" bson:"payment_required"`
	PaymentComplete        bool        `json:"payment_complete" bson:"payment_complete"`
	PaymentExpireTimestamp interface{} `json:"payment_expire_timestamp" bson:"payment_expire_timestamp"` // return null - Todo: check
	RateLimitMinute        int         `json:"rate_limit_minute" bson:"rate_limit_minute"`
	EndpointWhitelist      interface{} `json:"endpoint_whitelist" bson:"endpoint_whitelist"` // return null
	ConvertRestriction     int         `json:"convert_restriction" bson:"convert_restriction"`
	LimitDailyHard         int         `json:"limit_daily_hard" bson:"limit_daily_hard"`
	LimitMonthlyHard       int         `json:"limit_monthly_hard" bson:"limit_monthly_hard"`
	BillingRateRearly      int         `json:"billing_rate_yearly" bson:"billing_rate_yearly"`
	BillingRateMonthly     int         `json:"billing_rate_monthly" bson:"billing_rate_monthly"`
}

type SearchParams struct {
	Cookie string
}

//https://sandbox.coinmarketcap.com/api/accounts/my/plan/stats
