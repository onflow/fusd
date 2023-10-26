# DEPRECATED (DO NOT USE IT)

This repository exists only for historical purpose 

Flow deprecated use of FUSD in favour or USDC

Check [https://www.circle.com/en/usdc](https://www.circle.com/en/usdc) for more info.

# Flow USD (FUSD)

Flow USD (FUSD) is a stablecoin on Flow, issued by [Prime Trust](https://www.primetrust.com/), that is backed 1:1 by the US Dollar on Flow Mainnet.

This repository contains the Cadence source code for the FUSD contract and accompanying transactions.

You can read more about FUSD in the [FUSD documentation](https://developers.flow.com/flow/fusd).

## Contract Addresses

|Name|Testnet|Mainnet|Sandboxnet|
|----|-------|-------|----------|
|`FUSD`|`0xe223d8a629e49c68`|`0x3c5959b568896393`|`0x6c52cbc80f034d1b`|
|`FungibleToken`|`0x9a0766d93b6608b7`|`0xf233dcee88fe0abe`|`0xe20612a0776ca4bf`|

## Local Development & Setup

If you want to make changes on the `FUSD` contract and test it on emulator, type the following command on a terminal:

```bash
flow emulator --storage-limit=false --contracts -v
```

Let's create the 3 accounts listed on `flow.json` and update the `address` and `key` key/value pairs. (`fusd-account`, `minter-account`, `recipient-account`).

```bash
flow keys generate --output=json
# Copy the public key and paste it below
flow accounts create --key $public_key
```

Do the same for all 3 accounts. The generated addresses will be the same, only the `private` key will be different and needs to be substituted. Let's add some variables to help us with running scripts & transactions.

```bash
export FUSD_ACCOUNT=0x01cf0e2f2f715450

export MINTER_ACCOUNT=0x179b6b1cb6755e31

export RECIPIENT_ACCOUNT=0xf3fcd2c1a78f5eee
```

Now we can deploy `FUSD` on emulator, with:

```bash
flow project deploy --network=emulator
```

To setup a `FUSD.Vault` resource for any account, run:

```bash
flow transactions send ./transactions/setup_fusd_vault.cdc --network=emulator --signer=fusd-account

flow transactions send ./transactions/setup_fusd_vault.cdc --network=emulator --signer=minter-account

flow transactions send ./transactions/setup_fusd_vault.cdc --network=emulator --signer=recipient-account
```

To check if an account is properly setup, run:

```bash
flow scripts execute ./transactions/scripts/check_fusd_vault_setup.cdc $FUSD_ACCOUNT --network=emulator
```

We can mint some `FUSD` tokens and send it to the `fusd-account`:

```bash
flow transactions send ./transactions/minter/setup_fusd_minter.cdc --network=emulator --signer=minter-account

flow transactions send ./transactions/admin/deposit_fusd_minter.cdc $MINTER_ACCOUNT --network=emulator --signer=fusd-account

flow transactions send ./transactions/minter/mint_fusd.cdc 15000.0 $FUSD_ACCOUNT --network=emulator --signer=minter-account
```

With the following scripts, we can observe the effects of the above transactions:

```bash
flow scripts execute ./transactions/scripts/get_fusd_supply.cdc --network=emulator

# => Output:
Result: 15000.00000000

flow scripts execute ./transactions/scripts/get_fusd_balance.cdc $FUSD_ACCOUNT --network=emulator

# => Output:
Result: 15000.00000000
```

Transferring `FUSD` between accounts, can be achieved with:

```bash
flow transactions send ./transactions/transfer_fusd.cdc 1000.0 $RECIPIENT_ACCOUNT --network=emulator --signer=fusd-account

flow scripts execute ./transactions/scripts/get_fusd_balance.cdc $RECIPIENT_ACCOUNT --network=emulator

# => Output:
Result: 1000.00000000
```

To make use of the `FungibleTokenMetadataViews` contract, we can run the following script:

```bash
flow scripts execute ./transactions/scripts/get_view_types.cdc $RECIPIENT_ACCOUNT --network=emulator
```

The output will look something like this:

```cadence
Result: [
    Type<A.01cf0e2f2f715450.FungibleTokenMetadataViews.FTView>(),
    Type<A.01cf0e2f2f715450.FungibleTokenMetadataViews.FTDisplay>(),
    Type<A.01cf0e2f2f715450.FungibleTokenMetadataViews.FTVaultData>()
]
```

We can get the `FungibleTokenMetadataViews.FTDisplay` View, with:

```bash
flow scripts execute ./transactions/scripts/get_display_view.cdc $RECIPIENT_ACCOUNT --network=emulator
```

The output will look something like this:

```cadence
Result: A.01cf0e2f2f715450.FungibleTokenMetadataViews.FTDisplay(
    name: "Flow USD Fungible Token",
    symbol: "FUSD",
    description: "Flow USD (FUSD) is a stablecoin on Flow, issued by Prime Trust, that is backed 1:1 by the US Dollar on Flow Mainnet.",
    externalURL: A.f8d6e0586b0a20c7.MetadataViews.ExternalURL(
        url: "https://developers.flow.com/flow/fusd"
    ),
    logos: A.f8d6e0586b0a20c7.MetadataViews.Medias(
        items: [
            A.f8d6e0586b0a20c7.MetadataViews.Media(file: A.f8d6e0586b0a20c7.MetadataViews.HTTPFile(url: "https://global-uploads.webflow.com/60f008ba9757da0940af288e/6141ac940d6edec559d7f959_aAZ9V3yL_400x400.jpeg"), mediaType: "image/jpeg")
        ]
    ),
    socials: {
        "twitter": A.f8d6e0586b0a20c7.MetadataViews.ExternalURL(
            url: "https://twitter.com/flow_blockchain"
        )
    }
)
```
