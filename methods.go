package sqlite_qb

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var parameterize *regexp.Regexp

func init() {
	parameterize = regexp.MustCompile("(ʭ.*?ʭ)")
}

func toSQL(data map[string]string) (string, []interface{}, error) {
	var args []interface{}
	replacer := func(match string) string {
		cleaned := strings.Replace(match, "ʭ", "", -1)
		args = append(args, cleaned)

		return "?"
	}

	var fullStmt string

	sel, selOk := data["select"]
	if selOk {
		from, ok := data["from"]
		if !ok {
			return "", nil, errors.New("select statement missing a from")
		}

		fullStmt = fmt.Sprintf("%s %s", sel, from)

		where, ok := data["where"]
		if ok {
			fullStmt = fmt.Sprintf("%s %s", fullStmt, where)
		}

		// Parameterize
		paramed := parameterize.ReplaceAllStringFunc(fullStmt, replacer)

		if strings.Contains(paramed, "ʭ") {
			return "", nil, errors.New("illegal character")
		}

		return paramed, args, nil
	}

	return "", nil, nil
}

const (
	Eq = iota
	Ne
	Lt
	Gt
	Le
	Ge
)

func Compare(operator int, data map[string]interface{}) string {
	var stmts []string

	var opr string
	switch operator {
	case 0:
		opr = "="
	case 1:
		opr = "<>"
	case 2:
		opr = "<"
	case 3:
		opr = ">"
	case 4:
		opr = "<="
	case 5:
		opr = ">="
	}

	for k, v := range data {
		switch v.(type) {
		case string:
			stmts = append(stmts, fmt.Sprintf("'%s' %s 'ʭ%vʭ'", k, opr, v))
		default:
			stmts = append(stmts, fmt.Sprintf("'%s' %s ʭ%vʭ", k, opr, v))
		}

	}

	return strings.Join(stmts, " AND ")
}

func Like(data map[string]string) string {
	var stmts []string

	for k, v := range data {
		stmts = append(stmts, fmt.Sprintf("'%s' LIKE 'ʭ%%%v%%ʭ'", k, v))
	}

	return strings.Join(stmts, " AND ")
}

func Between(data map[string][]interface{}) string {
	var stmts []string
	for k, v := range data {
		stmts = append(stmts, fmt.Sprintf("'%s' BETWEEN ʭ%vʭ and ʭ%vʭ", k, v[0], v[1]))
	}

	return strings.Join(stmts, " AND ")
}

func In(data map[string][]interface{}) string {
	var stmts []string
	for k, v := range data {
		var list []string
		for _, z := range v {
			switch z.(type) {
			case string:
				list = append(list, fmt.Sprintf("'ʭ%sʭ'", z))
			default:
				list = append(list, fmt.Sprintf("ʭ%vʭ", z))
			}
		}

		stmts = append(stmts, fmt.Sprintf("'%s' IN (%s)", k, strings.Join(list, ",")))
	}

	return strings.Join(stmts, " AND ")
}
