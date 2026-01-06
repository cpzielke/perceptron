module TestLearn(
    testLearn, testInitialWeights, testDeltaOut, testDeltaOut2, testDeltaNeuron, testDeltaInner, testBackPropagate,
    testGD, testLE, testGate, testEpic, testLEgate, testLearnGate,
    testFFnet, testBPnet, testGDnet, testLEnet, testEpicNet, testLearnNet
)where

import Learn
import Algorithms
import Network
import LogicExamples
import Perceptron

testLearn :: [[[Weight]]]
testLearn = learn 1 andNetwork (weights andNetwork)
 
testInitialWeights :: [[[Weight]]]
testInitialWeights =  initialWeights nandNetwork 2

testDeltaOut :: [Float]
testDeltaOut =  deltaOut (head (layers (andNetwork))) (expectation (head (exs (andNetwork)))) [1]

testDeltaOut2 :: [Float] -> [Float]
testDeltaOut2 outs =  deltaOut (head (layers (andNetwork))) [1, 1] outs

testDeltaNeuron :: Bool
testDeltaNeuron = deltaNeuron 0 [[1,2,3,4], [5,6,7,8]] [2,3] == 17

testLayer :: Layer
testLayer = Layer { nodes = 3, algorithm = classifier }

testDeltaInner :: Bool
testDeltaInner =  deltaInner 0 testLayer [[1,2,3,4], [5,6,7,8]] [2,3] [10,20,30] == [17, 22, 27]

testBackPropagate :: [[Float]]
testBackPropagate =  backPropagate (reverse (layers nandNetwork)) (expectation example11) (reverse (weights nandNetwork)) [[0], [1], [1, 1]]
    where [_, _, _, example11] = exs nandNetwork

testGD :: [[[Weight]]]
testGD =  gradientDescent (learningRate nandNetwork) (reverse (weights nandNetwork)) [[0], [1], [1, 1]] [[0], [0]]

testLE :: [[[Weight]]]
testLE =  learnExample nandNetwork (reverse (weights nandNetwork)) (head (exs nandNetwork))

testLEs :: Network -> [Example]
testLEs    network =  filter (\e -> learnExample network (reverse (weights network)) e /= (reverse (weights network))) (exs network)

testEpic :: Network -> [[[Weight]]] -> [[[Weight]]]
testEpic    network    nws          =  learnEpic network nws

testGate :: Network -> [Input] -> [[[Weight]]] -> [Output]
testGate    network    inputs     nws          = eval2 network nws inputs

testLEgate :: Network -> [[[Weight]]] -> [[[Weight]]]
testLEgate    network    rnws         =  learnExample network rnws (head (exs network))

testLearnGate :: Network -> [[[Weight]]] -> [[[Weight]]]
testLearnGate    network    rnws         =  learn 1 network rnws


testNetwork :: Network
testNetwork =  xorNetwork

networkWeights :: [[[Weight]]]
networkWeights =  [[[0.2, 0.2, 0.2]], [[0.3, 0.3, 0.3], [0.4, 0.4, 0.4]]]

testNetExample :: Example
testNetExample =  head (exs testNetwork)

testFFnet :: [[Weight]]
testFFnet =  feedForward testNetwork [0, 0] networkWeights

testBPnet :: [[Float]]
testBPnet =  backPropagate (reverse (layers testNetwork)) (expectation testNetExample) networkWeights [[1], [0.574442516811659, 0.5986876601124521], [0, 0]]

testGDnet :: [[[Weight]]]
testGDnet =  gradientDescent (learningRate testNetwork) networkWeights [[1], [0.574442516811659, 0.5986876601124521], [0, 0]] [[1], [0]]

testLEnet :: [[[Weight]]]
testLEnet =  learnExample testNetwork networkWeights testNetExample

testEpicNet :: [[[Weight]]]
testEpicNet =  learnEpic testNetwork networkWeights

testLearnNet :: [[[Weight]]]
testLearnNet =  learn 1000 testNetwork networkWeights
