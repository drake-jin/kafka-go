#!/bin/bash

KAFKA_HOME="./kafka"

$KAFKA_HOME/bin/zookeeper-server-start.sh $KAFKA_HOME/config/zookeeper.properties
