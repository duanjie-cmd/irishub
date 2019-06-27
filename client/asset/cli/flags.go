package cli

import (
	"github.com/irisnet/irishub/app/v1/asset"
	flag "github.com/spf13/pflag"
)

const (
	FlagFamily         = "family"
	FlagSource         = "source"
	FlagGateway        = "gateway"
	FlagSymbol         = "symbol"
	FlagSymbolAtSource = "symbol-at-source"
	FlagName           = "name"
	FlagDecimal        = "decimal"
	FlagSymbolMinAlias = "symbol-min-alias"
	FlagInitialSupply  = "initial-supply"
	FlagMaxSupply      = "max-supply"
	FlagMintable       = "mintable"

	FlagOwner    = "owner"
	FlagMoniker  = "moniker"
	FlagIdentity = "identity"
	FlagDetails  = "details"
	FlagWebsite  = "website"
	FlagTo       = "to"

	FlagSubject = "subject"
	FlagID      = "id"
)

var (
	FsEditToken            = flag.NewFlagSet("", flag.ContinueOnError)
	FsAssetIssue           = flag.NewFlagSet("", flag.ContinueOnError)
	FsGatewayCreate        = flag.NewFlagSet("", flag.ContinueOnError)
	FsGatewayEdit          = flag.NewFlagSet("", flag.ContinueOnError)
	FsGatewayOwnerTransfer = flag.NewFlagSet("", flag.ContinueOnError)
	FsFeeQuery             = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsAssetIssue.String(FlagFamily, "", "the asset family, valid values can be fungible and non-fungible")
	FsAssetIssue.String(FlagSource, "", "the asset source, valid values can be native, external and gateway")
	FsAssetIssue.String(FlagGateway, "", "the gateway name of gateway asset. required if --source=gateway")
	FsAssetIssue.String(FlagSymbol, "", "the asset symbol. Once created, it cannot be modified")
	FsAssetIssue.String(FlagSymbolAtSource, "", "the source symbol of a gateway or external asset")
	FsAssetIssue.String(FlagName, "", "the asset name, e.g. IRIS Network")
	FsAssetIssue.Uint8(FlagDecimal, 0, "the asset decimal. The maximum value is 18")
	FsAssetIssue.String(FlagSymbolMinAlias, "", "the asset symbol minimum alias")
	FsAssetIssue.Uint64(FlagInitialSupply, 0, "the initial supply token of asset")
	FsAssetIssue.Uint64(FlagMaxSupply, asset.MaximumAssetMaxSupply, "the max supply token of asset")
	FsAssetIssue.Bool(FlagMintable, false, "whether the asset can be minted, default false")

	FsGatewayCreate.String(FlagMoniker, "", "the unique gateway name")
	FsGatewayCreate.String(FlagIdentity, "", "the gateway identity")
	FsGatewayCreate.String(FlagDetails, "", "the gateway description")
	FsGatewayCreate.String(FlagWebsite, "", "the external website")

	FsGatewayEdit.String(FlagMoniker, asset.DoNotModify, "the unique gateway name")
	FsGatewayEdit.String(FlagIdentity, asset.DoNotModify, "the gateway identity")
	FsGatewayEdit.String(FlagDetails, asset.DoNotModify, "the gateway description")
	FsGatewayEdit.String(FlagWebsite, asset.DoNotModify, "the external website")

	FsGatewayOwnerTransfer.String(FlagMoniker, "", "the unique name of the gateway to be transferred")
	FsGatewayOwnerTransfer.String(FlagTo, "", "the new owner")

	FsFeeQuery.String(FlagSubject, "", "the fee type to be queried")
	FsFeeQuery.String(FlagMoniker, "", "the gateway name, required if the subject is gateway")
	FsFeeQuery.String(FlagID, "", "the token id, required if the subject is token")

	FsEditToken.String(FlagName, "[do-not-modify]", "the token name, e.g. IRIS Network")
	FsEditToken.String(FlagSymbolAtSource, "[do-not-modify]", "the source symbol of a gateway or external token")
	FsEditToken.String(FlagSymbolMinAlias, "[do-not-modify]", "the token symbol minimum alias")
	FsEditToken.Uint64(FlagMaxSupply, 0, "the max supply of token")
	FsEditToken.Bool(FlagMintable, false, "whether the token can be minted, default false")
}