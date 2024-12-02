// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package sqlite

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gutil"
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
	checkViewSQL := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='view' AND name=%s", d.QuoteString(view))
	result, err = d.DoSelect(ctx, link, checkViewSQL)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("view %s does not exist", view)
	}

	// If the view exists, get its column information
	result, err = d.DoSelect(ctx, link, fmt.Sprintf(`PRAGMA TABLE_INFO(%s)`, d.QuoteWord(view)))
	if err != nil {
		return nil, err
	}

	fields = make(map[string]*gdb.TableField)
	for i, m := range result {
		fields[m["name"].String()] = &gdb.TableField{
			Index:   i,
			Name:    m["name"].String(),
			Type:    m["type"].String(),
			Key:     "", // Views in SQLite don't have primary keys
			Default: m["dflt_value"].Val(),
			Null:    !m["notnull"].Bool(),
		}
	}
	return fields, nil
}
