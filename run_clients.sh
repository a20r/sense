
for i in $(eval echo {1..$1})
do
    python sensipy/client.py &
done

python sensipy/client.py

wait
