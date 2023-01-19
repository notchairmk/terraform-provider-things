package things

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dogResource() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for dog",

		CreateContext: resourceDogCreate,
		ReadContext:   resourceDogRead,
		UpdateContext: resourceDogUpdate,
		DeleteContext: resourceDogDelete,

		Schema: map[string]*schema.Schema{
			"legs": {
				// This description is used by the documentation generator and the language server.
				Description: "Number of dog legs",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"tags": {
				Description: "Dog tags",
				Type:        schema.TypeMap,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tags_computed": {
				Description: "Dog tags - computed",
				Type:        schema.TypeMap,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceDogCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the things configure method
	// client := meta.(*apiClient)

	id, err := uuid.NewV7()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id.String())

	// write logs using the tflog package
	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
	// for more information
	tflog.Trace(ctx, fmt.Sprintf("created a dog with id %s", id.String()))

	return nil
}

func resourceDogRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the things configure method
	// client := meta.(*apiClient)

	tflog.Debug(ctx, "setting empty dog tags")
	forceTags := map[string]interface{}{}
	d.Set("tags", forceTags)
	return nil
}

func resourceDogUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the things configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceDogDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the things configure method
	// client := meta.(*apiClient)

	return nil
}
