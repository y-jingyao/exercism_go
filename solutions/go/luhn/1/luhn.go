package luhn

import "strings"

func Valid(id string) bool {
	//panic("Please implement the Valid function")
    id = strings.ReplaceAll(id," ", "")
    if len(id) <= 1 {
        return false
    }
    for _,s := range id{
        if s>'9' || s<'0' {
            return false
        }
    }
    sum := 0
    for i:=1;;i++ {
        if len(id)-i <0 {
            break
        }
        tmp := int(id[len(id)-i] -'0')
        if i % 2 == 0 {
        	tmp *= 2
        	if tmp > 9 {
            	tmp -= 9
        	}
        }
        sum += tmp
    } 
    if sum %10 == 0 {
        return true
    } else {
        return false
    }
    
}
