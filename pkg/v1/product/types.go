package product

import (
	"github.com/connellrobert/printify-go/pkg/v1/common"
)

// The Product resource lets you list, create, update, delete and publish products to a store.
type Product struct {
	// A unique string identifier for the product. Each id is unique across the Printify system.
	Id string `json:"id"`
	// The name of the product.
	Title string `json:"title"`
	// A description of the product. Supports HTML formatting.
	Description string `json:"description"`
	// Tags are also published to sales channel.
	Tags []string `json:"tags"`
	// Options are read only values and describes product options. There can be up to 3 options for a product.
	Options []ProductOptions `json:"options"`
	// A list of all product variants, each representing a different version of the product. But during product creation, only the variant id and price are necessary. See variant properties for reference.
	Variants []Variant `json:"variants"`
	// Mock-up images are read only values. The mock-up images are grouped by variants and position. See mock-up image properties for reference.
	Images []MockupImage `json:"images"`
	// The date and time when a product was created.
	CreatedAt string `json:"created_at"`
	// The date and time when a product was last updated.
	UpdateAt string `json:"update_at"`
	// Used for publishing. Visibility in sales channel. Can be true or false, defaults to true.
	Visible bool `json:"visible"`
	// Required when creating a product, but is read only after. See catalog for how to get blueprint_id.
	BlueprintId int `json:"blueprint_id"`
	// Required when creating a product, but is read only after. See catalog for how to get print_provider_id.
	PrintProviderId int `json:"print_provider_id"`
	// User id that a product belongs to.
	UserId int `json:"user_id"`
	// Shop id that a product belongs to.
	ShopId int `json:"shop_id"`
	// All print area values are required. Each variant has a print area attached to it. Each print area has placeholders which represent printable areas on a product. For example the front of the t-shirt, back of the t-shirt etc. Each placeholder has images and their positions, where they need to be printed in the printable area. See placeholder properties for reference.
	PrintAreas []PrintArea `json:"print_areas"`
	// "print_on_side" key is used to set the type of side printing for canvases. There are three possible values:
	// "regular" - to extend print area to the sides of canvas
	// "mirror" - to keep original print area and mirror it to the sides
	// "off" - stop printing on sides
	// Note: API documentation states this is a object, but actual results show it as a list. We're using what the
	//		API is actually returning.
	PrintDetails []common.PrintDetails `json:"print_details"`
	// Updated by sales channel with publishing succeeded endpoint. Id and handle are external references in the sales channel. See publishing succeeded endpoint for more reference.
	// Shipping Template ID is optional and can be passed during product creation or update.
	External PublishReference `json:"external"`
	// A product is locked during publishing. Locked products can't be updated until unlocked.
	IsLocked bool `json:"is_locked"`
	// Flag to indicate if product could be eligible for Printify Express Delivery, depending on selection of its variant.
	IsPrintifyExpressEligible bool `json:"is_printify_express_eligible"`
	// Flag to indicate if product could be eligible for Economy Shipping, depending on selection of its variant.
	IsEconomyShippingEligible bool `json:"is_economy_shipping_eligible"`
	// Flag to indicate if Printify Express Delivery is enabled for the product. Printify Express Delivery can be enabled only for eligible products (see is_printify_express_eligible flag). Defaults to false.
	IsPrintifyExpressEnabled bool `json:"is_printify_express_enabled"`
	// Flag to indicate if Economy Shipping is enabled for the product. Economy Shipping can be enabled only for eligible products (see is_economy_shipping_eligible flag). Defaults to false.
	IsEconomyShippingEnabled bool `json:"is_economy_shipping_enabled"`
	// Lists product properties specific to the sales channel associated with the product, if the sales channel has such custom properties, the attributes are listed in the array and may be actionable, but for all custom integrations, it will either be null or an empty array.
	SalesChannelProperties []interface{} `json:"sales_channel_properties"`
}

// Options are read only values and describes product options. There can be up to 3 options for a product.
type ProductOptions struct {
	Name   string               `json:"name"`
	Type   string               `json:"type"`
	Values []ProductOptionValue `json:"values"`
}

type ProductOptionValue struct {
	Id     int      `json:"id"`
	Title  string   `json:"title"`
	Colors []string `json:"colors"`
}

