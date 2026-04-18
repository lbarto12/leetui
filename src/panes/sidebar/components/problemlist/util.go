package problemlist

import (
	"leetui/src/lib/graphqlapi/models"

	"charm.land/bubbles/v2/table"
)

func ConvertProblemsToTableRows(problems []models.Problem) []table.Row {
	ptabledata := []table.Row{}
	for _, problem := range problems {
		ptabledata = append(ptabledata, table.Row{
			problem.ID, problem.Difficulty, problem.Title,
		})
	}
	return ptabledata
}
