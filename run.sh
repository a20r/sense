
go run src/sense/entry/*.go &
go run src/sense/worker/*.go -port=8001 &
go run src/sense/worker/*.go -port=8002

wait

