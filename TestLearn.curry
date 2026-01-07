module TestLearn(
    testLearn, testInitialWeights, testBackPropagate,
    testGD, testLE, testGate, testEpic, testLEgate, testLearnGate,
    testFFnet, testBPnet, testGDnet, testLEnet, testEpicNet, testLearnNet, testNetWs, testNetwork, verifyNet
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

testFFnet :: Bool
testFFnet =  feedForward testNetwork [0, 0] networkWeights == [[1], [0.574442516811659, 0.5986876601124521], [0, 0]]

testBPnet :: Bool
testBPnet =  backPropagate (reverse (layers testNetwork)) (expectation testNetExample) networkWeights [[1], [0.574442516811659, 0.5986876601124521], [0, 0]] == [[0], [0, 0]]

testGDnet :: Bool
testGDnet =  gradientDescent (learningRate testNetwork) networkWeights [[1], [0.574442516811659, 0.5986876601124521], [0, 0]] [[0], [0, 0]] == networkWeights

testLEnet :: Bool
testLEnet =  learnExample testNetwork networkWeights testNetExample == networkWeights

testEpicNet :: [[[Weight]]]
testEpicNet =  learnEpic testNetwork networkWeights

testLearnNet :: [[[Weight]]]
testLearnNet =  learn 5000 testNetwork networkWeights

testNetWs :: [[[Weight]]] -> [Example]
testNetWs    nws          =  filter (\e -> eval2 testNetwork nws (inputs e) /= expectation e) (exs testNetwork)

verifyNet :: [Example]
verifyNet =  testNetWs (reverse testLearnNet)
