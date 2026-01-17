module Train (trainEpic, train, trainNew, trainLoop, trainUntil) where

import Debug.Trace
import Network
import Perceptron
import Learn

trainEpic :: Network -> [[[Weight]]] -> [[[Weight]]]
trainEpic    network    rnws         =  foldl (learnExample network) rnws (exs network)

train :: Int -> Network -> [[[Weight]]] -> [[[Weight]]]
train    count  network    rnws         =  let newNWs = trainEpic network rnws
                                           in  if count <= 0 then rnws else train (count - 1) network newNWs

trainUntil :: Network -> [[[Weight]]] -> [[[Weight]]]
trainUntil    network    rnws         =  until (passRNetWs network) (trainEpic network) rnws

trainNew :: Network -> IO [[[Weight]]]
trainNew    network =  do putStrLn "Start"
                          trainLoop network (revInitialWeights network)

trainLoop :: Network -> [[[Weight]]] -> IO [[[Weight]]]
trainLoop    network    rnws         =  do  let newRnws = train 5000 network rnws
                                            putStr "W: "
                                            print newRnws
                                            if passRNetWs network newRnws then return newRnws else trainLoop network newRnws

-----------------------------------------------------------------------------------------------------------------
