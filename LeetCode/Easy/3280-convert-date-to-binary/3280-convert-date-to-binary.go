func convertDateToBinary(date string) string {
    year := date[0:4]
    month := date[5:7]
    day := date[8:10]

    intYear, _ := strconv.Atoi(year)
    intMonth, _ := strconv.Atoi(month)
    intDay, _ := strconv.Atoi(day)

    var binYear string 
    var binMonth string 
    var binDay string 

    for (intYear != 0) {
        binYear = strconv.Itoa(intYear % 2) + binYear
        intYear = intYear / 2
    }

    for (intMonth != 0) {
        binMonth = strconv.Itoa(intMonth % 2) + binMonth
        intMonth = intMonth / 2
    }

    for (intDay != 0) {
        binDay = strconv.Itoa(intDay % 2) + binDay
        intDay = intDay / 2
    }
    return binYear + "-" + binMonth + "-" + binDay
}