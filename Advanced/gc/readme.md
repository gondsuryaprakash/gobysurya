# Garbage Collector in Golang

### The basic principles to reduce GC pressure
 -  allocate less short-lived memory blocks on heap (garbage production speed control).
 -  create less pointers (scan work load control).
 