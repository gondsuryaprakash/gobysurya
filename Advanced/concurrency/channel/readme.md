# Channel



### Buffered Channel 

A buffered channel is a type of channel which have capacity to store data within it. 

Declaration of channel 

``` 
 bufferedChannel:= make(chan int , n)
```


 ### UnBuffered Channel 

 If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready.

 In UnBuffered Channel there is no requirement of capacity. In this channel when the message is emitted thre should be atleast a reciver 

 ```
  unBufferedChannel := make(chan int)  // No need of capacity. 
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

