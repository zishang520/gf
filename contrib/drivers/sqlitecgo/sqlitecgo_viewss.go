package sqlitecgo

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

	// Querying the SQLITE_MASTER table for views (TYPE='view')
	result, err = d.DoSelect(
		ctx,
		link,
		`SELECT NAME FROM SQLITE_MASTER WHERE TYPE='view' ORDER BY NAME`,
	)
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
