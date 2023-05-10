package keeper

import (
	"bytes"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ibc "github.com/cosmos/ibc-go/v3/modules/core"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v3/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
	ibcdmtypes "github.com/cosmos/ibc-go/v3/modules/light-clients/01-furyint/types"

	"github.com/gridfoundation/gridiron/x/rollapp/types"
)

var _ ibc.IBCMsgI = Keeper{}

func (k Keeper) CreateClientValidate(
	ctx sdk.Context,
	clientState exported.ClientState,
	consensusState exported.ConsensusState,
) error {
	// filter only rollapp chains
	chainID := clientState.GetChainID()
	if isFuryint, err := k.isRollappChain(ctx, clientState.ClientType(), chainID); !isFuryint || err != nil {
		return err
	}
	// get application stateRoot
	stateRoot := consensusState.GetRoot().GetHash()
	// get height
	height := clientState.GetLatestHeight().GetRevisionHeight()
	// check stateRoot validity
	return k.validateStateRoot(ctx, chainID, height, stateRoot)
}

// CreateClient defines a rpc handler method for MsgCreateClient.
func (k Keeper) CreateClient(goCtx context.Context, msg *clienttypes.MsgCreateClient) (*clienttypes.MsgCreateClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// unpack message data
	clientState, err := clienttypes.UnpackClientState(msg.ClientState)
	if err != nil {
		return nil, err
	}
	consensusState, err := clienttypes.UnpackConsensusState(msg.ConsensusState)
	if err != nil {
		return nil, err
	}

	// check validity
	if err := k.CreateClientValidate(ctx, clientState, consensusState); err != nil {
		return nil, err
	}

	// pass to ibc keeper
	return k.ibcKeeper.CreateClient(goCtx, msg)
}

func (k Keeper) UpdateClientValidate(
	ctx sdk.Context,
	clientID string,
	header exported.Header,
) error {
	// filter only rollapp chains
	chainID := header.GetChainID()
	if isFuryint, err := k.isRollappChain(ctx, header.ClientType(), chainID); !isFuryint || err != nil {
		return err
	}
	furyHeader := header.(*ibcdmtypes.Header)
	// get application stateRoot
	stateRoot := furyHeader.Header.GetAppHash()
	// get height
	height := furyHeader.Header.Height
	// check stateRoot validity
	return k.validateStateRoot(ctx, chainID, uint64(height), stateRoot)
}

// UpdateClient defines a rpc handler method for MsgUpdateClient.
func (k Keeper) UpdateClient(goCtx context.Context, msg *clienttypes.MsgUpdateClient) (*clienttypes.MsgUpdateClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// unpack message data
	header, err := clienttypes.UnpackHeader(msg.Header)
	if err != nil {
		return nil, err
	}

	// check validity
	if err := k.UpdateClientValidate(ctx, msg.ClientId, header); err != nil {
		return nil, err
	}

	// pass to ibc keeper
	return k.ibcKeeper.UpdateClient(goCtx, msg)
}

func (k Keeper) UpgradeClientValidate(
	ctx sdk.Context,
	clientID string,
	upgradedClient exported.ClientState,
	upgradedConsState exported.ConsensusState,
	proofUpgradeClient,
	proofUpgradeConsState []byte,
) error {
	// filter only rollapp chains
	chainID := upgradedClient.GetChainID()
	if isFuryint, err := k.isRollappChain(ctx, upgradedClient.ClientType(), chainID); !isFuryint || err != nil {
		return err
	}
	// get application stateRoot
	stateRoot := upgradedConsState.GetRoot().GetHash()
	// get height
	height := upgradedClient.GetLatestHeight().GetRevisionHeight()
	// check stateRoot validity
	return k.validateStateRoot(ctx, chainID, height, stateRoot)
}

