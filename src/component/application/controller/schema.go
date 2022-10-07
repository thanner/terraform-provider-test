package applicationProvider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var (
	applicationSchema = map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
)
