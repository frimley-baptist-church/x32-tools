
usage:
	@echo Usage
	@echo
	@echo make store
	@echo make zero
	@echo make restore

store:
	x32-snapshot/x32-snapshot --store snapshot.json x32

zero:
	x32-zero-main/x32-zero-main x32

restore:
	x32-restore-main/x32-restore-main --file snapshot.json x32
