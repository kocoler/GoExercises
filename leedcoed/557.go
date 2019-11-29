func reverseWords(s string) string {
    var a []byte
    var b []byte
    for i := 0 ; i < len(s) ; i++ {
        if s[i] == 32  {
            for j := len(b)-1 ; j >= 0 ; j-- {
                a = append(a,b[j])
            }
            a = append(a,' ')
            var c []byte
            b = c
        }else {
            b = append(b,s[i])
        }
        if i+1 == len(s) {
            for j := len(b)-1 ; j >= 0 ; j-- {
                a = append(a,b[j])
            }
        }
    }
    return string(a)
}