// UpgradeClient defines a rpc handler method for MsgUpgradeClient.
func (k Keeper) UpgradeClient(goCtx context.Context, msg *clienttypes.MsgUpgradeClient) (*clienttypes.MsgUpgradeClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// unpack message data
	clientState, err := clienttypes.UnpackClientState(msg.ClientState)
	if err != nil {
		return nil, err
	}
	upgradedConsState, err := clienttypes.UnpackConsensusState(msg.ConsensusState)
	if err != nil {
		return nil, err
	}

	// check validity
	if err := k.UpgradeClientValidate(ctx, msg.ClientId, clientState, upgradedConsState, msg.ProofUpgradeClient, msg.ProofUpgradeConsensusState); err != nil {
		return nil, err
	}

	// pass to ibc keeper
	return k.ibcKeeper.UpgradeClient(goCtx, msg)
}

func (k Keeper) SubmitMisbehaviourValidate(
	ctx sdk.Context,
	misbehaviour exported.Misbehaviour,
) error {
	// filter only rollapp chains
	chainID := misbehaviour.GetChainID()
	if isFuryint, err := k.isRollappChain(ctx, misbehaviour.ClientType(), chainID); !isFuryint || err != nil {
		return err
	}

	furyHeader1 := misbehaviour.(*ibcdmtypes.Misbehaviour).Header1
	furyHeader2 := misbehaviour.(*ibcdmtypes.Misbehaviour).Header2
	// get application stateRoot
	stateRoot1 := furyHeader1.Header.GetAppHash()
	stateRoot2 := furyHeader2.Header.GetAppHash()
	// get height
	height1 := furyHeader1.Header.Height
	height2 := furyHeader2.Header.Height

	// check stateRoot validity
	if err := k.validateStateRoot(ctx, chainID, uint64(height1), stateRoot1); err != nil {
		return err
	}
	return k.validateStateRoot(ctx, chainID, uint64(height2), stateRoot2)
}

// SubmitMisbehaviour defines a rpc handler method for MsgSubmitMisbehaviour.
func (k Keeper) SubmitMisbehaviour(goCtx context.Context, msg *clienttypes.MsgSubmitMisbehaviour) (*clienttypes.MsgSubmitMisbehaviourResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// unpack message data
	misbehaviour, err := clienttypes.UnpackMisbehaviour(msg.Misbehaviour)
	if err != nil {
		return nil, err
	}

	// check validity
	if err := k.SubmitMisbehaviourValidate(ctx, misbehaviour); err != nil {
		return nil, err
	}

	// pass to ibc keeper
	return k.ibcKeeper.SubmitMisbehaviour(goCtx, msg)
}

// ConnectionOpenInit defines a rpc handler method for MsgConnectionOpenInit.
func (k Keeper) ConnectionOpenInit(goCtx context.Context, msg *connectiontypes.MsgConnectionOpenInit) (*connectiontypes.MsgConnectionOpenInitResponse, error) {
	println("my ConnectionOpenInit")
	return k.ibcKeeper.ConnectionOpenInit(goCtx, msg)
}

// ConnectionOpenTry defines a rpc handler method for MsgConnectionOpenTry.
func (k Keeper) ConnectionOpenTry(goCtx context.Context, msg *connectiontypes.MsgConnectionOpenTry) (*connectiontypes.MsgConnectionOpenTryResponse, error) {
	println("my ConnectionOpenTry")
	return k.ibcKeeper.ConnectionOpenTry(goCtx, msg)
}

// ConnectionOpenAck defines a rpc handler method for MsgConnectionOpenAck.
func (k Keeper) ConnectionOpenAck(goCtx context.Context, msg *connectiontypes.MsgConnectionOpenAck) (*connectiontypes.MsgConnectionOpenAckResponse, error) {
	println("my ConnectionOpenAck")
	return k.ibcKeeper.ConnectionOpenAck(goCtx, msg)
}

