import NonFungibleToken from 0xf8d6e0586b0a20c7
import AllDay from 0xf8d6e0586b0a20c7
// This transaction configures an account to hold AllDay NFTs.

transaction {
    prepare(signer: AuthAccount) {
        // if the account doesn't already have a collection
        if signer.borrow<&AllDay.Collection>(from: AllDay.CollectionStoragePath) == nil {

            // create a new empty collection
            let collection <- AllDay.createEmptyCollection()
            
            // save it to the account
            signer.save(<-collection, to: AllDay.CollectionStoragePath)

            // create a public capability for the collection
            signer.link<&AllDay.Collection{NonFungibleToken.CollectionPublic, AllDay.MomentNFTCollectionPublic}>(
                AllDay.CollectionPublicPath,
                target: AllDay.CollectionStoragePath
            )
        }
    }
}
