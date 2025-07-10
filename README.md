inspiration: [Go vs Node Benchmark](https://itnext.io/performance-benchmark-node-js-vs-go-9dbad158c3b0)

my blog: [The difference between compile time and run time — Why GO is faster than NodeJS](https://medium.com/@mokmaard646/the-difference-between-compile-time-and-run-time-why-go-is-faster-than-node-f8ae86a7e045)


First, navigate to the node directory, install the dependencies and run node server:

```
cd node
npm i
node server.js
```

Then, in another terminal, navigate to the go directory and run the Go server:

```
cd go
go get github.com/gofiber/fiber/v2
go run main.go
```

Finally, run the test script to make requests to both servers:

```
bash test-curl.sh
```

results in the terminal will show the responses from both servers.
You should see output similar to:

```
Running Go request...
GO Benchmark: s ms
Sum: xxx

Running Node request...
Node Benchmark: s ms
Sum: xxx
```