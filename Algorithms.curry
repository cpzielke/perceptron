module Algorithms (Act, logistic, classifier,
                   Algorithm, Algorithm (activation), Algorithm (activation_der), Algorithm (initWeights)
) where

type Act = Float -> Float

data Algorithm = Algorithm { activation :: Act, activation_der :: Act, initWeights :: (Int -> Int -> Float) }

e :: Float
e = 2.718281828459045

classifier_act :: Act
classifier_act    value = if value >= 0.0 then 1.0 else 0.0

classifier_der :: Act
classifier_der    _     = 1

classifier_iw :: Int -> Int     -> Float
classifier_iw    inputs outputs =  scaling (sqrt (7 / ((toFloat inputs) + (toFloat outputs))))

classifier :: Algorithm
classifier =  Algorithm { activation = classifier_act, activation_der = classifier_der, initWeights = classifier_iw }


logistic_act :: Act
logistic_act    value =  1 / (1 + e ** (-value))

logistic_der :: Act
logistic_der    value =  value - value ** 2

logistic_iw :: Int -> Int     -> Float
logistic_iw    inputs outputs =  scaling (sqrt (6 / ((toFloat inputs) + (toFloat outputs))))

logistic :: Algorithm
logistic =  Algorithm { activation = logistic_act, activation_der = logistic_der, initWeights = logistic_iw }

scaling :: Float -> Float
scaling    value =  (((2 * value) * 10.0 + 1) - value * 10.0) / 50.0