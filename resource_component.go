package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	gostatuspage "github.com/yannh/terraform-provider-statuspage/go-statuspage"
)

func resourceComponentCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*gostatuspage.Client)
	id, err := client.CreateComponent(d.Get("page_id").(string), &gostatuspage.Component{
		Name:               d.Get("name").(string),
		Description:        d.Get("description").(string),
		GroupID:            d.Get("group_id").(string),
		OnlyShowIfDegraded: d.Get("only_show_if_degraded").(bool),
		Status:             d.Get("status").(string),
		Showcase:           d.Get("showcase").(bool),
	})
	if err != nil {
		log.Printf("[WARN] Statuspage Failed creating component: %s\n", err)
		return err
	}

	log.Printf("[INFO] Statuspage Created: %s\n", id)
	d.SetId(id)

	return resourceComponentRead(d, m)
}

func resourceComponentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*gostatuspage.Client)
	component, err := client.GetComponent(d.Get("page_id").(string), d.Id())
	if component == nil {
		d.SetId("")
		return nil
	}

	if err != nil {
		return err
	}

	d.Set("name", component.Name)
	d.Set("description", component.Description)
	d.Set("group_id", component.GroupID)
	d.Set("only_show_if_degraded", component.OnlyShowIfDegraded)
	d.Set("status", component.Status)
	d.Set("showcase", component.Showcase)

	return nil
}

func resourceComponentUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*gostatuspage.Client)
	componentID := d.Id()

	err := client.UpdateComponent(
		d.Get("page_id").(string),
		componentID,
		&gostatuspage.Component{
			Name:               d.Get("name").(string),
			Description:        d.Get("description").(string),
			GroupID:            d.Get("group_id").(string),
			OnlyShowIfDegraded: d.Get("only_show_if_degraded").(bool),
			Status:             d.Get("status").(string),
			Showcase:           d.Get("showcase").(bool),
		},
	)
	if err != nil {
		log.Printf("[WARN] Statuspage Failed creating component: %s\n", err)
		return err
	}

	d.SetId(componentID)

	return resourceComponentRead(d, m)
}

func resourceComponentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*gostatuspage.Client)

	return client.DeleteComponent("8l7wwkjhvgg7", d.Id())
}

func resourceComponent() *schema.Resource {
	return &schema.Resource{
		Create: resourceComponentCreate,
		Read:   resourceComponentRead,
		Update: resourceComponentUpdate,
		Delete: resourceComponentDelete,

		Schema: map[string]*schema.Schema{
			"page_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					allowed_values := []string{"operational", "under_maintenance", "degraded_performance", "partial_outage", "major_outage", ""}
					for _, allowed_value := range allowed_values {
						if v == allowed_value {
							return nil, nil
						}
					}

					errs = append(errs, fmt.Errorf("%q must be one of %v, got: %s", key, allowed_values, v))
					return nil, errs
				},
			},
			"only_show_if_degraded": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"group_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"showcase": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}
