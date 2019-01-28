// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package imagesearch

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/imagesearch"

const (
	DefaultEndpoint = original.DefaultEndpoint
)

type Currency = original.Currency

const (
	AED Currency = original.AED
	AFN Currency = original.AFN
	ALL Currency = original.ALL
	AMD Currency = original.AMD
	ANG Currency = original.ANG
	AOA Currency = original.AOA
	ARS Currency = original.ARS
	AUD Currency = original.AUD
	AWG Currency = original.AWG
	AZN Currency = original.AZN
	BAM Currency = original.BAM
	BBD Currency = original.BBD
	BDT Currency = original.BDT
	BGN Currency = original.BGN
	BHD Currency = original.BHD
	BIF Currency = original.BIF
	BMD Currency = original.BMD
	BND Currency = original.BND
	BOB Currency = original.BOB
	BOV Currency = original.BOV
	BRL Currency = original.BRL
	BSD Currency = original.BSD
	BTN Currency = original.BTN
	BWP Currency = original.BWP
	BYR Currency = original.BYR
	BZD Currency = original.BZD
	CAD Currency = original.CAD
	CDF Currency = original.CDF
	CHE Currency = original.CHE
	CHF Currency = original.CHF
	CHW Currency = original.CHW
	CLF Currency = original.CLF
	CLP Currency = original.CLP
	CNY Currency = original.CNY
	COP Currency = original.COP
	COU Currency = original.COU
	CRC Currency = original.CRC
	CUC Currency = original.CUC
	CUP Currency = original.CUP
	CVE Currency = original.CVE
	CZK Currency = original.CZK
	DJF Currency = original.DJF
	DKK Currency = original.DKK
	DOP Currency = original.DOP
	DZD Currency = original.DZD
	EGP Currency = original.EGP
	ERN Currency = original.ERN
	ETB Currency = original.ETB
	EUR Currency = original.EUR
	FJD Currency = original.FJD
	FKP Currency = original.FKP
	GBP Currency = original.GBP
	GEL Currency = original.GEL
	GHS Currency = original.GHS
	GIP Currency = original.GIP
	GMD Currency = original.GMD
	GNF Currency = original.GNF
	GTQ Currency = original.GTQ
	GYD Currency = original.GYD
	HKD Currency = original.HKD
	HNL Currency = original.HNL
	HRK Currency = original.HRK
	HTG Currency = original.HTG
	HUF Currency = original.HUF
	IDR Currency = original.IDR
	ILS Currency = original.ILS
	INR Currency = original.INR
	IQD Currency = original.IQD
	IRR Currency = original.IRR
	ISK Currency = original.ISK
	JMD Currency = original.JMD
	JOD Currency = original.JOD
	JPY Currency = original.JPY
	KES Currency = original.KES
	KGS Currency = original.KGS
	KHR Currency = original.KHR
	KMF Currency = original.KMF
	KPW Currency = original.KPW
	KRW Currency = original.KRW
	KWD Currency = original.KWD
	KYD Currency = original.KYD
	KZT Currency = original.KZT
	LAK Currency = original.LAK
	LBP Currency = original.LBP
	LKR Currency = original.LKR
	LRD Currency = original.LRD
	LSL Currency = original.LSL
	LYD Currency = original.LYD
	MAD Currency = original.MAD
	MDL Currency = original.MDL
	MGA Currency = original.MGA
	MKD Currency = original.MKD
	MMK Currency = original.MMK
	MNT Currency = original.MNT
	MOP Currency = original.MOP
	MRO Currency = original.MRO
	MUR Currency = original.MUR
	MVR Currency = original.MVR
	MWK Currency = original.MWK
	MXN Currency = original.MXN
	MXV Currency = original.MXV
	MYR Currency = original.MYR
	MZN Currency = original.MZN
	NAD Currency = original.NAD
	NGN Currency = original.NGN
	NIO Currency = original.NIO
	NOK Currency = original.NOK
	NPR Currency = original.NPR
	NZD Currency = original.NZD
	OMR Currency = original.OMR
	PAB Currency = original.PAB
	PEN Currency = original.PEN
	PGK Currency = original.PGK
	PHP Currency = original.PHP
	PKR Currency = original.PKR
	PLN Currency = original.PLN
	PYG Currency = original.PYG
	QAR Currency = original.QAR
	RON Currency = original.RON
	RSD Currency = original.RSD
	RUB Currency = original.RUB
	RWF Currency = original.RWF
	SAR Currency = original.SAR
	SBD Currency = original.SBD
	SCR Currency = original.SCR
	SDG Currency = original.SDG
	SEK Currency = original.SEK
	SGD Currency = original.SGD
	SHP Currency = original.SHP
	SLL Currency = original.SLL
	SOS Currency = original.SOS
	SRD Currency = original.SRD
	SSP Currency = original.SSP
	STD Currency = original.STD
	SYP Currency = original.SYP
	SZL Currency = original.SZL
	THB Currency = original.THB
	TJS Currency = original.TJS
	TMT Currency = original.TMT
	TND Currency = original.TND
	TOP Currency = original.TOP
	TRY Currency = original.TRY
	TTD Currency = original.TTD
	TWD Currency = original.TWD
	TZS Currency = original.TZS
	UAH Currency = original.UAH
	UGX Currency = original.UGX
	USD Currency = original.USD
	UYU Currency = original.UYU
	UZS Currency = original.UZS
	VEF Currency = original.VEF
	VND Currency = original.VND
	VUV Currency = original.VUV
	WST Currency = original.WST
	XAF Currency = original.XAF
	XCD Currency = original.XCD
	XOF Currency = original.XOF
	XPF Currency = original.XPF
	YER Currency = original.YER
	ZAR Currency = original.ZAR
	ZMW Currency = original.ZMW
)

