package prices

import (
	"fmt"
	"price-calculator/conversions"
	"price-calculator/iomanager"
)

type TaxPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversions.StringsToFloat(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadData()

	if err != nil {
		errorChan <- err
		return
	}
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxedPrice)
	}
	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)
	doneChan <- true
}

func NewTaxPriceJob(io iomanager.IOManager, taxRate float64) *TaxPriceJob {
	return &TaxPriceJob{
		IOManager:   io,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
