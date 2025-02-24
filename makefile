set-app-path:
	sudo chmod +x ./shell/set.sh
	zsh ./shell/set.sh
	#sh ./shell/set.sh
update-go-nest:
	sudo chmod +x ./shell/update-lib.sh
	zsh ./shell/update-lib.sh
	#sh ./shell/update-lib.sh
test_app:
	go test ./test/... -v -bench . -failfast -cover -count=1

