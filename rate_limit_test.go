package coinmarketcap

import (
	"reflect"
	"testing"
)

func TestCalculateCreditLimitWithPlan(t *testing.T) {
	type args struct {
		opt *PlanOptions
	}
	tests := []struct {
		name string
		args args
		want *CreditLimit
	}{
		{"Test case 1",
			args{
				opt: &PlanOptions{PlanID: 0, Keys: []string{"1", "2"}}},
			&CreditLimit{
				DailyCreditLimit:   400,
				MonthlyCreditLimit: 12000,
				APICallRateLimit:   20,
			}},
		{"Test case 2",
			args{
				opt: &PlanOptions{PlanID: 0, Keys: []string{}}},
			nil},
		{"Test case 3",
			args{
				opt: &PlanOptions{PlanID: 1, Keys: []string{"2", "3", "4"}}},
			&CreditLimit{
				DailyCreditLimit:   6000,
				MonthlyCreditLimit: 180000,
				APICallRateLimit:   90,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateCreditLimitWithPlan(tt.args.opt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateCreditLimitWithPlan() = %v, want %v", got, tt.want)
			}
		})
	}
}
