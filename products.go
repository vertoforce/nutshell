package nutshell

type Product struct {
	Stub

	CreatedTime  string `json:"createdTime"`
	Relationship interface{}
	ModifiedTime string `json:"modifiedTime"`
	Price        struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	}
	Quantity float64 `json:"quantity"`
}

type ProductsList []Product

// func (c *Client) GetProduct(id int) (*Product, error) {
// 	con := &Product{}

// 	err := c.rpc.Call("getProduct", map[string]interface{}{"productId": id}, con)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return con, nil
// }
