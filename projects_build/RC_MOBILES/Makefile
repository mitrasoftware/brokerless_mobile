build:
	GOOS=linux CGO_ENABLED=0 GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go

	# zip deployment.zip bootstrap

	zip function.zip bootstrap
	# aws lambda update-function-code --function-name brokerless-backend --zip-file fileb://function.zip
