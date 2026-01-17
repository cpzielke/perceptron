module TestTrain (testTrain, testEpic, testEpicNet, testTrainNet, testNetWs, verifyNet, w, networkWeights) where

import Network
import LogicExamples
import Perceptron
import Train


testTrain :: [[[Weight]]]
testTrain = train 2 andNetwork (weights andNetwork)

testEpic :: Network -> [[[Weight]]] -> [[[Weight]]]
testEpic    network    nws          =  trainEpic network nws


testNetwork :: Network
testNetwork =  xorNetwork

networkWeights :: [[[Weight]]]
networkWeights =  [[[0.2, 0.2, 0.2]], [[0.3, 0.3, 0.3], [0.4, 0.4, 0.4]]]

testEpicNet :: [[[Float]]]
testEpicNet =  trainEpic testNetwork networkWeights

testTrainNet :: [[[Weight]]]
testTrainNet =  train 5000 testNetwork networkWeights

testNetWs :: [[[Weight]]] -> [Example]
testNetWs    rnws          =  evalRNetWs testNetwork rnws

verifyNet :: [Example]
verifyNet =  testNetWs testTrainNet

w :: IO ()
w =  do let result = testEpicNet
        putStr "testEpicNet = "
        print result