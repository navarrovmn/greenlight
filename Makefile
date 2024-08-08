run-greenlight:
	go run ./cmd/api --smtp-username ${GREENLIGHT_SMTP_USERNAME} --smtp-password ${GREENLIGHT_SMTP_PASSWORD} -cors-trusted-origins="http://localhost:9000 http://localhost:9001"