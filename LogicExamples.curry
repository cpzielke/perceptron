module LogicExamples (andNetwork, orNetwork, notNetwork, nandNetwork, xorNetwork, xorplusNetwork) where
import Network
import Algorithms

lEARNING_RATE :: Float -- Lower is slower
lEARNING_RATE =  0.1

-- AND
andExamples :: [Example]
andExamples = [
    Example { name = "FF", inputs = [0, 0], expectation = [0] },
    Example { name = "FT", inputs = [0, 1], expectation = [0] },
    Example { name = "TF", inputs = [1, 0], expectation = [0] },
    Example { name = "TT", inputs = [1, 1], expectation = [1] }
  ]

andLayers :: [Layer]
andLayers = [
    Layer { nodes = 1, algorithm = classifier }
  ]

andWeights :: [[[Float]]]
andWeights = [[[1, 1, -1.5]]]

andNetwork :: Network
andNetwork = Network { weights = andWeights, exs = andExamples, layers = andLayers, learningRate = lEARNING_RATE }

-- OR
orExamples :: [Example]
orExamples = [
    Example { name = "FF", inputs = [0, 0], expectation = [0] },
    Example { name = "FT", inputs = [0, 1], expectation = [1] },
    Example { name = "TF", inputs = [1, 0], expectation = [1] },
    Example { name = "TT", inputs = [1, 1], expectation = [1] }
  ]

orLayers :: [Layer]
orLayers = [
    Layer { nodes = 1, algorithm = classifier }
  ]

orWeights :: [[[Float]]]
orWeights = [[[1, 1, -0.5]]]

orNetwork :: Network
orNetwork = Network { weights = orWeights, exs = orExamples, layers = orLayers, learningRate = lEARNING_RATE }

-- NOT
notExamples :: [Example]
notExamples = [
    Example { name = "F", inputs = [0], expectation = [1] },
    Example { name = "T", inputs = [1], expectation = [0] }
  ]

notLayers :: [Layer]
notLayers = [
    Layer { nodes = 1, algorithm = classifier }
  ]

notWeights :: [[[Float]]]
notWeights = [[[-1, 0.5]]]

notNetwork :: Network
notNetwork = Network { weights = notWeights, exs = notExamples, layers = notLayers, learningRate = lEARNING_RATE }

-- NAND
nandExamples :: [Example]
nandExamples = [
    Example { name = "FF", inputs = [0, 0], expectation = [1] },
    Example { name = "FT", inputs = [0, 1], expectation = [1] },
    Example { name = "TF", inputs = [1, 0], expectation = [1] },
    Example { name = "TT", inputs = [1, 1], expectation = [0] }
  ]

nandLayers :: [Layer]
nandLayers = [
    Layer { nodes = 1, algorithm = logistic },
    Layer { nodes = 1, algorithm = classifier }
  ]

nandWeights :: [[[Float]]]
nandWeights = [
    [
        [1, 1, -1.5]
    ],
    [
        [-1, 0.5]
    ]
  ]

nandNetwork :: Network
nandNetwork = Network { weights = nandWeights, exs = nandExamples, layers = nandLayers, learningRate = lEARNING_RATE }


-- XOR
xorExamples :: [Example]
xorExamples = [
    Example { name = "FF", inputs = [0, 0], expectation = [0] },
    Example { name = "FT", inputs = [0, 1], expectation = [1] },
    Example { name = "TF", inputs = [1, 0], expectation = [1] },
    Example { name = "TT", inputs = [1, 1], expectation = [0] }
  ]

xorLayers :: [Layer]
xorLayers = [
    Layer { nodes = 2, algorithm = logistic },
    Layer { nodes = 1, algorithm = classifier }
  ]

xorWeights :: [[[Float]]]
xorWeights = [[[-1.4371652540903184, -1.7011433422421145, -1.1245228358001023], [0.8418978570520899, 0.7372312749809187, -1.7623473531669684]], [[-1.1349023507532465, -0.8502081202751894, 0.4]]]

xorNetwork :: Network
xorNetwork = Network { weights = xorWeights, exs = xorExamples, layers = xorLayers, learningRate = lEARNING_RATE }

-- XORPLUS
xorplusExamples :: [Example]
xorplusExamples = [
    Example { name = "FF", inputs = [0, 0], expectation = [0, 0] },
    Example { name = "FT", inputs = [0, 1], expectation = [1, 0] },
    Example { name = "TF", inputs = [1, 0], expectation = [1, 1] },
    Example { name = "TT", inputs = [1, 1], expectation = [0, 0] }
  ]

xorplusLayers :: [Layer]
xorplusLayers = [
    Layer { nodes = 3, algorithm = logistic },
    Layer { nodes = 2, algorithm = classifier }
  ]

xorplusWeights :: [[[Float]]]
xorplusWeights = [[[]]]

xorplusNetwork :: Network
xorplusNetwork = Network { weights = xorplusWeights, exs = xorplusExamples, layers = xorplusLayers, learningRate = lEARNING_RATE }
