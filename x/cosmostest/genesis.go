package cosmostest

import (
	"fmt"

	"github.com/boodyvo/cosmostest/x/cosmostest/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState struct {
	WhoIsRecords []types.WhoIs `json:"whois_records"`
}

func NewGenesisState(whoIsRecords []types.WhoIs) GenesisState {
	return GenesisState{WhoIsRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.WhoIsRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid WhoIsRecord: Value: %s. Error: Missing Owner", record.Value)
		}
		if record.Value.IsNil() {
			return fmt.Errorf("invalid WhoIsRecord: Owner: %s. Error: Missing Value", record.Owner)
		}
		if record.Price == nil {
			return fmt.Errorf("invalid WhoIsRecord: Value: %s. Error: Missing Price", record.Value)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		WhoIsRecords: []types.WhoIs{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, record := range data.WhoIsRecords {
		keeper.SetWhoIs(ctx, record.Value.String(), record)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []types.WhoIs
	iterator := k.GetNamesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		name := string(iterator.Key())
		whois := k.GetWhoIs(ctx, name)
		records = append(records, whois)

	}
	return GenesisState{WhoIsRecords: records}
}
