
go run src/sense/broker/*.go &
let range=8002+$1
for i in $(eval echo {8002..$range})
do
    go run src/sense/worker/*.go -port=$i &
done

go run src/sense/worker/*.go -port=8001
wait

