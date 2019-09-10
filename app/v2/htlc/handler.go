package htlc

import (
	"encoding/hex"

	sdk "github.com/irisnet/irishub/types"
)

// NewHandler handles all "htlc" messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreateHTLC:
			return handleMsgCreateHTLC(ctx, k, msg)
		case MsgClaimHTLC:
			return handleMsgClaimHTLC(ctx, k, msg)
		case MsgRefundHTLC:
			return handleMsgRefundHTLC(ctx, k, msg)
		default:
			return sdk.ErrTxDecode("invalid message parsed in HTLC module").Result()
		}
	}
}

// handleMsgCreateHTLC handles MsgCreateHTLC
func handleMsgCreateHTLC(ctx sdk.Context, k Keeper, msg MsgCreateHTLC) sdk.Result {
	secret := make([]byte, 32)
	expireHeight := msg.TimeLock + uint64(ctx.BlockHeight())
	state := StateOpen

	htlc := NewHTLC(msg.Sender, msg.Receiver, msg.ReceiverOnOtherChain, msg.OutAmount, msg.InAmount, secret, msg.Timestamp, expireHeight, state)
	secretHashLock, _ := hex.DecodeString(msg.SecretHashLock)

	tags, err := k.CreateHTLC(ctx, htlc, secretHashLock)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Tags: tags,
	}
}

// handleMsgClaimHTLC handles MsgClaimHTLC
func handleMsgClaimHTLC(ctx sdk.Context, k Keeper, msg MsgClaimHTLC) sdk.Result {
	secret, _ := hex.DecodeString(msg.Secret)
	secretHash, _ := hex.DecodeString(msg.SecretHashLock)
	tags, err := k.ClaimHTLC(ctx, secret, secretHash)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Tags: tags,
	}
}

// handleMsgRefundHTLC handles MsgRefundHTLC
func handleMsgRefundHTLC(ctx sdk.Context, k Keeper, msg MsgRefundHTLC) sdk.Result {
	secretHash, _ := hex.DecodeString(msg.SecretHashLock)
	tags, err := k.RefundHTLC(ctx, secretHash)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Tags: tags,
	}
}