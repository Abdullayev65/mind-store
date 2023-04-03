package repo

import "github.com/uptrace/bun"

func SetIfNotNil[T any](dst, src *T) {
	if src != nil {
		*dst = *src
	}
}

func WhereGroupAnd(q *bun.SelectQuery, query string, args ...interface{}) {
	q.WhereGroup(" AND ", func(subQuery *bun.SelectQuery) *bun.SelectQuery {
		return subQuery.Where(query, args...)
	})
}
