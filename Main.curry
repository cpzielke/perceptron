module Main where

import Network
import Perceptron
import Train
import Learn
import Algorithms
import LogicExamples

-- ADD
addExamples :: [Example]
addExamples = [
    Example { name = "00N", inputs = [0, 0, 0], expectation = [0, 0] },
    Example { name = "01N", inputs = [0, 1, 0], expectation = [1, 0] },
    Example { name = "10N", inputs = [1, 0, 0], expectation = [1, 0] },
    Example { name = "11N", inputs = [1, 1, 0], expectation = [0, 1] },
    Example { name = "00C", inputs = [0, 0, 1], expectation = [1, 0] },
    Example { name = "01C", inputs = [0, 1, 1], expectation = [0, 1] },
    Example { name = "10C", inputs = [1, 0, 1], expectation = [0, 1] },
    Example { name = "11C", inputs = [1, 1, 1], expectation = [1, 1] }
  ]

addLayers :: [Layer]
addLayers = [
    Layer { nodes = 3, algorithm = logistic },
    Layer { nodes = 2, algorithm = classifier }
  ]

addWeights :: [[[Float]]]
addWeights = [[[1, 1, -1.5]]]

addNetwork :: Network
addNetwork = Network { weights = addWeights, exs = addExamples, layers = addLayers, learningRate = 0.1 }


main :: IO ()
main =  do  putStrLn "Training network..."
            finalWeights <- trainNew addNetwork
            putStrLn "Final weights:"
            print finalWeights
            