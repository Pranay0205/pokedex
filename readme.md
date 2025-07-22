# Pokedex CLI

A command-line Pokedex application built in Go that allows you to explore Pokemon locations, catch Pokemon, and manage your collection using the [PokeAPI](https://pokeapi.co/).

## Features

- 🗺️ **Explore Locations**: Browse through Pokemon location areas
- 🔍 **Location Details**: Discover which Pokemon can be found in specific areas
- ⚡ **Catch Pokemon**: Attempt to catch Pokemon with realistic success rates
- 📱 **Inspect Pokemon**: View detailed stats, types, and information about caught Pokemon
- 📚 **Pokedex**: Keep track of your Pokemon collection
- ⚡ **Caching**: Built-in HTTP response caching for improved performance

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Pranay0205/pokedexcli.git
cd pokedexcli
```

2. Build the application:
```bash
go build -o pokedexcli
```

3. Run the application:
```bash
./pokedexcli
```

## Usage

Once started, you'll see the Pokedex prompt. Here are the available commands:

### Available Commands

| Command | Description | Usage |
|---------|-------------|-------|
| `help` | Displays all available commands | `help` |
| `exit` | Exits the Pokedex application | `exit` |
| `map` | Displays next 20 location areas | `map` |
| `mapb` | Displays previous 20 location areas | `mapb` |
| `explore` | Explores a location area for Pokemon | `explore <location-name>` |
| `catch` | Attempts to catch a Pokemon | `catch <pokemon-name>` |
| `inspect` | Shows details of a caught Pokemon | `inspect <pokemon-name>` |
| `pokedex` | Lists all caught Pokemon | `pokedex` |

### Example Session

```bash
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon:
1) tentacool
2) tentacruel
3) staryu
...

Pokedex > catch pidgey
Throwing a Pokeball at pidgey...
pidgey was caught!
You may now inspect it with the inspect command.

Pokedex > inspect pidgey
Name: pidgey
Height: 3
Weight: 18
Stats:
  -hp: 40
  -attack: 45
  -defense: 40
  -special-attack: 35
  -special-defense: 35
  -speed: 56
Types:
  - normal
  - flying

Pokedex > pokedex
Your Pokedex:
 - pidgey
```

## Technical Features

### Caching System
- **HTTP Response Caching**: All API responses are cached with configurable TTL
- **Thread-Safe**: Uses mutex locks for concurrent access
- **Automatic Cleanup**: Background goroutine removes expired entries

### Error Handling
- Comprehensive input validation
- Graceful API error handling
- User-friendly error messages

### Testing
Run the test suite:
```bash
go test ./...
```

## Configuration

The application uses the following default settings:
- **HTTP Timeout**: 5 seconds
- **Cache TTL**: 5 minutes
- **API Base URL**: `https://pokeapi.co/api/v2`

## Dependencies

- **Go 1.24.4+**
- **Standard Library Only**: No external dependencies required
- **PokeAPI**: Free RESTful Pokemon API

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the [MIT License](LICENSE).

## Acknowledgments

- [PokeAPI](https://pokeapi.co/) for providing the Pokemon data
- The Go community for excellent documentation and examples

---

*Gotta catch 'em all! 🎮*