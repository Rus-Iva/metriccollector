build-server:
	go build -o cmd/server/server cmd/server/*.go
build-agent:
	go build -o cmd/agent/agent cmd/agent/*.go

test_1:
	metricstest -test.v -test.run=^TestIteration1$$ -binary-path=cmd/server/server
test_2:
	metricstest -test.v -test.run=^TestIteration2B$$ \
                -source-path=. \
                -agent-binary-path=cmd/agent/agent
test_3:
	metricstest -test.v -test.run=^TestIteration3B$ \
                -source-path=. \
                -agent-binary-path=cmd/agent/agent \
                -binary-path=cmd/server/server