# Pokedex

A command-line Pokedex application written in Go that allows you to explore Pokemon location areas, catch Pokemon, and view your collected Pokemon with help from boot.dev.

## Description

This Pokedex application uses the [PokeAPI](https://pokeapi.co/) to fetch Pokemon data and provides an interactive command-line interface for exploration and Pokemon collection. You can navigate through different location areas, discover which Pokemon can be found in each area, attempt to catch them, and inspect detailed information about your caught Pokemon.

## Features

- Explore Pokemon location areas in the Pokemon world
- View Pokemon available in each location area
- Attempt to catch Pokemon with a probability-based system
- View your collection of caught Pokemon
- Inspect detailed information about caught Pokemon including stats and types
- Command-line based interface with simple navigation
- Caching system for improved performance

## Installation

### Prerequisites

- Go 1.16 or higher installed on your system

### Building from Source

1. Clone the repository:
   ```
   git clone https://github.com/Skufu/Pokedex.git
   cd Pokedex
   ```

2. Build the application:
   ```
   go build
   ```

## Usage

Run the application:

```
./Pokedex
```

This will start the Pokedex REPL (Read-Eval-Print Loop) interface, where you can enter commands to interact with the application.

## Available Commands

- `help`: Displays a help message with available commands
- `exit`: Exit the Pokedex application
- `map`: Display names of 20 location areas in the Pokemon world
- `mapb`: Display the previous 20 location areas
- `explore [location]`: Display a list of Pokemon in a specific location area
- `catch [pokemon]`: Attempt to catch a Pokemon by name
- `pokedex`: View your caught Pokemon
- `inspect [pokemon]`: View detailed information about a caught Pokemon

## Examples

```
Pokedex > help
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Display names of 20 location areas in the Pokemon world
mapb: Display the previous 20 location areas
explore: Display a list of Pokémon in a location area
catch: Attempt to catch a Pokémon by name
pokedex: View your caught Pokémon
inspect: View detailed information about a caught Pokémon

Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league
...

Pokedex > explore eterna-city-area
Exploring eterna-city-area...
Found Pokémon:
 - budew
 - roselia
 - chansey
 ...

Pokedex > catch roselia
Throwing a Pokeball at roselia...
roselia was caught!

Pokedex > inspect roselia
Name: roselia
Height: 3
Weight: 20
Stats:
  -hp: 50
  -attack: 60
  -defense: 45
  ...
Types:
  - grass
  - poison
```


## Acknowledgements

- [PokeAPI](https://pokeapi.co/) for providing the Pokemon data API
- [boot.dev](https://www.boot.dev/lessons/dff17f87-1ce8-43ce-a43b-2cb611ce76f1) for providing the guide for building this
