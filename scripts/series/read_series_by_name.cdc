import AllDay from 0xf8d6e0586b0a20c7 
// This script returns a Series struct for the given name,
// if it exists

pub fun main(seriesName: String): AllDay.SeriesData {
    return AllDay.getSeriesDataByName(name: seriesName)
}

