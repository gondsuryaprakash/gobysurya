# Bit Magic 

- To check number is even or odd 
```go
   if x&1 == 0 {
      fmt.Println("Even")
   }
   if x&1 == 1 {
    fmt.Println("Odd")
   }
   // It is very fast rather than x%2 == 0 or 1
```

- To Check number is power of two or not 

```go

   if x & (x-1) ==0 {
     fmt.Println("Power of two")
   }
   /*
   4 = 100
   3 = 011
   ----------
   &   000 
   4 is power of two 
   Edge Case for x == 0 it will not work 
   */
```