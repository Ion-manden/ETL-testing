{-# LANGUAGE StrictData #-}
{-# LANGUAGE OverloadedStrings #-}

module Product
    ( Product (..)
    , Price (..)
    , decodeTopLevel
    ) where

import Data.Aeson
import Data.Aeson.Types (emptyObject)
import Data.ByteString.Lazy (ByteString)
import Data.HashMap.Strict (HashMap)
import Data.Text (Text)
import Data.Vector (Vector)

data Product = Product
    { productIDProduct :: Integer
    , nameProduct :: Text
    , descriptionProduct :: Text
    , pricesProduct :: Vector Price
    , createdProduct :: Text
    , createdFormatProduct :: Text
    } deriving (Show)

data Price = Price
    { countryPrice :: String
    , pricePrice :: Float
    } deriving (Show)

decodeTopLevel :: ByteString -> Maybe Product
decodeTopLevel = decode

instance ToJSON Product where
    toJSON (Product productIDProduct nameProduct descriptionProduct pricesProduct createdProduct createdFormatProduct) =
        object
        [ "Id" .= productIDProduct
        , "Name" .= nameProduct
        , "Description" .= descriptionProduct
        , "Prices" .= pricesProduct
        , "Created" .= createdProduct
        , "CreatedFormat" .= createdFormatProduct
        ]

instance FromJSON Product where
    parseJSON (Object v) = Product
        <$> v .: "Id"
        <*> v .: "Name"
        <*> v .: "Description"
        <*> v .: "Prices"
        <*> v .: "Created"
        <*> v .: "CreatedFormat"

instance ToJSON Price where
    toJSON (Price countryPrice pricePrice) =
        object
        [ "Country" .= countryPrice
        , "Price" .= pricePrice
        ]

instance FromJSON Price where
    parseJSON (Object v) = Price
        <$> v .: "Country"
        <*> v .: "Price"
