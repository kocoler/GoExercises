package main

func myAtoi(str string) int {
    a := 0
    flag := 0
    i:=0
    if len(str)==0{
        return a
    }else{
        for str[i]==' ' {
            i++
            if i>=len(str){
            return a
        }
        }
        switch str[i]{
            case '-': i++;flag = 1
            case '+': i++
        }
         /*if str[i]=='-'{ 
            flag = 1
             i++
        }
        if str[i]=='+'&&flag==0 {
            i++
        }*/
        for i<len(str) {
            if str[i]>'9'||str[i]<'0'{
                break
            }
            a = a*10+(int(str[i])-48)
            if a>=2147483648 && flag==1 {
                a = 2147483648
                break
            }
            if a>=2147483648 && flag==0{
                a = 2147483647
                break
            }
            i++
    }
        if flag==1{
            a = -a
        }
    }
    return a
}
