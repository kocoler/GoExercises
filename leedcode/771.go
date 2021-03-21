package main

func numJewelsInStones(J string, S string) int {
    if len(J) == 0 || len (S) == 0 {
        return 0
    }
    a := make(map[byte]int)
    for i := 0 ; i < len(J) ; i++ {
        a[J[i]] = 1
    }
    sum := 0
    for j := 0 ; j < len(S) ; j++ {
        if a[S[j]] == 1 {
            sum ++
        }
    }
    return sum
}
