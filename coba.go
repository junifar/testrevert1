package testrever1

type (
	MyStruct struct {
		Name    string
		Address string
	}

	Districts struct {
		DistrictID   int64  `db:"district_id"`
		DistrictName string `db:"district_name"`
		test         string
	}
)
