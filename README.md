# Simple Key-Value Store in Go

This is a simple key-value store implementation in Go. It provides basic functionalities to store, retrieve, update, and delete key-value pairs. The store supports concurrency control to ensure thread safety and enables persistence to a JSON file for data storage across sessions.

## Features

- **CRUD Operations**: Create, Read, Update, and Delete key-value pairs.
- **Concurrency Control**: Ensures thread safety with mutexes.
- **Persistence**: Data is saved to a JSON file for persistent storage.
- **Querying**: Basic querying functionality to search by keys or values.

## Getting Started

### Prerequisites

- Go installed on your machine.
- (Optional) Git installed to clone the repository.

### Installation

You can install the key-value store by cloning the repository:

```bash
git clone https://github.com/your-username/key-value-store.git

