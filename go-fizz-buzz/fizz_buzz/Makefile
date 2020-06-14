test:
	go test
bench:
	go test --bench=.

## You'll need graphviz for this step `apt-get install graphviz`
bench-graph:
	go test --bench=. -cpuprofile cpu.prof
	go tool pprof -svg cpu.prof > cpu.svg


