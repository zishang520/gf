// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package oracle

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gutil"
)

var (
	viewFieldsSqlTmp = `
SELECT 
    COLUMN_NAME AS FIELD, 
    CASE   
    WHEN (DATA_TYPE='NUMBER' AND NVL(DATA_SCALE,0)=0) THEN 'INT'||'('||DATA_PRECISION||','||DATA_SCALE||')'
    WHEN (DATA_TYPE='NUMBER' AND NVL(DATA_SCALE,0)>0) THEN 'FLOAT'||'('||DATA_PRECISION||','||DATA_SCALE||')'
    WHEN DATA_TYPE='FLOAT' THEN DATA_TYPE||'('||DATA_PRECISION||','||DATA_SCALE||')' 
    ELSE DATA_TYPE||'('||DATA_LENGTH||')' END AS TYPE,NULLABLE
FROM USER_TAB_COLUMNS 
WHERE TABLE_NAME = '%s' 
  AND EXISTS (SELECT 1 FROM USER_VIEWS WHERE VIEW_NAME = '%s')
ORDER BY COLUMN_ID
`
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
		structureSql = fmt.Sprintf(viewFieldsSqlTmp, strings.ToUpper(view), strings.ToUpper(view))
	)
	if link, err = d.SlaveLink(usedSchema); err != nil {
		return nil, err
	}

	result, err = d.DoSelect(ctx, link, structureSql)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("view %s does not exist in schema %s", view, usedSchema)
	}

	fields = make(map[string]*gdb.TableField)
	for i, m := range result {
		isNull := false
		if m["NULLABLE"].String() == "Y" {
			isNull = true
		}

		fields[m["FIELD"].String()] = &gdb.TableField{
			Index: i,
			Name:  m["FIELD"].String(),
			Type:  m["TYPE"].String(),
			Null:  isNull,
		}
	}
	return fields, nil
}
