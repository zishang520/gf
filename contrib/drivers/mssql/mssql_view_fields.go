// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package mssql

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gutil"
)

var (
	viewFieldsSqlTmp = `
SELECT 
    a.name Field,
    CASE b.name 
        WHEN 'datetime' THEN 'datetime'
        WHEN 'numeric' THEN b.name + '(' + convert(varchar(20), a.xprec) + ',' + convert(varchar(20), a.xscale) + ')' 
        WHEN 'char' THEN b.name + '(' + convert(varchar(20), a.length)+ ')'
        WHEN 'varchar' THEN b.name + '(' + convert(varchar(20), a.length)+ ')'
        ELSE b.name + '(' + convert(varchar(20),a.length)+ ')' END AS Type,
    CASE WHEN a.isnullable=1 THEN 'YES' ELSE 'NO' end AS [Null],
    '' AS [Key],
    '' AS Extra,
    isnull(e.text,'') AS [Default],
    isnull(g.[value],'') AS [Comment]
FROM syscolumns a
LEFT  JOIN systypes b ON a.xtype=b.xtype AND a.xusertype=b.xusertype
INNER JOIN sysobjects d ON a.id=d.id AND d.xtype='V' AND d.name<>'dtproperties'
LEFT  JOIN syscomments e ON a.cdefault=e.id
LEFT  JOIN sys.extended_properties g ON a.id=g.major_id AND a.colid=g.minor_id
LEFT  JOIN sys.extended_properties f ON d.id=f.major_id AND f.minor_id =0
WHERE d.name='%s'
ORDER BY a.id,a.colorder
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
		result     gdb.Result
		link       gdb.Link
		usedSchema = gutil.GetOrDefaultStr(d.GetSchema(), schema...)
	)
	if link, err = d.SlaveLink(usedSchema); err != nil {
		return nil, err
	}

	// First, check if the view exists
	checkViewSQL := fmt.Sprintf("SELECT 1 FROM INFORMATION_SCHEMA.VIEWS WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s'", usedSchema, view)
	checkResult, err := d.DoSelect(ctx, link, checkViewSQL)
	if err != nil {
		return nil, err
	}
	if len(checkResult) == 0 {
		return nil, fmt.Errorf("view %s does not exist in schema %s", view, usedSchema)
	}

	structureSql := fmt.Sprintf(viewFieldsSqlTmp, view)
	result, err = d.DoSelect(ctx, link, structureSql)
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
