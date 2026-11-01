func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    str1 := ""
    str2 := ""

    for _, word := range word1 {
        str1 += word
    }

    for _, word := range word2 {
        str2 += word
    }
    
    return str1 == str2
}