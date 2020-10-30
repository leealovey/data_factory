package mongoop

// DayAccount test
type DayAccount struct {
	ID                      string      `json:"_id"`
	AbnormalRateOfReturn    float64     `json:"abnormal_rate_of_return"`
	Accounts                interface{} `json:"accounts"`
	Available               float64     `json:"available"`
	Balance                 float64     `json:"balance"`
	BenchmarkCumulativeRate float64     `json:"benchmark_cumulative_rate"`
	BenchmarkRateOfReturn   float64     `json:"benchmark_rate_of_return"`
	CloseProfit             int64       `json:"close_profit"`
	Commission              float64     `json:"commission"`
	Currency                string      `json:"currency"`
	DailyRateOfReturn       float64     `json:"daily_rate_of_return"`
	DayProfit               float64     `json:"day_profit"`
	Deposit                 int64       `json:"deposit"`
	FloatProfit             int64       `json:"float_profit"`
	FrozenCommission        int64       `json:"frozen_commission"`
	FrozenMargin            int64       `json:"frozen_margin"`
	FrozenPremium           int64       `json:"frozen_premium"`
	Margin                  float64     `json:"margin"`
	PositionProfit          float64     `json:"position_profit"`
	PreBalance              float64     `json:"pre_balance"`
	Premium                 int64       `json:"premium"`
	RiskRatio               float64     `json:"risk_ratio"`
	StaticBalance           float64     `json:"static_balance"`
	StrategyCumulativeRate  float64     `json:"strategy_cumulative_rate"`
	StrategyID              string      `json:"strategy_id"`
	Timestamp               int64       `json:"timestamp"`
	TotalDeposit            int64       `json:"total_deposit"`
	TotalProfit             float64     `json:"total_profit"`
	TotalWithdraw           int64       `json:"total_withdraw"`
	TradeDay                string      `json:"trade_day"`
	WinOrLoseFlag           int64       `json:"win_or_lose_flag"`
	Withdraw                int64       `json:"withdraw"`
}
