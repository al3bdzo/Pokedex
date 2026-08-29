package main


func main() {
	conf := &config{
		supportedCommands: getCommands(),
	}
	repl(conf)
}


