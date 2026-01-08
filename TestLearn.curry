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
testNetExample =  let [_, _, _, x] = exs testNetwork
                  in x

testFFnet :: [[Float]]
testFFnet =  feedForward testNetwork (inputs testNetExample) networkWeights
                 -- == [[1], [0.7109495026250039, 0.7685247834990178], [1, 1]]

testBPnet :: [[Float]]
testBPnet =  backPropagate (reverse (layers testNetwork)) (expectation testNetExample) networkWeights [[1], [0.7, 0.7], [1, 1]]
                 -- == [[1], [0.04200000000000001, 0.04200000000000001]]

testGDnet :: [[[Float]]]
testGDnet =  gradientDescent (learningRate testNetwork) networkWeights [[1], [0.7, 0.7], [1, 1]] [[1], [0.1, 0.1]]
                 -- == [[[0.13, 0.13, 0.1]], [[0.29, 0.29, 0.29], [0.39, 0.39, 0.39]]]

testLEnet :: Bool
testLEnet =  learnExample testNetwork networkWeights testNetExample == networkWeights

testEpicNet :: [[[Float]]]
testEpicNet =  learnEpic testNetwork networkWeights


testLearnNet :: [[[Weight]]]
testLearnNet =  learn 5000 testNetwork networkWeights

testNetWs :: [[[Weight]]] -> [Example]
testNetWs    nws          =  filter (\e -> eval2 testNetwork nws (inputs e) /= expectation e) (exs testNetwork)

verifyNet :: [Example]
verifyNet =  testNetWs (reverse testLearnNet)
