package statuspage

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	sp "github.com/yannh/statuspage-go-sdk"
)

func resourceIncidentTemplateCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)

	// convert []interface{} to []string
	var componentIDs []string
	for _, c := range d.Get("component_ids").([]interface{}) {
		componentIDs = append(componentIDs, c.(string))
	}

	incidentTemplate, err := sp.CreateIncidentTemplate(
		client,
		d.Get("page_id").(string),
		&sp.IncidentTemplate{
			Name:         d.Get("name").(string),
			GroupID:      d.Get("group_id").(string),
			UpdateStatus: d.Get("update_status").(string),
			Title:        d.Get("title").(string),
			Body:         d.Get("body").(string),
			ComponentIDs: componentIDs,
			ShouldTweet:  d.Get("should_tweet").(bool),
		},
	)
	if err != nil {
		log.Printf("[WARN] Statuspage Failed creating incident template: %s\n", err)
		return err
	}

	log.Printf("[INFO] Statuspage Created: %s\n", incidentTemplate.ID)
	d.SetId(incidentTemplate.ID)

	return resourceIncidentTemplateRead(d, m)
}

func resourceIncidentTemplateRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	incidentTemplate, err := sp.GetIncidentTemplate(client, d.Get("page_id").(string), d.Id())
	if err != nil {
		log.Printf("[ERROR] Statuspage could not find incident template with ID: %s\n", d.Id())
		return err
	}

	if incidentTemplate == nil {
		log.Printf("[INFO] Statuspage could not find incident template with ID: %s\n", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", incidentTemplate.Name)
	d.Set("group_id", incidentTemplate.GroupID)
	d.Set("update_status", incidentTemplate.UpdateStatus)
	d.Set("title", incidentTemplate.Title)
	d.Set("body", incidentTemplate.Body)
	d.Set("component_ids", incidentTemplate.ComponentIDs)
	d.Set("should_tweet", incidentTemplate.ShouldTweet)

	return nil
}

func resourceIncidentTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	incidentTemplateID := d.Id()

	// convert []interface{} to []string
	var componentIDs []string
	for _, c := range d.Get("components_ids").([]interface{}) {
		componentIDs = append(componentIDs, c.(string))
	}

	_, err := sp.UpdateIncidentTemplate(
		client,
		d.Get("page_id").(string),
		incidentTemplateID,
		&sp.IncidentTemplate{
			Name:         d.Get("name").(string),
			GroupID:      d.Get("group_id").(string),
			UpdateStatus: d.Get("update_status").(string),
			Title:        d.Get("title").(string),
			Body:         d.Get("body").(string),
			ComponentIDs: componentIDs,
			ShouldTweet:  d.Get("should_tweet").(bool),
		},
	)
	if err != nil {
		log.Printf("[WARN] Statuspage Failed creating incident template: %s\n", err)
		return err
	}

	d.SetId(incidentTemplateID)

	return resourceIncidentTemplateRead(d, m)
}

func resourceIncidentTemplateDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)

	return sp.DeleteIncidentTemplate(client, d.Get("page_id").(string), d.Id())
}

func resourceIncidentTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceIncidentTemplateCreate,
		Read:   resourceIncidentTemplateRead,
		Update: resourceIncidentTemplateUpdate,
		Delete: resourceIncidentTemplateDelete,

		Schema: map[string]*schema.Schema{
			"page_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"group_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "no_group",
			},
			"update_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "resolved",
			},
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"body": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"component_ids": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				DefaultFunc: func() (interface{}, error) {
					return []string{}, nil
				},
			},
			"should_tweet": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
