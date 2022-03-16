func fib(n int) int {
    if n <=1{
        return n
    }
    res := make([]int , n+1)
    res[0], res[1] = 0, 1
    for i:=2;i<=n;i++{
        res[i] = res[i-1] + res[i-2]
    }
    return res[n]
}