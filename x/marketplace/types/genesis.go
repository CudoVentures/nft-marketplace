package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

func NewGenesisState(listings []Listing, listingCount uint64) *GenesisState {
	return &GenesisState{
		Listings:     listings,
		ListingCount: listingCount,
	}
}

func ValidateGenesis(data GenesisState) error {
	for _, l := range data.Listings {
		if l.GetOwner().Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing onft owner")
		}

		if err := ValidateId(l.GetId()); err != nil {
			return err
		}

		if err := ValidatePrice(l.Price); err != nil {
			return err
		}
	}
	return nil
}
