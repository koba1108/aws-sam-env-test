dev:
	sam build -p && sam local start-api -p 3001 -n env.json