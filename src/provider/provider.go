package provider

import (
	applicationProvider "terraform-provider-test/src/component/application/controller"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"application": applicationProvider.Resource(),
		},
	}
}
