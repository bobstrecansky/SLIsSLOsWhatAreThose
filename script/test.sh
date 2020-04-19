#!/bin/bash

for i in {1..100}; do curl localhost:2112/fast_response; done
curl localhost:2112/slow_response
curl localhost:2112/error_response
