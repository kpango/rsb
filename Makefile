GO_VERSION:=$(shell go version)

.PHONY: all clean bench bench-all profile lint test contributors update install

all: clean install lint test bench

clean:
	rm -rf ./*.log
	rm -rf ./*.svg
	rm -rf ./go.mod
	rm -rf ./go.sum
	rm -rf bench
	rm -rf pprof
	rm -rf vendor


bench: clean init
	go test -count=1 -run=NONE -bench . -benchmem

init:
	GO111MODULE=on go mod init
	GO111MODULE=on go mod vendor
	sleep 3

profile: clean init
	rm -rf bench
	mkdir bench
	mkdir pprof
	\
	go test -count=10 -run=NONE -bench=BenchmarkKnuth -benchmem -o pprof/knuth-test.bin -cpuprofile pprof/cpu-knuth.out -memprofile pprof/mem-knuth.out
	go tool pprof --svg pprof/knuth-test.bin pprof/cpu-knuth.out > cpu-knuth.svg
	go tool pprof --svg pprof/knuth-test.bin pprof/mem-knuth.out > mem-knuth.svg
	go-torch -f bench/cpu-knuth-graph.svg pprof/knuth-test.bin pprof/cpu-knuth.out
	go-torch --alloc_objects -f bench/mem-knuth-graph.svg pprof/knuth-test.bin pprof/mem-knuth.out
	\
	go test -count=10 -run=NONE -bench=BenchmarkShufflePartial -benchmem -o pprof/shuffule-partial-test.bin -cpuprofile pprof/cpu-shuffule-partial.out -memprofile pprof/mem-shuffule-partial.out
	go tool pprof --svg pprof/shuffule-partial-test.bin pprof/mem-shuffule-partial.out > mem-shuffule-partial.svg
	go tool pprof --svg pprof/shuffule-partial-test.bin pprof/cpu-shuffule-partial.out > cpu-shuffule-partial.svg
	go-torch -f bench/cpu-shuffule-partial-graph.svg pprof/shuffule-partial-test.bin pprof/cpu-shuffule-partial.out
	go-torch --alloc_objects -f bench/mem-shuffule-partial-graph.svg pprof/shuffule-partial-test.bin pprof/mem-shuffule-partial.out
	\
	mv ./*.svg bench/

lint:
	gometalinter --enable-all . | rg -v comment

test: clean init
	GO111MODULE=on go test --race -v $(go list ./... | rg -v vendor)

contributors:
	git log --format='%aN <%aE>' | sort -fu > CONTRIBUTORS
