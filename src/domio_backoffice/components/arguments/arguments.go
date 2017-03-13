package arguments

import (
    "flag"
    "os"
    "fmt"
    "domio_backoffice/components/config"
    "domio_backoffice/components/server"
)

type Arguments struct {
    Command string
}

func ProcessArguments() (Arguments, error) {
    var args = Arguments{}

    if len(os.Args) == 1 {
        showHelpAndExit()
    }

    switch os.Args[1] {

    case "init":
        args = processInitArguments()

    case "start":
        args = processStartArguments()

    default:
        fmt.Printf("%q is not valid command.\n", os.Args[1])
        os.Exit(2)
    }

    return args, nil
}

func showHelpAndExit() {
    fmt.Println("usage: domio <command> [<args>]")
    fmt.Println("Commands are: ")
    fmt.Println(" init   Init with new config file")
    fmt.Println(" start  Start server")
    os.Exit(1)

}

func processInitArguments() Arguments {
    var args = Arguments{
        Command: "init",
    }

    initCommand := flag.NewFlagSet("init", flag.ExitOnError)

    filenameFlag := initCommand.String("file", "config.json", "config file absolute path")
    webPortFlag := initCommand.Uint("port", 8090, "Port for the HTTP server to run on")
    envFlag := initCommand.String("env", "development", "Environment name: development, testing, production")
    apiUrl := initCommand.String("api-url", "https://api.domio.in/", "Url where API is served")
    useDart := initCommand.Bool("use-dart", false, "Use Dart instead of JS")

    initCommand.Parse(os.Args[2:])

    if initCommand.Parsed() {
        config.InitConfigFile(filenameFlag, webPortFlag, envFlag, apiUrl, useDart)
    }
    return args
}

func processStartArguments() Arguments {
    var args = Arguments{
        Command: "start",
    }

    startCommand := flag.NewFlagSet("start", flag.ExitOnError)
    //recipientFlag := sendCommand.String("recipient", "", "Recipient of your message")
    //messageFlag := sendCommand.String("message", "", "Text message")
    //log.Print(*startCommand)

    startCommand.Parse(os.Args[2:])

    config.LoadConfig()

    if startCommand.Parsed() {
        server.Start()
    }
    return args
}
