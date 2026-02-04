module Main (main, addNetwork) where

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
    Layer { nodes = 4, algorithm = logistic },
    Layer { nodes = 2, algorithm = classifier }
  ]

addWeights :: [[[Float]]]
addWeights = [
    [
      --[-0.14181385681495548, 0.18101883140728042, 0.2667752875156593, -2.235642747338454],
      --[1.0237803190746038, 0.4392286431233121, 1.0346636324225131, -2.736708827041986],
      --[-1.8682894731204003, -1.7149844573679054, -1.4160263070279422, -1.1095255779746231],
      --[-1.6601931463867132, -1.1745242464599182, -1.5306735213471891, 1.1162976192236396]
    ],
    [
      --[0.1621758672231053, 1.1312442911868033, -1.6679616858844406, 1.0869127409386288, -0.5],
      --[1.2882529085015741, 0.5111566473716767, -0.0021684209201429843, -1.609594197729459, 0.1]
    ]
  ]


addNetwork :: Network
addNetwork = Network { weights = addWeights, exs = addExamples, layers = addLayers, learningRate = 0.1 }

main :: IO [[[Weight]]]
--main =  do putStrLn "Training ADD Network..." 
--           let revWeights = trainNew addNetwork
--           putStrLn ("Trained Weights: " ++ show revWeights)
--           return revWeights

main =  do putStrLn "Testing dotProduct..."
           let result = dotProduct [1.0, 2.0, 3.0] [0.5, 0.5, 0.5, 2.0]
           putStrLn ("Dot Product Result: " ++ show result)
           return []

