// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package mysql

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gutil"
)

var (
	viewFieldsSqlByMariadb = `
SELECT
    c.COLUMN_NAME AS 'Field',
    ( CASE WHEN ch.CHECK_CLAUSE LIKE 'json_valid%%' THEN 'json' ELSE c.COLUMN_TYPE END ) AS 'Type',
    c.COLLATION_NAME AS 'Collation',
    c.IS_NULLABLE AS 'Null',
    c.COLUMN_KEY AS 'Key',
    ( CASE WHEN c.COLUMN_DEFAULT = 'NULL' OR c.COLUMN_DEFAULT IS NULL THEN NULL ELSE c.COLUMN_DEFAULT END) AS 'Default',
    c.EXTRA AS 'Extra',
    c.PRIVILEGES AS 'Privileges',
    c.COLUMN_COMMENT AS 'Comment' 
FROM
    information_schema.COLUMNS AS c
    LEFT JOIN information_schema.CHECK_CONSTRAINTS AS ch ON c.TABLE_NAME = ch.TABLE_NAME 
    AND c.COLUMN_NAME = ch.CONSTRAINT_NAME 
WHERE
    c.TABLE_SCHEMA = '%s' 
    AND c.TABLE_NAME = '%s'
    AND c.TABLE_NAME IN (SELECT TABLE_NAME FROM information_schema.VIEWS WHERE TABLE_SCHEMA = '%s')
ORDER BY c.ORDINAL_POSITION`
)

func init() {
	var err error
	viewFieldsSqlByMariadb, err = gdb.FormatMultiLineSqlToSingle(viewFieldsSqlByMariadb)
	if err != nil {
		panic(err)
	}
}

// ViewFields retrieves and returns the fields' information of specified view of current schema.
func (d *Driver) ViewFields(ctx context.Context, view string, schema ...string) (fields map[string]*gdb.TableField, err error) {
	var (
		result        gdb.Result
		link          gdb.Link
		usedSchema    = gutil.GetOrDefaultStr(d.GetSchema(), schema...)
		viewFieldsSql string
	)
	if link, err = d.SlaveLink(usedSchema); err != nil {
		return nil, err
	}
	dbType := d.GetConfig().Type
	switch dbType {
	case "mariadb":
		viewFieldsSql = fmt.Sprintf(viewFieldsSqlByMariadb, usedSchema, view, usedSchema)
	default:
		// Check if the view exists
		checkViewSql := fmt.Sprintf("SELECT TABLE_NAME FROM information_schema.VIEWS WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s'", usedSchema, view)
		checkResult, err := d.DoSelect(ctx, link, checkViewSql)
		if err != nil {
			return nil, err
		}
		if len(checkResult) == 0 {
			return nil, fmt.Errorf("view %s does not exist in schema %s", view, usedSchema)
		}
		viewFieldsSql = fmt.Sprintf(`SHOW FULL COLUMNS FROM %s`, d.QuoteWord(view))
	}

	result, err = d.DoSelect(ctx, link, viewFieldsSql)
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gdb.TableField)
	for i, m := range result {
		fields[m["Field"].String()] = &gdb.TableField{
			Index:   i,
			Name:    m["Field"].String(),
			Type:    m["Type"].String(),
			Null:    m["Null"].Bool(),
			Key:     m["Key"].String(),
			Default: m["Default"].Val(),
			Extra:   m["Extra"].String(),
			Comment: m["Comment"].String(),
		}
	}
	return fields, nil
}
