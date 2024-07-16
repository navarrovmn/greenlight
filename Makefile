run-greenlight:
	go run ./cmd/api --smtp-username ${GREENLIGHT_SMTP_USERNAME} --smtp-password ${GREENLIGHT_SMTP_PASSWORD}