module Learn (initialWeights, revInitialWeights, learnExample, backPropagate, gradientDescent) where
import Algorithms
import Network
import LogicExamples
import Perceptron

type Delta = Float

deltaOut :: Layer -> [Output] -> [Output]   -> [Delta]
deltaOut    _        []          []         =  []
deltaOut    layer    (ex:exs)    (out:outs) =  let act = activation_der (algorithm layer)
                                               in  (out - ex) * (act out) : deltaOut layer exs outs

deltaNeuron :: Int -> [[Weight]] -> [Delta] -> Delta
deltaNeuron    pos    lws           deltas  =  dotProduct (map (!! pos) lws) deltas

deltaInner :: Int -> Layer -> [[Weight]] -> [Delta] -> [Output] -> [Delta]
deltaInner    _      _        _             _          []       = []
deltaInner    index  layer    lws           deltas     (o:os)   = let algo = activation_der (algorithm layer)
                                                                      d = (deltaNeuron index lws deltas) * algo o
                                                                  in  d : deltaInner (index+1) layer lws deltas os

backPropagate :: [Layer] -> [Output] -> [[[Weight]]] -> [[Output]] -> [[Delta]]
backPropagate    (rl:rls)   exs         rnws            (ol:ols)   =  backProp rls rnws (deltaOut rl exs ol) ols
    where backProp []         _           deltas _            = [deltas]
          backProp (rl':rls') (lws:rnws') deltas (outl:outls) = let deltasLower = deltaInner 0 rl' lws deltas outl
                                                                in  deltas : backProp rls' rnws' deltasLower outls

adjustNeuron :: Float -> [Weight] -> [Input] -> Delta -> [Weight]
adjustNeuron    lr       [bias]      []         delta =  [bias - lr * delta]
adjustNeuron    lr       (w:ws)      (i:is)     delta =  w - lr * delta * i : adjustNeuron lr ws is delta

adjustLayer :: Float -> [[Weight]] -> [Input] -> [Delta] -> [[Weight]]
adjustLayer    _        []             _         _       =  []
adjustLayer    lr       (nws:lws)     is         (d:ds)  =  let ws = adjustNeuron lr nws is d
                                                            in ws : adjustLayer lr lws is ds

gradientDescent :: Float -> [[[Weight]]] -> [[Input]] -> [[Delta]] -> [[[Weight]]]
gradientDescent    _        []              _            _         =  []
gradientDescent    lr       (lws:nws)       (_:is:lis)   (ds:lds)  =  let ws = adjustLayer lr lws is ds
                                                                      in ws : gradientDescent lr nws (is:lis) lds

learnExample :: Network -> [[[Weight]]] -> Example -> [[[Weight]]]
learnExample    network    rnws            ex      =  let ns     = feedForward network (inputs ex) rnws
                                                          rls    = reverse (layers network)
                                                          lr     = learningRate network
                                                          deltas = backPropagate rls (expectation ex) rnws ns
                                                      in  gradientDescent lr rnws ns deltas

revInitialWeights :: Network -> [[[Weight]]]
revInitialWeights    network =  reverse (initialWeights network (length (inputs (head (exs network)))))

initialWeights :: Network -> Int -> [[[Weight]]]
initialWeights    network    ins =  initLWs ins (layers network)
    where initLWs _    []             = []
          initLWs ins' (layer:layers) = let ns = (nodes layer)
                                        in initNWs ins' ns : initLWs ns layers
              where iw = initWeights (algorithm layer)
                    initNWs ins'' ns' = if ns' <= 0 then [] else initNWs' ins'' ns'
                    where initNWs' ins''' ns'' = initWs ins''' : initNWs ins''' (ns'' - 1)
                        where initWs ins'''' = replicate (ins''''+1) (iw ins'''' (nodes layer))


---------------------------------------------------------------------------------------
-- Tests

testDeltaOut :: [Float]
testDeltaOut =  deltaOut (head (layers (andNetwork))) (expectation (head (exs (andNetwork)))) [1]

testDeltaOut2 :: [Float] -> [Float]
testDeltaOut2 outs =  deltaOut (head (layers (andNetwork))) [1, 1] outs

testDeltaNeuron :: Bool
testDeltaNeuron = deltaNeuron 0 [[1,2,3,4], [5,6,7,8]] [2,3] == 17

testLayer :: Layer
testLayer = Layer { nodes = 3, algorithm = classifier }

testDeltaInner :: Bool
testDeltaInner =  deltaInner 0 testLayer [[1,2,3,4], [5,6,7,8]] [2,3] [10,20,30] == [17, 22, 27]

