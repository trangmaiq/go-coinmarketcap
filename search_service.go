package coinmarketcap

import (
	"context"
)

const (
	pathSearchInfo   = "accounts/my/plan/info"
	pathSearchStatus = "accounts/my/plan/stats"
)

type SearchService service

type SearchResult struct {
	ID   *string `json:"_id,omitempty" bson:"_id"`
	Name *struct {
		First  *string `json:"first,omitempty" bson:"first,omitempty"`
		Middle *string `json:"middle,omitempty" bson:"middle,omitempty"`
		Last   *string `json:"last,omitempty" bson:"last,omitempty"`
	} `json:"name,omitempty" bson:"name,omitempty"`
	SignupPlanID *int         `json:"signup_plan_id,omitempty" bson:"signup_plan_id,omitempty"`
	Notes        *interface{} `json:"notes,omitempty" bson:"notes,omitempty"` // return []null - Todo: Check
	Status       *struct {
		Log *interface{} `json:"log,omitempty" bson:"log,omitempty"` // return []null - Todo: Check
	} `json:"status,omitempty" bson:"status,omitempty"`
	TimeCreated *string  `json:"timeCreated,omitempty" bson:"timeCreated,omitempty"`
	User        *User    `json:"user,omitempty" bson:"user,omitempty"`
	KeyPlan     *KeyPlan `json:"keyPlan,omitempty" bson:"keyPlan,omitempty"`

	Day          *DailyCredits  `json:"day,omitempty" bson:"day,omitempty"`
	Yesterday    *DailyCredits  `json:"yesterday,omitempty" bson:"yesterday,omitempty"`
	Month        *DailyCredits  `json:"month,omitempty" bson:"month,omitempty"`
	UniqueIps    []*string      `json:"unique_ips,omitempty" bson:"unique_ips,omitempty"`
	LastApiCalls []*LastApiCall `json:"last_api_calls,omitempty" bson:"last_api_calls,omitempty"`
}

type LastApiCall struct {
	Date     *string `json:"date,omitempty" bson:"date,omitempty"`
	IP       *string `json:"ip,omitempty" bson:"ip,omitempty"`
	HttpCode *string `json:"httpCode,omitempty" bson:"httpCode,omitempty"`
	URL      *string `json:"url,omitempty" bson:"url,omitempty"`
	Credits  *int    `json:"credits,omitempty" bson:"credits,omitempty"`
	Elapsed  *int    `json:"elapsed,omitempty" bson:"elapsed,omitempty"`
}

type User struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type KeyPlan struct {
	Key       *string `json:"key,omitempty" bson:"key,omitempty"`
	Enabled   *bool   `json:"enabled,omitempty" bson:"enabled,omitempty"`
	Timestamp *struct {
		Created  *string      `json:"created,omitempty" bson:"created,omitempty"`
		Updated  *string      `json:"updated,omitempty" bson:"updated,omitempty"`
		LastUsed *interface{} // sanbox return null - Todo: check
	} `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Plan *Plan `json:"plan,omitempty" bson:"plan,omitempty"`
}

type Plan struct {
	PlanID                 *int         `json:"plan_id,omitempty" bson:"plan_id,omitempty"`
	Label                  *string      `json:"label,omitempty" bson:"label,omitempty"`
	Description            *string      `json:"description,omitempty" bson:"description,omitempty"`
	LimitDaily             *int         `json:"limit_daily,omitempty" bson:"limit_daily,omitempty"`
	LimitMonthly           *int         `json:"limit_monthly,omitempty" bson:"limit_monthly,omitempty"`
	HistoricalAccess       *int         `json:"historical_access,omitempty" bson:"historical_access,omitempty"`
	PaymentRequired        *bool        `json:"payment_require,omitemptyd" bson:"payment_required,omitempty"`
	PaymentComplete        *bool        `json:"payment_complete,omitempty" bson:"payment_complete,omitempty"`
	PaymentExpireTimestamp *interface{} `json:"payment_expire_timestamp,omitempty" bson:"payment_expire_timestamp,omitempty"` // return null - Todo: check
	RateLimitMinute        *int         `json:"rate_limit_minute,omitempty" bson:"rate_limit_minute,omitempty"`
	EndpointWhitelist      *interface{} `json:"endpoint_whitelist,omitempty" bson:"endpoint_whitelist,omitempty"` // return null
	ConvertRestriction     *int         `json:"convert_restriction,omitempty" bson:"convert_restriction,omitempty"`
	LimitDailyHard         *int         `json:"limit_daily_hard,omitempty" bson:"limit_daily_hard,omitempty"`
	LimitMonthlyHard       *int         `json:"limit_monthly_hard,omitempty" bson:"limit_monthly_hard,omitempty"`
	BillingRateRearly      *int         `json:"billing_rate_yearly,omitempty" bson:"billing_rate_yearly,omitempty"`
	BillingRateMonthly     *int         `json:"billing_rate_monthly,omitempty" bson:"billing_rate_monthly,omitempty"`
}

type DailyCredits struct {
	CreditsUsed      *int `json:"credits_used,omitempty" bson:"credits_used,omitempty"`
	TotalCallsCount  *int `json:"total_calls_count,omitempty" bson:"total_calls_count,omitempty"`
	UniqueCallsCount *int `json:"unique_calls_count,omitempty" bson:"unique_calls_count,omitempty"`
}

type SearchParams struct {
	Cookie string
}

func (ss *SearchService) AccountInformation(ctx context.Context, params SearchParams) (*SearchResult, *Response, error) {
	u := pathSearchInfo

	req, errNewReq := ss.client.NewSearchRequest("GET", u, nil)
	if errNewReq != nil {
		return nil, nil, errNewReq
	}

	req.Header.Set("Cookie", params.Cookie)

	search := &SearchResult{}
	resp, errResp := ss.client.Do(ctx, req, search)
	if errResp != nil {
		return nil, nil, errResp
	}

	return search, resp, nil
}

func (ss *SearchService) AccountStatus(ctx context.Context, params SearchParams) (*SearchResult, *Response, error) {
	u := pathSearchStatus

	req, errNewReq := ss.client.NewSearchRequest("GET", u, nil)
	if errNewReq != nil {
		return nil, nil, errNewReq
	}

	req.Header.Set("Cookie", params.Cookie)

	search := &SearchResult{}
	resp, errResp := ss.client.Do(ctx, req, search)
	if errResp != nil {
		return nil, nil, errResp
	}
	return search, resp, nil
}
