func defangIPaddr(address string) string {
    if len(address) == 0 {
        return ""
    }
    var a []byte
    for i := 0 ; i < len(address) ; i++ {
        if address[i] == '.' {
            a = append(a,'[')
            a = append(a,'.')
            a = append(a,']')
        }else {
            a = append(a,address[i])
    }
    }
    var c = string(a)
    return c
}
