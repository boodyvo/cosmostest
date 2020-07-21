package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MinNamePrice is Initial Starting Price for a name that was never previously owned
var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

// WhoIs is a struct that contains all the metadata of a name
type WhoIs struct {
	Value sdk.Dec        `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}

// NewWhois returns a new WhoIs with the minprice as the price
func NewWhoIs() WhoIs {
	return WhoIs{
		Price: MinNamePrice,
	}
}

// implement fmt.Stringer
func (w WhoIs) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Price: %s`, w.Owner, w.Value, w.Price))
}
