package catalog

import (
	"fmt"

	"github.com/connellrobert/printify-go/pkg/common"
)

// Client defines catalog operations and enables dependency injection.
type Client interface {
	ListBlueprints() ([]Blueprint, error)
	GetBlueprint(id int) (*Blueprint, error)
	ListPrintProvidersByBlueprint(id int) ([]PrintProvider, error)
	ListVariantsByBlueprintPrintProvider(idOne int, idTwo int) ([]Variant, error)
	GetShippingInformation(idOne int, idTwo int) (*Shipping, error)
	ListPrintProviders() ([]PrintProvider, error)
	GetPrintProvider(id int) (*PrintProvider, error)
}

type client struct {
	c *common.Client
}

// NewClient creates a catalog client implementation backed by common.Client.
func NewClient(c *common.Client) Client {
	return &client{c: c}
}

func (cl *client) ListBlueprints() ([]Blueprint, error) {
	return ListBlueprints(cl.c)
}

func (cl *client) GetBlueprint(id int) (*Blueprint, error) {
	return GetBlueprint(cl.c, id)
}

func (cl *client) ListPrintProvidersByBlueprint(id int) ([]PrintProvider, error) {
	return ListPrintProvidersByBlueprint(cl.c, id)
}

func (cl *client) ListVariantsByBlueprintPrintProvider(idOne int, idTwo int) ([]Variant, error) {
	return ListVariantsByBlueprintPrintProvider(cl.c, idOne, idTwo)
}

func (cl *client) GetShippingInformation(idOne int, idTwo int) (*Shipping, error) {
	return GetShippingInformation(cl.c, idOne, idTwo)
}

func (cl *client) ListPrintProviders() ([]PrintProvider, error) {
	return ListPrintProviders(cl.c)
}

func (cl *client) GetPrintProvider(id int) (*PrintProvider, error) {
	return GetPrintProvider(cl.c, id)
}

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
	// ListBlueprints calls GET /v1/catalog/blueprints.json and returns all available product blueprints.
	//
	// Signature:
	//	func(c *common.Client) ([]Blueprint, error)
	ListBlueprints = common.ListResources[Blueprint](LIST_BLUEPRINTS_ENDPOINT)
	// GetBlueprint calls GET /v1/catalog/blueprints/{blueprintId}.json and returns a single blueprint.
	//
	// Signature:
	//	func(c *common.Client, id int) (*Blueprint, error)
	// Parameter mapping:
	//	id -> {blueprintId}
	//
	// The blueprint id can be discovered from Printify's catalog UI or by calling ListBlueprints.
	GetBlueprint = common.GetResourceById[Blueprint, int](GET_BLUEPRINT_ENDPOINT)
	// ListPrintProvidersByBlueprint calls GET /v1/catalog/blueprints/{blueprintId}/print_providers.json.
	//
	// Signature:
	//	func(c *common.Client, id int) ([]PrintProvider, error)
	// Parameter mapping:
	//	id -> {blueprintId}
	//
	// The blueprint id can be discovered from Printify's catalog UI or by calling ListBlueprints.
	ListPrintProvidersByBlueprint = common.ListResourceWithId[PrintProvider, int](LIST_PRINT_PROVIDERS_BY_BLUEPRINT_ENDPOINT)
	// ListVariantsByBlueprintPrintProvider calls
	// GET /v1/catalog/blueprints/{blueprintId}/print_providers/{printProviderId}/variants.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int) ([]Variant, error)
	// Parameter mapping:
	//	idOne -> {blueprintId}
	//	idTwo -> {printProviderId}
	//
	// The blueprint id can be discovered from Printify's catalog UI or by calling ListBlueprints.
	// The print provider id can be discovered by calling ListPrintProvidersByBlueprint(idOne)
	// or ListPrintProviders when you already know the provider.
	ListVariantsByBlueprintPrintProvider = common.ListResourceWithTwoID[Variant, int, int](LIST_VARIANTS_BY_BLUEPRINT_PRINT_PROVIDER_ENDPOINT)
	// GetShippingInformation calls
	// GET /v1/catalog/blueprints/{blueprintId}/print_providers/{printProviderId}/shipping.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int) (*Shipping, error)
	// Parameter mapping:
	//	idOne -> {blueprintId}
	//	idTwo -> {printProviderId}
	//
	// The blueprint id can be discovered from Printify's catalog UI or by calling ListBlueprints.
	// The print provider id can be discovered by calling ListPrintProvidersByBlueprint(idOne)
	// or ListPrintProviders when you already know the provider.
	GetShippingInformation = common.GetResourceWithTwoId[Shipping, int, int](GET_SHIPPING_INFORMATION_ENDPOINT)
	// ListPrintProviders calls GET /v1/catalog/print_providers.json and returns all print providers.
	//
	// Signature:
	//	func(c *common.Client) ([]PrintProvider, error)
	ListPrintProviders = common.ListResources[PrintProvider](LIST_PRINT_PROVIDERS_ENDPOINT)
	// GetPrintProvider calls GET /v1/catalog/print_providers/{printProviderId}.json.
	//
	// Signature:
	//	func(c *common.Client, id int) (*PrintProvider, error)
	// Parameter mapping:
	//	id -> {printProviderId}
	//
	// The print provider id can be discovered from Printify's catalog UI or by calling ListPrintProviders.
	GetPrintProvider = common.GetResourceById[PrintProvider, int](GET_PRINT_PROVIDER_ENDPOINT)
)
