# IOI Playground
## Requirements
- Build Essential
- Docker

## Instruction

1. To build this and run this just **$make** in command line

| Command | Used |
| ------- | ---- |
| make | build image and run container |
| make all | build image and run container |
| make run | run image as container |
| make build | build image |
| make stop | stop container |
| make status | check container status |
| make logs | loggs container |
| make clean | stop container and clear image |
| make re | rebuild image and run container |

> Then it can be access via ***localhost:8080***  
***test.md*** is available test json to use

### Routes

| Route | Used |
| ------- | ---- |
| / | test connection |
| /run_c | put c code in to running queue and return job_id |
| /job_id | job_id to get result of code (status, stdout, stderr) |