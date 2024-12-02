// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package dm

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gutil"
)

const (
	viewFieldsSqlTmp = `
SELECT * FROM ALL_TAB_COLUMNS 
WHERE Table_Name = '%s' 
  AND OWNER = '%s' 
  AND Table_Name IN (SELECT VIEW_NAME FROM ALL_VIEWS WHERE OWNER = '%s')
`
)

// ViewFields retrieves and returns the fields' information of specified view of current schema.
func (d *Driver) ViewFields(
	ctx context.Context, view string, schema ...string,
) (fields map[string]*gdb.TableField, err error) {
	var (
		result gdb.Result
		link   gdb.Link
		// When no schema is specified, the configuration item is returned by default
		usedSchema = gutil.GetOrDefaultStr(d.GetSchema(), schema...)
	)
	// When usedSchema is empty, return the default link
	if link, err = d.SlaveLink(usedSchema); err != nil {
		return nil, err
	}

	// Check if the view exists
	checkViewSQL := fmt.Sprintf(
		"SELECT 1 FROM ALL_VIEWS WHERE VIEW_NAME = '%s' AND OWNER = '%s'",
		strings.ToUpper(view),
		strings.ToUpper(usedSchema),
	)
	checkResult, err := d.DoSelect(ctx, link, checkViewSQL)
	if err != nil {
		return nil, err
	}
	if len(checkResult) == 0 {
		return nil, fmt.Errorf("view %s does not exist in schema %s", view, usedSchema)
	}

	// The link has been distinguished and no longer needs to judge the owner
	result, err = d.DoSelect(
		ctx, link,
		fmt.Sprintf(
			viewFieldsSqlTmp,
			strings.ToUpper(view),
			strings.ToUpper(usedSchema),
			strings.ToUpper(usedSchema),
		),
	)
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gdb.TableField)
	for i, m := range result {
		// m[NULLABLE] returns "N" "Y"
		// "N" means not null
		// "Y" means could be null
		var nullable bool
		if m["NULLABLE"].String() != "N" {
			nullable = true
		}
		fields[m["COLUMN_NAME"].String()] = &gdb.TableField{
			Index:   i,
			Name:    m["COLUMN_NAME"].String(),
			Type:    m["DATA_TYPE"].String(),
			Null:    nullable,
			Default: m["DATA_DEFAULT"].Val(),
			// Key and Extra are typically not applicable for views
			// Comment: m["Comment"].String(), // Add if comment information is available for views
		}
	}
	return fields, nil
}