type ErrorCode = original.ErrorCode

const (
	InsufficientAuthorization ErrorCode = original.InsufficientAuthorization
	InvalidAuthorization      ErrorCode = original.InvalidAuthorization
	InvalidRequest            ErrorCode = original.InvalidRequest
	None                      ErrorCode = original.None
	RateLimitExceeded         ErrorCode = original.RateLimitExceeded
	ServerError               ErrorCode = original.ServerError
)

type ErrorSubCode = original.ErrorSubCode

const (
	AuthorizationDisabled   ErrorSubCode = original.AuthorizationDisabled
	AuthorizationExpired    ErrorSubCode = original.AuthorizationExpired
	AuthorizationMissing    ErrorSubCode = original.AuthorizationMissing
	AuthorizationRedundancy ErrorSubCode = original.AuthorizationRedundancy
	Blocked                 ErrorSubCode = original.Blocked
	HTTPNotAllowed          ErrorSubCode = original.HTTPNotAllowed
	NotImplemented          ErrorSubCode = original.NotImplemented
	ParameterInvalidValue   ErrorSubCode = original.ParameterInvalidValue
	ParameterMissing        ErrorSubCode = original.ParameterMissing
	ResourceError           ErrorSubCode = original.ResourceError
	UnexpectedError         ErrorSubCode = original.UnexpectedError
)

type Freshness = original.Freshness

const (
	Day   Freshness = original.Day
	Month Freshness = original.Month
	Week  Freshness = original.Week
)

type ImageAspect = original.ImageAspect

const (
	All    ImageAspect = original.All
	Square ImageAspect = original.Square
	Tall   ImageAspect = original.Tall
	Wide   ImageAspect = original.Wide
)

type ImageColor = original.ImageColor

const (
	Black      ImageColor = original.Black
	Blue       ImageColor = original.Blue
	Brown      ImageColor = original.Brown
	ColorOnly  ImageColor = original.ColorOnly
	Gray       ImageColor = original.Gray
	Green      ImageColor = original.Green
	Monochrome ImageColor = original.Monochrome
	Orange     ImageColor = original.Orange
	Pink       ImageColor = original.Pink
	Purple     ImageColor = original.Purple
	Red        ImageColor = original.Red
	Teal       ImageColor = original.Teal
	White      ImageColor = original.White
	Yellow     ImageColor = original.Yellow
)

type ImageContent = original.ImageContent

const (
	Face     ImageContent = original.Face
	Portrait ImageContent = original.Portrait
)

type ImageCropType = original.ImageCropType

const (
	Rectangular ImageCropType = original.Rectangular
)

type ImageInsightModule = original.ImageInsightModule

const (
	ImageInsightModuleAll                ImageInsightModule = original.ImageInsightModuleAll
	ImageInsightModuleBRQ                ImageInsightModule = original.ImageInsightModuleBRQ
	ImageInsightModuleCaption            ImageInsightModule = original.ImageInsightModuleCaption
	ImageInsightModuleCollections        ImageInsightModule = original.ImageInsightModuleCollections
	ImageInsightModulePagesIncluding     ImageInsightModule = original.ImageInsightModulePagesIncluding
	ImageInsightModuleRecipes            ImageInsightModule = original.ImageInsightModuleRecipes
	ImageInsightModuleRecognizedEntities ImageInsightModule = original.ImageInsightModuleRecognizedEntities
	ImageInsightModuleRelatedSearches    ImageInsightModule = original.ImageInsightModuleRelatedSearches
	ImageInsightModuleShoppingSources    ImageInsightModule = original.ImageInsightModuleShoppingSources
	ImageInsightModuleSimilarImages      ImageInsightModule = original.ImageInsightModuleSimilarImages
	ImageInsightModuleSimilarProducts    ImageInsightModule = original.ImageInsightModuleSimilarProducts
	ImageInsightModuleTags               ImageInsightModule = original.ImageInsightModuleTags
)

