module Train (trainEpic, train, trainNew, trainLoop, trainUntil) where

import Debug.Trace
import Network
import Perceptron
import Learn

trainEpic :: Network -> [[[Weight]]] -> [[[Weight]]]
trainEpic    network    rnws         =  foldl (learnExample network) rnws (exs network)

train :: Int -> Network -> [[[Weight]]] -> [[[Weight]]]
train    count  network    rnws         =  let newNWs = trainEpic network rnws
                                           in  if count <= 0 || rnws == newNWs then rnws else train (count - 1) network newNWs

trainUntil :: Network -> [[[Weight]]] -> [[[Weight]]]
trainUntil    network    rnws         =  until (passRNetWs network) (trainEpic network) rnws

trainNew :: Network -> [[[Weight]]]
trainNew    network = trainLoop network (revInitialWeights network)

trainLoop :: Network -> [[[Weight]]] -> [[[Weight]]]
trainLoop    network    rnws         =  let newRnws = train 5000 network rnws
                                        in  if passRNetWs network newRnws then 
                                            newRnws 
                                        else 
                                            trainLoop network newRnws

-----------------------------------------------------------------------------------------------------------------
