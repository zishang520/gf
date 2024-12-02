package pgsql

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gutil"
)

var (
	viewsSqlTmp = `
SELECT
	c.relname
FROM
	pg_class c
INNER JOIN pg_namespace n ON
	c.relnamespace = n.oid
WHERE
	n.nspname = '%s'
	AND c.relkind = 'v'
	%s
ORDER BY
	c.relname
`
)

func init() {
	var err error
	viewsSqlTmp, err = gdb.FormatMultiLineSqlToSingle(viewsSqlTmp)
	if err != nil {
		panic(err)
	}
}

// Views retrieves and returns the views of the current schema.
// It's mainly used in cli tool chain for automatically generating the models.
func (d *Driver) Views(ctx context.Context, schema ...string) (views []string, err error) {
	var (
		result     gdb.Result
		usedSchema = gutil.GetOrDefaultStr(d.GetConfig().Namespace, schema...)
	)
	if usedSchema == "" {
		usedSchema = defaultSchema
	}
	// DO NOT use `usedSchema` as parameter for function `SlaveLink`.
	link, err := d.SlaveLink(schema...)
	if err != nil {
		return nil, err
	}

	useRelpartbound := ""
	if gstr.CompareVersion(d.version(ctx, link), "10") >= 0 {
		useRelpartbound = "AND c.relpartbound IS NULL"
	}

	var query = fmt.Sprintf(
		viewsSqlTmp,
		usedSchema,
		useRelpartbound,
	)

	query, _ = gregex.ReplaceString(`[\n\r\s]+`, " ", gstr.Trim(query))
	result, err = d.DoSelect(ctx, link, query)
	if err != nil {
		return
	}
	for _, m := range result {
		for _, v := range m {
			views = append(views, v.String())
		}
	}
	return
}
