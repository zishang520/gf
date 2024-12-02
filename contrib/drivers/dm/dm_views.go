package dm

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

const (
	viewsSqlTmp = `SELECT * FROM ALL_VIEWS`
)

// Views retrieves and returns the views of the current schema.
// It's mainly used in cli tool chain for automatically generating the models.
func (d *Driver) Views(ctx context.Context, schema ...string) (views []string, err error) {
	var result gdb.Result
	// When schema is empty, return the default link
	link, err := d.SlaveLink(schema...)
	if err != nil {
		return nil, err
	}
	// The link has been distinguished and no longer needs to judge the owner
	result, err = d.DoSelect(ctx, link, viewsSqlTmp)
	if err != nil {
		return
	}
	for _, m := range result {
		if v, ok := m["VIEW_NAME"]; ok {
			views = append(views, v.String())
		}
	}
	return
}
