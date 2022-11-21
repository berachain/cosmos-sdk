package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the account contract that must be fulfilled when
// creating a x/bank keeper.
type AccountKeeper interface {
	NewAccount(sdk.Context, sdk.AccountI) sdk.AccountI
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) sdk.AccountI

	GetAccount(ctx sdk.Context, addr sdk.AccAddress) sdk.AccountI
	GetAllAccounts(ctx sdk.Context) []sdk.AccountI
	HasAccount(ctx sdk.Context, addr sdk.AccAddress) bool
	SetAccount(ctx sdk.Context, acc sdk.AccountI)

	IterateAccounts(ctx sdk.Context, process func(sdk.AccountI) bool)

	ValidatePermissions(macc types.ModuleAccountI) error

	GetModuleAddress(moduleName string) sdk.AccAddress
	GetModuleAddressAndPermissions(moduleName string) (addr sdk.AccAddress, permissions []string)
	GetModuleAccountAndPermissions(ctx sdk.Context, moduleName string) (types.ModuleAccountI, []string)
	GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI
	SetModuleAccount(ctx sdk.Context, macc types.ModuleAccountI)
	GetModulePermissions() map[string]types.PermissionsForAddress
}
