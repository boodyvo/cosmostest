package keeper

import (
	"fmt"

	"github.com/boodyvo/cosmostest/x/cosmostest/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the cosmostest Querier
const (
	QueryResolve = "resolve"
	QueryWhois   = "whois"
	QueryNames   = "names"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryResolve:
			return queryResolve(ctx, path[1:], req, keeper)
		case QueryWhois:
			return queryWhoIs(ctx, path[1:], req, keeper)
		case QueryNames:
			return queryNames(ctx, req, keeper)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown query endpoint")
		}
	}
}

// nolint: unparam
func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	value := keeper.ResolveName(ctx, path[0])

	if value.IsNil() {
		return []byte{}, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not resolve name")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: value})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// nolint: unparam
func queryWhoIs(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	whoIs := keeper.GetWhoIs(ctx, path[0])

	res, err := codec.MarshalJSONIndent(keeper.cdc, whoIs)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryNames(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	var namesList types.QueryResNames

	iterator := keeper.GetNamesIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		name := string(iterator.Key())
		value := keeper.ResolveName(ctx, name)
		fmt.Println("resolving", value, name)

		namesList = append(namesList, types.QueryValue{
			Name:  string(iterator.Key()),
			Value: value,
		})
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, namesList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
