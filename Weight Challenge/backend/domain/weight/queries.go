package weight

const (
	QueryInsertWeightData = `
		INSERT INTO weight_data (
			date,
			min,
			max
		) VALUES (
			$1,
			$2,
			$3
		);
	`

	QueryEditWeightData = `
		UPDATE weight_data 
			SET
				date = $1,
				min = $2,
				max = $3
			WHERE id = $4;
	`

	QueryGetAllWeightData = `
		SELECT
			id,
			date,
			min,
			max
		FROM
			weight_data
		ORDER BY date DESC;
	`

	QueryDeleteAll = `
		DELETE FROM weight_data
	`
)
