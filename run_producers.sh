
for i in $(eval echo {1..$1})
do
    python sensipy/producer.py &
done
    python sensipy/producer.py

wait

