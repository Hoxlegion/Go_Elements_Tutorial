package iteration

func Iter(char string, count int) string {
    var result string

    for i := range count {
        result += char
        i++
    }
    return result
}
