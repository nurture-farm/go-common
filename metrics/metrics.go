/*
 *  Copyright 2023 NURTURE AGTECH PVT LTD
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
	"context"
)

type Helper struct {
	SERVICE_NAME string
	DATABASE string
}

func (m *Helper) PushToSummarytMetrics() func(*prometheus.SummaryVec,string,*error,context.Context) {
	startTime := time.Now()
	return func(request *prometheus.SummaryVec,methodName string,err *error,ctx context.Context) {

		if *err != nil {
			request.WithLabelValues(m.SERVICE_NAME,methodName,"ko").Observe(float64(time.Now().Sub(startTime).Milliseconds()))
		} else {
			request.WithLabelValues(m.SERVICE_NAME,methodName,"ok").Observe(float64(time.Now().Sub(startTime).Milliseconds()))
		}
	}
}

func (m *Helper) PushToErrorCounterMetrics() func(*prometheus.CounterVec,error,context.Context) {
	return func(request *prometheus.CounterVec,err error,ctx context.Context) {
		request.WithLabelValues(m.SERVICE_NAME,m.DATABASE,err.Error()).Inc()
	}
}