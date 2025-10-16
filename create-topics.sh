docker exec -it kafka kafka-topics.sh \
  --create \
  --bootstrap-server localhost:9092 \
  --topic event.completed 
