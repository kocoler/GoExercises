func isPalindrome(x int) bool {
    if(x<0){
        return false
    }
    a := x
    num:=0
    if(x>0){    
        for x>0{
            num=num*10+x%10
            x=x/10
        }
    }
    return a==num
}
