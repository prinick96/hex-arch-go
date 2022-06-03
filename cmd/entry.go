package cmd

import (
	"context"
	"hex-arch-go/env"
	"hex-arch-go/internal/database"
	"hex-arch-go/server"
)

// Start the server
func Start() {
	// App context
	ctx := context.Background()

	// env config
	_env := env.GetEnv(".env.development")

	// Run database with env config
	//db := database.NewMySQLDatabase(ctx, _env).ConnectDB() // or work with mysql
	db := database.NewCockRoachDatabase(ctx, _env).ConnectDB()
	defer db.Close()

	// Run server with context, database
	//server.NewGinServer(ctx, db, _env.SERVER_PORT).Run() // with Gin for example
	server.NewEchoServer(ctx, db, _env.SERVER_PORT).Run()
}
