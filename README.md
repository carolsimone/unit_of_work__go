# Unit of Work Pattern Example

A practical implementation of the Unit of Work (UoW) pattern for managing database transactions across multiple tables in Go.

## Overview

This project demonstrates how to ensure data consistency when working with related database operations. The UoW pattern guarantees that changes to multiple tables are committed together or rolled back entirely if any operation fails.

### Database Schema

The example uses two PostgreSQL tables:

**`scheduler_status`**
- Manages scheduler lifecycle with UUID primary keys
- Tracks scheduler name, state, and timestamps
- Uses timezone-aware timestamps for audit trails

**`k8s_deployment_status`**
- Monitors Kubernetes deployment states
- Foreign key relationship to `scheduler_status`
- Records deployment state transitions with timestamps

## Quick Start

1. **Prerequisites**: Docker and Docker Compose installed

2. **Start local environment**:
```bash
   ./local-development-start.sh
```

3. **Run tests**:
```bash
   make test
```

The tests verify that the UoW pattern correctly handles both successful transactions and rollback scenarios.