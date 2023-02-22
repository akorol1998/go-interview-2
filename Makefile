run:
	go run main.go
bench:
	go test -bench=. ./pkg/ttl -benchtime=2s -benchmem -count=5
	go test -bench=. ./pkg/ttl -benchtime=500x -benchmem
	go test -bench=. ./pkg/ttl -benchtime=1000x -benchmem