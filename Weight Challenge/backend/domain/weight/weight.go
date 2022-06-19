package weight

import (
	"database/sql"
)

type Domain struct {
	db *sql.DB
}

type DomainItf interface {
	CreateWeightData(data WeightData) error
	EditWeightData(data WeightData) error
	GetAllWeightData() (result []WeightData, err error)
	ResetWeightData() error
}

func NewDomain(db *sql.DB) *Domain {
	return &Domain{
		db: db,
	}
}

func (d *Domain) CreateWeightData(data WeightData) error {
	_, err := d.db.Exec(QueryInsertWeightData, data.Date, data.Min, data.Max)

	if err != nil {
		return err
	}

	return nil
}

func (d *Domain) EditWeightData(data WeightData) error {
	_, err := d.db.Exec(QueryEditWeightData, data.Date, data.Min, data.Max, data.ID)
	if err != nil {
		return err
	}

	return nil
}

func (d *Domain) GetAllWeightData() (result []WeightData, err error) {
	rows, err := d.db.Query(QueryGetAllWeightData)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		weight := WeightData{}
		err = rows.Scan(&weight.ID, &weight.Date, &weight.Min, &weight.Max)
		if err != nil {
			return result, err
		}

		result = append(result, weight)
	}

	return result, nil
}

func (d *Domain) ResetWeightData() error {
	_, err := d.db.Exec(QueryDeleteAll)
	if err != nil {
		return err
	}

	return nil
}