type ImageLicense = original.ImageLicense

const (
	ImageLicenseAll                ImageLicense = original.ImageLicenseAll
	ImageLicenseAny                ImageLicense = original.ImageLicenseAny
	ImageLicenseModify             ImageLicense = original.ImageLicenseModify
	ImageLicenseModifyCommercially ImageLicense = original.ImageLicenseModifyCommercially
	ImageLicensePublic             ImageLicense = original.ImageLicensePublic
	ImageLicenseShare              ImageLicense = original.ImageLicenseShare
	ImageLicenseShareCommercially  ImageLicense = original.ImageLicenseShareCommercially
)

type ImageSize = original.ImageSize

const (
	ImageSizeAll       ImageSize = original.ImageSizeAll
	ImageSizeLarge     ImageSize = original.ImageSizeLarge
	ImageSizeMedium    ImageSize = original.ImageSizeMedium
	ImageSizeSmall     ImageSize = original.ImageSizeSmall
	ImageSizeWallpaper ImageSize = original.ImageSizeWallpaper
)

type ImageType = original.ImageType

const (
	AnimatedGif ImageType = original.AnimatedGif
	Clipart     ImageType = original.Clipart
	Line        ImageType = original.Line
	Photo       ImageType = original.Photo
	Shopping    ImageType = original.Shopping
	Transparent ImageType = original.Transparent
)

type ItemAvailability = original.ItemAvailability

const (
	Discontinued        ItemAvailability = original.Discontinued
	InStock             ItemAvailability = original.InStock
	InStoreOnly         ItemAvailability = original.InStoreOnly
	LimitedAvailability ItemAvailability = original.LimitedAvailability
	OnlineOnly          ItemAvailability = original.OnlineOnly
	OutOfStock          ItemAvailability = original.OutOfStock
	PreOrder            ItemAvailability = original.PreOrder
	SoldOut             ItemAvailability = original.SoldOut
)

type SafeSearch = original.SafeSearch

const (
	Moderate SafeSearch = original.Moderate
	Off      SafeSearch = original.Off
	Strict   SafeSearch = original.Strict
)

type Type = original.Type

const (
	TypeAggregateRating Type = original.TypeAggregateRating
	TypePropertiesItem  Type = original.TypePropertiesItem
	TypeRating          Type = original.TypeRating
)

type TypeBasicResponseBase = original.TypeBasicResponseBase

const (
	TypeAggregateOffer         TypeBasicResponseBase = original.TypeAggregateOffer
	TypeAnswer                 TypeBasicResponseBase = original.TypeAnswer
	TypeCollectionPage         TypeBasicResponseBase = original.TypeCollectionPage
	TypeCreativeWork           TypeBasicResponseBase = original.TypeCreativeWork
	TypeErrorResponse          TypeBasicResponseBase = original.TypeErrorResponse
	TypeIdentifiable           TypeBasicResponseBase = original.TypeIdentifiable
	TypeImageGallery           TypeBasicResponseBase = original.TypeImageGallery
	TypeImageInsights          TypeBasicResponseBase = original.TypeImageInsights
	TypeImageObject            TypeBasicResponseBase = original.TypeImageObject
	TypeImages                 TypeBasicResponseBase = original.TypeImages
	TypeIntangible             TypeBasicResponseBase = original.TypeIntangible
	TypeMediaObject            TypeBasicResponseBase = original.TypeMediaObject
	TypeNormalizedRectangle    TypeBasicResponseBase = original.TypeNormalizedRectangle
	TypeOffer                  TypeBasicResponseBase = original.TypeOffer
	TypeOrganization           TypeBasicResponseBase = original.TypeOrganization
	TypePerson                 TypeBasicResponseBase = original.TypePerson
	TypeRecipe                 TypeBasicResponseBase = original.TypeRecipe
	TypeRecognizedEntity       TypeBasicResponseBase = original.TypeRecognizedEntity
	TypeRecognizedEntityRegion TypeBasicResponseBase = original.TypeRecognizedEntityRegion
	TypeResponse               TypeBasicResponseBase = original.TypeResponse
	TypeResponseBase           TypeBasicResponseBase = original.TypeResponseBase
	TypeSearchResultsAnswer    TypeBasicResponseBase = original.TypeSearchResultsAnswer
	TypeStructuredValue        TypeBasicResponseBase = original.TypeStructuredValue
	TypeThing                  TypeBasicResponseBase = original.TypeThing
	TypeTrendingImages         TypeBasicResponseBase = original.TypeTrendingImages
	TypeWebPage                TypeBasicResponseBase = original.TypeWebPage
)

