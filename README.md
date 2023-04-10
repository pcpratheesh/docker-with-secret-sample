/# docker-with-secret-sample
A sample docker application for demonstrating docker secrets

## Follow the commands to run ande deploy the service

    docker image build -t docker-secret-reader .

    docker run -p 8080:8080 docker-secret-reader

    docker-compose up --build



## How to create new secrets
    // create new secrets
    docker secret create demo_secrets sample_secret.txt

    docker secret rm db_host

    echo "mysecretusername" | docker secret create db_user -

    echo "db-host" | docker secret create db_host -



