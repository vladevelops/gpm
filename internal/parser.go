package internal

import "errors"

func ArgumentsFormatter(args []string, COMMANDS_ARRAY []string) ([]validCommand, error) {
	baseHelpString := `
gpm is a tool for managing the instalation for go packeges

Usage: 
	gpm <command> [value]

The commands are:

	i  <install>                  install package
	id <install deep>             serch n packages and let you choose with to install
	mx <use with install deep>    max length list for serch results
`
	if len(args) == 0 {
		// fmt.Println("here")
		return nil, errors.New(baseHelpString)
	}
	command, err := parser(args, COMMANDS_ARRAY, baseHelpString)
	if err != nil {
		return nil, err
	}
	return command, err
}

type validCommand struct {
	Commad string
	Value  string
}

func parser(args []string, COMMANDS_ARRAY []string, baseHelpString string) ([]validCommand, error) {
	var validCommandsArray []validCommand

	for i, token := range args {
		for _, command_main := range COMMANDS_ARRAY {
			if token == command_main {
				selectedValue := args[i+1]

				for _, command := range COMMANDS_ARRAY {
					if selectedValue == command {
						return nil, errors.New("can not run the command whithout values\n" + baseHelpString)
					}
				}
				c := validCommand{command_main, selectedValue}
				validCommandsArray = append(validCommandsArray, c)

			}
		}

	}

	if len(validCommandsArray) == 0 {
		return nil, errors.New("no valid command found\n" + baseHelpString)

	}
	return validCommandsArray, nil
}
