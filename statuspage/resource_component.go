package statuspage

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	sp "github.com/yannh/statuspage-go-sdk"
)

func s(s string) *string {
	return &s
}

func resourceComponentCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	onlyShowIfDegraded := d.Get("only_show_if_degraded").(bool)
	status := d.Get("status").(string)
	showcase := d.Get("showcase").(bool)
	groupId := d.Get("group_id").(string)
	inputComponent := &sp.Component{
		Name:               &name,
		Description:        &description,
		OnlyShowIfDegraded: &onlyShowIfDegraded,
		Status:             &status,
		Showcase:           &showcase,
	}
	if groupId != "" {
		inputComponent.GroupID = &groupId
	}
	component, err := sp.CreateComponent(client, d.Get("page_id").(string), inputComponent)
	if err != nil {
		log.Printf("[WARN] Statuspage Failed creating component: %s\n", err)
		return err
	}

	log.Printf("[INFO] Statuspage Created: %s\n", *component.ID)
	d.SetId(*component.ID)

	return resourceComponentRead(d, m)
}

func resourceComponentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	component, err := sp.GetComponent(client, d.Get("page_id").(string), d.Id())
	if err != nil {
		log.Printf("[ERROR] Statuspage could not find component with ID: %s\n", d.Id())
		return err
	}

	if component == nil {
		log.Printf("[INFO] Statuspage could not find component with ID: %s\n", d.Id())
		d.SetId("")
		return nil
	}

	log.Printf("[INFO] Statuspage read: %s\n", *component.ID)

	d.Set("name", component.Name)
	d.Set("description", component.Description)
	d.Set("group_id", component.GroupID)
	d.Set("only_show_if_degraded", component.OnlyShowIfDegraded)
	d.Set("status", component.Status)
	d.Set("showcase", component.Showcase)
	d.Set("automation_email", component.AutomationEmail)

	return nil
}

func resourceComponentUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	componentID := d.Id()

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	onlyShowIfDegraded := d.Get("only_show_if_degraded").(bool)
	status := d.Get("status").(string)
	showcase := d.Get("showcase").(bool)

	_, err := sp.UpdateComponent(
		client,
		d.Get("page_id").(string),
		componentID,
		&sp.Component{
			Name:               &name,
			Description:        &description,
			OnlyShowIfDegraded: &onlyShowIfDegraded,
			Status:             &status,
			Showcase:           &showcase,
		},
	)
	if err != nil {
		log.Printf("[WARN] Statuspage Failed updating component: %s\n", err)
		return err
	}

	d.SetId(componentID)

	return resourceComponentRead(d, m)
}

func resourceComponentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	pageID := d.Get("page_id").(string)

	return sp.DeleteComponent(client, pageID, d.Id())
}

func resourceComponentImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	if len(strings.Split(d.Id(), "/")) != 2 {
		return []*schema.ResourceData{}, fmt.Errorf("[ERROR] Invalid resource format: %s. Please use 'page-id/component-id'", d.Id())
	}
	pageID := strings.Split(d.Id(), "/")[0]
	componentID := strings.Split(d.Id(), "/")[1]
	log.Printf("[INFO] Importing Component %s from Page %s", componentID, pageID)

	d.Set("page_id", pageID)
	d.SetId(componentID)
	err := resourceComponentRead(d, m)

	return []*schema.ResourceData{d}, err
}

func resourceComponent() *schema.Resource {
	return &schema.Resource{
		Create: resourceComponentCreate,
		Read:   resourceComponentRead,
		Update: resourceComponentUpdate,
		Delete: resourceComponentDelete,
		Importer: &schema.ResourceImporter{
			State: resourceComponentImport,
		},

		Schema: map[string]*schema.Schema{
			"page_id": {
				Type:        schema.TypeString,
				Description: "the ID of the page this component belongs to",
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Display Name for the component",
				Required:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "More detailed description for the component",
				Optional:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Status of component",
				Optional:    true,
				ValidateFunc: validation.StringInSlice(
					[]string{"operational", "under_maintenance", "degraded_performance", "partial_outage", "major_outage", ""},
					false,
				),
				Default: "operational",
			},
			"only_show_if_degraded": {
				Type:        schema.TypeBool,
				Description: "Should this component be shown component only if in degraded state",
				Optional:    true,
			},
			"showcase": {
				Type:        schema.TypeBool,
				Description: "Should this component be showcased",
				Optional:    true,
				Default:     true,
			},
			"automation_email": {
				Type:        schema.TypeString,
				Description: "Email address to send automation events to",
				Computed:    true,
			},
			"group_id": {
				Type:        schema.TypeString,
				Description: "Component Group Id",
				Optional:    true,
				Default:     "",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if new == "" {
					  return true
					}
					return false
				},
			},
		},
	}
}
