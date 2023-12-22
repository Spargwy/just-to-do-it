package service

import (
	"fmt"
	"strings"
)

func buildWhereConditionFromParams(filterParams map[string][]string) string {
	whereCondition := strings.Builder{}

	paramsLength := len(filterParams)
	i := 1
	for param, value := range filterParams {
		if i == paramsLength {
			whereCondition.WriteString(fmt.Sprintf("%s = '%s'", param, value[0]))
		} else {
			whereCondition.WriteString(fmt.Sprintf("%s = '%s' and ", param, value[0]))
		}
		i++
	}

	return whereCondition.String()
}
