package service

import (
	"fmt"
	"strings"
)

// Not secure because of sql injection. TO-DO fix
func buildWhereConditionFromParams(filterParams map[string][]string) string {
	conds := []string{}

	allowedParams := map[string]interface{}{
		"id":                  nil,
		"created_at":          nil,
		"parent_task_id":      nil,
		"creater_id":          nil,
		"responsible_user_id": nil,
		"title":               nil,
		"description":         nil,
		"status":              nil,
		"task_group_id":       nil,
		"priority":            nil,
		"estimate_time":       nil,
		"time_spent":          nil,
		"deleted_at":          nil,
		"archived":            nil,
	}

	for param := range filterParams {
		_, ok := allowedParams[param]
		if !ok {
			continue
		}

		conds = append(conds, param+fmt.Sprintf(" = :%s", param))
	}

	return strings.Join(conds, " and ")
}
