package statuspage

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	sp "github.com/yannh/statuspage-go-sdk"
)

func resourceMetricCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	metric, err := sp.CreateMetric(
		client,
		d.Get("page_id").(string),
		d.Get("metrics_provider_id").(string),
		&sp.Metric{
			Name:               d.Get("name").(string),
			MetricIdentifier:   d.Get("metric_identifier").(string),
			Transform:          d.Get("transform").(string),
			Suffix:             d.Get("suffix").(string),
			YAxisMin:           d.Get("y_axis_min").(int32),
			YAxisMax:           d.Get("y_axis_max").(int32),
			YAxisHidden:        d.Get("y_axis_hidden").(bool),
			Display:            d.Get("display").(bool),
			DecimalPlaces:      d.Get("decimal_places").(int32),
			TooltipDescription: d.Get("tooltip_description").(string),
		},
	)
	if err != nil {
		log.Printf("[WARN] Statuspage Failed creating metric: %s\n", err)
		return err
	}

	log.Printf("[INFO] Statuspage Created metric: %s\n", metric.ID)
	d.SetId(metric.ID)

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

	log.Printf("[INFO] Statuspage read metric: %s\n", metric.ID)

	d.Set("name", metric.Name)
	d.Set("metric_identifier", metric.MetricIdentifier)
	d.Set("transform", metric.Transform)
	d.Set("suffix", metric.Suffix)
	d.Set("y_axis_min", metric.YAxisMin)
	d.Set("y_axis_max", metric.YAxisMax)
	d.Set("y_axis_hidden", metric.YAxisHidden)
	d.Set("display", metric.Display)
	d.Set("decimal_places", metric.DecimalPlaces)
	d.Set("tooltip_description", metric.TooltipDescription)

	return nil
}

func resourceMetricUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	metricID := d.Id()

	_, err := sp.UpdateMetric(
		client,
		d.Get("page_id").(string),
		metricID,
		&sp.Metric{
			Name:               d.Get("name").(string),
			MetricIdentifier:   d.Get("metric_identifier").(string),
			Transform:          d.Get("transform").(string),
			Suffix:             d.Get("suffix").(string),
			YAxisMin:           d.Get("y_axis_min").(int32),
			YAxisMax:           d.Get("y_axis_max").(int32),
			YAxisHidden:        d.Get("y_axis_hidden").(bool),
			Display:            d.Get("display").(bool),
			DecimalPlaces:      d.Get("decimal_places").(int32),
			TooltipDescription: d.Get("tooltip_description").(string),
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

func resourceMetric() *schema.Resource {
	return &schema.Resource{
		Create: resourceMetricCreate,
		Read:   resourceMetricRead,
		Update: resourceMetricUpdate,
		Delete: resourceMetricDelete,

		Schema: map[string]*schema.Schema{
			"page_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"metrics_provider_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"transform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice(
					[]string{"average", "count", "max", "min", "sum"},
					false,
				),
			},
			"suffix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"y_axis_min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"y_axis_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"y_axis_hidden": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"decimal_places": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tooltip_description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
