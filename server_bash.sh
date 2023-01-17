cd server


server_pids=()

go build server.go 

./server -ip 127.0.0.20 -port 8000 &
server_pid=$!
server_pids+=($server_pid)
./server -ip 127.0.0.30 -port 8000 &
server_pid=$!
server_pids+=($server_pid)
./server -ip 127.0.0.31 -port 8001 &
server_pid=$!
server_pids+=($server_pid)
./server -ip 127.0.0.32 -port 8002 &
server_pid=$!
server_pids+=($server_pid)
./server -ip 127.0.0.20 -port 8001 &
server_pid=$!
server_pids+=($server_pid)
./server -ip 127.0.0.33 -port 8000 &
server_pid=$!
server_pids+=($server_pid)
./server -ip 127.0.0.1 -port 8081 &
server_pid=$!
server_pids+=($server_pid)

cd ..

cd client

go build client.go
./client -write -id 101 -name abc -low 0 -mid 10 -high 100 -host 127.0.0.20 -port 8000

./client -read -id 101 -host 127.0.0.30 -port 8000 

for pid in ${server_pids[@]}; do
    # echo $pid
    kill $pid
done


