module Perceptron (eval, eval2, dotProduct, networkDotProduct, feedForward) where
import Network
import Algorithms

dotProduct :: [Float] -> [Float]  -> Output
dotProduct    (_:_:_)    []       =  failed
dotProduct    []         (_:_:_)  =  failed
dotProduct    []         []       =  0
dotProduct    []         [bias]   =  bias
dotProduct    (i:is)     (w:ws)   =  i * w + dotProduct is ws

layerDotProduct :: Act -> [Input] -> [[Weight]] -> [Output]
layerDotProduct    _      _          []         =  []
layerDotProduct    act    is         (ws:lws)   =  act (dotProduct is ws) : layerDotProduct act is lws

networkDotProduct :: Network -> [Input] -> [[[Weight]]] -> [[Output]]
networkDotProduct    network    is         nws          =  is : accum (layers network) is nws
    where accum (_:_)  _   []         = failed
          accum []     _   _          = []
          accum (l:ls) is' (lws:nws') = let act = activation (algorithm l)
                                            nextis = layerDotProduct act is' lws
                                        in  nextis : accum ls nextis nws'

feedForward :: Network -> [Input] -> [[[Weight]]] -> [[Output]]
feedForward    network    is         rnws         =  reverse (networkDotProduct network is (reverse rnws))

eval :: Network -> [Input] -> [Output]
eval    network    is      =  let (out:_) = reverse (networkDotProduct network is (weights network))
                              in out

eval2 :: Network -> [[[Weight]]] -> [Input] -> [Output]
eval2    network    nws             is      =  let (out:_) = reverse (networkDotProduct network is nws)
                                               in out
