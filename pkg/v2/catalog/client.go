package catalog

import (
	"fmt"

	"github.com/connellrobert/printify-go/pkg/common"
)

// Client defines v2 catalog shipping operations and enables dependency injection.
type Client interface {
	GetShippingForVariantsOfBlueprintById(idOne int, idTwo int) (*ShippingInfo, error)
	GetShippingStandardInfoForVariantsOfBlueprintById(idOne int, idTwo int) (*ShippingInfo, error)
	GetShippingPriorityInfoForVariantsOfBlueprintById(idOne int, idTwo int) (*ShippingInfo, error)
	GetShippingExpressInfoForVariantsOfBlueprintById(idOne int, idTwo int) (*ShippingInfo, error)
	GetShippingEconomyInfoForVariantsOfBlueprintById(idOne int, idTwo int) (*ShippingInfo, error)
}

type client struct {
	c *common.Client
}

// NewClient creates a v2 catalog client implementation backed by common.Client.
func NewClient(c *common.Client) Client {
	return &client{c: c}
}

func (cl *client) GetShippingForVariantsOfBlueprintById(idOne int, idTwo int) (*ShippingInfo, error) {
	return GetShippingForVariantsOfBlueprintById(cl.c, idOne, idTwo)
}

func (cl *client) GetShippingStandardInfoForVariantsOfBlueprintById(idOne int, idTwo int) (*ShippingInfo, error) {
	return GetShippingStandardInfoForVariantsOfBlueprintById(cl.c, idOne, idTwo)
}

func (cl *client) GetShippingPriorityInfoForVariantsOfBlueprintById(idOne int, idTwo int) (*ShippingInfo, error) {
	return GetShippingPriorityInfoForVariantsOfBlueprintById(cl.c, idOne, idTwo)
}

func (cl *client) GetShippingExpressInfoForVariantsOfBlueprintById(idOne int, idTwo int) (*ShippingInfo, error) {
	return GetShippingExpressInfoForVariantsOfBlueprintById(cl.c, idOne, idTwo)
}

func (cl *client) GetShippingEconomyInfoForVariantsOfBlueprintById(idOne int, idTwo int) (*ShippingInfo, error) {
	return GetShippingEconomyInfoForVariantsOfBlueprintById(cl.c, idOne, idTwo)
}

var (
	ENDPOINT                                                            = "/v2/catalog"
	GET_SHIPPING_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT               = fmt.Sprintf("%s/blueprints/%%d/print_providers/%%d/shipping.json", ENDPOINT)
	GET_SHIPPING_STANDARD_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT = fmt.Sprintf("%s/blueprints/%%d/print_providers/%%d/shipping/standard.json", ENDPOINT)
	GET_SHIPPING_PRIORITY_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT = fmt.Sprintf("%s/blueprints/%%d/print_providers/%%d/shipping/priority.json", ENDPOINT)
	GET_SHIPPING_EXPRESS_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT  = fmt.Sprintf("%s/blueprints/%%d/print_providers/%%d/shipping/express.json", ENDPOINT)
	GET_SHIPPING_ECONOMY_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT  = fmt.Sprintf("%s/blueprints/%%d/print_providers/%%d/shipping/economy.json", ENDPOINT)
)

var (
	// GetShippingForVariantsOfBlueprintById calls
	// GET /v2/catalog/blueprints/{blueprintId}/print_providers/{printProviderId}/shipping.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int) (*ShippingInfo, error)
	// Parameter mapping:
	//	idOne -> {blueprintId}
	//	idTwo -> {printProviderId}
	//
	// blueprintId can be discovered with v1 catalog.ListBlueprints.
	// printProviderId can be discovered with v1 catalog.ListPrintProvidersByBlueprint(idOne).
	GetShippingForVariantsOfBlueprintById = common.GetResourceWithTwoId[ShippingInfo, int, int](GET_SHIPPING_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT)
	// GetShippingStandardInfoForVariantsOfBlueprintById calls
	// GET /v2/catalog/blueprints/{blueprintId}/print_providers/{printProviderId}/shipping/standard.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int) (*ShippingInfo, error)
	// Parameter mapping:
	//	idOne -> {blueprintId}
	//	idTwo -> {printProviderId}
	//
	// blueprintId can be discovered with v1 catalog.ListBlueprints.
	// printProviderId can be discovered with v1 catalog.ListPrintProvidersByBlueprint(idOne).
	GetShippingStandardInfoForVariantsOfBlueprintById = common.GetResourceWithTwoId[ShippingInfo, int, int](GET_SHIPPING_STANDARD_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT)
	// GetShippingPriorityInfoForVariantsOfBlueprintById calls
	// GET /v2/catalog/blueprints/{blueprintId}/print_providers/{printProviderId}/shipping/priority.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int) (*ShippingInfo, error)
	// Parameter mapping:
	//	idOne -> {blueprintId}
	//	idTwo -> {printProviderId}
	//
	// blueprintId can be discovered with v1 catalog.ListBlueprints.
	// printProviderId can be discovered with v1 catalog.ListPrintProvidersByBlueprint(idOne).
	GetShippingPriorityInfoForVariantsOfBlueprintById = common.GetResourceWithTwoId[ShippingInfo, int, int](GET_SHIPPING_PRIORITY_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT)
	// GetShippingExpressInfoForVariantsOfBlueprintById calls
	// GET /v2/catalog/blueprints/{blueprintId}/print_providers/{printProviderId}/shipping/express.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int) (*ShippingInfo, error)
	// Parameter mapping:
	//	idOne -> {blueprintId}
	//	idTwo -> {printProviderId}
	//
	// blueprintId can be discovered with v1 catalog.ListBlueprints.
	// printProviderId can be discovered with v1 catalog.ListPrintProvidersByBlueprint(idOne).
	GetShippingExpressInfoForVariantsOfBlueprintById = common.GetResourceWithTwoId[ShippingInfo, int, int](GET_SHIPPING_EXPRESS_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT)
	// GetShippingEconomyInfoForVariantsOfBlueprintById calls
	// GET /v2/catalog/blueprints/{blueprintId}/print_providers/{printProviderId}/shipping/economy.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int) (*ShippingInfo, error)
	// Parameter mapping:
	//	idOne -> {blueprintId}
	//	idTwo -> {printProviderId}
	//
	// blueprintId can be discovered with v1 catalog.ListBlueprints.
	// printProviderId can be discovered with v1 catalog.ListPrintProvidersByBlueprint(idOne).
	GetShippingEconomyInfoForVariantsOfBlueprintById = common.GetResourceWithTwoId[ShippingInfo, int, int](GET_SHIPPING_ECONOMY_INFO_FOR_VARIANTS_OF_BLUEPRINT_BY_ID_ENDPOINT)
)
