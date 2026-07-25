glog:
	git log --oneline --all --graph

gsave:
	git add .
	git commit -m "$(MESSAGE)"
	git push

gsave-new:
	git add .
	git commit -m "$(MESSAGE)"
	git push -u origin "$(NEW_BRANCH)"

gnew-branch:
	git switch "$(SOURCE_BRANCH)"
	git pull
	git switch -c "$(NEW_BRANCH)"

extract-audio:
	if not exist "%USERPROFILE%\Music\converted" mkdir "%USERPROFILE%\Music\converted"
	ffmpeg -i "$(VIDEO)" -vn -b:a 192k ../Music/converted/"$(AUDIO)".mp3

drestart:
	docker compose down
	docker compose up --build -d