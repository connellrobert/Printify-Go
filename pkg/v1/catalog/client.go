package catalog

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/connellrobert/printify-go/pkg/common"
)

var (
	ENDPOINT                                           = "/v1/catalog"
	LIST_BLUEPRINTS_ENDPOINT                           = fmt.Sprintf("%s/blueprints.json", ENDPOINT)
	GET_BLUEPRINT_ENDPOINT                             = fmt.Sprintf("%s/blueprints/%%d.json", ENDPOINT)
	LIST_PRINT_PROVIDERS_BY_BLUEPRINT_ENDPOINT         = fmt.Sprintf("%s/blueprints/%%d/print_providers.json", ENDPOINT)
	LIST_VARIANTS_BY_BLUEPRINT_PRINT_PROVIDER_ENDPOINT = fmt.Sprintf("%s/blueprints/%%d/print_providers/%%d/variants.json", ENDPOINT)
	GET_SHIPPING_INFORMATION_ENDPOINT                  = fmt.Sprintf("%s/blueprints/%%d/print_providers/%%d/shipping.json", ENDPOINT)
	LIST_PRINT_PROVIDERS_ENDPOINT                      = fmt.Sprintf("%s/print_providers.json", ENDPOINT)
	GET_PRINT_PROVIDER_ENDPOINT                        = fmt.Sprintf("%s/print_providers/%%d.json", ENDPOINT)
)

func ListBlueprints(c *common.Client) ([]Blueprint, error) {
	req, err := http.NewRequest("GET", c.Host+LIST_BLUEPRINTS_ENDPOINT, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var blueprints []Blueprint
	err = json.NewDecoder(resp.Body).Decode(&blueprints)
	if err != nil {
		return nil, err
	}

	return blueprints, nil
}

func GetBlueprint(c *common.Client, id int) (*Blueprint, error) {
	req, err := http.NewRequest("GET", c.Host+fmt.Sprintf(GET_BLUEPRINT_ENDPOINT, id), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var blueprint Blueprint
	err = json.NewDecoder(resp.Body).Decode(&blueprint)
	if err != nil {
		return nil, err
	}

	return &blueprint, nil
}

func ListPrintProvidersByBlueprint(c *common.Client, id int) ([]PrintProvider, error) {
	req, err := http.NewRequest("GET", c.Host+fmt.Sprintf(LIST_PRINT_PROVIDERS_BY_BLUEPRINT_ENDPOINT, id), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var printProviders []PrintProvider
	err = json.NewDecoder(resp.Body).Decode(&printProviders)
	if err != nil {
		return nil, err
	}

	return printProviders, nil
}

func ListVariantsByBlueprintPrintProvider(c *common.Client, blueprintId, printProviderId int) ([]Variant, error) {
	req, err := http.NewRequest("GET", c.Host+fmt.Sprintf(LIST_VARIANTS_BY_BLUEPRINT_PRINT_PROVIDER_ENDPOINT, blueprintId, printProviderId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var variants []Variant
	err = json.NewDecoder(resp.Body).Decode(&variants)
	if err != nil {
		return nil, err
	}

	return variants, nil
}

func GetShippingInformation(c *common.Client, blueprintId, printProviderId int) (*Shipping, error) {
	req, err := http.NewRequest("GET", c.Host+fmt.Sprintf(GET_SHIPPING_INFORMATION_ENDPOINT, blueprintId, printProviderId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var shipping Shipping
	err = json.NewDecoder(resp.Body).Decode(&shipping)
	if err != nil {
		return nil, err
	}

	return &shipping, nil
}

func ListPrintProviders(c *common.Client) ([]PrintProvider, error) {
	req, err := http.NewRequest("GET", c.Host+LIST_PRINT_PROVIDERS_ENDPOINT, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var printProviders []PrintProvider
	err = json.NewDecoder(resp.Body).Decode(&printProviders)
	if err != nil {
		return nil, err
	}

	return printProviders, nil
}

func GetPrintProvider(c *common.Client, id int) (*PrintProvider, error) {
	req, err := http.NewRequest("GET", c.Host+fmt.Sprintf(GET_PRINT_PROVIDER_ENDPOINT, id), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var printProvider PrintProvider
	err = json.NewDecoder(resp.Body).Decode(&printProvider)
	if err != nil {
		return nil, err
	}

	return &printProvider, nil
}
