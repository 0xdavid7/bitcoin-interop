package types

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: Params{},
	}
}

func (data GenesisState) Validate() error {
	// Add validation logic
	return nil
}

func NewGenesisState(params Params) *GenesisState {
	return &GenesisState{
		Params: params,
	}
}
