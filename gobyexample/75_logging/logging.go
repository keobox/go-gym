package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	log.Println("standard logger")

	// set binary flags for us accuracy
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// set binary flags for file and line
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// custom logger (can be passed around)
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// you can set a prefix
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// loggers can have custom targets
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	buflog.Println("hello")

	fmt.Print("from buflog:", buf.String())

	// slog is used for structured logging
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	// for example, logging in json
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	myslog.Info("hello again", "key", "val", "age", 25)
}
