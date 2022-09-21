import AllDay from 0xf8d6e0586b0a20c7
// This script returns all the Edition structs.
// This will be *long*.

pub fun main(query: String): [UInt64] {
    let editionIDs: [UInt64] = []
    var id: UInt64 = 1

    while id < AllDay.nextEditionID {
		var edition: AllDay.EditionData = AllDay.getEditionData(id: id)

		if edition.tier == query {
			editionIDs.append(edition.id)
		}
        id = id + 1
    }

    return editionIDs 
}

