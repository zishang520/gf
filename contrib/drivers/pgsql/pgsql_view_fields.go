// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package pgsql

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gutil"
)

var (
	viewFieldsSqlTmp = `
SELECT a.attname AS field, t.typname AS type, a.attnotnull as null,
       '' as key, -- Views don't have keys
       ic.column_default as default_value, b.description as comment,
       coalesce(character_maximum_length, numeric_precision, -1) as length,
       numeric_scale as scale
FROM pg_attribute a
         LEFT JOIN pg_class c ON a.attrelid = c.oid
         LEFT JOIN pg_description b ON a.attrelid = b.objoid AND a.attnum = b.objsubid
         LEFT JOIN pg_type t ON a.atttypid = t.oid
         LEFT JOIN information_schema.columns ic ON ic.column_name = a.attname AND ic.table_name = c.relname
WHERE c.relname = '%s' AND c.relkind = 'v' AND a.attisdropped IS FALSE AND a.attnum > 0
ORDER BY a.attnum`
)

func init() {
	var err error
	viewFieldsSqlTmp, err = gdb.FormatMultiLineSqlToSingle(viewFieldsSqlTmp)
	if err != nil {
		panic(err)
	}
}

// ViewFields retrieves and returns the fields' information of specified view of current schema.
func (d *Driver) ViewFields(ctx context.Context, view string, schema ...string) (fields map[string]*gdb.TableField, err error) {
	var (
		result       gdb.Result
		link         gdb.Link
		usedSchema   = gutil.GetOrDefaultStr(d.GetSchema(), schema...)
		structureSql = fmt.Sprintf(viewFieldsSqlTmp, view)
	)
	if link, err = d.SlaveLink(usedSchema); err != nil {
		return nil, err
	}

	// First, check if the view exists
	checkViewSQL := fmt.Sprintf("SELECT 1 FROM information_schema.views WHERE table_name = '%s' AND table_schema = '%s'", view, usedSchema)
	result, err = d.DoSelect(ctx, link, checkViewSQL)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("view %s does not exist in schema %s", view, usedSchema)
	}

	// If the view exists, get its column information
	result, err = d.DoSelect(ctx, link, structureSql)
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gdb.TableField)
	for index, m := range result {
		name := m["field"].String()
		// Filter duplicated fields.
		if _, ok := fields[name]; ok {
			continue
		}
		fields[name] = &gdb.TableField{
			Index:   index,
			Name:    name,
			Type:    m["type"].String(),
			Null:    !m["null"].Bool(),
			Key:     "", // Views don't have keys
			Default: m["default_value"].Val(),
			Comment: m["comment"].String(),
		}
	}
	return fields, nil
}
