package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QueryResResolve Queries Result Payload for a resolve query
type QueryResResolve struct {
	Value sdk.Dec `json:"value"`
}

// implement fmt.Stringer
func (r QueryResResolve) String() string {
	return r.Value.String()
}

type QueryValue struct {
	Name  string  `json:"name"`
	Value sdk.Dec `json:"value"`
}

func (qv QueryValue) String() string {
	res, err := json.Marshal(qv)
	if err != nil {
		return err.Error()
	}
	return string(res)
}

// QueryResNames Queries Result Payload for a names query
type QueryResNames []QueryValue

// implement fmt.Stringer
func (n QueryResNames) String() string {
	res := ""
	for i, qv := range n {
		if i == len(n)-1 {
			res += qv.String()
		} else {
			res += fmt.Sprintf("%s,", qv.String())
		}
	}
	return res
}
