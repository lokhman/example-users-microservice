# Example Users Microservice
This is the example microservice written in Go, that manipulates users in the database and
notifies external services on various actions.

## About
This project was designed to show example architecture of a microservice, that can be used
in a distributed environment. It does not provide a fully featured production ready
functionality, it presents my personal understanding of how I would design this kind of
software.

Selection of components was driven by simplicity of their integration, own preferences and
a final conciseness of the solution. For example, as a messaging service I used [NSQ][nsq]
as a compact, reliable and DevOps-friendly platform (also written in Go) with useful admin
web user interface. Depending on the requirements NSQ can be easily replaced with any
other broker, like [NATS][nats], [Kafka][kafka], [RabbitMQ][rabbitmq], etc.

As a web framework for RESTful API I used [GIN][gin], as one of the simplest and fully
featured web frameworks for Go, which I know quite well as a user and contributor.
[GORM][gorm] is an easily integrated abstraction layer with handy ORM functionality,
whereas PostgreSQL is just my personal choice of DBMS.

## Requirements
In order to start the microservice for development and debug you need to have Docker
installed locally. To production the microservice can be deployed using any orchestration
software (e.g. [Kubernetes][kubernetes], [Amazon ECS][ecs], etc). To support scalability
the `app` service can be scaled to multiple nodes behind a load balancer. Deployment to
production is out of scope of this project.

## How to start
To start the microservice, first unpack the snapshot to the local directory and run the
following command in the command line interface:

    $ docker-compose up

Docker will download and configure the following services within the container:

  - `app`: RESTful API application service (golang:alpine)
  - `database`: PostgreSQL database service (postgres:alpine)
  - `nsqlookupd`: NSQ topology daemon service (nsqio/nsq)
  - `nsqd`: NSQ messaging daemon service (nsqio/nsq)
  - `nsqadmin`: NSQ web UI (nsqio/nsq)

By default `app` service runs [auto-rebuild tool][gin-auto], which is useful for
development and debugging. One-off command to run the application on its own is the
following:

    $ docker-compose run --service-ports app go run .

When container is up and running, the following endpoints will be available.

### RESTful API
| Method | URL                              | Description         |
|--------|----------------------------------|---------------------|
| GET    | http://localhost:8000/           | Health check        |
| GET    | http://localhost:8000/users      | List users          |
| POST   | http://localhost:8000/users      | Create new user     |
| GET    | http://localhost:8000/users/{id} | View user details   |
| PUT    | http://localhost:8000/users/{id} | Update user details |
| DELETE | http://localhost:8000/users/{id} | Delete user         |

### Swagger documentation
URL: http://localhost:8000/docs/index.html

### NSQ admin UI
URL: http://localhost:4171/

## Testing
There are two types of tests provided in the project: unit tests and integration tests.
Since this microservice was designed as a self-sufficient isolated container, the CI
process may consider it as a standalone testing environment, which can be created and
destroyed.

### Unit tests
Example unit tests are provided only for some functionality in `./common` package. To
run unit tests with Docker run the following command in the command line interface:

    $ docker-compose run app go test -v ./...

### Integration tests
Integration tests are located in `main_test.go` file and assess all API endpoints using
database and NSQ servers from the container. Since we assume that the container can be
deployed to a testing environment, database and NSQ services are safe to use for mocking.
To run integration tests with Docker run the following command in the command line
interface:

    $ docker-compose run -e GIN_MODE=test app go test -tags=integration -failfast -v

## TODO
As was mentioned above, this project is not production ready, and was written simply to
demonstrate overall design and integration of separate components. It means, that there
are things, that are missing or not properly implemented. At the end of the day it's just
an example.

My TODO list includes but not limited to the following:

  - Extend unit and integration tests with new cases. Add coverage report.
  - Add NSQ subscriber to integration tests to assess message delivery.
  - Refactor code to smaller functions for better unit testing (e.g. in API "controller").
  - Add better error handling and logging (currently we rely on GIN).
  - Improve data validation and reporting.
  - Revise documentation.

You can also find many considerations and assumptions in code comments.

## License
This code is available under the MIT license. LICENSE file describes this in detail.

[nsq]: https://nsq.io/
[nats]: https://nats.io/
[kafka]: https://kafka.apache.org/
[rabbitmq]: https://www.rabbitmq.com/
[gin]: https://gin-gonic.github.io/gin/
[gorm]: http://gorm.io/
[kubernetes]: https://kubernetes.io/
[ecs]: https://aws.amazon.com/ecs/
[gin-auto]: https://github.com/codegangsta/gin
