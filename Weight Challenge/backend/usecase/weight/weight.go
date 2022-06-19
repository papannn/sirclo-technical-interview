package ucweight

import (
	"github.com/backend/domain/weight"
)

type Usecase struct {
	weightDomain weight.DomainItf
}

func NewUsecase(weightDomain weight.DomainItf) *Usecase {
	return &Usecase{
		weightDomain: weightDomain,
	}
}

func (uc *Usecase) CreateWeight(weightData weight.WeightData) error {
	err := uc.weightDomain.CreateWeightData(weightData)
	if err != nil {
		return err
	}

	return nil
}

func (uc *Usecase) GetAllWeightData() (resp WeightResp, err error) {
	weightData, err := uc.weightDomain.GetAllWeightData()
	if err != nil {
		return resp, err
	}

	if len(weightData) == 0 {
		weightData = make([]weight.WeightData, 0)
	}

	return BuildWeightResp(weightData), err
}

func (uc *Usecase) EditWeightData(weightData weight.WeightData) error {
	err := uc.weightDomain.EditWeightData(weightData)
	if err != nil {
		return err
	}

	return nil
}

// E2E Testing helper
func (uc *Usecase) Reset() error {
	err := uc.weightDomain.ResetWeightData()
	if err != nil {
		return err
	}

	return nil
}
