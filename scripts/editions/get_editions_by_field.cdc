import AllDay from 0xf8d6e0586b0a20c7

pub fun main(): [UInt64] {
    let editionIDs: [UInt64] = []
    var id: UInt64 = 1

    while id < AllDay.nextEditionID {
		var edition: AllDay.EditionData = AllDay.getEditionData(id: id)

		if edition.tier == "Legendary" {
			editionIDs.append(edition.id)
		}
        id = id + 1
    }

    return editionIDs 
}

