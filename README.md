# golang-fetch-currency-sample

This repository is implementation for currency fetcher with golang


## handle with concurreny on golang

1. goroutine with waitGroup
waitGroup will make sure the execution time matched

for shared resource need to use mutex to handle race condition to limit one time one thread 

2. goroutine with channel

with channel we could shared data by communication

