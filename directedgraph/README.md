# directedgraph

Golang implementation of DirectedGraph type.

    host$ export GOPATH=~/workspace
    host$ mkdir -p $GOPATH
    host$ cd $GOPATH

    host$ go get github.com/chrisshiels/golang/directedgraph/...

    host$ go test -v github.com/chrisshiels/golang/directedgraph/...
    === RUN   Test_ShouldInstantiate
    --- PASS: Test_ShouldInstantiate (0.00s)
    === RUN   Test_ShouldManageVertices
    --- PASS: Test_ShouldManageVertices (0.00s)
    === RUN   Test_ShouldManageEdges
    --- PASS: Test_ShouldManageEdges (0.00s)
    === RUN   Test_ShouldIgnoreSubsequentlyAddingTheSameVertex
    --- PASS: Test_ShouldIgnoreSubsequentlyAddingTheSameVertex (0.00s)
    === RUN   Test_ShouldTopologicallySortAnEmptyGraph
    --- PASS: Test_ShouldTopologicallySortAnEmptyGraph (0.00s)
    === RUN   Test_ShouldTopologicallySortASimpleAcyclicGraph
    --- PASS: Test_ShouldTopologicallySortASimpleAcyclicGraph (0.00s)
    === RUN   Test_ShouldNotDetectCyclesInAnEmptyGraph
    --- PASS: Test_ShouldNotDetectCyclesInAnEmptyGraph (0.00s)
    === RUN   Test_ShouldNotDetectCyclesWhenTheGraphIsACyclic
    --- PASS: Test_ShouldNotDetectCyclesWhenTheGraphIsACyclic (0.00s)
    === RUN   Test_ShouldDetectCyclesFromTheLastVertexToTheFirstVertex
    --- PASS: Test_ShouldDetectCyclesFromTheLastVertexToTheFirstVertex (0.00s)
    === RUN   Test_ShouldDetectCyclesFromWithinTheVertexChain
    --- PASS: Test_ShouldDetectCyclesFromWithinTheVertexChain (0.00s)
    PASS
    ok  	github.com/chrisshiels/golang/directedgraph	0.002s
