# Pub/Sub Implementations

In order for scoreboard to store points against uers it needs to subscribe to
the SOON\_ FM event system.

At time of writting this is a simple Redis Pub/Sub service. Other implementations
may follow, such as Apache Kafka.

## Events

* `play`: The play event is used to add a point to a user
* `stop`: The stop event is used to remove a point from a user, this is because
  stop is equivilant to skip.

## Implementations

* `redis`: The `scoreboard/pubsub/redis` package provides a Redis Pub/Sub implementation.
