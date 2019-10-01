package main

import "fmt"

func quicksort(arr [] int,l int,r int){
     if l>=r {
        return
  }
     tmp:=l
     for i:=l+1;i<=r;i++{
      if arr[l]>=arr[i]{
       tmp++;
       arr[i],arr[tmp]=arr[tmp],arr[i]
       }
     }
   arr[l],arr[tmp]=arr[tmp],arr[l]
   quicksort(arr,l,tmp-1)
   quicksort(arr,tmp+1,r) 
        
}

func main() {
    a:=[] int {2,1,3,4,6,2}
    quicksort(a,0,len(a)-1)
    fmt.Printf("%d",a)
}

