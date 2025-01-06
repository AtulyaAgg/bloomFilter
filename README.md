# Bloom Filter Implementation in Go

This repository contains a Go implementation of a **Bloom Filter**—a space-efficient probabilistic data structure designed to test whether an element is a member of a set. The implementation uses the MurmurHash hashing algorithm to demonstrate how false-positive rates can vary depending on the number of hash functions.


## Features

- **Efficient Filtering**: Implements Bloom Filter with customizable sizes.
- **Multiple Hash Functions**: Uses MurmurHash with different seeds.
- **UUID Generation**: Generates random UUIDs for simulating dataset entries.
- **False Positive Analysis**: Computes and logs false-positive rates.


## Prerequisites

- Go (Golang) installed
- Modules `github.com/spaolacci/murmur3` and `github.com/google/uuid installed`
    Install using:
    ```go get github.com/spaolacci/murmur3
        go get github.com/google/uuid ```

## Key Components

### 1. Hash Function (`murmurhash`)

- Utilizes MurmurHash3 with various seeds.
- Scales hashing performance by preinitializing hashers.

### 2. Bloom Filter

- Methods:
    - `Add(key string, numHashfns int)`: Adds a key to the filter.
    - `Exists(key string, numHashfns int)`: Checks for key presence.
- Initializes with a configurable size for space efficiency.

### 3. Simulation & Metrics
- Generates datasets of UUIDs:
    - `database_exists` stores keys known to exist in the dataset.
    - `database_notexists` simulates non-existent keys.
- Tests Bloom Filter's effectiveness for varying hash function counts.

### 4. Output
- Displays false-positive rates for each configuration of hash functions.


## How to Run

1. Clone the repository:
```
    git clone https://github.com/username/bloom-filter-go.git
    cd bloom-filter-go
```

2. Execute the program:
```
    go run main.go
```

3. Observe the output:
- The program prints the false-positive rate for each hash function count.


## Enhancements & Use Cases
- **Enhancements**:
    - Support for more hashing algorithms.
    - Serialization for saving/loading filters.
- **Use Cases**:
    - Probabilistic membership testing.
    - Large-scale data deduplication.
    - Spam filtering in messaging systems.


## Contributing
Contributions are welcome! Feel free to open issues or submit PRs.    
