func reverse(x int) int {
    a := 0
    for x!=0{
        a=a*10+x%10
        x=x/10
    }
    if a>2147483647 || a<(-2147483648) {
        return 0
    }
    return a
}
