#!/bin/sh

# Set parameters

CHAIN_ID=${CHAIN_ID:-"local-testnet"}
MONIKER_NAME=${MONIKER_NAME:-"local"}
KEY_NAME=${KEY_NAME:-"local-user"}

NODE_AMOUNT=${NODE_AMOUNT:-"20000000ufury"}
NODE_STAKING_AMOUNT=${NODE_STAKING_AMOUNT:-"10000000ufury"}

query_balance() {
 NODE_ACCOUNT="$(furyd keys show -a "$KEY_NAME" --keyring-backend test)"
    echo "Current balance of the full node account on chain[$NODE_ACCOUNT]: "
    furyd q bank balances "$NODE_ACCOUNT"

    echo "Make sure the sequencer account [$NODE_ACCOUNT] is funded"
    echo "From within the hub node: \"furyd tx bank send $KEY_NAME $NODE_ACCOUNT $NODE_AMOUNT --keyring-backend test\""
    read -r -p "Press to continue..."
}

create_validator() {
    echo "Current balance of the full node account on chain[$NODE_ACCOUNT]: "
    furyd q bank balances "$NODE_ACCOUNT"

    echo `# ------------------- Running create-validator transaction ------------------- #`
    furyd tx staking create-validator \
        --amount "$NODE_STAKING_AMOUNT" \
        --commission-max-change-rate "0.1" \
        --commission-max-rate "0.20" \
        --commission-rate "0.1" \
        --min-self-delegation "1" \
        --details "validators write bios too" \
        --pubkey=$(furyd tendermint show-validator) \
        --moniker "2ndmoniker" \
        --chain-id "$CHAIN_ID" \
        --gas-prices 0.025ufury \
        --from "$KEY_NAME" \
        --keyring-backend test
}


query_balance
create_validator