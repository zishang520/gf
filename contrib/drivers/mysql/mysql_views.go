package mysql

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

// Views retrieves and returns the views of the current schema.
// It's mainly used in cli tool chain for automatically generating the models.
func (d *Driver) Views(ctx context.Context, schema ...string) (views []string, err error) {
	var result gdb.Result
	link, err := d.SlaveLink(schema...)
	if err != nil {
		return nil, err
	}
	// Querying the INFORMATION_SCHEMA.VIEWS table to get the list of views
	result, err = d.DoSelect(ctx, link, `SELECT TABLE_NAME FROM INFORMATION_SCHEMA.VIEWS WHERE TABLE_SCHEMA = DATABASE()`)
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