type AggregateOffer = original.AggregateOffer
type AggregateRating = original.AggregateRating
type Answer = original.Answer
type BaseClient = original.BaseClient
type BasicAnswer = original.BasicAnswer
type BasicCollectionPage = original.BasicCollectionPage
type BasicCreativeWork = original.BasicCreativeWork
type BasicIdentifiable = original.BasicIdentifiable
type BasicIntangible = original.BasicIntangible
type BasicMediaObject = original.BasicMediaObject
type BasicOffer = original.BasicOffer
type BasicPropertiesItem = original.BasicPropertiesItem
type BasicRating = original.BasicRating
type BasicResponse = original.BasicResponse
type BasicResponseBase = original.BasicResponseBase
type BasicSearchResultsAnswer = original.BasicSearchResultsAnswer
type BasicStructuredValue = original.BasicStructuredValue
type BasicThing = original.BasicThing
type BasicWebPage = original.BasicWebPage
type CollectionPage = original.CollectionPage
type CreativeWork = original.CreativeWork
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type Identifiable = original.Identifiable
type ImageGallery = original.ImageGallery
type ImageInsights = original.ImageInsights
type ImageInsightsImageCaption = original.ImageInsightsImageCaption
type ImageObject = original.ImageObject
type ImageTagsModule = original.ImageTagsModule
type Images = original.Images
type ImagesClient = original.ImagesClient
type ImagesImageMetadata = original.ImagesImageMetadata
type ImagesModule = original.ImagesModule
type InsightsTag = original.InsightsTag
type Intangible = original.Intangible
type MediaObject = original.MediaObject
type NormalizedRectangle = original.NormalizedRectangle
type Offer = original.Offer
type Organization = original.Organization
type Person = original.Person
type PivotSuggestions = original.PivotSuggestions
type PropertiesItem = original.PropertiesItem
type Query = original.Query
type Rating = original.Rating
type Recipe = original.Recipe
type RecipesModule = original.RecipesModule
type RecognizedEntitiesModule = original.RecognizedEntitiesModule
type RecognizedEntity = original.RecognizedEntity
type RecognizedEntityGroup = original.RecognizedEntityGroup
type RecognizedEntityRegion = original.RecognizedEntityRegion
type RelatedCollectionsModule = original.RelatedCollectionsModule
type RelatedSearchesModule = original.RelatedSearchesModule
type Response = original.Response
type ResponseBase = original.ResponseBase
type SearchResultsAnswer = original.SearchResultsAnswer
type StructuredValue = original.StructuredValue
type Thing = original.Thing
type TrendingImages = original.TrendingImages
type TrendingImagesCategory = original.TrendingImagesCategory
type TrendingImagesTile = original.TrendingImagesTile
type WebPage = original.WebPage

func New() BaseClient {
	return original.New()
}
func NewImagesClient() ImagesClient {
	return original.NewImagesClient()
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleCurrencyValues() []Currency {
	return original.PossibleCurrencyValues()
}
func PossibleErrorCodeValues() []ErrorCode {
	return original.PossibleErrorCodeValues()
}
func PossibleErrorSubCodeValues() []ErrorSubCode {
	return original.PossibleErrorSubCodeValues()
}
func PossibleFreshnessValues() []Freshness {
	return original.PossibleFreshnessValues()
}
func PossibleImageAspectValues() []ImageAspect {
	return original.PossibleImageAspectValues()
}
func PossibleImageColorValues() []ImageColor {
	return original.PossibleImageColorValues()
}
func PossibleImageContentValues() []ImageContent {
	return original.PossibleImageContentValues()
}
func PossibleImageCropTypeValues() []ImageCropType {
	return original.PossibleImageCropTypeValues()
}
func PossibleImageInsightModuleValues() []ImageInsightModule {
	return original.PossibleImageInsightModuleValues()
}
func PossibleImageLicenseValues() []ImageLicense {
	return original.PossibleImageLicenseValues()
}
func PossibleImageSizeValues() []ImageSize {
	return original.PossibleImageSizeValues()
}
func PossibleImageTypeValues() []ImageType {
	return original.PossibleImageTypeValues()
}
func PossibleItemAvailabilityValues() []ItemAvailability {
	return original.PossibleItemAvailabilityValues()
}
func PossibleSafeSearchValues() []SafeSearch {
	return original.PossibleSafeSearchValues()
}
func PossibleTypeBasicResponseBaseValues() []TypeBasicResponseBase {
	return original.PossibleTypeBasicResponseBaseValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
