SHELL=/bin/sh

SERVER_PORT=9999
ADDRESS="localhost:${SERVER_PORT}"
TEMP_FILE=tmpfile

build-server:
	go build -o cmd/server/server cmd/server/*.go
build-agent:
	go build -o cmd/agent/agent cmd/agent/*.go

test_1:
	metricstest -test.v -test.run=^TestIteration1$$ -binary-path=cmd/server/server
test_2a:
	metricstest -test.v -test.run=^TestIteration2A$ \
                -source-path=. \
                -agent-binary-path=cmd/agent/agent
test_2b:
	metricstest -test.v -test.run=^TestIteration2B$ \
                -source-path=. \
                -agent-binary-path=cmd/agent/agent
test_3:
	metricstest -test.v -test.run=^TestIteration3B$ \
                -source-path=. \
                -agent-binary-path=cmd/agent/agent \
                -binary-path=cmd/server/server
test_4:
	metricstest -test.v -test.run=^TestIteration4$$ \
	-agent-binary-path=cmd/agent/agent \
	-binary-path=cmd/server/server \
	-server-port=${SERVER_PORT} \
	-source-path=.
test_5:
	metricstest -test.v -test.run=^TestIteration5$$ \
	-agent-binary-path=cmd/agent/agent \
	-binary-path=cmd/server/server \
	-server-port=${SERVER_PORT} \
	-source-path=.