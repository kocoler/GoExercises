package main

func romanToInt(s string) int {
    num:=0
	for i:=len(s)-1;i>=0;i--{
		switch {
		case s[i]=='I': num++
		case s[i]=='V':if i>=1&&s[i-1]=='I'{num+=4;i--}else {num+=5}
		case s[i]=='X':if i>=1&&s[i-1]=='I'{num+=9;i--}else {num+=10}
		case s[i]=='L':if i>=1&&s[i-1]=='X'{num+=40;i--}else {num+=50}
		case s[i]=='C':if i>=1&&s[i-1]=='X'{num+=90;i--}else {num+=100}
		case s[i]=='D':if i>=1&&s[i-1]=='C'{num+=400;i--}else {num+=500}
		case s[i]=='M':if i>=1&&s[i-1]=='C'{num+=900;i--}else {num+=1000}
		//default: continue
		}
	}
	return num
}
