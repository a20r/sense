
let range=9002+$1
for i in $(eval echo {9002..$range})
do
    go run src/sense/worker/*.go -port=$i &
done

go run src/sense/worker/*.go -port=9001
wait

