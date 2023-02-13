# Wadhwani AI Assignment

## Components

## Setup

## Testing

## Improvements

- File Upload can use a file storage system (S3)
  - Separate storage for files
  - pipeline to parse files and send messages to the service to add/update data
- Parallelism
  - Use goroutines and sync groups to parallelize some functionalities (dto, csv parsing)
