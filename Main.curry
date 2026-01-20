module Main where

import Network
import Perceptron
import Train
import Learn
import Algorithms
import LogicExamples

main :: IO ()
main =  do  putStrLn "Training network..."
            finalWeights <- trainNew xorNetwork
            putStrLn "Final weights:"
            print finalWeights