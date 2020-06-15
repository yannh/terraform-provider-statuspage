package statuspage

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	sp "github.com/yannh/statuspage-go-sdk"
)

func resourceMetricsProviderCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	email := d.Get("email").(string)
	password := d.Get("password").(string)
	aPIKey := d.Get("api_key").(string)
	aPIToken := d.Get("api_token").(string)
	applicationKey := d.Get("application_key").(string)
	baseUri := d.Get("metric_base_uri").(string)
	t := d.Get("type").(string)

	mp, err := sp.CreateMetricsProvider(
		client,
		d.Get("page_id").(string),
		&sp.MetricsProvider{
			Email:          &email,
			Password:       &password,
			APIKey:         &aPIKey,
			APIToken:       &aPIToken,
			ApplicationKey: &applicationKey,
			MetricBaseUri:  &baseUri,
			Type:           &t,
		},
	)
	if err != nil {
		log.Printf("[WARN] Statuspage Failed creating metrics provider: %s\n", err)
		return err
	}

	log.Printf("[INFO] Statuspage Created metrics provider: %s\n", *mp.ID)
	d.SetId(*mp.ID)

	return resourceMetricsProviderRead(d, m)
}

func resourceMetricsProviderRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	mp, err := sp.GetMetricsProvider(client, d.Get("page_id").(string), d.Id())
	if err != nil {
		log.Printf("[ERROR] Statuspage could not find metrics provider with ID: %s\n", d.Id())
		return err
	}

	if mp == nil {
		log.Printf("[INFO] Statuspage could not find metrics provider with ID: %s\n", d.Id())
		d.SetId("")
		return nil
	}

	log.Printf("[INFO] Statuspage read metrics provider: %s\n", *mp.ID)

	d.Set("email", mp.Email)
	d.Set("type", mp.Type)

	return nil
}

func resourceMetricsProviderUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	metricsProviderID := d.Id()

	email := d.Get("email").(string)
	password := d.Get("password").(string)
	aPIKey := d.Get("api_key").(string)
	aPIToken := d.Get("api_token").(string)
	applicationKey := d.Get("application_key").(string)
	baseUri := d.Get("metric_base_uri").(string)
	t := d.Get("type").(string)

	_, err := sp.UpdateMetricsProvider(
		client,
		d.Get("page_id").(string),
		metricsProviderID,
		&sp.MetricsProvider{
			Email:          &email,
			Password:       &password,
			APIKey:         &aPIKey,
			APIToken:       &aPIToken,
			ApplicationKey: &applicationKey,
			MetricBaseUri:  &baseUri,
			Type:           &t,
		},
	)
	if err != nil {
		log.Printf("[WARN] Statuspage Failed creating metrics provider: %s\n", err)
		return err
	}

	d.SetId(metricsProviderID)

	return resourceMetricsProviderRead(d, m)
}

func resourceMetricsProviderDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	return sp.DeleteMetricsProvider(client, d.Get("page_id").(string), d.Id())
}

func resourceMetricsProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourceMetricsProviderCreate,
		Read:   resourceMetricsProviderRead,
		Update: resourceMetricsProviderUpdate,
		Delete: resourceMetricsProviderDelete,

		Schema: map[string]*schema.Schema{
			"page_id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The ID of the page this metric provider belongs to",
				Required:    true,
			},
			"email": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Required by the Librato and Pingdom type metrics providers",
				Optional:    true,
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Required by the Pingdom-type metrics provider",
				Optional:    true,
			},
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Required by the Datadog and NewRelic type metrics providers",
				Optional:    true,
			},
			"api_token": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Required by the Librato and Datadog type metrics providers",
				Optional:    true,
			},
			"application_key": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Required by the Pingdom-type metrics provider",
				Optional:    true,
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "One of 'Pingdom', 'NewRelic', 'Librato', 'Datadog', or 'Self'",
				Required:    true,
				ValidateFunc: validation.StringInSlice(
					[]string{"Pingdom", "NewRelic", "Librato", "Datadog", "Self"},
					false,
				),
			},
			"metric_base_uri": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Required by the NewRelic-type metrics provider",
				Optional:    true,
			},
		},
	}
}
