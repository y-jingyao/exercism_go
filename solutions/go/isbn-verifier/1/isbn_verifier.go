package isbnverifier

import "strings"

func IsValidISBN(isbn string) bool {
	//panic("Please implement the IsValidISBN function")
    risbn := strings.ReplaceAll(isbn, "-", "")
    if len(risbn) != 10 {
        return false
    }
    ans := 0
    for i, c := range risbn {
        if i != 9 && c =='X' {
            return false
        }
        var tmp int
        if c == 'X' {
            tmp = 10
        } else {
        tmp = int(c - '0')
        }
        ans += tmp * (10-i)
    }
    if ans % 11 ==0 {
        return true
    } else {
        return false
    }
}
