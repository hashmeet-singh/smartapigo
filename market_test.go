package smartapigo

import (
	"testing"
)

func (ts *TestSuite) TestGetLTP(t *testing.T) {
	t.Parallel()
	params := LTPParams{
		Exchange:      "NSE",
		TradingSymbol: "SBIN-EQ",
		SymbolToken:   "3045",
	}
	ltp, err := ts.TestConnect.GetLTP(params)
	if err != nil {
		t.Errorf("Error while fetching LTP. %v", err)
	}

	if ltp.Exchange == "" {
		t.Errorf("Error while exchange in LTP. %v", err)
	}

}

func (ts *TestSuite) TestGetHistoricalData(t *testing.T) {
	t.Parallel()
	params := HistoricalDataParams{
		Exchange:    "NSE",
		SymbolToken: "15083",
		Interval:    "FIVE_MINUTE",
		FromDate:    "2021-11-09 09:15",
		ToDate:      "2021-11-09 12:45",
	}
	historicalData, err := ts.TestConnect.GetHistoricalData(params)
	if err != nil {
		t.Errorf("Error while fetching historicalData. %v", err)
	}
	if len(historicalData) <= 0 {
		t.Errorf("Error while getting historicalData. %v", err)
	}

}
