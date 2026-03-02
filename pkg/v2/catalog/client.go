package catalog

import (
	"fmt"

	"github.com/connellrobert/printify-go/pkg/common"
)

var (
	ENDPOINT                                                            = "/v2/catalog"
	GET_SHIPPING_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT               = fmt.Sprintf("%s/blueprints/%%d/print_providers/%%d/shipping.json", ENDPOINT)
	GET_SHIPPING_STANDARD_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT = fmt.Sprintf("%s/blueprints/%%d/print_providers/%%d/shipping/standard.json", ENDPOINT)
	GET_SHIPPING_PRIORITY_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT = fmt.Sprintf("%s/blueprints/%%d/print_providers/%%d/shipping/priority.json", ENDPOINT)
	GET_SHIPPING_EXPRESS_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT  = fmt.Sprintf("%s/blueprints/%%d/print_providers/%%d/shipping/express.json", ENDPOINT)
	GET_SHIPPING_ECONOMY_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT  = fmt.Sprintf("%s/blueprints/%%d/print_providers/%%d/shipping/economy.json", ENDPOINT)
)

var (
	GetShippingForVariantsOfBlueprintById             = common.GetResourceWithTwoId[ShippingInfo, int, int](GET_SHIPPING_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT)
	GetShippingStandardInfoForVariantsOfBlueprintById = common.GetResourceWithTwoId[ShippingInfo, int, int](GET_SHIPPING_STANDARD_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT)
	GetShippingPriorityInfoForVariantsOfBlueprintById = common.GetResourceWithTwoId[ShippingInfo, int, int](GET_SHIPPING_PRIORITY_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT)
	GetShippingExpressInfoForVariantsOfBlueprintById  = common.GetResourceWithTwoId[ShippingInfo, int, int](GET_SHIPPING_EXPRESS_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT)
	GetShippingEconomyInfoForVariantsOfBlueprintById  = common.GetResourceWithTwoId[ShippingInfo, int, int](GET_SHIPPING_ECONOMY_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT)
)
