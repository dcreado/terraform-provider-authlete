package provider

import (
	"github.com/authlete/authlete-go/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func createBackchannelDeliverySchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
			ValidateFunc: validation.StringInSlice([]string{
				string(types.DeliveryMode_POLL),
				string(types.DeliveryMode_PING),
				string(types.DeliveryMode_PUSH),
			}, false),
		},
	}
}

func createDeliveryModeSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Required: false,
		Computed: true,
		ValidateFunc: validation.StringInSlice([]string{
			string(types.DeliveryMode_POLL),
			string(types.DeliveryMode_PING),
			string(types.DeliveryMode_PUSH),
		}, false),
	}
}

func mapDeliveryModeToDto(v interface{}) types.DeliveryMode {
	return types.DeliveryMode(v.(string))
}

func mapBackchannelDelivery(vals []interface{}) []types.DeliveryMode {

	values := make([]types.DeliveryMode, len(vals))

	for i, v := range vals {
		values[i] = mapDeliveryModeToDto(v)
	}

	return values
}

func mapBackchannelDeliveryFromDTO(vals *[]types.DeliveryMode) []interface{} {

	if vals != nil {
		entries := make([]interface{}, len(*vals), len(*vals))
		for i, v := range *vals {
			entries[i] = v
		}
		return entries
	}
	return make([]interface{}, 0)
}
