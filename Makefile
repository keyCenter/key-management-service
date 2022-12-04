BIN_FILE=KMS

hello:
	echo "Hello"

clean:
	rm -f ${BIN_FILE}

build:
	go build -o ${BIN_FILE}

run:
	./${BIN_FILE}

start:
	nohup make run >> /home/work/log/${BIN_FILE}.log 2>&1 &

stop:
	pidof ./${BIN_FILE} | xargs kill -9

restart: stop start

c_run: clean build run

c_start: clean build start

c_restart: clean build restart