// ConnectionOpenConfirm defines a rpc handler method for MsgConnectionOpenConfirm.
func (k Keeper) ConnectionOpenConfirm(goCtx context.Context, msg *connectiontypes.MsgConnectionOpenConfirm) (*connectiontypes.MsgConnectionOpenConfirmResponse, error) {
	println("my ConnectionOpenConfirm")
	return k.ibcKeeper.ConnectionOpenConfirm(goCtx, msg)
}

// ChannelOpenInit defines a rpc handler method for MsgChannelOpenInit.
// ChannelOpenInit will perform 04-channel checks, route to the application
// callback, and write an OpenInit channel into state upon successful execution.
func (k Keeper) ChannelOpenInit(goCtx context.Context, msg *channeltypes.MsgChannelOpenInit) (*channeltypes.MsgChannelOpenInitResponse, error) {
	println("my ChannelOpenInit")
	return k.ibcKeeper.ChannelOpenInit(goCtx, msg)
}

// ChannelOpenTry defines a rpc handler method for MsgChannelOpenTry.
// ChannelOpenTry will perform 04-channel checks, route to the application
// callback, and write an OpenTry channel into state upon successful execution.
func (k Keeper) ChannelOpenTry(goCtx context.Context, msg *channeltypes.MsgChannelOpenTry) (*channeltypes.MsgChannelOpenTryResponse, error) {
	println("my ChannelOpenTry")
	return k.ibcKeeper.ChannelOpenTry(goCtx, msg)
}

// ChannelOpenAck defines a rpc handler method for MsgChannelOpenAck.
// ChannelOpenAck will perform 04-channel checks, route to the application
// callback, and write an OpenAck channel into state upon successful execution.
func (k Keeper) ChannelOpenAck(goCtx context.Context, msg *channeltypes.MsgChannelOpenAck) (*channeltypes.MsgChannelOpenAckResponse, error) {
	println("my ChannelOpenAck")
	return k.ibcKeeper.ChannelOpenAck(goCtx, msg)
}

// ChannelOpenConfirm defines a rpc handler method for MsgChannelOpenConfirm.
// ChannelOpenConfirm will perform 04-channel checks, route to the application
// callback, and write an OpenConfirm channel into state upon successful execution.
func (k Keeper) ChannelOpenConfirm(goCtx context.Context, msg *channeltypes.MsgChannelOpenConfirm) (*channeltypes.MsgChannelOpenConfirmResponse, error) {
	println("my ChannelOpenConfirm")
	return k.ibcKeeper.ChannelOpenConfirm(goCtx, msg)
}

// ChannelCloseInit defines a rpc handler method for MsgChannelCloseInit.
func (k Keeper) ChannelCloseInit(goCtx context.Context, msg *channeltypes.MsgChannelCloseInit) (*channeltypes.MsgChannelCloseInitResponse, error) {
	println("my ChannelCloseInit")
	return k.ibcKeeper.ChannelCloseInit(goCtx, msg)
}

// ChannelCloseConfirm defines a rpc handler method for MsgChannelCloseConfirm.
func (k Keeper) ChannelCloseConfirm(goCtx context.Context, msg *channeltypes.MsgChannelCloseConfirm) (*channeltypes.MsgChannelCloseConfirmResponse, error) {
	println("my ChannelCloseConfirm")
	return k.ibcKeeper.ChannelCloseConfirm(goCtx, msg)
}

// RecvPacket defines a rpc handler method for MsgRecvPacket.
func (k Keeper) RecvPacket(goCtx context.Context, msg *channeltypes.MsgRecvPacket) (*channeltypes.MsgRecvPacketResponse, error) {
	println("my RecvPacket")
	return k.ibcKeeper.RecvPacket(goCtx, msg)
}

// Timeout defines a rpc handler method for MsgTimeout.
func (k Keeper) Timeout(goCtx context.Context, msg *channeltypes.MsgTimeout) (*channeltypes.MsgTimeoutResponse, error) {
	println("my Timeout")
	return k.ibcKeeper.Timeout(goCtx, msg)
}

