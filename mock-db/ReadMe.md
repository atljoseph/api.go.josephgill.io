# mock-db

This is a mock database used for development.

Until migrated and populated, it is basically an emtpy shell of a mysql database.

# Commands

- To run with Docker locally, run `docker-compose up`, optionally provide the `--build` argument.
- To stop the container, hit `Ctrl + c`.
- To kill the container, run `docker-compose down`.
- Sometimes it is useful to run `docker system prune --force` to reset your docker environment. It will remove all containers not running, as well as all docker networks.
