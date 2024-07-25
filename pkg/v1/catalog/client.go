package catalog

import (
	"fmt"

	"github.com/connellrobert/printify-go/pkg/common"
)

var (
	ENDPOINT                                           = "/v1/catalog"
	GET_BLUEPRINT_ENDPOINT                             = fmt.Sprintf("%s/blueprints/%%d.json", ENDPOINT)
	LIST_PRINT_PROVIDERS_BY_BLUEPRINT_ENDPOINT         = fmt.Sprintf("%s/blueprints/%%d/print_providers.json", ENDPOINT)
	LIST_VARIANTS_BY_BLUEPRINT_PRINT_PROVIDER_ENDPOINT = fmt.Sprintf("%s/blueprints/%%d/print_providers/%%d/variants.json", ENDPOINT)
	GET_SHIPPING_INFORMATION_ENDPOINT                  = fmt.Sprintf("%s/blueprints/%%d/print_providers/%%d/shipping.json", ENDPOINT)
	LIST_PRINT_PROVIDERS_ENDPOINT                      = fmt.Sprintf("%s/print_providers.json", ENDPOINT)
	GET_PRINT_PROVIDER_ENDPOINT                        = fmt.Sprintf("%s/print_providers/%%d.json", ENDPOINT)
	LIST_BLUEPRINTS_ENDPOINT                           = fmt.Sprintf("%s/blueprints.json", ENDPOINT)
)

var (
	ListBlueprints                       = common.ListResources[Blueprint](LIST_BLUEPRINTS_ENDPOINT)
	GetBlueprint                         = common.GetResourceById[Blueprint, int](GET_BLUEPRINT_ENDPOINT)
	ListPrintProvidersByBlueprint        = common.ListResourceWithId[PrintProvider, int](LIST_PRINT_PROVIDERS_BY_BLUEPRINT_ENDPOINT)
	ListVariantsByBlueprintPrintProvider = common.ListResourceWithTwoID[Variant, int, int](LIST_VARIANTS_BY_BLUEPRINT_PRINT_PROVIDER_ENDPOINT)
	GetShippingInformation               = common.GetResourceWithTwoId[Shipping, int, int](GET_SHIPPING_INFORMATION_ENDPOINT)
	ListPrintProviders                   = common.ListResources[PrintProvider](LIST_PRINT_PROVIDERS_ENDPOINT)
	GetPrintProvider                     = common.GetResourceById[PrintProvider, int](GET_PRINT_PROVIDER_ENDPOINT)
)
