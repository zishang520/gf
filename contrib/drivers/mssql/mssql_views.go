package mssql

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

const (
	viewsSqlTmp = `SELECT NAME FROM SYSOBJECTS WHERE XTYPE='V' AND STATUS >= 0 ORDER BY NAME`
)

// Views retrieves and returns the views of the current schema.
// It's mainly used in cli tool chain for automatically generating the models.
func (d *Driver) Views(ctx context.Context, schema ...string) (views []string, err error) {
	var result gdb.Result
	link, err := d.SlaveLink(schema...)
	if err != nil {
		return nil, err
	}

	result, err = d.DoSelect(ctx, link, viewsSqlTmp)
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
