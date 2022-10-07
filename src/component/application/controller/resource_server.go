package applicationProvider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Update: update,
		Delete: delete,
		Exists: exists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: applicationSchema,
	}
}

func create(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func read(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func update(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func delete(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func exists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return true, nil
}
