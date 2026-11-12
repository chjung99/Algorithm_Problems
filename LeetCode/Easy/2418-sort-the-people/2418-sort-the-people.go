type Person struct {
    Name string
    Height int
}

func sortPeople(names []string, heights []int) []string {
    n := len(names)
    people := make([]Person, 0)

    for i := 0; i < n; i++ {
        people = append(people, Person{Name: names[i], Height: heights[i]})
    }

    slices.SortFunc(people, func(a, b Person) int {
        return -cmp.Compare(a.Height, b.Height)
    })

    sortedNames := make([]string, 0)
    for _, person := range people {
        sortedNames = append(sortedNames, person.Name)
    }

    return sortedNames
}