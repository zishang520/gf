// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//

package gdb

import "github.com/gogf/gf/v2/os/gmetric"

type localMetricManager struct {
	SqlDuration            gmetric.UpDownCounter
	SqlTotal               gmetric.Counter
	SqlActive              gmetric.UpDownCounter
	ConnMaxOpenConnections gmetric.UpDownCounter
	ConnOpenConnections    gmetric.UpDownCounter
	ConnInUse              gmetric.UpDownCounter
	ConnIdle               gmetric.UpDownCounter
	ConnWaitCount          gmetric.Counter
	ConnMaxIdleClosed      gmetric.Counter
	ConnMaxIdleTimeClosed  gmetric.Counter
	ConnMaxLeftTimeClosed  gmetric.Counter
}

const (
	metricAttrKeyOrmSqlType     = "orm.sql.type"
	metricAttrKeyOrmConfigType  = "orm.config.type"
	metricAttrKeyOrmConfigGroup = "orm.config.group"
	metricAttrKeyOrmConfigName  = "orm.config.name"
	metricAttrKeyNetHostAddress = "net.host.address"
	metricAttrKeyNetHostPort    = "net.host.port"
)
