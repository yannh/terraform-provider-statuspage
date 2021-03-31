package statuspage

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	sp "github.com/yannh/statuspage-go-sdk"
)

func resourceMetricCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	name := d.Get("name").(string)
	metricIdentifier := d.Get("metric_identifier").(string)
	transform := d.Get("transform").(string)
	suffix := d.Get("suffix").(string)
	yAxisMin := d.Get("y_axis_min").(float64)
	yAxisMax := d.Get("y_axis_max").(float64)
	yAxisHidden := d.Get("y_axis_hidden").(bool)
	display := d.Get("display").(bool)
	decimalPlaces := d.Get("decimal_places").(int)
	tooltipDescription := d.Get("tooltip_description").(string)

	metric, err := sp.CreateMetric(
		client,
		d.Get("page_id").(string),
		d.Get("metrics_provider_id").(string),
		&sp.Metric{
			Name:               &name,
			MetricIdentifier:   &metricIdentifier,
			Transform:          &transform,
			Suffix:             &suffix,
			YAxisMin:           &yAxisMin,
			YAxisMax:           &yAxisMax,
			YAxisHidden:        &yAxisHidden,
			Display:            &display,
			DecimalPlaces:      &decimalPlaces,
			TooltipDescription: &tooltipDescription,
		},
	)
	if err != nil {
		log.Printf("[WARN] Statuspage Failed creating metric: %s\n", err)
		return err
	}

	log.Printf("[INFO] Statuspage Created metric: %s\n", *metric.ID)
	d.SetId(*metric.ID)

	return resourceMetricRead(d, m)
}

func resourceMetricRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	metric, err := sp.GetMetric(client, d.Get("page_id").(string), d.Id())
	if err != nil {
		log.Printf("[ERROR] Statuspage could not find metric with ID: %s\n", d.Id())
		return err
	}

	if metric == nil {
		log.Printf("[INFO] Statuspage could not find metric with ID: %s\n", d.Id())
		d.SetId("")
		return nil
	}

	log.Printf("[INFO] Statuspage read metric: %s\n", *metric.ID)

	d.Set("name", *metric.Name)
	d.Set("metric_identifier", *metric.MetricIdentifier)
	if metric.Transform != nil {
		// statuspage.io api does not return transform for GET metric operations
		// See https://developer.statuspage.io/#operation/getPagesPageIdMetricsMetricId
		d.Set("transform", *metric.Transform)
	}
	d.Set("suffix", *metric.Suffix)
	d.Set("y_axis_min", *metric.YAxisMin)
	d.Set("y_axis_max", *metric.YAxisMax)
	d.Set("y_axis_hidden", *metric.YAxisHidden)
	d.Set("display", *metric.Display)
	d.Set("decimal_places", *metric.DecimalPlaces)

	if metric.TooltipDescription != nil {
		d.Set("tooltip_description", *metric.TooltipDescription)
	}

	return nil
}

func resourceMetricUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	metricID := d.Id()

	name := d.Get("name").(string)
	metricIdentifier := d.Get("metric_identifier").(string)
	transform := d.Get("transform").(string)
	suffix := d.Get("suffix").(string)
	yAxisMin := d.Get("y_axis_min").(float64)
	yAxisMax := d.Get("y_axis_max").(float64)
	yAxisHidden := d.Get("y_axis_hidden").(bool)
	display := d.Get("display").(bool)
	decimalPlaces := d.Get("decimal_places").(int)
	tooltipDescription := d.Get("tooltip_description").(string)

	_, err := sp.UpdateMetric(
		client,
		d.Get("page_id").(string),
		metricID,
		&sp.Metric{
			Name:               &name,
			MetricIdentifier:   &metricIdentifier,
			Transform:          &transform,
			Suffix:             &suffix,
			YAxisMin:           &yAxisMin,
			YAxisMax:           &yAxisMax,
			YAxisHidden:        &yAxisHidden,
			Display:            &display,
			DecimalPlaces:      &decimalPlaces,
			TooltipDescription: &tooltipDescription,
		},
	)
	if err != nil {
		log.Printf("[WARN] Statuspage Failed creating metric: %s\n", err)
		return err
	}

	d.SetId(metricID)

	return resourceMetricRead(d, m)
}

func resourceMetricDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)

	return sp.DeleteMetric(client, d.Get("page_id").(string), d.Id())
}

func resourceMetricImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	if len(strings.Split(d.Id(), "/")) != 2 {
		return []*schema.ResourceData{}, fmt.Errorf("[ERROR] Invalid resource format: %s. Please use 'page-id/metric-id'", d.Id())
	}
	pageID := strings.Split(d.Id(), "/")[0]
	metricID := strings.Split(d.Id(), "/")[1]
	log.Printf("[INFO] Importing Metric %s from Page %s", metricID, pageID)

	d.Set("page_id", pageID)
	d.SetId(metricID)
	err := resourceMetricRead(d, m)

	return []*schema.ResourceData{d}, err
}

func resourceMetric() *schema.Resource {
	return &schema.Resource{
		Create: resourceMetricCreate,
		Read:   resourceMetricRead,
		Update: resourceMetricUpdate,
		Delete: resourceMetricDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMetricImport,
		},

		Schema: map[string]*schema.Schema{
			"page_id": {
				Type:        schema.TypeString,
				Description: "The ID of the page this metric belongs to",
				Required:    true,
				ForceNew:    true,
			},
			"metrics_provider_id": {
				Type:        schema.TypeString,
				Description: "ID of the metric provider",
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Display name for the metric",
				Optional:    true,
			},
			"metric_identifier": {
				Type:        schema.TypeString,
				Description: "The identifier used to look up the metric data from the provider",
				Optional:    true,
			},
			"transform": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The transform to apply to metric before pulling into Statuspage. One of: 'average', 'count', 'max', 'min', 'sum', 'response_time' or 'uptime'",
				ValidateFunc: validation.StringInSlice(
					[]string{"average", "count", "max", "min", "sum", "response_time", "uptime"},
					false,
				),
			},
			"suffix": {
				Type:        schema.TypeString,
				Description: "Suffix to describe the units on the graph",
				Optional:    true,
			},
			"y_axis_min": {
				Type:        schema.TypeFloat,
				Description: "The lower bound of the y axis",
				Optional:    true,
			},
			"y_axis_max": {
				Type:        schema.TypeFloat,
				Description: "The upper bound of the y axis",
				Optional:    true,
			},
			"y_axis_hidden": {
				Type:        schema.TypeBool,
				Description: "Should the values on the y axis be hidden on render",
				Optional:    true,
			},
			"display": {
				Type:        schema.TypeBool,
				Description: "Should the metric be displayed",
				Optional:    true,
			},
			"decimal_places": {
				Type:        schema.TypeInt,
				Description: "How many decimal places to render on the graph",
				Optional:    true,
			},
			"tooltip_description": {
				Type:        schema.TypeString,
				Description: "Tooltip for the metric",
				Optional:    true,
			},
		},
	}
}
