package statuspage

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	sp "github.com/yannh/statuspage-go-sdk"
)

func resourceComponentGroupCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)

	tfComponents := d.Get("components").(*schema.Set).List()
	components := make([]string, len(tfComponents))
	for i, tfComponent := range tfComponents {
		components[i] = tfComponent.(string)
	}

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	componentGroup, err := sp.CreateComponentGroup(
		client,
		d.Get("page_id").(string),
		&sp.ComponentGroup{
			Name:        &name,
			Description: &description,
			Components:  components,
		},
	)

	if err != nil {
		log.Printf("[WARN] Statuspage Failed creating component group: %s\n", err)
		return err
	}

	log.Printf("[INFO] Statuspage Created component group: %s\n", *componentGroup.ID)
	d.SetId(*componentGroup.ID)

	return resourceComponentGroupRead(d, m)
}

func resourceComponentGroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	componentGroup, err := sp.GetComponentGroup(client, d.Get("page_id").(string), d.Id())
	if err != nil {
		log.Printf("[ERROR] Statuspage could not find component group with ID: %s\n", d.Id())
		return err
	}

	if componentGroup == nil {
		log.Printf("[INFO] Statuspage could not find component group with ID: %s\n", d.Id())
		d.SetId("")
		return nil
	}

	log.Printf("[INFO] Statuspage read component group: %s\n", *componentGroup.ID)

	d.Set("name", componentGroup.Name)
	d.Set("description", componentGroup.Description)
	d.Set("components", componentGroup.Components)

	return nil
}

func resourceComponentGroupUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)
	componentGroupID := d.Id()

	tfComponents := d.Get("components").(*schema.Set).List()
	components := make([]string, len(tfComponents))
	for i, tfComponent := range tfComponents {
		components[i] = tfComponent.(string)
	}

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	_, err := sp.UpdateComponentGroup(
		client,
		d.Get("page_id").(string),
		componentGroupID,
		&sp.ComponentGroup{
			Name:        &name,
			Description: &description,
			Components:  components,
		},
	)
	if err != nil {
		log.Printf("[WARN] Statuspage Failed creating component group: %s\n", err)
		return err
	}

	d.SetId(componentGroupID)

	return resourceComponentGroupRead(d, m)
}

func resourceComponentGroupDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*sp.Client)

	return sp.DeleteComponentGroup(client, d.Get("page_id").(string), d.Id())
}

func resourceComponentGroupImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	if len(strings.Split(d.Id(), "/")) != 2 {
		return []*schema.ResourceData{}, fmt.Errorf("[ERROR] Invalid resource format: %s. Please use 'page-id/component-group-id'", d.Id())
	}
	pageID := strings.Split(d.Id(), "/")[0]
	componentGroupID := strings.Split(d.Id(), "/")[1]
	log.Printf("[INFO] Importing Component Group %s from Page %s", componentGroupID, pageID)

	d.Set("page_id", pageID)
	d.SetId(componentGroupID)
	err := resourceComponentGroupRead(d, m)

	return []*schema.ResourceData{d}, err
}

func resourceComponentGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceComponentGroupCreate,
		Read:   resourceComponentGroupRead,
		Update: resourceComponentGroupUpdate,
		Delete: resourceComponentGroupDelete,
		Importer: &schema.ResourceImporter{
			State: resourceComponentGroupImport,
		},

		Schema: map[string]*schema.Schema{
			"page_id": {
				Type:        schema.TypeString,
				Description: "the ID of the page this component group belongs to",
				Required:    true,
				ForceNew:    true,
			},
			"components": {
				Type:        schema.TypeSet,
				Description: "An array with the IDs of the components in this group",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Display name for this component group",
				Required:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "More detailed description for this component group",
				Optional:    true,
			},
		},
	}
}
