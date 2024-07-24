package order

type Order struct {
	// A unique string identifier for the order. Each id is unique across the Printify system.
	Id string `json:"id"`
	// The delivery details of the order's recipient.
	AddressTo Address `json:"address_to"`
	// A list of all line items in the order. See line item properties for reference.
	LineItems []LineItem `json:"line_items"`
	// Other data about the order. See OrderMetadata properties for reference.
	Metadata OrderMetadata `json:"metadata"`
	// Retail price in cents, integer value.
	TotalPrice int `json:"total_price"`
	// Shipping price in cents, integer value.
	TotalShipping int `json:"total_shipping"`
	// Tax cost in cents, integer value.
	TotalTax int `json:"total_tax"`
	// Production status of the entire order in string format, it can be any of the following:
	// Status	Definition
	// pending	An order is created in the pending status. Orders should not stay in this status for a long time.
	// on-hold
	// The order is ready to handle user actions. Users can edit orders in this status. The order gets this status at later stages for various reasons:

	// Line items are discontinued during the order processing
	// Line items go out of stock during the order processing
	// The order contain items with shipping restricted line items
	// User needs to take actions upon orders that got the status during the submission.

	// sending-to-production	Order is picked for sending to production. The status should be changed when the system receives updates from print providers.
	// in-production	Order has been received by the print providers successfully. Print providers start the fulfillment process on their side.
	// canceled	The order is canceled. No actions can be taken towards the order by the user.
	// fulfilled	All line items are fulfilled by all print providers. The last status for successful order fulfillment.
	// partially-fulfilled	At least one of the line items is fulfilled, but not all of them.
	// payment-not-received	The order could not be charged during the submission in our system. It can be retried by the merchant. The order waits for merchant actions.
	// had-issues	If an order encounters any issue during its lifecycle it will get this status. For instance, the shipping address is invalid.
	Status string `json:"status"`
	// Method of shipping.
	// "1" is for standard shipping
	// "2" is for priority shipping
	// "3" is for Printify Express shipping
	// "4" is for Economy shipping
	ShippingMethod int `json:"shipping_method"`
	// Boolean value that indicates if the order is using Printify Express shipping.
	IsPrintifyExpress bool `json:"is_printify_express"`
	// Boolean value that indicates if the order is using Economy shipping.
	IsEconomyShipping bool `json:"is_economy_shipping"`
	// Tracking details of the order after fulfillment. See shipment properties for reference.
	Shipments []Shipment `json:"shipments"`
	// The date and time the order was created. It is stored in ISO date format.
	CreatedAt string `json:"created_at"`
	// The date and time the order was sent to production. It is stored in ISO date format.
	SentToProductionAt string `json:"sent_to_production_at"`
	// The date and time the order was fulfilled. It is stored in ISO date format.
	FulfilledAt string `json:"fulfilled_at"`
	// Printify Connect data containing link to the order in the Printify Connect page and the unique hash for the order.
	// More about Printify Connect can be read in our Help Center or Blog.
	PrintifyConnect PrintifyConnect `json:"printify_connect"`
}

