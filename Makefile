# Development purposes

PID = /tmp/serving.pid

# Serve task will run fswatch monitor and performs restart task if any source file changed. Before serving it will execute start task.
serve: start
	fswatch -or --event=Updated . | xargs -n1 -I {} make restart

kill:
	-kill `pstree -p \`cat $(PID)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"`

# Before task will only prints message. Actually, it is not necessary. You can remove it, if you want.
before:
	@echo "STOPPED" && printf '%*s\n' "40" '' | tr ' ' -

start:
	./scripts/development.sh & echo $$! > $(PID)

# Restart task will execute kill, before and start tasks in strict order and prints message.
restart: kill before start
	@echo "STARTED" && printf '%*s\n' "40" '' | tr ' ' -

.PHONY: serve restart kill before start