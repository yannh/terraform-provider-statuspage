package statuspage

import (
	"fmt"
)

type MetricsProvider struct {
	Email          *string `json:"email,omitempty"`
	Password       *string `json:"password,omitempty"`
	APIKey         *string `json:"api_key,omitempty"`
	APIToken       *string `json:"api_token,omitempty"`
	ApplicationKey *string `json:"application_key,omitempty"`
	MetricBaseUri  *string `json:"metric_base_uri,omitempty"`
	Type           *string `json:"type,omitempty"`
}

type MetricsProviderFull struct {
	MetricsProvider
	ID                 *string `json:"id"`
	Name               *string `json:"name"`
	MetricIdentifider  *string `json:"metric_identifier"`
	Suffix             *string `json:"suffix"`
	Display            *string `json:"display"`
	ToolTipDescription *string `json:"tooltip_description"`
	YAxisMin           *string `json:"y_axis_min"`
	YAxisMax           *string `json:"y_axis_max"`
	DecimalPlaces      *string `json:"decimal_places"`
	Disabled           *bool   `json:"disabled"`
	AccountID          *string `json:"account_id"`
	LastRevalidatedAt  *string `json:"last_revalidated_at"`
	CreatedAt          *string `json:"created_at"`
	UpdatedAt          *string `json:"updated_at"`
}

func (mp *MetricsProvider) validate() error {
	switch *mp.Type {
	case "Pingdom":
		if *mp.Email == "" {
			return fmt.Errorf("parameter email is required for Pingdom Metrics Provider")
		}
		if *mp.Password == "" {
			return fmt.Errorf("parameter password is required for Pingdom Metrics Provider")
		}
		if *mp.ApplicationKey == "" {
			return fmt.Errorf("parameter application_key is required for Pingdom Metrics Provider")
		}
	case "NewRelic":
		if *mp.APIKey == "" {
			return fmt.Errorf("parameter api_key is required for NewRelic Metrics Provider")
		}
		if *mp.MetricBaseUri == "" {
			return fmt.Errorf("parameter metric_base_uri is required for NewRelic Metrics Provider")
		}
	case "Librato":
		if *mp.Email == "" {
			return fmt.Errorf("parameter email is required for Librato Metrics Provider")
		}
		if *mp.APIToken == "" {
			return fmt.Errorf("parameter api_token is required for Librato Metrics Provider")
		}
	case "Datadog":
		if *mp.APIKey == "" {
			return fmt.Errorf("parameter api_key is required for Datadog Metrics Provider")
		}
		if *mp.ApplicationKey == "" {
			return fmt.Errorf("parameter application_key is required for Datadog Metrics Provider")
		}
	}

	return nil
}

func CreateMetricsProvider(client *Client, pageID string, metricsProvider *MetricsProvider) (*MetricsProviderFull, error) {
	if err := metricsProvider.validate(); err != nil {
		return nil, err
	}

	var mp MetricsProviderFull
	err := createResource(
		client,
		pageID,
		"metrics_provider",
		struct {
			MetricsProvider *MetricsProvider `json:"metrics_provider"`
		}{metricsProvider},
		&mp,
	)

	return &mp, err
}

func GetMetricsProvider(client *Client, pageID string, metricsProviderID string) (*MetricsProviderFull, error) {
	var mp MetricsProviderFull
	err := readResource(client, pageID, metricsProviderID, "metrics_provider", &mp)

	return &mp, err
}

func UpdateMetricsProvider(client *Client, pageID, metricsProviderID string, metricsProvider *MetricsProvider) (*MetricsProviderFull, error) {
	if err := metricsProvider.validate(); err != nil {
		return nil, err
	}
	var mp MetricsProviderFull
	err := updateResource(
		client,
		pageID,
		"metrics_provider",
		metricsProviderID,
		struct {
			MetricsProvider *MetricsProvider `json:"metrics_provider"`
		}{metricsProvider},
		&mp,
	)
	return &mp, err
}

func DeleteMetricsProvider(client *Client, pageID, metricsProviderID string) (err error) {
	return deleteResource(client, pageID, "metrics_provider", metricsProviderID)
}
