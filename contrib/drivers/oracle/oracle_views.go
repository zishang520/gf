package oracle

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

const (
	viewsSqlTmp = `SELECT VIEW_NAME FROM USER_VIEWS ORDER BY VIEW_NAME`
)

// Views retrieves and returns the views of the current schema.
// It's mainly used in cli tool chain for automatically generating the models.
// Note that it ignores the parameter `schema` in Oracle database, as it is not necessary.
func (d *Driver) Views(ctx context.Context, schema ...string) (views []string, err error) {
	var result gdb.Result
	// DO NOT use `usedSchema` as parameter for function `SlaveLink` since Oracle uses USER views.
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
