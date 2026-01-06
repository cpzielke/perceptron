module Network (
   Input, Weight, Output,
   Example, Example (name), Example (inputs), Example (expectation),
   Layer, Layer (nodes), Layer (algorithm),
   Network, Network (weights), Network (exs), Network (layers), Network (learningRate)
) where

import Algorithms

type Input = Float
type Output = Float
type Weight = Float

data Example = Example { name :: String, inputs :: [Input], expectation :: [Output] }

data Layer = Layer { nodes :: Int, algorithm :: Algorithm }

data Network = Network { weights :: [[[Weight]]], exs :: [Example], layers :: [Layer], learningRate :: Float }
