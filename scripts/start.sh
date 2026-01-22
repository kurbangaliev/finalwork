#!/bin/bash
# Start the first server
./server &
# Start the second server
./admin-server &

# Wait for any process to exit
wait -n

# Exit with status of process that exited first
exit $?