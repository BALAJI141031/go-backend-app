server: 
	nodemon --watch './**/*.go' --signal SIGTERM --exec APP_ENV=local 'go' run main.go