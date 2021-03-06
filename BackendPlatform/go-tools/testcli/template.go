package testcli

var (
	yml = `version: "3.0"

services:
  db:
    image: mysql:5.6
    ports:
    - 3306
    environment:
    - MYSQL_ROOT_PASSWORD=root
    - MYSQL_DATABASE=%s
    - TZ=Asia/Shanghai
    volumes:
    - .:/docker-entrypoint-initdb.d
    command: [
      '--character-set-server=utf8',
      '--collation-server=utf8_unicode_ci'
    ]

  redis:
    image: redis
    ports:
      - 6379
  zookeeper:
    image: wurstmeister/zookeeper
    ports:
      - 2181
  kafka:
    image: wurstmeister/kafka
    depends_on: [ zookeeper ]
    ports:
      - 9092
    environment:
      KAFKA_ADVERTISED_HOST_NAME: "%s"
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_CREATE_TOPICS: "%s"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock`
)
