// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package clickhouse

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gutil"
)

const (
	viewFieldsColumns = `name,position,default_expression,comment,type`
)

// ViewFields retrieves and returns the fields' information of specified view of current schema.
func (d *Driver) ViewFields(ctx context.Context, view string, schema ...string) (fields map[string]*gdb.TableField, err error) {
	var (
		result     gdb.Result
		link       gdb.Link
		usedSchema = gutil.GetOrDefaultStr(d.GetSchema(), schema...)
	)
	if link, err = d.SlaveLink(usedSchema); err != nil {
		return nil, err
	}

	// First, check if the view exists
	checkViewSQL := fmt.Sprintf("SELECT name FROM system.tables WHERE database = '%s' AND name = '%s' AND engine = 'View'", usedSchema, view)
	checkResult, err := d.DoSelect(ctx, link, checkViewSQL)
	if err != nil {
		return nil, err
	}
	if len(checkResult) == 0 {
		return nil, fmt.Errorf("view %s does not exist in schema %s", view, usedSchema)
	}

	var (
		getColumnsSql = fmt.Sprintf(
			"SELECT %s FROM system.columns WHERE database = '%s' AND table = '%s'",
			viewFieldsColumns, usedSchema, view,
		)
	)
	result, err = d.DoSelect(ctx, link, getColumnsSql)
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gdb.TableField)
	for _, m := range result {
		var (
			isNull    = false
			fieldType = m["type"].String()
		)
		// in clickhouse, field type like is Nullable(int)
		fieldsResult, _ := gregex.MatchString(`^Nullable\((.*?)\)`, fieldType)
		if len(fieldsResult) == 2 {
			isNull = true
			fieldType = fieldsResult[1]
		}
		position := m["position"].Int()
		if result[0]["position"].Int() != 0 {
			position -= 1
		}
		fields[m["name"].String()] = &gdb.TableField{
			Index:   position,
			Name:    m["name"].String(),
			Default: m["default_expression"].Val(),
			Comment: m["comment"].String(),
			Type:    fieldType,
			Null:    isNull,
		}
	}
	return fields, nil
}
