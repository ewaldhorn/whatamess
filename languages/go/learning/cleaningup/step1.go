package main

func main() {
	if err := run(); err != nil {
	  fmt.Fprintf(os.Stderr, "%s\n", err)
	  os.Exit(exitFail)
	}
  }
  
  func run() error {
	db, err := psql.Connect(...)
	if err != nil {
	  return fmt.Errorf("connecting to db: %w", err)
	}
	app := App{db}
	return app.Start()
  }