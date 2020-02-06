package nutshell

// Add ...
func (c *Client) Add(a, b int) (int, error) {
	args := map[string]interface{}{
		"num1": a,
		"num2": b,
	}
	var ans int
	err := c.rpc.Call("add", args, &ans)
	return ans, err
}
