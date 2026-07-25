package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	//panic("Please implement the CollatzConjecture function")
    ans := 0
    for {
        if n <=0 {
            return 0,errors.New("n <=0 ")
        }
        if n == 1{
            return ans, nil
        } else if n%2 == 0{
            n = n/2
            ans++
        } else {
            n = n * 3 + 1
            ans++
        }
    }
}
