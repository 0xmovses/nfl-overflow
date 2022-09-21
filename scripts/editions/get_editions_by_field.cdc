import AllDay from 0xf8d6e0586b0a20c7
// This script returns all the Edition structs.
// This will be *long*.

pub fun main(query: String): [AllDay.EditionData] {
    let editions: [AllDay.EditionData] = []
    var id: UInt64 = 1

    while id < AllDay.nextEditionID {
		var edition: AllDay.EditionData = AllDay.getEditionData(id: id)

		if edition.tier == query {
			editions.append(edition)
		}
        id = id + 1
    }
    return editions
}

