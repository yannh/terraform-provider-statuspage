package statuspage

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	sp "github.com/yannh/statuspage-go-sdk"
)

func resourceMetricsProviderCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	mp, err := sp.CreateMetricsProvider(
		client,
		d.Get("page_id").(string),
		&sp.MetricsProvider{
			Email:          d.Get("email").(string),
			Password:       d.Get("password").(string),
			APIKey:         d.Get("api_key").(string),
			APIToken:       d.Get("api_token").(string),
			ApplicationKey: d.Get("application_key").(string),
			Type:           d.Get("type").(string),
		},
	)
	if err != nil {
		log.Printf("[WARN] Statuspage Failed creating metrics provider: %s\n", err)
		return err
	}

	log.Printf("[INFO] Statuspage Created metrics provider: %s\n", mp.ID)
	d.SetId(mp.ID)

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

	log.Printf("[INFO] Statuspage read metrics provider: %s\n", mp.ID)

	d.Set("email", mp.Email)
	d.Set("password", mp.Password)
	d.Set("api_key", mp.APIKey)
	d.Set("api_token", mp.APIToken)
	d.Set("application_key", mp.ApplicationKey)
	d.Set("type", mp.Type)

	return nil
}

func resourceMetricsProviderUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	metricsProviderID := d.Id()

	_, err := sp.UpdateMetricsProvider(
		client,
		d.Get("page_id").(string),
		metricsProviderID,
		&sp.MetricsProvider{
			Email:          d.Get("email").(string),
			Password:       d.Get("password").(string),
			APIKey:         d.Get("api_key").(string),
			APIToken:       d.Get("api_token").(string),
			ApplicationKey: d.Get("application_key").(string),
			Type:           d.Get("type").(string),
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"api_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"api_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice(
					[]string{"Pingdom", "NewRelic", "Librato", "Datadog", "Self"},
					false,
				),
			},
		},
	}
}
