package statuspage

type Metric struct {
<<<<<<< HEAD
	Name               *string  `json:"name"`
	MetricIdentifier   *string  `json:"metric_identifier,omitempty"`
	Transform          *string  `json:"transform,omitempty"`
	Suffix             *string  `json:"suffix,omitempty"`
	YAxisMin           *float64 `json:"y_axis_min,omitempty"`
	YAxisMax           *float64 `json:"y_axis_max,omitempty"`
	YAxisHidden        *bool    `json:"y_axis_hidden,omitempty"`
	Display            *bool    `json:"display,omitempty"`
	DecimalPlaces      *int     `json:"decimal_places,omitempty"`
	TooltipDescription *string  `json:"tooltip_description,omitempty"`
=======
	Name               string  `json:"name"`
	MetricIdentifier   string  `json:"metric_identifier,omitempty"`
	Transform          string  `json:"transform,omitempty"`
	Suffix             string  `json:"suffix,omitempty"`
	YAxisMin           float64 `json:"y_axis_min,omitempty"`
	YAxisMax           float64 `json:"y_axis_max,omitempty"`
	YAxisHidden        bool    `json:"y_axis_hidden,omitempty"`
	Display            bool    `json:"display,omitempty"`
	DecimalPlaces      int     `json:"decimal_places,omitempty"`
	TooltipDescription string  `json:"tooltip_description,omitempty"`
>>>>>>> dda28c7436493399fdfd1b1aed0ce0f1345df307
}

type MetricFull struct {
	Metric
<<<<<<< HEAD
	ID                *string `json:"id"`
	MetricsProviderID *string `json:"metrics_provider_id"`
	MetricsDisplayID  *string `json:"metrics_display_id"`
	Backfilled        *bool   `json:"backfilled"`
	MostRecentDataAt  *string `json:"most_recent_data_at"`
	CreatedAt         *string `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"`
=======
	ID                string `json:"id"`
	MetricsProviderID string `json:"metrics_provider_id"`
	MetricsDisplayID  string `json:"metrics_display_id"`
	Backfilled        bool   `json:"backfilled"`
	MostRecentDataAt  string `json:"most_recent_data_at"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
>>>>>>> dda28c7436493399fdfd1b1aed0ce0f1345df307
}

func CreateMetric(client *Client, pageID, metricsProviderID string, metric *Metric) (*MetricFull, error) {
	var m MetricFull
	err := createResourceCustomURL(
		client,
		"/pages/"+pageID+"/metrics_providers/"+metricsProviderID+"/metrics",
		struct {
			Metric *Metric `json:"metric"`
		}{metric},
		&m,
	)

	return &m, err
}

func GetMetric(client *Client, pageID, metricID string) (*MetricFull, error) {
	var m MetricFull
	err := readResource(client, pageID, metricID, "metric", &m)

	return &m, err
}

func UpdateMetric(client *Client, pageID, metricID string, metric *Metric) (*MetricFull, error) {
	var m MetricFull

	err := updateResource(
		client,
		pageID,
		"metric",
		metricID,
		struct {
			Metric *Metric `json:"metric"`
		}{metric},
		&m,
	)

	return &m, err
}

func DeleteMetric(client *Client, pageID, metricID string) (err error) {
	return deleteResource(client, pageID, "metric", metricID)
}
