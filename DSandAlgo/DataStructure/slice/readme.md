# Slice 

### Slice Passing as an argument : 

 1. Case 1 : Change in outer function 

 ```
 func change(abc []int) {

    for i:= range abc {
        abc[i] = 5
    }
    fmt.Println(abc)
 }

 func main() {
    abc:= []int{1,2,3}
    change(abc)

    fmt.Println(abc)
 }

 // 

 ```

 