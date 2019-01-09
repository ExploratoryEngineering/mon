package mon

//
//Copyright 2018 Telenor Digital AS
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
import (
	"expvar"
)

type timeseriesCounter struct {
	name             string
	minuteTimeSeries *TimeSeries
	total            *expvar.Int
}

func (c *timeseriesCounter) init() {
	expvar.Publish(c.name+".minute", c.minuteTimeSeries)
	expvar.Publish(c.name+".total", c.total)
}

func (c *timeseriesCounter) Increment() {
	c.minuteTimeSeries.Increment()
	c.total.Add(1)
}
func newTimeseriesCounter(name string) *timeseriesCounter {
	ret := &timeseriesCounter{name, NewTimeSeries(Minutes), expvar.NewInt(name)}
	ret.init()
	return ret
}

type histogramCounter struct {
	name      string
	histogram *Histogram
	gauge     *AverageGauge
}

func (h *histogramCounter) init() {
	expvar.Publish(h.name+".histogram", h.histogram)
	expvar.Publish(h.name+".average", h.gauge)
}

// Add adds a new sample to the histogram counter
func (h *histogramCounter) Add(value float64) {
	h.histogram.Add(value)
	h.gauge.Add(value)
}

func newHistogramCounter(name string) *histogramCounter {
	ret := &histogramCounter{name, NewHistogram(), NewAverageGauge(1000)}
	ret.init()
	return ret
}
