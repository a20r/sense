
for i in {2..2}
do
    echo ${i}0
    bash run_clients.sh ${i}0 & # > data/client_${i}0.txt &
    bash run_producers.sh ${i}0 > data/producer_${i}0.txt

    wait
done