// TimeoutOnClose defines a rpc handler method for MsgTimeoutOnClose.
func (k Keeper) TimeoutOnClose(goCtx context.Context, msg *channeltypes.MsgTimeoutOnClose) (*channeltypes.MsgTimeoutOnCloseResponse, error) {
	println("my TimeoutOnClose")
	return k.ibcKeeper.TimeoutOnClose(goCtx, msg)
}

// Acknowledgement defines a rpc handler method for MsgAcknowledgement.
func (k Keeper) Acknowledgement(goCtx context.Context, msg *channeltypes.MsgAcknowledgement) (*channeltypes.MsgAcknowledgementResponse, error) {
	println("my Acknowledgement")
	return k.ibcKeeper.Acknowledgement(goCtx, msg)
}

// isRollappChain checks that the clientType is Furyint
// and that the rollapp exists
func (k Keeper) isRollappChain(
	ctx sdk.Context,
	clientType string,
	chainID string,
) (bool, error) {
	// rollapp client type is Furyint
	isFuryint := clientType == exported.Furyint
	// rollappId is the rollapp chainId
	_, isFound := k.rollappKeeper.GetRollapp(ctx, chainID)
	// if the client type isn't furyint and there is no such rollapp
	// we can be sure that the chain isn't a rollapp
	if !isFuryint && !isFound {
		return false, nil
	}
	// client type is furyint and we know this rollapp
	if isFuryint && isFound {
		return true, nil
	}
	// client type is furyint but no such rollapp
	if isFuryint && !isFound {
		return true, sdkerrors.Wrapf(types.ErrUnknownRollappID, "rollappID: %s", chainID)
	}
	// client type is not furyint but the chain is a rollapp
	return false, sdkerrors.Wrapf(types.ErrInvalidClientType, "clientType: %s", clientType)
}

// validateStateRoot load the stateInfo and verify the state was finalized
// and that the stateRoot is matching to the one which stored
func (k Keeper) validateStateRoot(ctx sdk.Context, rollappId string, height uint64, stateRoot []byte) error {
	stateInfo, err := k.rollappKeeper.FindStateInfoByHeight(ctx, rollappId, height)
	if err != nil {
		return err
	}
	if stateInfo.GetStatus() != types.STATE_STATUS_FINALIZED {
		return sdkerrors.Wrapf(types.ErrHeightStateNotFainalized, "rollappID: %s, height=%d", rollappId, height)
	}
	BlockDescriptionIndex := int(height - stateInfo.StartHeight)
	if BlockDescriptionIndex < 0 || BlockDescriptionIndex >= len(stateInfo.BDs.BD) {
		return sdkerrors.Wrapf(sdkerrors.ErrLogic,
			"searching height=%d, found stateInfo.StartHeight=%d, stateInfo.NumBlocks=%d, len(stateInfo.BDs.BD)=%d",
			height, stateInfo.StartHeight, stateInfo.NumBlocks, len(stateInfo.BDs.BD))
	}
	blockDescription := stateInfo.BDs.BD[BlockDescriptionIndex]
	if blockDescription.Height != height {
		return sdkerrors.Wrapf(sdkerrors.ErrLogic,
			"searching height=%d, found stateInfo.StartHeight=%d, stateInfo.NumBlocks=%d, len(stateInfo.BDs.BD)=%d, but BD[%d].Height=%d",
			height, stateInfo.StartHeight, stateInfo.NumBlocks, len(stateInfo.BDs.BD), BlockDescriptionIndex, blockDescription.Height)
	}
	if !bytes.Equal(stateRoot, blockDescription.StateRoot) {
		return sdkerrors.Wrapf(types.ErrInvalidAppHash, "rollappID: %s, appHash=%d", rollappId, blockDescription.StateRoot)
	}
	return nil
}