// A list of all product variants, each representing a different version of the product. But during product creation, only the variant id and price are necessary.
type Variant struct {
	// A unique int identifier for the product variant from the blueprint. Each id is unique across the Printify system. See catalog for instructions on how to get variant ids.
	Id int `json:"id"`
	// Optional unique string identifier for the product variant. If one is not provided, one will be generated by Printify.
	Sku string `json:"sku"`
	// Price in cents, integer value.
	Price int `json:"price"`
	// Product variant's fulfillment cost in cents, integer value.
	Cost int `json:"cost"`
	// Variant title.
	Title string `json:"title"`
	// Weight in grams for a product variant
	Grams int `json:"grams"`
	// Used for publishing, the value is true if one has the variant in question selected as an offering and wants it published.
	IsEnabled bool `json:"is_enabled"`
	// Only one variant can be default. Used when publishing to sales channel. Default variant's image will be the title image of the product.
	IsDefault bool `json:"is_default"`
	// Actual stock status of the variant, if false, the variant is out of stock and vice versa.
	IsAvailable bool `json:"is_available"`
	// Flag to indicate if product's variant is eligible for Printify Express delivery.
	IsPrintifyExpressEligible bool `json:"is_printify_express_eligible"`
	// Reference to options by id.
	Options []int `json:"options"`
}

// Mock-up images are read only values. The mock-up images are grouped by variants and position.
type MockupImage struct {
	// Url of a mock-up image.
	Src string `json:"src"`
	// Array of integer ids for variants illustrated by the mock-up image.
	VariantIds []int `json:"variant_ids"`
	// Camera position of a mockup (i.e. what part of the product is being displayed).
	Position string `json:"position"`
	// Used by the sales channel. If set to true, The specific mockup is the title image. Can be used to decide the first image displayed when a product's page is accessed.
	IsDefault bool `json:"is_default"`
}

// All print area values are required. Each variant has a print area attached to it. Each print area has placeholders which represent printable areas on a product. For example the front of the t-shirt, back of the t-shirt etc. Each placeholder has images and their positions, where they need to be printed in the printable area. See placeholder properties for reference.
type PrintArea struct {
	VariantIds   []int         `json:"variant_ids"`
	Placeholders []Placeholder `json:"placeholders"`
}

type Placeholder struct {
	// See blueprint placeholder properties from the catalog endpoint for reference on how to get positions.
	Position string `json:"position"`
	// Array of images. See image properties for reference.
	Images []Image `json:"images"`
}

type Image struct {
	// See upload images for reference on how to upload images and get all needed properties.
	Id string `json:"id"`
	// Name of an image file.
	Name string `json:"name"`
	// Type of an image. Valid image types are image/png, image/jpg, image/jpeg.
	Type string `json:"type"`
	// Float value for image height in pixels. See upload images for reference on how to upload images and get all needed properties.
	Height float64 `json:"height"`
	// Float value for image width in pixels. See upload images for reference on how to upload images and get all needed properties.
	Width float64 `json:"width"`
	// Float value to position an image on the x axis. See image positioning for reference on how to position images.
	X float64 `json:"x"`
	// Float value to position an image on the y axis. See image positioning for reference on how to position images.
	Y float64 `json:"y"`
	// Float value to scale image up or down. See image positioning for reference on how to position images.
	Scale float64 `json:"scale"`
	// Integer value used to rotate image. See image positioning for reference on how to position images.
	Angle int `json:"angle"`
}

// The "publish" button in the Printify app only locks the product on the Printify app and triggers the product:publish:started event if you are subscribed to it, see See Product events for reference. To publish a product, you need to create it manually on your store from the data you can obtain from the product resource or develop a system to automate that. Once done, you can use the Publish succeeded endpoint or Publish failed endpoint to unlock the product.
type Publish struct {
	// Used by the sales channel. If set to false, Images will not be published, and existing images will be used instead.
	Images bool `json:"images"`
	// Used by the sales channel. If set to false, product variations will not be published.
	Variants bool `json:"variants"`
	// Used by sales channel. If set to false, product title will not be updated.
	Title bool `json:"title"`
	// Used by sales channel. If set to false, product description will not be updated.
	Description bool `json:"description"`
	// Used by sales channel. If set to false, product tags will not be updated.
	Tags bool `json:"tags"`
	// Used by Walmart sales channel only. If set to false, product key features will not be updated.
	KeyFeatures bool `json:"key_features"`
	// Used by Etsy and Walmart sales channels only. If set to false, product shipping template will not be updated.
	ShippingTemplate bool `json:"shipping_template"`
}

// Updated by sales channel with publishing succeeded endpoint. Id and handle are external references in the sales channel. See publishing succeeded endpoint for more reference.
// Shipping Template ID is optional and can be passed during product creation or update.
type PublishReference struct {
	Id                 string `json:"id"`
	Handle             string `json:"handle"`
	ShippingTemplateId string `json:"shipping_template_id,omitempty"`
}

type PublishFailedRequest struct {
	Reason string `json:"reason"`
}
