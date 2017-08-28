# cockroachdb
Testing cockroachdb along with a simple docker client.

## Setup
1. Build the docker image with **docker/haproxy/build_docker.sh**
2. Create the docker network with **scripts/create_network.sh**
3. Run the master cockroachdb node with **scripts/run_master.sh**
4. Join a second node to the master node with **scripts/run_slave.sh slave1**
5. Join a third node to the master node with **scripts/run_slave.sh slave2**
6. Note: More nodes can be entered but the **haproxy.cfg** file needs to be updated.
7. Start the haproxy with **scripts/run_haproxy.sh**

## Run
Build and run with the parameters:

-db_url=postgres-connection-url

-init to initialize the database with some values (example: postgres://root@<url>:<port>/MovieDB?sslmode=disable)

Stop and start three nodes one at a time. HAProxy will detect and stop the traffic to the downed node until its started again. And the inserted values will be distributed along the three nodes.
