module Train (trailingExpected, precedingExpected, train) where
import Learn

trailingExpected :: [Float] -> [Float] -> [Float]
trailingExpected [] _                       = failed
trailingExpected ds ws | is ++ [exp] =:= ds = learn ws is exp
     where is, exp free

precedingExpected :: [Float] -> [Float] -> [Float]
precedingExpected []       _  = failed
precedingExpected (exp:is) ws = learn ws is exp

epoch :: ([Float] -> [Float] -> [Float]) -> [[Float]] -> Int -> [Float] -> [Float]
epoch expected dataset _ ws = foldr expected ws dataset

train :: ([Float] -> [Float] -> [Float]) -> [[Float]] -> [Int] -> [Float]
train _ [] _  = failed
train expected (input:inputs) epochList = foldr (epoch expected (input:inputs)) (initialWeights input) epochList