type Address struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Region    string `json:"region"`
	Address1  string `json:"address1"`
	City      string `json:"city"`
	Zip       string `json:"zip"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Country   string `json:"country"`
	Company   string `json:"company"`
}

type LineItem struct {
	// A unique string identifier for the product. Each id is unique across the Printify system.
	ProductId string `json:"product_id"`
	// A unique int identifier for the product variant from the blueprint. Each id is unique across the Printify system.
	VariantId string `json:"variant_id"`
	// Describes the number of said product ordered as an integer.
	Quantity int `json:"quantity"`
	// A unique int identifier for the print provider. Each id is unique across the Printify system.
	PrintProviderId int `json:"print_provider_id"`
	// Product variant's fulfillment cost in cents, integer value.
	Cost int `json:"cost"`
	// Product variant's shipment cost in cents, integer value.
	ShippingCost int `json:"shipping_cost"`
	// Specific line item fulfillment status:
	// Status	Definition
	// on-hold
	// The item waits for user actions. Items get this status when:

	// The order is created.
	// The order fails at one of the checks in the submission process.
	// The order fails with payments.
	// in-production	Print provider received the item for production.
	// sending-to-production	Says that the order is picked for sending to production. The status is changed once the system receives updates from print providers.
	// has-issues	Line item encountered an issue during its lifecycle.
	// fulfilled	The item is fulfilled by the print provider.
	// canceled	The item is canceled.
	Status string `json:"status"`
	// Other details about the specific product variant. See line item metadata properties for reference.
	Metadata LineItemMetadata `json:"metadata"`
	// The date and time the product variant was sent to production. It is stored in ISO date format.
	SentToProductionAt string `json:"sent_to_production_at"`
	// The date and time the product variant was fulfilled. It is stored in ISO date format.
	FulfilledAt string `json:"fulfilled_at"`
}

type LineItemMetadata struct {
	// The name of the product.
	Title string `json:"title"`
	// Retail price in cents, integer value.
	Price int `json:"price"`
	// Name of the product variant.
	VariantLabel string `json:"variant_label"`
	// A unique string identifier for the product variant.
	Sku string `json:"sku"`
	// Location of print provider handling fulfillment.
	Country string `json:"country"`
}

type OrderMetadata struct {
	// Describes the order type, can be external, manual, or sample.
	OrderType string `json:"order_type"`
	// A unique integer identifier for the order in the external sales channel.
	ShopOrderId int `json:"shop_order_id"`
	// A unique string identifier for the order in the external sales channel.
	ShopOrderLabel string `json:"shop_order_label"`
	// The date and time the order was fulfilled. It is stored in ISO date format.
	ShopFulfilledAt string `json:"shop_fulfilled_at"`
}

type Shipment struct {
	// Name of the shipping courier used to deliver the order to its recipient.
	Carrier string `json:"carrier"`
	// A unique string tracking number from the shipping courier used to track the status of the shipment.
	Number string `json:"number"`
	// A unique string tracking link from the shipping courier used to track the status of the shipment.
	Url string `json:"url"`
	// The date and time the order was delivered. It is stored in ISO date format.
	DeliveredAt string `json:"delivered_at"`
}

type PrintifyConnect struct {
	Url string `json:"url"`
	Id  string `json:"id"`
}

type OrderSubmission struct {
	// A unique string identifier from the sales channel specifying the order name or id.
	ExternalId string `json:"external_id"`
	// Optional value to specify order label instead of using "external_id"
	Label string `json:"label"`
	// Required for ordering existing products. Provide the product_id (Printify Product ID), variant_id (selected variant, e.g. 'White / XXL') and desired item quantity. If creating a product from the order is required, then additional attributes will need to be provided, specifically the blueprint_id and print_areas.
	LineItems []OrderSubmissionLineItem `json:"line_items"`
	// Required to specify what method of shipping is desired, "1" means standard shipping, "2" means priority shipping, "3" means printify express shipping and "4" means economy shipping. It is stored as an integer.
	ShippingMethod int `json:"shipping_method"`
	// A boolean for choosing whether or not to receive email notifications after an order is shipped.
	SendShippingNotification bool `json:"send_shipping_notification"`
	// The delivery details of the order's recipient.
	AddressTo Address `json:"address_to"`
}

type OrderSubmissionLineItem struct {
	PrintProviderId int                         `json:"print_provider_id"`
	BlueprintId     int                         `json:"blueprint_id"`
	VariantId       int                         `json:"variant_id"`
	PrintAreas      map[string][]PrintAreaValue `json:"print_areas"`
}

type PrintAreaValue struct {
	Src   string  `json:"src"`
	Scale float64 `json:"scale"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Angle float64 `json:"angle"`
}

type ShipmentCalculationRequest struct {
	LineItems []LineItem `json:"line_items"`
	AddressTo Address    `json:"address_to"`
}

type ShipmentCalculationResponse struct {
	Standard        int `json:"standard"`
	Express         int `json:"express"`
	Priority        int `json:"priority"`
	PrintifyExpress int `json:"printify_express"`
	Economy         int `json:"economy"`
}
