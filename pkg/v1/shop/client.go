package shop

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/connellrobert/printify-go/pkg/common"
)

var (
	ENDPOINT             = "/v1/shops"
	LIST_SHOPS_ENDPOINT  = ENDPOINT + ".json"
	DELETE_SHOP_ENDPOINT = ENDPOINT + "/%d.json"
)

func ListShops(c *common.Client) ([]Shop, error) {
	req, err := http.NewRequest("GET", c.Host+LIST_SHOPS_ENDPOINT, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var shops []Shop
	err = json.NewDecoder(resp.Body).Decode(&shops)
	if err != nil {
		return nil, err
	}
	return shops, nil
}

func DeleteShop(c *common.Client, shopId int) error {
	req, err := http.NewRequest("DELETE", c.Host+fmt.Sprintf(DELETE_SHOP_ENDPOINT, shopId), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
