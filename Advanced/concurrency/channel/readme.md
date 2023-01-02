# Channel



### Buffered Channel 

A buffered channel is a type of channel which have capacity to store data within it. 

Declaration of channel 

``` 
 bufferedChannel:= make(chan int , n)
```

### Directional Channel 

Send and Recive data through channel

 **Sends Data Only** 

```
  out chan<- int 
```

**Recieve Channel**

```
  in <-chan int
```

