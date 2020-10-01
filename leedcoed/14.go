
func longestCommonPrefix(strs []string) string {
    if len(strs)==0{
        var aa string
        return aa
    }
    if len(strs)==1{
        //aa := strs[0]
        return strs[0]
    }
    a := 0
	flag:=0
    s := 0    
    for j := 0;j<len(strs);j++{
        if len(strs[j])==0 {
            //var aa string
            return strs[j]
        }
        if len(strs[s])>len(strs[j]){
            s = j
        }
    }
    b:= strs[s]
    for i := 0;i<len(b);i++{
        a++
		for j := 0;j<len(strs);j++{
			c := strs[j]
			if(b[i]!=c[i]){
                a--
                flag = 1
                break
			}
		}
        if flag==1{
            break
        }
        
	}
    if a<0{
        a = 0
    }
	e := strs[s]
	return e[:a]
}


 

