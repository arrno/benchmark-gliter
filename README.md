## Benchmarking gliter

This repo is only for end-to-end benchmark testing of gliter.

[gliter repo located here.](https://github.com/arrno/gliter)

## Method

The test is to fetch N pages of 100 records from a database, do arbitrary transformations on the records, then store the results back into the database.

There are two implementations of the Database interface:

-   `Firestore` is implemented to test against a real remote datastore with real network latency
-   `MockDB` is a mock database that allows us to simulate different latency conditions

The results displayed here are from the later method (`MockDB`).

## Results

### Summarized

We start with very low simulated latency. Incrementally, we turn up latency and processed page count.

Unsurprisingly, the pipeline version always clocks in at ~50% the duration of the sequential version. This is expected because we are fetching and storing concurrently. This allows us to wait for reads while we are also waiting for writes.

### Test conditions

I'm testing on a 2021 MacBook Pro, v Monterey (12.0.1), M1 Pro chip.

```
goos: darwin
goarch: arm64
pkg: github.com/arrno/benchmark-gliter
cpu: Apple M1 Pro
```

### Measurements

**10p/20ms**

-   Page size: 100
-   Pages: 10
-   Latency: 20ms

```
% go test -v
=== RUN   TestSequential
--- PASS: TestSequential (0.44s)
=== RUN   TestPipeline
--- PASS: TestPipeline (0.25s)
PASS
ok      github.com/arrno/benchmark-gliter       1.147s
```

**100p/20ms**

-   Page size: 100
-   Pages: 100
-   Latency: 20ms

```
% go test -v
=== RUN   TestSequential
--- PASS: TestSequential (4.24s)
=== RUN   TestPipeline
--- PASS: TestPipeline (2.14s)
PASS
ok      github.com/arrno/benchmark-gliter       6.831s
```

**10p/300ms**

-   Page size: 100
-   Pages: 10
-   Latency: 300ms

```
% go test -v
=== RUN   TestSequential
--- PASS: TestSequential (6.33s)
=== RUN   TestPipeline
--- PASS: TestPipeline (3.61s)
PASS
ok      github.com/arrno/benchmark-gliter       10.377s
```

**50p/300ms**

-   Page size: 100
-   Pages: 50
-   Latency: 300ms

```
% go test -v
=== RUN   TestSequential
--- PASS: TestSequential (30.42s)
=== RUN   TestPipeline
--- PASS: TestPipeline (15.66s)
PASS
ok      github.com/arrno/benchmark-gliter       46.636s
```

**100p/500ms**

-   Page size: 100
-   Pages: 100
-   Latency: 500ms

```
% go test -v
=== RUN   TestSequential
--- PASS: TestSequential (100.73s)
=== RUN   TestPipeline
--- PASS: TestPipeline (51.11s)
PASS
ok      github.com/arrno/benchmark-gliter       152.289s

```

### Next steps

If we want to push further, we can use `InParallel` utility to further parallelize either at the pipeline level or at the fetch/store function level.

For example, this is our store func:

```go
store := func(data []DocBundle) ([]DocBundle, error) {
	sim.SetData(data)
	return nil, nil
}
```

If we break data into N chunks within this function, we could do N calls to `SetData` concurrently with `InParallel` to further drive down total latency.

This of course will work until we hit the rate limit of our datastore. [A full example of this strategy is documented here.](https://github.com/arrno/gliter/blob/main/cmd/pipeline_fan_out.go)
