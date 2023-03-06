# Golang Interview Questions & Answers

> This is Golang Interview Questions from Beginer to High level 


### Table of Contents

| No. | Questions                                                                                                                                                                                                                        |
| --- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|     | **Basic  Golang**                                                                                                                                                                                                                   |
| 1   | [What is Golang?](#what-is-Golang)                                                                                                                                                                                                 |
| 2   | [What are the major features of Golang?](#what-are-the-major-features-of-Golang)                                                                                                                                                   |
| 3   | [Is Golang an interpreter or compiler?](#Is-Golang-an-interpreter-or-compiler)                                                                                                                                                                                                    |
| 4   | [What is Garbage Collector in Golang?](#What-is-Garbage-Collector-in-Golang)                                                                                                                
| 5   | [What is runtime in Golang?](#What-is-runtime-in-Golang)      
| 6   | [What is $GOPATH? ](#What-is-$GOPATH)        
| 7   | [What is src pkg and bin in Golang ?](#What-is-src-pkg-and-bin-in-Golang-?)   |
|     |              |
|     | **Goroutine and Channels**|
| 1   | [What is Concurrency?](#What-is-Concurrency)

 

## Basic Golang

1.  ### What is Golang?

    Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.  <div>Source : <a src="https://en.wikipedia.org/wiki/Go_(programming_language)">Wiki</a></div>

    **[Back to Top](#Basic-Golang)**

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

    
    **[Back to Top](#Basic-Golang)**

3.  ### Is Golang an interpreter or compiler?
    Go, is a compiled language. This means that Go source code is compiled into machine code before it is executed, resulting in faster execution times compared to interpreted languages.
    
    **[Back to Top](#Basic-Golang)**

4.  ### What is Garbage Collector in Golang?

    Garbage Collector (GC) in Golang is a built-in automatic memory management system that automatically frees up memory that is no longer being used by the program.

    In Golang, memory is allocated dynamically using the new and make keywords. When a program creates new objects or variables, memory is allocated on the heap. However, once these objects are no longer needed, the memory they occupy can be released back to the system for reuse.

    The Golang Garbage Collector periodically scans the program's heap to identify objects that are no longer in use and can be safely removed. The GC works by identifying and marking active objects and then reclaiming memory occupied by objects that are no longer in use. This process is transparent to the program, which does not need to explicitly manage memory deallocation.
     
    The Golang GC is designed to minimize pauses and to work efficiently in a concurrent environment, where multiple Goroutines (concurrent execution threads) may be allocating and deallocating memory simultaneously. Overall, the Garbage Collector in Golang helps to simplify memory management and reduces the risk of memory leaks or other memory-related issues in programs.   

    **[Back to Top](#Basic-Golang)**

5. ### What is runtime in Golang ? 

    Runtime in Golang is a collection of libraries and services that provide low-level infrastructure and support for running and managing Go programs.
    The Golang Runtime includes a number of key components, such as:

    **Garbage Collector**: The Garbage Collector is responsible for automatically managing memory allocation and deallocation in the program.

    **Scheduler**: The Scheduler is responsible for managing Goroutines, which are lightweight threads of execution that allow concurrent execution of multiple tasks within a program.

    **Memory Allocator**: The Memory Allocator is responsible for allocating and deallocating memory on the heap.

    **Channel Management**: Channels are a key feature of Golang that enable communication and synchronization between Goroutines. The Runtime includes libraries and services to manage channel creation, synchronization, and communication.

    **System Interface**: The Golang Runtime provides an interface between the program and the underlying operating system, allowing programs to interact with the system's resources, such as file systems, network interfaces, and hardware devices.

    Overall, the Golang Runtime provides a robust and efficient infrastructure that allows programs to run smoothly and efficiently, while minimizing the need for manual intervention or management by the programmer.

    **[Back to Top](#Basic-Golang)**

6. ### What is $GOPATH ? 
    
    GOPATH is an environment variable in Golang that specifies the root directory of the Go workspace. The Go workspace is a directory hierarchy that contains Go source code files, libraries, and executables.

    The GOPATH environment variable is used by the Go toolchain to locate and manage Go packages and their dependencies. When the Go compiler and related tools are invoked, they search for packages and libraries in the directories specified in GOPATH.

    **[Back to Top](#Basic-Golang)**

7. ### What is src pkg and bin in Golang ? 
    
    **[Back to Top](#Basic-Golang)**
          



## Goroutine and Channels

1. ### What is Concurrency? 
    Concurrency in Golang refers to the ability of a program to perform multiple tasks simultaneously, with each task running independently and concurrently. Golang has built-in features to support concurrency, including goroutines and channels.
    A goroutine is a lightweight thread managed by the Go runtime. Goroutines allow developers to write concurrent code without having to manage threads directly. A program can create many thousands of goroutines, which can all run concurrently.
    Channels provide a way for goroutines to communicate with each other and share data. A channel is a typed conduit through which values can be passed between goroutines. Channels can be used to synchronize the execution of goroutines, allowing them to coordinate their work and avoid race conditions.
    In Golang, concurrency is an important concept that allows developers to write highly efficient and scalable programs. By leveraging goroutines and channels, Golang programs can easily handle many simultaneous tasks and make the most of modern multicore processors.

    **[Back to Top](#Goroutine-andChannels)**


