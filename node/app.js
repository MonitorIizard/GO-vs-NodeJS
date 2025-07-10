const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => {
  let startTime = performance.now();

  let sum = 0;

  for(let i = 0; i < 1_000_000_000; i++) {
      sum += i;
  }

  let message = "Node Benchmark: " + (performance.now() - startTime) + "ms";
  message += "\nSum: " + sum;

  res.send(message);
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})