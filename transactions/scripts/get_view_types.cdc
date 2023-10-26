import FungibleToken from "../../contracts/FungibleToken.cdc"
import FUSD from "../../contracts/FUSD.cdc"
import MetadataViews from 0xf8d6e0586b0a20c7

pub fun main(address: Address): [Type] {
    let metadataRef = getAccount(address)
        .getCapability(/public/fusdMetadata)
        .borrow<&{MetadataViews.Resolver}>()
        ?? panic("Could not borrow a reference to the MetadataViews.Resolver")

    return metadataRef.getViews()
}
