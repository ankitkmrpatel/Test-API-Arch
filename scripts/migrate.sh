#!/bin/bash

# Run database migrations
echo "Starting database migration..."
migrate -path ./migrations -database "$DATABASE_URL" up
if [ $? -eq 0 ]; then
  echo "Migration completed successfully."
else
  echo "Migration failed."
  exit 1
fi
