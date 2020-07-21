package keeper

import (
	"github.com/boodyvo/cosmostest/x/cosmostest/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	CoinKeeper types.BankKeeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// Sets the entire WhoIs metadata struct for a name
func (k Keeper) SetWhoIs(ctx sdk.Context, name string, whoIs types.WhoIs) {
	if whoIs.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(whoIs))
}

// Gets the entire WhoIs metadata struct for a name
func (k Keeper) GetWhoIs(ctx sdk.Context, name string) types.WhoIs {
	store := ctx.KVStore(k.storeKey)
	if !k.IsNamePresent(ctx, name) {
		return types.NewWhoIs()
	}
	bz := store.Get([]byte(name))
	var whoIs types.WhoIs
	k.cdc.MustUnmarshalBinaryBare(bz, &whoIs)
	return whoIs
}

// Deletes the entire WhoIs metadata struct for a name
func (k Keeper) DeleteWhoIs(ctx sdk.Context, name string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(name))
}

// ResolveName - returns the sdk.Dec that the name resolves to
func (k Keeper) ResolveName(ctx sdk.Context, name string) sdk.Dec {
	return k.GetWhoIs(ctx, name).Value
}

// SetName - sets the value sdk.Dec that a name resolves to
func (k Keeper) SetName(ctx sdk.Context, name string, value sdk.Dec) {
	whoIs := k.GetWhoIs(ctx, name)
	whoIs.Value = value
	k.SetWhoIs(ctx, name, whoIs)
}

// HasOwner - returns whether or not the name already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
	return !k.GetWhoIs(ctx, name).Owner.Empty()
}

// GetOwner - get the current owner of a name
func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	return k.GetWhoIs(ctx, name).Owner
}

// SetOwner - sets the current owner of a name
func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	whoIs := k.GetWhoIs(ctx, name)
	whoIs.Owner = owner
	k.SetWhoIs(ctx, name, whoIs)
}

// GetPrice - gets the current price of a name
func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
	return k.GetWhoIs(ctx, name).Price
}

// SetPrice - sets the current price of a name
func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	whoIs := k.GetWhoIs(ctx, name)
	whoIs.Price = price
	k.SetWhoIs(ctx, name, whoIs)
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

// NewKeeper creates new instances of the cosmostest Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, coinKeeper types.BankKeeper) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		CoinKeeper: coinKeeper,
	}
}
