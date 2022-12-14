import FUSD from "../../contracts/FUSD.cdc"
import FungibleTokenMetadataViews from "../../contracts/FungibleTokenMetadataViews.cdc"
import MetadataViews from 0xf8d6e0586b0a20c7

pub fun main(address: Address): FungibleTokenMetadataViews.FTDisplay? {
    let metadataRef = getAccount(address)
        .getCapability(/public/fusdMetadata)
        .borrow<&{MetadataViews.Resolver}>()
        ?? panic("Could not borrow a reference to the MetadataViews.Resolver")

    let displayView = FungibleTokenMetadataViews.getFTDisplay(
        metadataRef
    )

    return displayView
}