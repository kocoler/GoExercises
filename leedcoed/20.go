func isValid(s string) bool {
    if len(s)==1 {
        return false
    }
    if len(s)==0 {
        return true
    }
    a := len(s)/2
    var b [3] int
    var c [3] int
    var e [] string
    d := 0
    for i := 0;i<len(s);i++ {
        switch {
            case s[i]=='(': b[0]++;e = append(e,"(");d++
            case s[i]=='[': b[1]++;e = append(e,"[");d++
            case s[i]=='{': b[2]++;e = append(e,"{");d++
            case s[i]==')': c[0]++;if c[0]>b[0] {return false};if e[d-1]=="(" {e = e[0:d-1];d--}else {return false}
            case s[i]==']': c[1]++;if c[1]>b[1] {return false};if e[d-1]=="[" {e = e[0:d-1];d--}else {return false}
            case s[i]=='}': c[2]++;if c[2]>b[2] {return false};if e[d-1]=="{" {e = e[0:d-1];d--}else {return false}
        }
        if b[0]+b[1]+b[2]>a{
            return false
        }
    }
    return true
}
