package phonenumber

import "errors"

func Number(phoneNumber string) (string, error) {
	//panic("Please implement the Number function")
    var fmtNumber []rune
	for _, r :=range phoneNumber {
        if r>='0' && r<='9'{
        	fmtNumber = append(fmtNumber, r)
        }
    }
    if len(fmtNumber)>11 || len(fmtNumber)<10 {
        return "", errors.New("wr fmt")
    } else if len(fmtNumber) == 11{
        if fmtNumber[0] != '1' {
            return "", errors.New("wr fmt")
        }
        fmtNumber = fmtNumber[1:]
    }
    if fmtNumber[0] < '2' || fmtNumber[3] < '2' {
        return "", errors.New("wr fmt")
    }
    return string(fmtNumber), nil
}

func AreaCode(phoneNumber string) (string, error) {
	//panic("Please implement the AreaCode function")
    fmtNum, err := Number(phoneNumber)
    if err != nil {
        return "", errors.New("wr fmt")
    }
    return fmtNum[:3], nil
}

func Format(phoneNumber string) (string, error) {
	//panic("Please implement the Format function")
        fmtNum, err := Number(phoneNumber)
    if err != nil {
        return "", errors.New("wr fmt")
    }
    return "("+fmtNum[:3]+") "+fmtNum[3:6]+"-"+fmtNum[6:], nil
}
