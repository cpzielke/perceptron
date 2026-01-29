module TestPerceptron(testLogic, testNand11, testNetwork, testEvalNetWs) where
import Algorithms
import Network
import LogicExamples
import Perceptron

testNetwork :: Network -> [Example]
testNetwork    network =  filter (\e -> eval network (inputs e) /= expectation e) (exs network)

testLogic :: [Example]
testLogic =  (testNetwork andNetwork) ++ (testNetwork orNetwork) ++ (testNetwork notNetwork) ++ (testNetwork nandNetwork)

testNand11 :: [[Float]]
testNand11 =  networkDotProduct nandNetwork (inputs example11) (weights nandNetwork)
    where [_, _, _, example11] = exs nandNetwork

testEvalNetWs :: [Example]
testEvalNetWs =  evalRNetWs nandNetwork (reverse (weights nandNetwork))
