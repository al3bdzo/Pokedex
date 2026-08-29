package main


func main() {
	conf := &config{
		supportedCommands: getCommands(),
		nextURL: "https://pokeapi.co/api/v2/location-area",
	}
	repl(conf)
}


