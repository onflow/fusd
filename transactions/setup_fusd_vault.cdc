// This transaction configures the signer's account with an empty FUSD vault.
//
// It also links the following capabilities:
//
// - FungibleToken.Receiver: this capability allows this account to accept FUSD deposits.
// - FungibleToken.Balance: this capability allows anybody to inspect the FUSD balance of this account.
// - MetadataViews.Resolver: this capability allows anybody to inspect the views supported by the FUSD vault of this account.
// - FungibleToken.Provider: this capability allows the account to make withdrawals from the FUSD vault.

import FungibleToken from "../contracts/FungibleToken.cdc"
import FUSD from "../contracts/FUSD.cdc"
import MetadataViews from 0xf8d6e0586b0a20c7

transaction {

    prepare(signer: AuthAccount) {

        // It's OK if the account already has a Vault, but we don't want to replace it
        if(signer.borrow<&FUSD.Vault>(from: /storage/fusdVault) != nil) {
            return
        }

        // Create a new FUSD Vault and put it in storage
        signer.save(<-FUSD.createEmptyVault(), to: /storage/fusdVault)

        // Create a public capability to the Vault that only exposes
        // the deposit function through the Receiver interface
        signer.link<&FUSD.Vault{FungibleToken.Receiver}>(
            /public/fusdReceiver,
            target: /storage/fusdVault
        )

        // Create a public capability to the Vault that only exposes
        // the balance field through the Balance interface.
        // NOTE: We leave this as is, for backwards compatibility.
        signer.link<&FUSD.Vault{FungibleToken.Balance}>(
            /public/fusdBalance,
            target: /storage/fusdVault
        )

        // Create a public capability to the Vault that only exposes
        // the balance field through the Balance interface and the
        // getViews(), resolveView() functions from MetadataViews.Resolver
        // interface
        signer.link<&FUSD.Vault{FungibleToken.Balance, MetadataViews.Resolver}>(
            /public/fusdMetadata,
            target: /storage/fusdVault
        )

        // Create a private capability to the Vault that only exposes
        // the withdraw function through the Provider interface
        signer.link<&FUSD.Vault{FungibleToken.Provider}>(
            /private/fusdProvider,
            target: /storage/fusdVault
        )
    }
}
