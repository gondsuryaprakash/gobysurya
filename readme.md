# Golang Interview Questions & Answers

> This is Golang Interview Questions from Beginer to High level 


### Table of Contents

| No. | Questions                                                                                                                                                                                                                        |
| --- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|     | [**Basic  Golang**                                                                                                                                                                                                                |
| 1   | [What is Golang?](#what-is-Golang)                                                                                                                                                                                                 |
| 2   | [What are the major features of Golang](#what-are-the-major-features-of-Golang)                                                                                                                                                   |
| 3   | [Is Golang an interpreter or compiler](#Is-Golang-an-interpreter-or-compiler)                                                                                                                                                                                                    |
| 4   | [What is Garbage Collector in Golang](#What-is-Garbage-Collector-in-Golang)                                                                                                                
| 5   | [What is runtime in Golang](#What-is-runtime-in-Golang)      
| 6   | [What is $GOPATH](#What-is-$GOPATH)        
| 7   | [What is src pkg and bin in Golang](#What-is-src-pkg-and-bin-in-Golang-?)   |
|     | **Goroutine and Channels**|
| 1   | [What is Concurrency](#What-is-Concurrency)|
| 2   | [What is difference between Concurrency and Parallelism](#What-is-difference-between-Concurrency-and-Parallelism)|
| 3   | [What is channel](#What-is-channel)|
| 4   | [What is Fan In and Fan Out Design Pattern](#What-is-Fan-In-and-Fan-Out-Design-Pattern?)|
| 5   | [What is worker pool design?](#What-is-worker-pool-design)|

 

## Basic Golang

1.  ### What is Golang?

    Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.  <div>Source : <a src="https://en.wikipedia.org/wiki/Go_(programming_language)">Wiki</a></div>

    **[Back to Top](#Table-of-Contents)**

2.  ### What are the major features of Golang?

    The major features of Golang are:
    - **Simplicity**: Golang was designed to be simple and easy to learn, making it a good language for beginners.

    - **Concurrency**: Golang provides built-in support for concurrency, which allows multiple tasks to run concurrently without the need for complex synchronization.

    - **Garbage Collection**: Golang has an efficient garbage collector that automatically frees up memory that is no longer being used.

    - **Fast Compilation**: Golang compiles very quickly, making it a good choice for large-scale projects that require frequent compilation.

    - **Static Typing:** Golang is a statically typed language, which means that variables are assigned a specific type at compile-time. This helps catch errors early on in the development process.

    - **Cross-platform Support**: Golang can be compiled to run on many different platforms, including Windows, macOS, and Linux.

    - **Standard Library**: Golang comes with a rich standard library that includes support for networking, cryptography, and many other common tasks.

    - **Error Handling**: Golang has a built-in error handling mechanism that allows for easy identification and handling of errors.

    - **Open Source**: Golang is an open-source language, which means that the source code is freely available and can be modified and distributed by anyone.

    
    **[Back to Top](#Table-of-Contents)**

3.  ### Is Golang an interpreter or compiler?
    Go, is a compiled language. This means that Go source code is compiled into machine code before it is executed, resulting in faster execution times compared to interpreted languages.
    
    **[Back to Top](#Table-of-Contents)**

4.  ### What is Garbage Collector in Golang?

    Garbage Collector (GC) in Golang is a built-in automatic memory management system that automatically frees up memory that is no longer being used by the program.

    In Golang, memory is allocated dynamically using the new and make keywords. When a program creates new objects or variables, memory is allocated on the heap. However, once these objects are no longer needed, the memory they occupy can be released back to the system for reuse.

    The Golang Garbage Collector periodically scans the program's heap to identify objects that are no longer in use and can be safely removed. The GC works by identifying and marking active objects and then reclaiming memory occupied by objects that are no longer in use. This process is transparent to the program, which does not need to explicitly manage memory deallocation.
     
    The Golang GC is designed to minimize pauses and to work efficiently in a concurrent environment, where multiple Goroutines (concurrent execution threads) may be allocating and deallocating memory simultaneously. Overall, the Garbage Collector in Golang helps to simplify memory management and reduces the risk of memory leaks or other memory-related issues in programs.   

    **[Back to Top](#Table-of-Contents)**

5. ### What is runtime in Golang ? 

    Runtime in Golang is a collection of libraries and services that provide low-level infrastructure and support for running and managing Go programs.
    The Golang Runtime includes a number of key components, such as:

    **Garbage Collector**: The Garbage Collector is responsible for automatically managing memory allocation and deallocation in the program.

    **Scheduler**: The Scheduler is responsible for managing Goroutines, which are lightweight threads of execution that allow concurrent execution of multiple tasks within a program.

    **Memory Allocator**: The Memory Allocator is responsible for allocating and deallocating memory on the heap.

    **Channel Management**: Channels are a key feature of Golang that enable communication and synchronization between Goroutines. The Runtime includes libraries and services to manage channel creation, synchronization, and communication.

    **System Interface**: The Golang Runtime provides an interface between the program and the underlying operating system, allowing programs to interact with the system's resources, such as file systems, network interfaces, and hardware devices.

    Overall, the Golang Runtime provides a robust and efficient infrastructure that allows programs to run smoothly and efficiently, while minimizing the need for manual intervention or management by the programmer.

    **[Back to Top](#Table-of-Contents)**

6. ### What is $GOPATH ? 
    
    GOPATH is an environment variable in Golang that specifies the root directory of the Go workspace. The Go workspace is a directory hierarchy that contains Go source code files, libraries, and executables.

    The GOPATH environment variable is used by the Go toolchain to locate and manage Go packages and their dependencies. When the Go compiler and related tools are invoked, they search for packages and libraries in the directories specified in GOPATH.

    **[Back to Top](#Table-of-Contents)**

7. ### What is src pkg and bin in Golang ? 
    
    **[Back to Top](#Table-of-Contents)**
          



## Goroutine and Channels

1. ### What is Concurrency? 
    Concurrency in Golang refers to the ability of a program to perform multiple tasks simultaneously, with each task running independently and concurrently. Golang has built-in features to support concurrency, including goroutines and channels.
    A goroutine is a lightweight thread managed by the Go runtime. Goroutines allow developers to write concurrent code without having to manage threads directly. A program can create many thousands of goroutines, which can all run concurrently.
    Channels provide a way for goroutines to communicate with each other and share data. A channel is a typed conduit through which values can be passed between goroutines. Channels can be used to synchronize the execution of goroutines, allowing them to coordinate their work and avoid race conditions.
    In Golang, concurrency is an important concept that allows developers to write highly efficient and scalable programs. By leveraging goroutines and channels, Golang programs can easily handle many simultaneous tasks and make the most of modern multicore processors.

   **[Back to Top](#Table-of-Contents)**

2. ### What is difference between Concurrency and Parallelism ?


    | Concurrency  |  Paralalism |
    | ------------- | ------------- |
    | Refers to the ability of a program to manage multiple tasks simultaneously, with each task running independently and concurrently.  | Refers to the ability of a program to execute multiple tasks simultaneously, using multiple processors or cores.  |
    | Allows a program to switch between different tasks as needed, and to make progress on multiple tasks simultaneously.  | Allows a program to make the most of modern hardware, by running tasks in parallel and completing them more quickly. |
    |Does not necessarily require multiple processors or cores; it can be achieved on a single processor through time-slicing or other techniques. |Requires multiple processors or cores to be effective.|
    |Can be used to create highly efficient and scalable programs.|Can be used to make the most of modern hardware and achieve high levels of performance.| 

    **[Back to Top](#Table-of-Contents)**

3. ### What is channel ? 
    A channel is a typed conduit through which values can be passed between `goroutines`. A channel provides a way for goroutines to communicate and synchronize their execution, allowing them to coordinate their work and share data.

    Channels can be created using the `make` function, with a type that specifies the type of values that can be passed through the channel. For example, the following code creates a channel of integers:
 
    ```go
       ch:= make(chan int)
    ```
    There are two types of channel - 
    #### Buffered Channel 
    - A buffered channel is asynchronous, meaning that the sender and receiver goroutines are not synchronized and send operations will not block as long as there is space in the buffer, but block once the buffer is full.
    - Buffered channels are typically used for many-to-one or many-to-many communication between goroutines, where multiple goroutines can send data to a single receiver without blocking.
    - Here's an example of using a buffered channel to implement a simple message queue:

    ```go 
    func main() {
    queue := make(chan string, 3) // buffer size of 3

    // Start a goroutine to process messages from the queue
    go func() {
        for {
            msg := <-queue // receive a message from the queue
            fmt.Println("Processing message:", msg)
            time.Sleep(time.Second) // simulate processing time
        }
    }()

    // Send some messages to the queue
    queue <- "Message 1"
    queue <- "Message 2"
    queue <- "Message 3"
    queue <- "Message 4" // blocks until a message is processed and space is freed up

    fmt.Println("All messages sent")
    }  
    ```


     #### UnBuffered Channel
    - An unbuffered channel is synchronous, meaning that the sender and receiver goroutines are synchronized and each send operation will block until there is a corresponding receive operation, and vice versa.
    - Unbuffered channels are typically used for one-to-one communication between goroutines, where one goroutine needs to wait for another goroutine to complete a specific task before continuing.
    - Here's an example of using an unbuffered channel for communication between two goroutines:

    ```go

    func worker(id int, jobs <-chan int, results chan<- int) {
        for j := range jobs {
            fmt.Printf("Worker %d started job %d\n", id, j)
            time.Sleep(time.Second) // simulate doing work
            fmt.Printf("Worker %d finished job %d\n", id, j)
            results <- j * 2
            } 
        }

    func main() {
        jobs := make(chan int)
        results := make(chan int)

        // Start some worker goroutines
        for w := 1; w <= 3; w++ {
            go worker(w, jobs, results)
        }

        // Send some jobs to the workers and collect the results
        for j := 1; j <= 5; j++ {
            jobs <- j
            result := <-results
            fmt.Printf("Result for job %d: %d\n", j, result)
        }
    }

    ```
  
  
   
4. ### What is Fan In and Fan Out Design Pattern? 

    #### Fan-in
    
    - `Fan-In` is a pattern where multiple goroutines send data to a single channel, which collects and combines the   data from all the goroutines.
    - This pattern is useful when you have multiple goroutines producing similar results that need to be processed together, such as reading data from multiple files or querying multiple APIs for data.
    - Here's an example of using the "Fan-In" pattern to read data from multiple files: 

    ```go
    func readFile(filename string, out chan<- string) {
        data, err := ioutil.ReadFile(filename)
            if err != nil {
                log.Fatal(err)
            }
        out <- string(data)
    }

    func main() {
        files := []string{"file1.txt", "file2.txt", "file3.txt"}

        // Create a channel to collect the data from the goroutines
        out := make(chan string)
        // Start a goroutine for each file to read the data and send it to the output channel
        for _, file := range files {
            go readFile(file, out)
        }

        // Collect the data from the output channel
        for i := 0; i < len(files); i++ {
            data := <-out
            fmt.Println(data)
        }
    }

    ```

    In this example, the readFile function is a worker function that reads the contents of a file and sends it to the out channel. The main function creates an output channel and starts a goroutine for each file to read the data and send it to the output channel. The main function then collects the data from the output channel and prints it to the console.

    #### Fan-Out

    - `Fan-Out` is a pattern where a single goroutine sends work to multiple worker goroutines, which each process the work and send the results to a single channel that collects the results.
    - This pattern is useful when you have a large amount of work that needs to be processed in parallel, such as processing data in batches or querying a large dataset.
    - Here's an example of using the "Fan-Out" pattern to process data in parallel:
    
    ```go
    func process(data []int, out chan<- int) {
        for _, num := range data {
            out <- num * 2
       }
    }

    func main() {
        data := []int{1, 2, 3, 4, 5}

        // Create an output channel to collect the results from the workers
        out := make(chan int)

        // Start some worker goroutines to process the data and send the results to the output channel
        for i := 0; i < 3; i++ {
            go process(data, out)
        }

        // Collect the results from the output channel
        results := make([]int, 0, len(data)*3)
        for i := 0; i < len(data)*3; i++ {
            result := <-out
            results = append(results, result)
        }

        fmt.Println(results)
    }



    ```
    
    In this example, the process function is a worker function that processes a slice of integers and sends the results to the out channel. The main function creates an output channel and starts multiple worker goroutines to process the data and send the results to the output channel. The main function then collects the results from the output channel and stores them in a slice. The results are then printed to the console.
   
5. ### What is worker pool design ? 
    The worker pool pattern is a popular concurrency design pattern used in Go (Golang) for efficient and concurrent processing of tasks. It involves creating a pool of goroutines (workers) that can pick up tasks from a queue and process them concurrently.

    Here are some key points of the worker pool pattern in Go:

    - A fixed number of worker goroutines are created upfront.
    - Each worker goroutine waits for tasks to be submitted to a task channel.
    - When a task is submitted, an idle worker picks it up and processes it.
    - Once a task is completed, the worker returns to the idle state and waits for another task.
    - The task channel can be buffered to allow for a certain number of tasks to be queued up before blocking.
    - The number of workers and buffer size can be tuned to optimize performance for the specific workload.
    
    Here is an example implementation of the worker pool pattern in Go:

    ```go

    package main

    import (
        "fmt"
        "sync"
    )

    func worker(id int, jobs <-chan int, results chan<- int) {
        for j := range jobs {
            fmt.Println("worker", id, "started  job", j)
            results <- j * 2
            fmt.Println("worker", id, "finished job", j)
        }
    }
 
    func main() {
        const numJobs = 10
        const numWorkers = 3

        jobs := make(chan int, numJobs)
        results := make(chan int, numJobs)

        // Spawn a fixed number of worker goroutines.
        var wg sync.WaitGroup
        for i := 1; i <= numWorkers; i++ {
            wg.Add(1)
            go func(id int) {
                defer wg.Done()
                worker(id, jobs, results)
        }(i)
        }

        // Submit tasks to the task channel.
        for j := 1; j <= numJobs; j++ {
            jobs <- j
        }
        close(jobs)

        // Collect the results from the result channel.
        for a := 1; a <= numJobs; a++ {
            <-results
        }

        // Wait for all workers to complete.
        wg.Wait()
    }

    ```
