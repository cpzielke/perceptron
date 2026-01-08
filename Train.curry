module Train (trainEpic, train) where

import Network
import Learn

trainEpic :: Network -> [[[Weight]]] -> [[[Weight]]]
trainEpic    network    rnws         =  foldl (learnExample network) rnws (exs network)

train :: Int -> Network -> [[[Weight]]] -> [[[Weight]]]
train    count  network    rnws         =  let newNWs = trainEpic network rnws
                                           in  if count <= 0 then rnws else train (count - 1) network newNWs



-----------------------------------------------------------------------------------------------------------------
