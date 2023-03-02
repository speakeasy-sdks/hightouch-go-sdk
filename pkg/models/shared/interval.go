package shared

type Interval struct {
	Quantity float64          `json:"quantity"`
	Unit     IntervalUnitEnum `json:"unit"`
}
