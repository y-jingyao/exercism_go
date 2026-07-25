package differenceofsquares

func SquareOfSum(n int) int {
	//panic("Please implement the SquareOfSum function")
    ans := 0
    for i:=0;i<=n;i++ {
        ans += i
    }
    return ans*ans
}

func SumOfSquares(n int) int {
	//panic("Please implement the SumOfSquares function")
    ans := 0
    for i:=0;i<=n;i++ {
        ans += i*i
    }
    return ans
}

func Difference(n int) int {
	//panic("Please implement the Difference function")
    return SquareOfSum(n) - SumOfSquares(n)
}
