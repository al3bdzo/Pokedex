# Pokedex CLI

A command-line Pokedex application built with Go, designed as a learning project for Go fundamentals. This project demonstrates core Go concepts including HTTP client development, concurrent programming, caching mechanisms, and CLI REPL implementation.

## Features

- **Interactive REPL**: Command-line interface for exploring Pokemon locations and catching Pokemon
- **PokeAPI Integration**: Fetch real Pokemon data from the [PokéAPI](https://pokeapi.co/)
- **Location Exploration**: Browse different locations and discover Pokemon
- **Pokemon Catching**: Attempt to catch Pokemon with a chance-based mechanic
- **Pokedex Tracking**: Keep track of all Pokemon you've caught
- **Caching System**: Built-in cache to avoid redundant API calls with automatic expiration

## Commands

- **help**: Display all available commands
- **map**: Show the next page of locations
- **mapb**: Show the previous page of locations
- **explore <location-name>**: explore a location and list Pokemons available in that location 
- **catch <pokemon-name>**: Attempt to catch a Pokemon
- **inspect <pokemon-name>**: Show info about a caught Pokemon
- **pokedex**: Display all Pokemon you've caught
- **exit**: Quit the application

## Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation & Setup

```bash
# Clone or navigate to the project directory
cd Pokedex

# Download dependencies
go mod download

# Run the application
go run .
```

## Usage Example

```
Pokedex > help
Welcome to the Pokedex!
Usage:

help: Displays help message
map: Shows locations
mapb: Shows previous locations
explore: Explores a location and lists Pokemons available 
catch: Attempt to catch a Pokemon
inspect: List information about a cuagt Pokemon
pokedex: Shows all caught Pokemon
exit: Exit the Pokedex

Pokedex > map
canalave-city, eterna-city, pastoria-city, ...

Pokedex > explore canalave-city-area
tentacool, staryu, magikarp, ...

Pokedex > catch staryu
Throwing a Pokeball at staryu...
staryu was caught!

Pokedex > inspect staryu
Name: staryu
Height: 8
...

Pokedex > pokedex
Your Pokedex:
  - pikachu
  - staryu

Pokedex > exit
```

## Project Structure

```
Pokedex/
├── main.go                          # Entry point & initialization
├── repl.go                          # CLI REPL loop & command routing
├── repl_test.go                     # REPL tests
├── command_help.go                  # Help command
├── command_map.go                   # Map navigation command
├── command_exit.go                  # Exit command
├── command_catch.go                 # Pokemon catching logic
├── command_explore.go               # Explore location command
├── command_inspect.go               # Inspect caught Pokemon command
├── command_pokedex.go               # Show caught Pokemon command
├── go.mod                           # Module definition
├── .gitignore                       # Git ignore rules
├── README.md                        # This file
└── internal/
    ├── pokeapi/                     # PokeAPI client package
    │   ├── client.go               # HTTP client initialization & setup
    │   ├── pokeapi.go              # Main API wrapper methods
    │   ├── location_list.go        # Location pagination methods
    │   ├── location_explore.go     # Location exploration methods
    │   ├── pokemon_catch.go        # Pokemon data fetching methods
    │   └── response_types.go       # API response data structures
    └── pokecache/                   # Caching mechanism package
        ├── types_cache.go          # Cache data structures & initialization
        ├── cache_methods.go        # Cache Add/Get operations & reapLoop
        └── cache_test.go           # Cache unit tests
```

## Key Learning Concepts

### 1. **HTTP Clients & APIs**
   - Creating and configuring HTTP clients
   - Parsing JSON responses
   - Error handling for network requests

### 2. **Concurrency**
   - Goroutines for background cache cleanup
   - Mutex locks for thread-safe cache operations
   - Race condition prevention

### 3. **Caching**
   - Time-based cache expiration (TTL)
   - Automatic cleanup of expired entries
   - Reducing redundant API calls

### 4. **CLI/REPL Development**
   - User input handling with bufio.Scanner
   - Command parsing and routing
   - Interactive command loop

### 5. **Type System**
   - Struct composition
   - Method receivers
   - Interface usage

### 6. **Standard Library**
   - `net/http` for API calls
   - `encoding/json` for JSON parsing
   - `sync` for concurrency primitives
   - `time` for duration and time operations

## How Caching Works

The cache automatically:
1. Stores API responses with timestamps
2. Wakes up periodically (every 5 minutes by default) to clean expired entries
3. Ensures thread-safe access using mutex locks
4. Prevents redundant API calls for recently fetched data

## Configuration

In `main.go`, you can configure:
- **Cache TTL**: Change the 5-minute cache duration
- **HTTP Timeout**: Adjust the API client timeout (currently 5 seconds)
