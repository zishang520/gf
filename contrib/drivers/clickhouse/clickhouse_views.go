package clickhouse

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
)

const (
	viewsSqlTmp = "SELECT name FROM system.tables WHERE database = '%s' AND engine = 'View'"
)

// Views retrieves and returns the views of the current schema.
// It's mainly used in cli tool chain for automatically generating the models.
func (d *Driver) Views(ctx context.Context, schema ...string) (views []string, err error) {
	var result gdb.Result
	link, err := d.SlaveLink(schema...)
	if err != nil {
		return nil, err
	}
	result, err = d.DoSelect(ctx, link, fmt.Sprintf(viewsSqlTmp, d.GetConfig().Name))
	if err != nil {
		return
	}
	for _, m := range result {
		views = append(views, m["name"].String())
	}
	return
}
