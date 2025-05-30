CERT_DIR=./app/cert
CERT_PATH=$(CERT_DIR)/cert.pem
KEY_PATH=$(CERT_DIR)/key.pem

cert:
	@if command -v ipconfig > /dev/null; then \
		IP=$$(ipconfig getifaddr en1 2>/dev/null || ipconfig getifaddr en0 2>/dev/null); \
	elif command -v hostname > /dev/null; then \
		IP=$$(hostname -I | awk '{print $$1}'); \
	else \
		IP=$$(ifconfig | grep "inet " | grep -v "127.0.0.1" | awk '{print $$2}' | head -n 1); \
	fi; \
	mkdir -p $(CERT_DIR); \
	mkcert -cert-file $(CERT_PATH) -key-file $(KEY_PATH) $$IP localhost 127.0.0.1 ::1;

