package statuspage

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func getOptionalBool(d *schema.ResourceData, key string) *bool {
	if value, ok := d.GetOk(key); ok {
		b := value.(bool)
		return &b
	}
	return nil
}

func getOptionalFloat(d *schema.ResourceData, key string) *float64 {
	if value, ok := d.GetOk(key); ok {
		f := value.(float64)
		return &f
	}
	return nil
}

func getOptionalInt(d *schema.ResourceData, key string) *int {
	if value, exists := d.GetOk(key); exists {
		i := value.(int)
		return &i
	}
	return nil
}

func getOptionalString(d *schema.ResourceData, key string) *string {
	if value, exists := d.GetOk(key); exists {
		s := value.(string)
		return &s
	}
	return nil
}
