# GoNATS
NATS messaging - GoLang quickstart project

### ENVS: create .env file with
    # default
    NATS_SERVER=0.0.0.0:4222
    # docker-compose 
    # NATS_SERVER=nats:4222
    
##### If local NATS server is available     
    $ make start_local 
##### Docker     
    $ make start_docker
##### Docker-Compose stack
    $ make start_compose
##### clean build
    $ make clean
    
### Please suggest improvements

Author: Sumit S. Agrawal