CC=go

heartbeat.so:
	$(CC) build --buildmode=plugin heartbeat.go

install: heartbeat.so
	mkdir -p /var/lib/nwatch/
	mv heartbeat.so /var/lib/nwatch/

.PHONY: clean

clean:
	rm -f heartbeat.so
