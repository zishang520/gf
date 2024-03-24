// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package otelmetric

import (
	"context"

	"go.opentelemetry.io/otel/metric"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gmetric"
)

// localHistogramPerformer is an implementer for interface HistogramPerformer.
type localHistogramPerformer struct {
	metric.Float64Histogram
	gmetric.MeterOption
	gmetric.MetricOption
	attributesOption metric.MeasurementOption
}

// newHistogramPerformer creates and returns a HistogramPerformer that truly takes action to implement Histogram.
func (l *localMeterPerformer) newHistogramPerformer(
	meter metric.Meter,
	metricName string,
	metricOption gmetric.MetricOption,
) (gmetric.HistogramPerformer, error) {
	histogram, err := meter.Float64Histogram(
		metricName,
		metric.WithDescription(metricOption.Help),
		metric.WithUnit(metricOption.Unit),
		metric.WithExplicitBucketBoundaries(metricOption.Buckets...),
	)
	if err != nil {
		return nil, gerror.WrapCodef(
			gcode.CodeInternalError,
			err,
			`create Float64Histogram "%s" failed with option: %s`,
			metricName, gjson.MustEncodeString(metricOption),
		)
	}
	return &localHistogramPerformer{
		Float64Histogram: histogram,
		MeterOption:      l.MeterOption,
		MetricOption:     metricOption,
		attributesOption: metric.WithAttributes(attributesToKeyValues(metricOption.Attributes)...),
	}, nil
}

// Record adds a single value to the histogram. The value is usually positive or zero.
func (l *localHistogramPerformer) Record(increment float64, option ...gmetric.Option) {
	l.Float64Histogram.Record(
		context.Background(),
		increment,
		l.generateRecordOptions(option...)...,
	)
}

func (l *localHistogramPerformer) generateRecordOptions(option ...gmetric.Option) []metric.RecordOption {
	var (
		dynamicOption          = getDynamicOptionByMetricOption(option...)
		recordOptions          = []metric.RecordOption{l.attributesOption}
		globalAttributesOption = getGlobalAttributesOption(gmetric.GetGlobalAttributesOption{
			Instrument:        l.MeterOption.Instrument,
			InstrumentVersion: l.MeterOption.InstrumentVersion,
		})
	)
	if globalAttributesOption != nil {
		recordOptions = append(recordOptions, globalAttributesOption)
	}
	if dynamicOption != nil {
		recordOptions = append(recordOptions, dynamicOption)
	}
	return recordOptions
}