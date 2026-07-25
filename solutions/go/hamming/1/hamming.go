package hamming

import "errors"

func Distance(a, b string) (int, error) {
	//panic("Implement the Distance function")
	if len(a) != len(b){
        return 0, errors.New("different len")
    }
    ans := 0
    for i:=0;i<len(a);i++ {
        if a[i] != b[i] {
            ans++
        }
    }
    return ans, nil
}
