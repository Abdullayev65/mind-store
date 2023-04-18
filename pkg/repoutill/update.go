package repoutill

import "fmt"

func UpdateSetColumn[C any](update *C, columnName string, setColumns *[]string, args *[]any, argNum *int) {
	if update == nil {
		return
	}

	*setColumns = append(*setColumns, fmt.Sprintf(columnName+"=$%d", *argNum))
	*args = append(*args, *update)
	*argNum++
}
