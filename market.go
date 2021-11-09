package smartapigo

import (
	"fmt"
	"net/http"
	"time"
)

// LTPResponse represents LTP API Response.
type LTPResponse struct {
	Exchange      string  `json:"exchange"`
	TradingSymbol string  `json:"tradingsymbol"`
	SymbolToken   string  `json:"symboltoken"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Close         float64 `json:"close"`
	Ltp           float64 `json:"ltp"`
}

// LTPParams represents parameters for getting LTP.
type LTPParams struct {
	Exchange      string `json:"exchange"`
	TradingSymbol string `json:"tradingsymbol"`
	SymbolToken   string `json:"symboltoken"`
}

type HistoricalData struct {
	Date   time.Time `json:"date"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Volume int       `json:"volume"`
}

type HistoricalDataParams struct {
	Exchange    string `json:"exchange"`
	SymbolToken string `json:"symboltoken"`
	Interval    string `json:"interval"`
	FromDate    string `json:"fromdate"`
	ToDate      string `json:"todate"`
}

// GetLTP gets Last Traded Price.
func (c *Client) GetLTP(ltpParams LTPParams) (LTPResponse, error) {
	var ltp LTPResponse
	params := structToMap(ltpParams, "json")
	err := c.doEnvelope(http.MethodPost, URILTP, params, nil, &ltp, true)
	return ltp, err
}

func (c *Client) formatHistoricalData(candles [][]interface{}) ([]HistoricalData, error) {
	var data []HistoricalData
	for _, candle := range candles {
		var (
			ds     string
			open   float64
			high   float64
			low    float64
			close  float64
			volume int
			ok     bool
		)
		if ds, ok = candle[0].(string); !ok {
			return data, NewError("DecodingException", fmt.Sprintf("Error decoding response `date`: %v", candle[0]), nil)
		}

		if open, ok = candle[1].(float64); !ok {
			return data, NewError("DecodingException", fmt.Sprintf("Error decoding response `open`: %v", candle[1]), nil)
		}

		if high, ok = candle[2].(float64); !ok {
			return data, NewError("DecodingException", fmt.Sprintf("Error decoding response `high`: %v", candle[2]), nil)
		}

		if low, ok = candle[3].(float64); !ok {
			return data, NewError("DecodingException", fmt.Sprintf("Error decoding response `low`: %v", candle[3]), nil)
		}

		if close, ok = candle[4].(float64); !ok {
			return data, NewError("DecodingException", fmt.Sprintf("Error decoding response `close`: %v", candle[4]), nil)
		}

		date, err := time.Parse("2006-01-02T15:04:05-07:00", ds)
		if err != nil {
			return data, NewError("DecodingException", fmt.Sprintf("Error decoding response: %v", err), nil)
		}

		data = append(data, HistoricalData{
			Date:   date,
			Open:   open,
			High:   high,
			Low:    low,
			Close:  close,
			Volume: volume,
		})

	}

	return data, nil
}

func (c *Client) GetHistoricalData(historicalDataParams HistoricalDataParams) ([]HistoricalData, error) {
	var resp [][]interface{}
	var data []HistoricalData
	params := structToMap(historicalDataParams, "json")
	err := c.doEnvelope(http.MethodPost, URIHistoricalData, params, nil, &resp, true)
	if err != nil {
		return data, nil
	}
	data, err = c.formatHistoricalData(resp)
	if err != nil {
		return data, err
	}

	return data, nil
}
