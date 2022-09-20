import AllDay from 0xf8d6e0586b0a20c7

transaction(
    name: String,
    metadata: {String: String}
   ) {
    // local variable for the admin reference
    let admin: &AllDay.Admin

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        self.admin = signer.borrow<&AllDay.Admin>(from: AllDay.AdminStoragePath)
            ?? panic("Could not borrow a reference to the AllDay Admin capability")
    }

    execute {
        let id = self.admin.createPlay(
            classification: name,
            metadata: metadata
        )

        log("====================================")
        log("New Play:")
        log("PlayID: ".concat(id.toString()))
        log("====================================")
    }
}

