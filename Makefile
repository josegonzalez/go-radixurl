test:
	REDIS_URL=redis://127.0.0.1:6379                     go test
	REDIS_URL=redis://127.0.0.1:6379/                    go test
	REDIS_URL=redis://127.0.0.1:6379/2                   go test
