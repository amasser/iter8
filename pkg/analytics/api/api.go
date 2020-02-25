/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package api

import (
	iter8v1alpha1 "github.com/iter8-tools/iter8-controller/pkg/apis/iter8/v1alpha1"
)

const (
	AnalyticsAPIPath string = "/api/v1/analytics/canary/"
)

// Request defines payload to analytics service
type Request struct {
	// Specifies the name of the experiment
	Name string `json:"name"`

	// Specifies a time interval and key-value pairs for retrieving and processing data pertaining to the baseline version
	Baseline Window `json:"baseline"`

	// Specifies a time interval and key-value pairs for retrieving and processing data pertaining to the candidate version
	Candidate Window `json:"candidate"`

	// State returned by the server on the previous call
	LastState interface{} `json:"_last_state"`

	// Parameters controlling the behavior of the analytics
	TrafficControl TrafficControl `json:"traffic_control"`
}

// Window specifies the time range and tags
type Window struct {
	// ISO8601 timestamp for the beginning of the time range of interest
	StartTime string `json:"start_time"`

	// ISO8601 timestamp for the end of the time range of interest; if omitted, current time is assumed
	EndTime string `json:"end_time"`

	// Key-value pairs identifying the data pertaining to a version
	Tags map[string]string `json:"tags"`
}

const (
	TCKeyMaxTrafficPercent string = "max_traffic_percent"

	TCKeySuccessCriteria string = "success_criteria"
)

type TrafficControl map[string]interface{}

const (
	// SCKeyMetricName specifies the name of the metric to which the criterion applies
	// example: iter8_latency
	SCKeyMetricName string = "metric_name"

	// SCKeyType stores the criterion type. Options:
	// "delta": compares the candidate against the baseline version with respect to the metric;
	// "threshold": checks the candidate with respect to the metric
	SCKeyType string = "type"

	// SCKeyIsCounter indicates whether this is a counter
	SCKeyIsCounter string = "is_counter"

	// SCKeyAbsentValue is default value when data source does not return one
	SCKeyAbsentValue string = "absent_value"

	// SCKeyTemplate specifies the query template for the metric
	SCKeyTemplate string = "metric_query_template"

	// SCKeySampleSizeTemplate specifies the query template for the sample size
	SCKeySampleSizeTemplate string = "metric_sample_size_query_template"

	// SCKeyValue is the value for the success criteria to check
	SCKeyValue string = "value"

	// SCKeyStopOnFailure indicates whether or not the experiment must finish if this criterion is not satisfied;
	// defaults to false
	SCKeyStopOnFailure string = "stop_on_failure"
)

// SuccessCriterion stores the fields as key-value pairs
type SuccessCriterion map[string]interface{}

// Response defines the response from analytics server
type Response struct {
	// Traffic recommendation for the baseline version
	Baseline MetricsTraffic `json:"baseline"`

	// Traffic recommendation for the candidate version
	Candidate MetricsTraffic `json:"candidate"`

	// Summary of the candidate assessment based on success criteria
	Assessment Assessment `json:"assessment"`

	// State returned by the server, to be passed on the next call
	LastState interface{} `json:"_last_state"`
}

// Assessment includes the assessment content
type Assessment struct {
	// Summary of the candidate assessment based on success criteria
	Summary iter8v1alpha1.Summary `json:"summary"`

	// Summary of results for each success criterion
	SuccessCriteria []SuccessCriterionOutput `json:"success_criteria"`
}

// SuccessCriterionOutput contains assessment for a specific metric
type SuccessCriterionOutput struct {
	// Name of the metric to which the criterion applies
	// example: iter8_latency
	MetricName string `json:"metric_name"`

	// Overall summary based on all success criteria
	Conclusions []string `json:"conclusions"`

	// Indicates whether or not the success criterion for the corresponding metric has been met
	SuccessCriteriaMet bool `json:"success_criteria_met"`

	// Indicates whether or not the experiment must be aborted on the basis of the criterion for this metric
	AbortExperiment bool `json:"abort_experiment"`
}

// MetricsTraffic specifies traffic recommendations and stores states information
type MetricsTraffic struct {
	// Traffic recommendation for the version
	TrafficPercentage float64 `json:"traffic_percentage"`

	// Sate returned by the server, to be passed on the next call
	LastState interface{} `json:"_last_state"`
}