# Description 
We will use a Unit of Work (UoW) pattern to manage the database operations, ensuring that changes to both tables are committed together.

These are the tables:
**scheduler_status table**
Stores information about schedulers with a unique ID, name, and current state
Tracks when each scheduler was created and last updated
Uses UUIDs as primary keys and includes timestamps with timezone support

**k8s_deployment_status table**
Tracks the deployment status of Kubernetes resources associated with schedulers
Links to the scheduler_status table through a foreign key relationship
Records the deployment state and timestamps for creation and updates
The script uses PostgreSQL-style syntax with UUID data types and timezone-aware timestamps. It's designed to monitor and track the lifecycle of schedulers and their corresponding Kubernetes deployments in a system.

# How To Run The Code

1. Clone the repository and install Docker and Docker Compose if you haven't already.
2. From the root directory of the project, run the following command to start the local development environment:
```bash
$ ./local-development-start.sh
```
3. Run the tests to ensure UoW runs correctly:
```bash
$ make test
```