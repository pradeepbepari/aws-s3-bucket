# prerequisite for these application 
-- you should have docker and docker compose installed in your machine

# to tun the application 
command : make up

# After building the container we need to create a bucket manually why because we are running s3 on localstack
command to create bucket : make create

# to stop the application
command : make down

# to check the bucket list
command : make list
