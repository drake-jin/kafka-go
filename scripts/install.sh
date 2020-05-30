#!/bin/bash

curl http://apache.mirror.cdnetworks.com/kafka/2.5.0/kafka_2.12-2.5.0.tgz -sLo ./kafka_2.12-2.5.0.tgz
tar xvf ./kafka_2.12-2.5.0.tgz
rm kafka_2.12-2.5.0.tgz
mv kafka_2.12-2.5.0 ../kafka