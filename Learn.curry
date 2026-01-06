module Learn (initialWeights, learn, learnEpic, learnExample, deltaOut, deltaNeuron, deltaInner, backPropagate, gradientDescent) where
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
deltaNeuron    pos    lws           deltas  =  let ws = map (!! pos) lws
                                               in  dotProduct ws deltas

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

--deltaInner :: Layer -> [[Weight]] -> [Delta] -> [Output] -> [Delta]
--deltaInner    _        _             _          []       =  []
--deltaInner    layer    (ws:lws)      deltas     (o:os)   =
--    let algo = activation_der (algorithm layer)
--    in  (deltaNeuron ws deltas) * (algo o) : deltaInner layer lws deltas os
--        where deltaNeuron _       []     = 0
--              deltaNeuron (w:ws') (d:ds) = w * d + deltaNeuron ws' ds

--backPropagate :: [Layer] -> [Output] -> [[[Weight]]] -> [[Output]] -> [[Delta]]
--backPropagate    (rl:rls)   exs         rnws            (ol:ols)   =  backProp rls rnws (deltaOut rl exs ol) ols
--    where backProp []         _           delta _            = [delta]
--          backProp (rl':rls') (lws:rnws') delta (outl:outls) = let deltaLower = deltaInner rl' lws delta outl
--                                                               in  delta : backProp rls' rnws' deltaLower outls

adjustNeuron :: Float -> [Weight] -> [Input] -> Delta -> [Weight]
adjustNeuron    lr       [bias]      []         delta =  [bias - lr * delta]
adjustNeuron    lr       (w:ws)      (i:is)     delta =  w - lr * delta * i : adjustNeuron lr ws is delta

adjustLayer :: Float -> [[Weight]] -> [Input] -> [Delta] -> [[Weight]]
adjustLayer    lr       (nws:lws)     is          (d:ds) =  let ws = adjustNeuron lr nws is d
                                                            in if lws == [] then
                                                                [ws]
                                                            else
                                                                ws : adjustLayer lr lws is ds

gradientDescent :: Float -> [[[Weight]]] -> [[Input]] -> [[Delta]] -> [[[Weight]]]
gradientDescent    lr       (lws:nws)       (_:is:lis)   (ds:lds)  =  let ws = adjustLayer lr lws is ds
                                                                      in if nws == [] then
                                                                          [ws]
                                                                      else
                                                                          ws : gradientDescent lr nws (is:lis) lds

learnExample :: Network -> [[[Weight]]] -> Example -> [[[Weight]]]
learnExample    network    rnws            ex      =  let ns     = feedForward network (inputs ex) rnws
                                                          rls    = reverse (layers network)
                                                          lr     = learningRate network
                                                          deltas = backPropagate rls (expectation ex) rnws ns
                                                      in  gradientDescent lr rnws ns deltas

learnExampleNetwork :: (Network -> [[[Weight]]] -> Example -> [[[Weight]]]) -> Network -> ([[[Weight]]] -> Example -> [[[Weight]]])
learnExampleNetwork    learningFunction                                        network =  learningFunction network

learnEpic :: Network -> [[[Weight]]] -> [[[Weight]]]
learnEpic    network    rnws         =  let lf = learnExampleNetwork learnExample network
                                        in  foldl lf rnws (exs network)

learn :: Int -> Network -> [[[Weight]]] -> [[[Weight]]]
learn    count  network    rnws         =  let newNWs = learnEpic network rnws
                                           in  if count <= 0 then rnws else learn (count - 1) network newNWs

initialWeights :: Network -> Int -> [[[Weight]]]
initialWeights    network    ins =  initLWs ins (layers network)
    where initLWs _    []             = []
          initLWs ins' (layer:layers) = let ns = (nodes layer)
                                        in initNWs ins' ns : initLWs ns layers
              where iw = initWeights (algorithm layer)
                    initNWs ins'' ns' = if ns' <= 0 then [] else initNWs' ins'' ns'
                    where initNWs' ins''' ns'' = initWs ins''' : initNWs ins''' (ns'' - 1)
                        where initWs ins'''' = replicate (ins''''+1) (iw ins'''' (nodes layer))
