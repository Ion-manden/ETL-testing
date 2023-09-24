{-# LANGUAGE OverloadedStrings #-}
{-# LANGUAGE StrictData #-}

module Main where

import Data.Aeson (FromJSON, ToJSON, decode, encode)
import qualified Data.ByteString.Lazy.Char8 as B
import Data.Text (Text)
import Data.Vector as V

import Data.Vector.Fusion.Bundle (flatten)
import GHC.Generics
import GHC.IO.Handle (Handle, hGetLine)
import GHC.IO.Handle.FD (withFile)
import GHC.IO.IOMode (IOMode (ReadMode))
import Product (Price, Product, pricesProduct, countryPrice)
import System.Environment
import System.IO

import qualified Data.Map.Strict as Map

getProductPrices :: Maybe Product -> V.Vector Price
getProductPrices (Just product) = pricesProduct product
getProductPrices _ = V.fromList []

parseJSONString :: String -> Maybe Product
parseJSONString jsonString = decode (B.pack jsonString)


countPricesByCountry :: V.Vector Price -> Map.Map String Int
countPricesByCountry prices = V.foldl' (\acc price -> Map.insertWith (+) (countryPrice price) 1 acc) Map.empty prices

processLine :: String -> IO (Map.Map String Int)
processLine line = do
  let product = parseJSONString line
  let prices = getProductPrices product
  let result = countPricesByCountry prices
  return result


processLines :: System.IO.Handle -> Map.Map String Int -> IO (Map.Map String Int)
processLines handle acc = do
  eof <- hIsEOF handle
  if eof
    then return acc
    else do
      line <- hGetLine handle
      countryCountMap <- processLine line

      let updatedMap = Map.unionWith (+) acc countryCountMap

      updatedMap `seq` processLines handle updatedMap


computeResult :: FilePath -> IO (Map.Map String Int)
computeResult file = do
  System.IO.withFile file ReadMode $ \handle -> 
    processLines handle Map.empty


main :: IO ()
main = do
  result <- computeResult "../../../data-generator/product-data.ndjson"

  let output = encode result
  B.putStrLn output

