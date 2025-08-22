package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// FuelEfficiency represents the FuelEfficiency schema from the OpenAPI specification
type FuelEfficiency struct {
	Annual_miles float64 `json:"annual_miles,omitempty"` // Annual Miles of Car
	City_mileage string `json:"city_mileage,omitempty"` // City Mileage in MPG
	Combined_mileage string `json:"combined_mileage,omitempty"` // Combined Mileage
	Highway_mileage string `json:"highway_mileage,omitempty"` // Highway Mileage in MPG
	Monthly_fuel_expense float64 `json:"monthly_fuel_expense,omitempty"` // Monthly fuel expense
}

// CompetitorsSameCars represents the CompetitorsSameCars schema from the OpenAPI specification
type CompetitorsSameCars struct {
	Avg_price float64 `json:"avg_price,omitempty"` // Price of year make model combination
	Name string `json:"name,omitempty"` // ymm_comb_name
	Avg_market_value float64 `json:"avg_market_value,omitempty"` // Estimated market value of year make model combination
	Avg_miles float64 `json:"avg_miles,omitempty"` // Mileage of year make model combination
}

// UKRVSearchFacets represents the UKRVSearchFacets schema from the OpenAPI specification
type UKRVSearchFacets struct {
	Year []FacetItem `json:"year,omitempty"`
	Fuel_type []FacetItem `json:"fuel_type,omitempty"`
	City []FacetItem `json:"city,omitempty"`
	Postal_code []FacetItem `json:"postal_code,omitempty"`
	Steering []FacetItem `json:"steering,omitempty"`
	Model []FacetItem `json:"model,omitempty"`
	MakeField []FacetItem `json:"make,omitempty"`
	Category []FacetItem `json:"category,omitempty"`
	Interior_color []FacetItem `json:"interior_color,omitempty"`
	Sub_category []FacetItem `json:"sub_category,omitempty"`
	Seller_name []FacetItem `json:"seller_name,omitempty"`
	Transmission []FacetItem `json:"transmission,omitempty"`
	Dealer_id []FacetItem `json:"dealer_id,omitempty"`
	Chassis []FacetItem `json:"chassis,omitempty"`
	Engine []FacetItem `json:"engine,omitempty"`
	Inventory_type []FacetItem `json:"inventory_type,omitempty"`
	Drivetrain []FacetItem `json:"drivetrain,omitempty"`
	Exterior_color []FacetItem `json:"exterior_color,omitempty"`
	Zip []FacetItem `json:"zip,omitempty"`
	Drive_type []FacetItem `json:"drive_type,omitempty"`
	Source []FacetItem `json:"source,omitempty"`
	Berths []FacetItem `json:"berths,omitempty"`
	State []FacetItem `json:"state,omitempty"`
}

// ComparisonPoint represents the ComparisonPoint schema from the OpenAPI specification
type ComparisonPoint struct {
	Dealer_indicator string `json:"dealer_indicator,omitempty"` // Deal Indicator
	Fair_deal_price float64 `json:"fair_deal_price,omitempty"` // Fair Deal Price
	Similar_vehicles_count float64 `json:"similar_vehicles_count,omitempty"` // Similar Vehicles Count
	Vin_price float64 `json:"vin_price,omitempty"` // Price for Vin
}

// HeavyEquipmentsSearchStats represents the HeavyEquipmentsSearchStats schema from the OpenAPI specification
type HeavyEquipmentsSearchStats struct {
	Price StatsItem `json:"price,omitempty"`
	Hours_used StatsItem `json:"hours_used,omitempty"`
	Miles StatsItem `json:"miles,omitempty"`
}

// RVSearchFacets represents the RVSearchFacets schema from the OpenAPI specification
type RVSearchFacets struct {
	Inventory_type []FacetItem `json:"inventory_type,omitempty"`
	Sleeps []FacetItem `json:"sleeps,omitempty"`
	Seller_name []FacetItem `json:"seller_name,omitempty"`
	Awnings []FacetItem `json:"awnings,omitempty"`
	Interior_color []FacetItem `json:"interior_color,omitempty"`
	Slideouts []FacetItem `json:"slideouts,omitempty"`
	Source []FacetItem `json:"source,omitempty"`
	Fuel_type []FacetItem `json:"fuel_type,omitempty"`
	State []FacetItem `json:"state,omitempty"`
	Model []FacetItem `json:"model,omitempty"`
	Engine []FacetItem `json:"engine,omitempty"`
	City []FacetItem `json:"city,omitempty"`
	Year []FacetItem `json:"year,omitempty"`
	Class []FacetItem `json:"class,omitempty"`
	Length []FacetItem `json:"length,omitempty"`
	MakeField []FacetItem `json:"make,omitempty"`
	Exterior_color []FacetItem `json:"exterior_color,omitempty"`
	Area []FacetItem `json:"area,omitempty"`
	Dealer_id []FacetItem `json:"dealer_id,omitempty"`
	Transmission []FacetItem `json:"transmission,omitempty"`
}

// UKRVSearchRangeFacets represents the UKRVSearchRangeFacets schema from the OpenAPI specification
type UKRVSearchRangeFacets struct {
	Exterior_length map[string]interface{} `json:"exterior_length,omitempty"`
	Interior_length map[string]interface{} `json:"interior_length,omitempty"`
	Miles map[string]interface{} `json:"miles,omitempty"`
	Price map[string]interface{} `json:"price,omitempty"`
	Width map[string]interface{} `json:"width,omitempty"`
}

// ListingFinance represents the ListingFinance schema from the OpenAPI specification
type ListingFinance struct {
	Estimated_monthly_payment float64 `json:"estimated_monthly_payment,omitempty"` // estimated monthly payment for this finance offer
	Loan_apr float64 `json:"loan_apr,omitempty"` // Loan APR for this finance offer
	Loan_term int `json:"loan_term,omitempty"` // Loan term for this finance offer
	Down_payment float64 `json:"down_payment,omitempty"` // Down payment for this finance offer
	Down_payment_percentage float64 `json:"down_payment_percentage,omitempty"` // down payment percentage for this finance offer
}

// HeavyEquipmentsSearchRangeFacets represents the HeavyEquipmentsSearchRangeFacets schema from the OpenAPI specification
type HeavyEquipmentsSearchRangeFacets struct {
	Miles map[string]interface{} `json:"miles,omitempty"`
	Price map[string]interface{} `json:"price,omitempty"`
	Hours_used map[string]interface{} `json:"hours_used,omitempty"`
}

// CompetitorsPoint represents the CompetitorsPoint schema from the OpenAPI specification
type CompetitorsPoint struct {
	Current_car CompetitorsCarDetails `json:"current_car,omitempty"`
	Same_car_models []CompetitorsSameCars `json:"same_car_models,omitempty"`
	Similar_car_models []CompetitorsSimilarCars `json:"similar_car_models,omitempty"`
}

// ListingNestMedia represents the ListingNestMedia schema from the OpenAPI specification
type ListingNestMedia struct {
	Photo_links []string `json:"photo_links,omitempty"` // A list of photo urls for the car
	Photo_links_cached []string `json:"photo_links_cached,omitempty"` // A list of cached photo urls for the car
}

// MotorcycleBaseListing represents the MotorcycleBaseListing schema from the OpenAPI specification
type MotorcycleBaseListing struct {
	First_seen_at int `json:"first_seen_at,omitempty"` // Listing first seen at first scraped timestamp
	Last_seen_at int `json:"last_seen_at,omitempty"` // Listing last seen at (most recent) timestamp
	Color string `json:"color,omitempty"` // Color of the motorcycle
	Dealer NestDealer `json:"dealer,omitempty"`
	First_seen_at_date string `json:"first_seen_at_date,omitempty"` // Listing first seen at first scraped date
	Inventory_type string `json:"inventory_type,omitempty"` // Inventory type of motorcycle
	Miles int `json:"miles,omitempty"` // Odometer reading / reported miles usage for the motorcycle
	Scraped_at float64 `json:"scraped_at,omitempty"` // Listing last seen at date timestamp
	Heading string `json:"heading,omitempty"` // Listing title as displayed on the source website
	Id string `json:"id,omitempty"` // Unique identifier representing a specific listing from the Marketcheck database
	Msrp int `json:"msrp,omitempty"` // MSRP for the motorcycle
	Scraped_at_date string `json:"scraped_at_date,omitempty"` // Listing last seen at date
	Dist float64 `json:"dist,omitempty"` // Distance of the motorcycle's location from the specified user lcoation
	Build MotorcycleBuild `json:"build,omitempty"` // Describes the Motorcycle specification
	Stock_no string `json:"stock_no,omitempty"` // Stock number of motorcycle in dealers inventory
	Last_seen_at_date string `json:"last_seen_at_date,omitempty"` // Listing last seen at (most recent) date
	Vin string `json:"vin,omitempty"` // VIN for the Motorcycle
	Dp_url string `json:"dp_url,omitempty"` // Details Page url of the specific motorcycle
	Price int `json:"price,omitempty"` // Asking price for the motorcycle
	Media ListingNestMedia `json:"media,omitempty"`
	Seller_type string `json:"seller_type,omitempty"` // Seller type for the motorcycle
	Source string `json:"source,omitempty"` // Source domain of the listing
}

// MotorcycleBuild represents the MotorcycleBuild schema from the OpenAPI specification
type MotorcycleBuild struct {
	Body_type string `json:"body_type,omitempty"` // Body type of the motorcycle
	Engine string `json:"engine,omitempty"` // Engine of the motorcycle
	Made_in string `json:"made_in,omitempty"` // Made in of the motorcycle
	MakeField string `json:"make,omitempty"` // Motorcycle Make
	Vehicle_type string `json:"vehicle_type,omitempty"` // Vehicle type of the motorcycle
	Year int `json:"year,omitempty"` // Year of the motorcycle
	Cylinders int `json:"cylinders,omitempty"` // No of cylinders of the motorcycle
	Drivetrain string `json:"drivetrain,omitempty"` // Drivetrain of the motorcycle
	Dry_weight string `json:"dry_weight,omitempty"` // Dry weight of motorcycle
	Trim string `json:"trim,omitempty"` // Trim of the motorcycle
	Fuel_type string `json:"fuel_type,omitempty"` // Fuel type of the motorcycle
	Model string `json:"model,omitempty"` // Motorcycle model
	Transmission string `json:"transmission,omitempty"` // Transmission of the motorcycle
}

// DealerReview represents the DealerReview schema from the OpenAPI specification
type DealerReview struct {
	State string `json:"state,omitempty"` // State of the dealer
	Latitude string `json:"latitude,omitempty"` // Latutide for the dealer location
	Longitude string `json:"longitude,omitempty"` // Longitude for the dealer location
	Name string `json:"name,omitempty"` // Name of the dealer
	Overall_reviews float64 `json:"overall_reviews,omitempty"` // Overall reviews of the dealership
	Street string `json:"street,omitempty"` // Street of the dealer
	City string `json:"city,omitempty"` // City of the dealer
	Review_components []ReviewComponents `json:"review_components,omitempty"`
	Zip string `json:"zip,omitempty"` // Zip of the dealer
}

// NeoVIN represents the NeoVIN schema from the OpenAPI specification
type NeoVIN struct {
	Model string `json:"model,omitempty"` // Model of the vehicle
	Package_code string `json:"package_code,omitempty"` // Extracted manufacturer package code related to selection of the vehicle version (where available / applicable)
	Version_confidence string `json:"version_confidence,omitempty"` // Identification of confidence related to the selection of the vehicle version
	Created_at int `json:"created_at,omitempty"` // vin first decoded at timestamp
	City_mpg float64 `json:"city_mpg,omitempty"` // City miles per gallon
	Highway_mpg float64 `json:"highway_mpg,omitempty"` // Highway miles per gallon
	Transmission_confidence string `json:"transmission_confidence,omitempty"` // Identification of confidence related to the selection of the installed transmission
	Version string `json:"version,omitempty"` // Version Name of the vehicle
	Vin string `json:"vin,omitempty"` // 17 char long VIN of the vehicle
	Created_at_date string `json:"created_at_date,omitempty"` // vin first decoded at date
	Transmission_description string `json:"transmission_description,omitempty"` // The description of the installed transmission
	Available_options_details map[string]interface{} `json:"available_options_details,omitempty"` // Listing of all options and packages with detailed equipment identified that could have been installed on the vehicle
	Installed_options_details map[string]interface{} `json:"installed_options_details,omitempty"` // Listing of all options and packages with detailed equipment identified that have been identified as installed on the vehicle
	Package_description string `json:"package_description,omitempty"` // Extracted manufacturer package code description (where available / applicable)
	Width float64 `json:"width,omitempty"` // Overall width of the vehicle (in)
	MakeField string `json:"make,omitempty"` // Make of the vehicle
	Installed_options_msrp float64 `json:"installed_options_msrp,omitempty"` // Calculated combination of the MSRP for the installed options and packages
	Combined_msrp float64 `json:"combined_msrp,omitempty"` // Calculated combination of base MSRP with the delivery charges and the installed options MSRP
	Options_packages string `json:"options_packages,omitempty"` // Option and package codes installed on the vehicle
	Interior_color map[string]interface{} `json:"interior_color,omitempty"` // The manufacturer defined option code and name for the interior color of the vehicle
	Transmission string `json:"transmission,omitempty"` // Identified installed Transmission of the vehicle
	Exterior_color map[string]interface{} `json:"exterior_color,omitempty"` // The manufacturer defined option code and name for the exterior color of the vehicle
	Body_type string `json:"body_type,omitempty"` // Body type of the vehicle
	Msrp float64 `json:"msrp,omitempty"` // Base MSRP as defined for the defined version of that vehicle with no options installed
	Engine string `json:"engine,omitempty"` // Base identification of the number of cylinders and configuration of the installed engine on the vehicle
	Trim string `json:"trim,omitempty"` // Trim of the vehicle
	Body_subtype string `json:"body_subtype,omitempty"` // Body subtype (cab) of the vehicle
	Seating_capacity float64 `json:"seating_capacity,omitempty"` // Identified seating capacity of the vehicle
	Trim_confidence string `json:"trim_confidence,omitempty"` // Identification of confidence related to the selection of the trim
	Updated_at int `json:"updated_at,omitempty"` // vin last decoded at timestamp
	Squish_vin string `json:"squish_vin,omitempty"` // Vin prefix identification used in the decoding of the vehicle
	Features map[string]interface{} `json:"features,omitempty"` // Normalized feature set of equipment identified that have been identified as installed on the vehicle, including whether that item was standard to the version of the vehicle or was added as an installed option
	Fuel_type string `json:"fuel_type,omitempty"` // Identified Fuel type of the vehicle
	Manufacturer_code string `json:"manufacturer_code,omitempty"` // Manufacturer code of the vehicle
	Installed_equipment map[string]interface{} `json:"installed_equipment,omitempty"` // Detailed equipment identified as installed on the vehicle, including whether that item was standard to the version of the vehicle or was added as an installed option
	Doors int `json:"doors,omitempty"` // Number of doors on the vehicle
	Height float64 `json:"height,omitempty"` // Overall height of the vehicle (in)
	Delivery_charges float64 `json:"delivery_charges,omitempty"` // Manufacturer defined delivery charges
	Drivetrain string `json:"drivetrain,omitempty"` // Identified drivetrain of the vehicle
	Length float64 `json:"length,omitempty"` // Overall length of the vehicle (in)
	Weight float64 `json:"weight,omitempty"` // Gross vehicle weight (lbs)
	Updated_at_date string `json:"updated_at_date,omitempty"` // vin last decoded at date
	Year int `json:"year,omitempty"` // Model Year of the vehicle
	Listing_confidence string `json:"listing_confidence,omitempty"` // Identification of primary source data used for trim and installed colors and options
	Decode_version int `json:"decode_version,omitempty"` // This will represent if previous vin decode values are updated
}

// SalesStats represents the SalesStats schema from the OpenAPI specification
type SalesStats struct {
	Population_standard_deviation float64 `json:"population_standard_deviation,omitempty"` // population_standard_deviation
	Variance float64 `json:"variance,omitempty"` // variance
	Mean float64 `json:"mean,omitempty"` // mean
	Trimmed_mean float64 `json:"trimmed_mean,omitempty"` // trimmed_mean
	Iqr float64 `json:"iqr,omitempty"` // iqr
	Median float64 `json:"median,omitempty"` // median
	Standard_deviation float64 `json:"standard_deviation,omitempty"` // standard_deviation
	Weighted_mean float64 `json:"weighted_mean,omitempty"` // weighted_mean
	Absolute_mean_deviation float64 `json:"absolute_mean_deviation,omitempty"` // absolute_mean_deviation
}

// VinReport represents the VinReport schema from the OpenAPI specification
type VinReport struct {
	Icon string `json:"icon,omitempty"` // icon url
	Some_text string `json:"some_text,omitempty"` // Some descriptional text
}

// BasePopular represents the BasePopular schema from the OpenAPI specification
type BasePopular struct {
	Price_stats SalesStats `json:"price_stats,omitempty"` // Sales stats response
	State string `json:"state,omitempty"` // State
	Counts string `json:"counts,omitempty"` // Number of sold listings of this model
	Inventorytype string `json:"inventoryType,omitempty"` // Inventory type
	City string `json:"city,omitempty"` // City
	Dom_stats SalesStats `json:"dom_stats,omitempty"` // Sales stats response
	MakeField string `json:"make,omitempty"` // Make name
	Model string `json:"model,omitempty"` // Model name
	Miles_stats SalesStats `json:"miles_stats,omitempty"` // Sales stats response
}

// UKRVSearchResponse represents the UKRVSearchResponse schema from the OpenAPI specification
type UKRVSearchResponse struct {
	Num_found int `json:"num_found,omitempty"` // The number of listings found
	Range_facets map[string]interface{} `json:"range_facets,omitempty"`
	Stats map[string]interface{} `json:"stats,omitempty"`
	Facets map[string]interface{} `json:"facets,omitempty"`
	Listings []UKRVBaseListing `json:"listings,omitempty"`
}

// Listing represents the Listing schema from the OpenAPI specification
type Listing struct {
	First_seen_at_mc_date string `json:"first_seen_at_mc_date,omitempty"` // Listing first seen at first scraped at MC date
	Ref_miles string `json:"ref_miles,omitempty"` // Last Odometer reading / reported miles usage for the car. If the asking miles value is not or is available then the last_miles could perhaps be used. last_miles is the miles for the car listed on the source website as of last_miles_dt date
	Extra ListingNestExtraAttributes `json:"extra,omitempty"`
	Stock_no string `json:"stock_no,omitempty"` // Stock number of car in dealers inventory
	First_seen_at_source_date string `json:"first_seen_at_source_date,omitempty"` // Listing first seen at source relisted date
	Inventory_type string `json:"inventory_type,omitempty"` // Inventory type of car
	First_seen_at int `json:"first_seen_at,omitempty"` // Listing first seen at first scraped timestamp
	Source string `json:"source,omitempty"` // Source domain of the listing
	Carfax_clean_title bool `json:"carfax_clean_title,omitempty"` // Flag to indicate whether listing is carfax_clean_title
	Scraped_at_date string `json:"scraped_at_date,omitempty"` // Listing last seen at date
	Media ListingNestMedia `json:"media,omitempty"`
	Seller_type string `json:"seller_type,omitempty"` // Seller type for the car
	Vehicle_registration_mark string `json:"vehicle_registration_mark,omitempty"` // Vehicle Registration Mark of the car
	First_seen_at_source int `json:"first_seen_at_source,omitempty"` // Listing first seen at source relisted timestamp
	Interior_color string `json:"interior_color,omitempty"` // Interior color of the car
	Price_change_percent float64 `json:"price_change_percent,omitempty"` // Percentage difference between the cars's current price and ref_price i.e. last reported price
	Vdp_url string `json:"vdp_url,omitempty"` // Vehicle Details Page url of the specific car
	Build Build `json:"build,omitempty"` // Describes the Car specification
	Id string `json:"id,omitempty"` // Unique identifier representing a specific listing from the Marketcheck database
	Dealer NestDealer `json:"dealer,omitempty"`
	Dom_180 int `json:"dom_180,omitempty"` // Days on Market value for the car based on current and last 6 month listings found in the Marketcheck database for this car
	Heading string `json:"heading,omitempty"` // Listing title as displayed on the source website
	Price int `json:"price,omitempty"` // Asking price for the car
	Last_seen_at int `json:"last_seen_at,omitempty"` // Listing last seen at (most recent) timestamp
	Last_seen_at_date string `json:"last_seen_at_date,omitempty"` // Listing last seen at (most recent) date
	Rank int `json:"rank,omitempty"` // Relative rank of the listing determined by ranking service
	Ref_miles_dt int `json:"ref_miles_dt,omitempty"` // The date at which the last miles was reported online. This is earlier to last_seen_date
	Dom_active int `json:"dom_active,omitempty"` // Days on Market value for the car based on current and last 30 days listings found in the Marketcheck database for this car
	First_seen_at_mc int `json:"first_seen_at_mc,omitempty"` // Listing first seen at first scraped at MC timestamp
	Msrp int `json:"msrp,omitempty"` // MSRP for the car
	Dom int `json:"dom,omitempty"` // Days on Market value for the car based on current and historical listings found in the Marketcheck database for this car
	Financing_options []ListingFinance `json:"financing_options,omitempty"` // Array of all finance offers for this listing
	Is_certified int `json:"is_certified,omitempty"` // Certified car
	Scraped_at int `json:"scraped_at,omitempty"` // Listing last seen at date timestamp
	Base_ext_color string `json:"base_ext_color,omitempty"` // Base exterior color of the car
	Base_int_color string `json:"base_int_color,omitempty"` // Base interior color of the car
	Data_source string `json:"data_source,omitempty"` // Data source of the listing
	Score float64 `json:"score,omitempty"` // Score of the ranked listing as per the ranking criteria determined by ranking service
	Carfax_1_owner bool `json:"carfax_1_owner,omitempty"` // Flag to indicate whether listing is carfax_1_owner
	Exterior_color string `json:"exterior_color,omitempty"` // Exterior color of the car
	First_seen_at_date string `json:"first_seen_at_date,omitempty"` // Listing first seen at first scraped date
	Miles int `json:"miles,omitempty"` // Odometer reading / reported miles usage for the car
	Ref_price_dt int `json:"ref_price_dt,omitempty"` // The date at which the last price was reported online. This is earlier to last_seen_date
	Leasing_options []ListingLease `json:"leasing_options,omitempty"` // Array of all finance offers for this listing
	Ref_price string `json:"ref_price,omitempty"` // Last reported price for the car. If the asking price value is not or is available then the last_price could perhaps be used. last_price is the price for the car listed on the source website as of last_price_dt date
	Vin string `json:"vin,omitempty"` // VIN for the car
}

// ListingDebugAttributes represents the ListingDebugAttributes schema from the OpenAPI specification
type ListingDebugAttributes struct {
	Scraped_at string `json:"scraped_at,omitempty"` // Scraped date and time
	Taxonomy_vin string `json:"taxonomy_vin,omitempty"` // Taxonomy vin
	Template_id float64 `json:"template_id,omitempty"` // Template id
	User_id float64 `json:"user_id,omitempty"` // User id
	Cycle_id float64 `json:"cycle_id,omitempty"` // Cycle id
	Robot_id float64 `json:"robot_id,omitempty"` // Robot id
}

// MotorcycleSearchFacets represents the MotorcycleSearchFacets schema from the OpenAPI specification
type MotorcycleSearchFacets struct {
	Trim []FacetItem `json:"trim,omitempty"`
	Model []FacetItem `json:"model,omitempty"`
	Seller_name []FacetItem `json:"seller_name,omitempty"`
	Cylinders []FacetItem `json:"cylinders,omitempty"`
	Source []FacetItem `json:"source,omitempty"`
	Vehicle_type []FacetItem `json:"vehicle_type,omitempty"`
	Body_type []FacetItem `json:"body_type,omitempty"`
	MakeField []FacetItem `json:"make,omitempty"`
	Transmission []FacetItem `json:"transmission,omitempty"`
	Engine []FacetItem `json:"engine,omitempty"`
	Inventory_type []FacetItem `json:"inventory_type,omitempty"`
	Color []FacetItem `json:"color,omitempty"`
	Drivetrain []FacetItem `json:"drivetrain,omitempty"`
	City []FacetItem `json:"city,omitempty"`
	Fuel_type []FacetItem `json:"fuel_type,omitempty"`
	State []FacetItem `json:"state,omitempty"`
	Dealer_id []FacetItem `json:"dealer_id,omitempty"`
	Year []FacetItem `json:"year,omitempty"`
}

// CarSearchStats represents the CarSearchStats schema from the OpenAPI specification
type CarSearchStats struct {
	Finance_down_payment_per StatsItem `json:"finance_down_payment_per,omitempty"`
	Lease_down_payment StatsItem `json:"lease_down_payment,omitempty"`
	Lease_term StatsItem `json:"lease_term,omitempty"`
	Miles StatsItem `json:"miles,omitempty"`
	Msrp StatsItem `json:"msrp,omitempty"`
	Dom_180 StatsItem `json:"dom_180,omitempty"`
	Dom_active StatsItem `json:"dom_active,omitempty"`
	Finance_loan_apr StatsItem `json:"finance_loan_apr,omitempty"`
	Lease_emp StatsItem `json:"lease_emp,omitempty"`
	Price StatsItem `json:"price,omitempty"`
	Finance_down_payment StatsItem `json:"finance_down_payment,omitempty"`
	Finance_emp StatsItem `json:"finance_emp,omitempty"`
	Finance_loan_term StatsItem `json:"finance_loan_term,omitempty"`
	Dom StatsItem `json:"dom,omitempty"`
}

// DepreciationPoint represents the DepreciationPoint schema from the OpenAPI specification
type DepreciationPoint struct {
	Similar_models []DepreciationStats `json:"similar_models,omitempty"`
}

// UKRVBaseListing represents the UKRVBaseListing schema from the OpenAPI specification
type UKRVBaseListing struct {
	Scraped_at_date string `json:"scraped_at_date,omitempty"` // Listing last seen at date
	Media ListingNestMedia `json:"media,omitempty"`
	Motorhome_build string `json:"motorhome_build,omitempty"` // motorhome_build of the listing
	Exterior_color string `json:"exterior_color,omitempty"` // Exterior color of the RV
	Last_seen_at_date string `json:"last_seen_at_date,omitempty"` // Listing last seen at (most recent) date
	Scraped_at float64 `json:"scraped_at,omitempty"` // Listing last seen at date timestamp
	Source string `json:"source,omitempty"` // Source domain of the listing
	Seller_type string `json:"seller_type,omitempty"` // Seller type for the RV
	Origin string `json:"origin,omitempty"` // origin of the listing
	Currency_indicator string `json:"currency_indicator,omitempty"` // currency_indicator of the listing
	Interior_color string `json:"interior_color,omitempty"` // Interior color of the RV
	Price int `json:"price,omitempty"` // Asking price for the RV
	Inventory_type string `json:"inventory_type,omitempty"` // Inventory type of RV
	Vdp_url string `json:"vdp_url,omitempty"` // Details Page url of the specific RV
	Last_seen_at int `json:"last_seen_at,omitempty"` // Listing last seen at (most recent) timestamp
	Msrp int `json:"msrp,omitempty"` // MSRP for the RV
	Dealer NestDealer `json:"dealer,omitempty"`
	First_seen_at int `json:"first_seen_at,omitempty"` // Listing first seen at first scraped timestamp
	Id string `json:"id,omitempty"` // Unique identifier representing a specific listing from the Marketcheck database
	Availability_status string `json:"availability_status,omitempty"` // availability_status of the listing
	Build UKRVBuild `json:"build,omitempty"` // Describes the RV specification
	Miles_indicator string `json:"miles_indicator,omitempty"` // miles_indicator of the listing
	Stock_no string `json:"stock_no,omitempty"` // Stock number of RV in dealers inventory
	Heading string `json:"heading,omitempty"` // Listing title as displayed on the source website
	Vin string `json:"vin,omitempty"` // VIN for the RV
	Miles int `json:"miles,omitempty"` // Odometer reading / reported miles usage for the RV
	Mot_expires string `json:"mot_expires,omitempty"` // mot_expires of the listing
	First_seen_at_date string `json:"first_seen_at_date,omitempty"` // Listing first seen at first scraped date
	Dist float64 `json:"dist,omitempty"` // Distance of the RV's location from the specified user lcoation
}

// HeavyEquipmentsBuild represents the HeavyEquipmentsBuild schema from the OpenAPI specification
type HeavyEquipmentsBuild struct {
	Category string `json:"category,omitempty"` // Heavy equipment category
	Condition string `json:"condition,omitempty"` // Heavy equipment condition
	MakeField string `json:"make,omitempty"` // Heavy equipment Make
	Model string `json:"model,omitempty"` // Heavy equipment model
	Drivetrain string `json:"drivetrain,omitempty"` // Drivetrain of the heavy equipment
	Engine string `json:"engine,omitempty"` // Engine of the heavy equipment
	Trim string `json:"trim,omitempty"` // Trim of the heavy equipment
	Body_type string `json:"body_type,omitempty"` // Body type of the heavy equipment
	Fuel_type string `json:"fuel_type,omitempty"` // Fuel type of the heavy equipment
	Sub_category string `json:"sub_category,omitempty"` // Heavy equipment sub_category
	Transmission string `json:"transmission,omitempty"` // Transmission of the heavy equipment
	Year int `json:"year,omitempty"` // Year of the heavy equipment
}

// DealersResponse represents the DealersResponse schema from the OpenAPI specification
type DealersResponse struct {
	Num_found int `json:"num_found,omitempty"` // number of dealers found
	Dealers []Dealer `json:"dealers,omitempty"`
}

// RVBaseListing represents the RVBaseListing schema from the OpenAPI specification
type RVBaseListing struct {
	Msrp int `json:"msrp,omitempty"` // MSRP for the RV
	Id string `json:"id,omitempty"` // Unique identifier representing a specific listing from the Marketcheck database
	Dp_url string `json:"dp_url,omitempty"` // Details Page url of the specific RV
	Seller_type string `json:"seller_type,omitempty"` // Seller type for the RV
	Source string `json:"source,omitempty"` // Source domain of the listing
	Miles int `json:"miles,omitempty"` // Odometer reading / reported miles usage for the RV
	First_seen_at int `json:"first_seen_at,omitempty"` // Listing first seen at first scraped timestamp
	Media ListingNestMedia `json:"media,omitempty"`
	Dealer NestDealer `json:"dealer,omitempty"`
	Dist float64 `json:"dist,omitempty"` // Distance of the RV's location from the specified user lcoation
	Exterior_color string `json:"exterior_color,omitempty"` // Exterior color of the RV
	Inventory_type string `json:"inventory_type,omitempty"` // Inventory type of RV
	Scraped_at_date string `json:"scraped_at_date,omitempty"` // Listing last seen at date
	Last_seen_at int `json:"last_seen_at,omitempty"` // Listing last seen at (most recent) timestamp
	Price int `json:"price,omitempty"` // Asking price for the RV
	Build RVBuild `json:"build,omitempty"` // Describes the RV specification
	First_seen_at_date string `json:"first_seen_at_date,omitempty"` // Listing first seen at first scraped date
	Vin string `json:"vin,omitempty"` // VIN for the RV
	Heading string `json:"heading,omitempty"` // Listing title as displayed on the source website
	Last_seen_at_date string `json:"last_seen_at_date,omitempty"` // Listing last seen at (most recent) date
	Scraped_at float64 `json:"scraped_at,omitempty"` // Listing last seen at date timestamp
	Stock_no string `json:"stock_no,omitempty"` // Stock number of RV in dealers inventory
	Interior_color string `json:"interior_color,omitempty"` // Interior color of the RV
}

// ListingLease represents the ListingLease schema from the OpenAPI specification
type ListingLease struct {
	Estimated_monthly_payment float64 `json:"estimated_monthly_payment,omitempty"` // estimated monthly payment for this lease offer
	Lease_term int `json:"lease_term,omitempty"` // Lease term for this lease offer
	Down_payment float64 `json:"down_payment,omitempty"` // Down payment for this lease offer
}

// Build represents the Build schema from the OpenAPI specification
type Build struct {
	Doors int `json:"doors,omitempty"` // No of doors of the car
	Model string `json:"model,omitempty"` // Car model
	Trim string `json:"trim,omitempty"` // Trim of the car
	MakeField string `json:"make,omitempty"` // Car Make
	Engine_size float64 `json:"engine_size,omitempty"` // Engine size of the car
	Highway_mpg int `json:"highway_mpg,omitempty"` // Highway Mileage in MPG
	Drivetrain string `json:"drivetrain,omitempty"` // Drivetrain of the car
	Steering_type string `json:"steering_type,omitempty"` // Steering type of the car
	Fuel_type string `json:"fuel_type,omitempty"` // Fuel type of the car
	Engine_block string `json:"engine_block,omitempty"` // Engine block of the car
	Std_seating string `json:"std_seating,omitempty"` // Std seating of the car
	Engine_aspiration string `json:"engine_aspiration,omitempty"` // Engine aspiration of the car
	Powertrain_type string `json:"powertrain_type,omitempty"` // powertrain_type of the car
	Year int `json:"year,omitempty"` // Year of the Car
	Overall_width string `json:"overall_width,omitempty"` // Overall width of the car
	City_mpg int `json:"city_mpg,omitempty"` // City Mileage in MPG
	Overall_length string `json:"overall_length,omitempty"` // Overall length of the car
	Short_trim string `json:"short_trim,omitempty"` // Short trim of the car
	Made_in string `json:"made_in,omitempty"` // Made in of the car
	Trim_r string `json:"trim_r,omitempty"` // Trim_r of the car
	Transmission string `json:"transmission,omitempty"` // Transmission of the car
	Engine string `json:"engine,omitempty"` // Engine of the car
	Antibrake_sys string `json:"antibrake_sys,omitempty"` // Antibrake system of the car
	Highway_miles string `json:"highway_miles,omitempty"` // Highway miles of the car
	Body_type string `json:"body_type,omitempty"` // Body type of the car
	City_miles string `json:"city_miles,omitempty"` // City miles of the car
	Vehicle_type string `json:"vehicle_type,omitempty"` // Vehicle type of the car
	Tank_size string `json:"tank_size,omitempty"` // Tank size of the car
	Engine_measure string `json:"engine_measure,omitempty"` // Engine block of the car
	Overall_height string `json:"overall_height,omitempty"` // Overall height of the car
	Body_subtype string `json:"body_subtype,omitempty"` // Body subtype of the car
	Cylinders int `json:"cylinders,omitempty"` // No of cylinders of the car
	Opt_seating string `json:"opt_seating,omitempty"` // opt seating of the car
}

// HeavyEquipmentsBaseListing represents the HeavyEquipmentsBaseListing schema from the OpenAPI specification
type HeavyEquipmentsBaseListing struct {
	Source string `json:"source,omitempty"` // Source domain of the listing
	Media ListingNestMedia `json:"media,omitempty"`
	Seller_type string `json:"seller_type,omitempty"` // Seller type for the heavy equipment
	Dp_url string `json:"dp_url,omitempty"` // Details Page url of the specific heavy equipment
	Exterior_color string `json:"exterior_color,omitempty"` // Exterior color of the heavy equipment
	Msrp int `json:"msrp,omitempty"` // MSRP for the heavy equipment
	Last_seen_at int `json:"last_seen_at,omitempty"` // Listing last seen at (most recent) timestamp
	Build HeavyEquipmentsBuild `json:"build,omitempty"` // Describes the Heavy Equipments specification
	Price int `json:"price,omitempty"` // Asking price for the heavy equipment
	Stock_no string `json:"stock_no,omitempty"` // Stock number of heavy equipment in dealers inventory
	First_seen_at int `json:"first_seen_at,omitempty"` // Listing first seen at first scraped timestamp
	Miles int `json:"miles,omitempty"` // Odometer reading / reported miles usage for the heavy equipment
	Scraped_at float64 `json:"scraped_at,omitempty"` // Listing last seen at date timestamp
	Last_seen_at_date string `json:"last_seen_at_date,omitempty"` // Listing last seen at (most recent) date
	Vin string `json:"vin,omitempty"` // VIN for the heavy equipment
	Dist float64 `json:"dist,omitempty"` // Distance of the quipments's location from the specified user lcoation
	Heading string `json:"heading,omitempty"` // Listing title as displayed on the source website
	Scraped_at_date string `json:"scraped_at_date,omitempty"` // Listing last seen at date
	First_seen_at_date string `json:"first_seen_at_date,omitempty"` // Listing first seen at first scraped date
	Inventory_type string `json:"inventory_type,omitempty"` // Inventory type of heavy equipment
	Interior_color string `json:"interior_color,omitempty"` // Interior color of the heavy equipment
	Dealer NestDealer `json:"dealer,omitempty"`
	Id string `json:"id,omitempty"` // Unique identifier representing a specific listing from the Marketcheck database
}

// HeavyEquipmentsSearchFacets represents the HeavyEquipmentsSearchFacets schema from the OpenAPI specification
type HeavyEquipmentsSearchFacets struct {
	Body_type []FacetItem `json:"body_type,omitempty"`
	Category []FacetItem `json:"category,omitempty"`
	State []FacetItem `json:"state,omitempty"`
	Model []FacetItem `json:"model,omitempty"`
	Seller_name []FacetItem `json:"seller_name,omitempty"`
	Fuel_type []FacetItem `json:"fuel_type,omitempty"`
	Year []FacetItem `json:"year,omitempty"`
	Transmission []FacetItem `json:"transmission,omitempty"`
	Sub_category []FacetItem `json:"sub_category,omitempty"`
	Exterior_color []FacetItem `json:"exterior_color,omitempty"`
	City []FacetItem `json:"city,omitempty"`
	Drivetrain []FacetItem `json:"drivetrain,omitempty"`
	Trim []FacetItem `json:"trim,omitempty"`
	Interior_color []FacetItem `json:"interior_color,omitempty"`
	Source []FacetItem `json:"source,omitempty"`
	Cylinders []FacetItem `json:"cylinders,omitempty"`
	Dealer_id []FacetItem `json:"dealer_id,omitempty"`
	Engine []FacetItem `json:"engine,omitempty"`
	Inventory_type []FacetItem `json:"inventory_type,omitempty"`
	MakeField []FacetItem `json:"make,omitempty"`
}

// RVListing represents the RVListing schema from the OpenAPI specification
type RVListing struct {
	Dealer NestDealer `json:"dealer,omitempty"`
	Extra ListingNestExtraAttributes `json:"extra,omitempty"`
	Scraped_at float64 `json:"scraped_at,omitempty"` // Listing last seen at date timestamp
	Build RVBuild `json:"build,omitempty"` // Describes the RV specification
	Dp_url string `json:"dp_url,omitempty"` // Details Page url of the specific RV
	Exterior_color string `json:"exterior_color,omitempty"` // Exterior color of the RV
	First_seen_at_date string `json:"first_seen_at_date,omitempty"` // Listing first seen at first scraped date
	Last_seen_at_date string `json:"last_seen_at_date,omitempty"` // Listing last seen at (most recent) date
	Vin string `json:"vin,omitempty"` // VIN for the RV
	Msrp int `json:"msrp,omitempty"` // MSRP for the RV
	Miles int `json:"miles,omitempty"` // Odometer reading / reported miles usage for the RV
	First_seen_at int `json:"first_seen_at,omitempty"` // Listing first seen at first scraped timestamp
	Seller_type string `json:"seller_type,omitempty"` // Seller type for the RV
	Source string `json:"source,omitempty"` // Source domain of the listing
	Interior_color string `json:"interior_color,omitempty"` // Interior color of the RV
	Price int `json:"price,omitempty"` // Asking price for the RV
	Heading string `json:"heading,omitempty"` // Listing title as displayed on the source website
	Id string `json:"id,omitempty"` // Unique identifier representing a specific listing from the Marketcheck database
	Last_seen_at int `json:"last_seen_at,omitempty"` // Listing last seen at (most recent) timestamp
	Inventory_type string `json:"inventory_type,omitempty"` // Inventory type of RV
	Media ListingNestMedia `json:"media,omitempty"`
	Stock_no string `json:"stock_no,omitempty"` // Stock number of RV in dealers inventory
	Scraped_at_date string `json:"scraped_at_date,omitempty"` // Listing last seen at date
}

// SearchAutoCompleteResponse represents the SearchAutoCompleteResponse schema from the OpenAPI specification
type SearchAutoCompleteResponse struct {
	Terms []map[string]interface{} `json:"terms,omitempty"`
}

// PopularityItem represents the PopularityItem schema from the OpenAPI specification
type PopularityItem struct {
	Color string `json:"color,omitempty"` // Color depending upon left and right analysis numbers
	Delta_percent float64 `json:"delta_percent,omitempty"` // Delta percent
	Difference float64 `json:"difference,omitempty"` // Difference depending upon left and right analysis
	Left float64 `json:"left,omitempty"` // Left side rating for above description (2016 ford F-150 Lariat)
	Right float64 `json:"right,omitempty"` // Right side rating for above description (All Other Cars)
	Text string `json:"text,omitempty"` // Description for which popularity should be show eg: 2016 ford F-150 Lariat vs All Other Cars
	Thumbs string `json:"thumbs,omitempty"` // Thumbs up/down depending upon left and right analysis numbers
}

// StatsItem represents the StatsItem schema from the OpenAPI specification
type StatsItem struct {
	Sum_of_squares float64 `json:"sum_of_squares,omitempty"` // sum_of_squares of the field
	Count int `json:"count,omitempty"` // count
	Mean float64 `json:"mean,omitempty"` // Mean of the field
	Median float64 `json:"median,omitempty"` // median of the field
	Missing int `json:"missing,omitempty"` // count of listings missing particular field
	Stddev float64 `json:"stddev,omitempty"` // stddev of the field
	Max int `json:"max,omitempty"` // Maximum value of the field
	Sum int `json:"sum,omitempty"` // Summation of all values of the field
	Min int `json:"min,omitempty"` // Minimum value of the field
}

// DealerLandingPage represents the DealerLandingPage schema from the OpenAPI specification
type DealerLandingPage struct {
	Html string `json:"html,omitempty"` // The HTML string for the listing web page
	Srp_url string `json:"srp_url,omitempty"` // The URL of the SRP
	Crawled_at float64 `json:"crawled_at,omitempty"` // The timestamp indicating the time when the SRP was cached
}

// CarSearchFacets represents the CarSearchFacets schema from the OpenAPI specification
type CarSearchFacets struct {
	Fuel_type []FacetItem `json:"fuel_type,omitempty"`
	Year []FacetItem `json:"year,omitempty"`
	Drivetrain []FacetItem `json:"drivetrain,omitempty"`
	Interior_color []FacetItem `json:"interior_color,omitempty"`
	Doors []FacetItem `json:"doors,omitempty"`
	Trim []FacetItem `json:"trim,omitempty"`
	Car_type []FacetItem `json:"car_type,omitempty"`
	Seller_name []FacetItem `json:"seller_name,omitempty"`
	Vehicle_type []FacetItem `json:"vehicle_type,omitempty"`
	Seller_name_o []FacetItem `json:"seller_name_o,omitempty"`
	Seller_type []FacetItem `json:"seller_type,omitempty"`
	MakeField []FacetItem `json:"make,omitempty"`
	Transmission []FacetItem `json:"transmission,omitempty"`
	State []FacetItem `json:"state,omitempty"`
	Carfax_1_owner []FacetItem `json:"carfax_1_owner,omitempty"`
	Trim_o []FacetItem `json:"trim_o,omitempty"`
	Engine_size []FacetItem `json:"engine_size,omitempty"`
	Body_type []FacetItem `json:"body_type,omitempty"`
	Model []FacetItem `json:"model,omitempty"`
	Body_subtype []FacetItem `json:"body_subtype,omitempty"`
	Engine []FacetItem `json:"engine,omitempty"`
	Engine_block []FacetItem `json:"engine_block,omitempty"`
	Source []FacetItem `json:"source,omitempty"`
	Carfax_clean_title []FacetItem `json:"carfax_clean_title,omitempty"`
	Base_interior_color []FacetItem `json:"base_interior_color,omitempty"`
	Engine_aspiration []FacetItem `json:"engine_aspiration,omitempty"`
	City []FacetItem `json:"city,omitempty"`
	Trim_r []FacetItem `json:"trim_r,omitempty"`
	Base_exterior_color []FacetItem `json:"base_exterior_color,omitempty"`
	Exterior_color []FacetItem `json:"exterior_color,omitempty"`
	Dealer_type []FacetItem `json:"dealer_type,omitempty"`
	Cylinders []FacetItem `json:"cylinders,omitempty"`
	Data_source []FacetItem `json:"data_source,omitempty"`
	Dealer_id []FacetItem `json:"dealer_id,omitempty"`
}

// Sales represents the Sales schema from the OpenAPI specification
type Sales struct {
	Model string `json:"model,omitempty"` // model
	Taxonomy_vin string `json:"taxonomy_vin,omitempty"` // taxonomy_vin
	Trim string `json:"trim,omitempty"` // trim
	Counts int `json:"counts,omitempty"` // Sales count
	Cpo int `json:"cpo,omitempty"` // cpo sales count
	Price_stats SalesStats `json:"price_stats,omitempty"` // Sales stats response
	Year string `json:"year,omitempty"` // year
	Miles_stats SalesStats `json:"miles_stats,omitempty"` // Sales stats response
	State string `json:"state,omitempty"` // State
	City string `json:"city,omitempty"` // City
	Dom_stats SalesStats `json:"dom_stats,omitempty"` // Sales stats response
	Non_cpo int `json:"non-cpo,omitempty"` // Non-cpo sales count
	Inventory_type string `json:"inventory_type,omitempty"` // inventory_type
	MakeField string `json:"make,omitempty"` // Make
}

// ReviewComponents represents the ReviewComponents schema from the OpenAPI specification
type ReviewComponents struct {
	User_rating float64 `json:"user_rating,omitempty"` // rating for dealer given by user ranging from 1-5
	Actual_review string `json:"actual_review,omitempty"` // review of car given by current user
	User_name string `json:"user_name,omitempty"` // Name of user who has given review
}

// UKSearchResponse represents the UKSearchResponse schema from the OpenAPI specification
type UKSearchResponse struct {
	Stats CarSearchStats `json:"stats,omitempty"`
	Facets UKCarSearchFacets `json:"facets,omitempty"`
	Listings []UKBaseListing `json:"listings,omitempty"`
	Num_found int `json:"num_found,omitempty"` // The number of listings found
	Range_facets CarSearchRangeFacets `json:"range_facets,omitempty"`
}

// RVBuild represents the RVBuild schema from the OpenAPI specification
type RVBuild struct {
	Fuel_type string `json:"fuel_type,omitempty"` // Fuel type of the RV
	Gvwr string `json:"gvwr,omitempty"` // Gross Vehicle Weight Ratio of RV
	Slideouts string `json:"slideouts,omitempty"` // Slideouts of the RV
	Area string `json:"area,omitempty"` // Area of the RV
	Length string `json:"length,omitempty"` // Length of the RV
	Made_in string `json:"made_in,omitempty"` // Made in of the RV
	MakeField string `json:"make,omitempty"` // RV Make
	Model string `json:"model,omitempty"` // RV model
	Transmission string `json:"transmission,omitempty"` // Transmission of the RV
	Engine string `json:"engine,omitempty"` // Engine of the RV
	Sleeps string `json:"sleeps,omitempty"` // Sleeps of the RV
	Year int `json:"year,omitempty"` // Year of the RV
	Class string `json:"class,omitempty"` // Class of the RV
}

// PricePrediction represents the PricePrediction schema from the OpenAPI specification
type PricePrediction struct {
	Specs PredictedSpecs `json:"specs,omitempty"` // Specs of predicted vehicle
	Predicted_price int `json:"predicted_price,omitempty"` // Actual predicted price of a car
	Price_range map[string]interface{} `json:"price_range,omitempty"` // Price range around the predicted price, estimating the error in predicted price
}

// SearchResponse represents the SearchResponse schema from the OpenAPI specification
type SearchResponse struct {
	Num_found int `json:"num_found,omitempty"` // The number of listings found
	Range_facets CarSearchRangeFacets `json:"range_facets,omitempty"`
	Stats CarSearchStats `json:"stats,omitempty"`
	Facets CarSearchFacets `json:"facets,omitempty"`
	Listings []BaseListing `json:"listings,omitempty"`
}

// ListingVDP represents the ListingVDP schema from the OpenAPI specification
type ListingVDP struct {
	Vdp_url string `json:"vdp_url,omitempty"` // The URL of the VDP
	Crawled_at float64 `json:"crawled_at,omitempty"` // The timestamp indicating the time when the VDP was cached
	Html string `json:"html,omitempty"` // The HTML string for the listing web page
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Code int `json:"code,omitempty"` // Error code
	Message string `json:"message,omitempty"` // Error message
}

// ClientFiltersRequest represents the ClientFiltersRequest schema from the OpenAPI specification
type ClientFiltersRequest struct {
	List []string `json:"list"`
}

// HeavyEquipmentsSearchResponse represents the HeavyEquipmentsSearchResponse schema from the OpenAPI specification
type HeavyEquipmentsSearchResponse struct {
	Listings []HeavyEquipmentsBaseListing `json:"listings,omitempty"`
	Num_found int `json:"num_found,omitempty"` // The number of listings found
	Range_facets map[string]interface{} `json:"range_facets,omitempty"`
	Stats map[string]interface{} `json:"stats,omitempty"`
	Facets map[string]interface{} `json:"facets,omitempty"`
}

// DealerRating represents the DealerRating schema from the OpenAPI specification
type DealerRating struct {
	Rating_components []RatingComponents `json:"rating_components,omitempty"`
	Street string `json:"street,omitempty"` // Street of the dealer
	City string `json:"city,omitempty"` // City of the dealer
	Name string `json:"name,omitempty"` // Name of the dealer
	Overall_rating float64 `json:"overall_rating,omitempty"` // Overall rating of the dealership on scale 1-5
	Zip string `json:"zip,omitempty"` // Zip of the dealer
	Longitude string `json:"longitude,omitempty"` // Longitude for the dealer location
	State string `json:"state,omitempty"` // State of the dealer
	Latitude string `json:"latitude,omitempty"` // Latutide for the dealer location
}

// MotorcycleSearchResponse represents the MotorcycleSearchResponse schema from the OpenAPI specification
type MotorcycleSearchResponse struct {
	Stats map[string]interface{} `json:"stats,omitempty"`
	Facets map[string]interface{} `json:"facets,omitempty"`
	Listings []MotorcycleBaseListing `json:"listings,omitempty"`
	Num_found int `json:"num_found,omitempty"` // The number of listings found
	Range_facets map[string]interface{} `json:"range_facets,omitempty"`
}

// BaseListing represents the BaseListing schema from the OpenAPI specification
type BaseListing struct {
	Model_code string `json:"model_code,omitempty"` // model_code of the listing
	Vin string `json:"vin,omitempty"` // VIN for the car
	Last_seen_at int `json:"last_seen_at,omitempty"` // Listing last seen at (most recent) timestamp
	Carfax_1_owner bool `json:"carfax_1_owner,omitempty"` // Flag to indicate whether listing is carfax_1_owner
	Build Build `json:"build,omitempty"` // Describes the Car specification
	First_seen_at_date string `json:"first_seen_at_date,omitempty"` // Listing first seen at first scraped date
	Scraped_at_date string `json:"scraped_at_date,omitempty"` // Listing last seen at date
	Msrp int `json:"msrp,omitempty"` // MSRP for the car
	Base_ext_color string `json:"base_ext_color,omitempty"` // Base exterior color of the car
	Id string `json:"id,omitempty"` // Unique identifier representing a specific listing from the Marketcheck database
	Last_seen_at_date string `json:"last_seen_at_date,omitempty"` // Listing last seen at (most recent) date
	Scraped_at float64 `json:"scraped_at,omitempty"` // Listing last seen at date timestamp
	Ref_price_dt int `json:"ref_price_dt,omitempty"` // The date at which the last price was reported online. This is earlier to last_seen_date
	Seller_type string `json:"seller_type,omitempty"` // Seller type for the car
	Dist float64 `json:"dist,omitempty"` // Distance of the car's location from the specified user lcoation
	Heading string `json:"heading,omitempty"` // Listing title as displayed on the source website
	In_transit bool `json:"in_transit,omitempty"` // in_transit of the listing
	Ref_price string `json:"ref_price,omitempty"` // Last reported price for the car. If the asking price value is not or is available then the last_price could perhaps be used. last_price is the price for the car listed on the source website as of last_price_dt date
	First_seen_at int `json:"first_seen_at,omitempty"` // Listing first seen at first scraped timestamp
	Ref_miles string `json:"ref_miles,omitempty"` // Last Odometer reading / reported miles usage for the car. If the asking miles value is not or is available then the last_miles could perhaps be used. last_miles is the miles for the car listed on the source website as of last_miles_dt date
	Stock_no string `json:"stock_no,omitempty"` // Stock number of car in dealers inventory
	Dom int `json:"dom,omitempty"` // Days on Market value for the car based on current and historical listings found in the Marketcheck database for this car
	Price_change_percent float64 `json:"price_change_percent,omitempty"` // Percentage difference between the cars's current price and ref_price i.e. last reported price
	Inventory_type string `json:"inventory_type,omitempty"` // Inventory type of car
	Is_translated bool `json:"is_translated,omitempty"` // is_translated of the listing
	Data_source string `json:"data_source,omitempty"` // Data source of the listing
	Dealer NestDealer `json:"dealer,omitempty"`
	Interior_color string `json:"interior_color,omitempty"` // Interior color of the car
	Ref_miles_dt int `json:"ref_miles_dt,omitempty"` // The date at which the last miles was reported online. This is earlier to last_seen_date
	Media ListingNestMedia `json:"media,omitempty"`
	Miles int `json:"miles,omitempty"` // Odometer reading / reported miles usage for the car
	Is_certified int `json:"is_certified,omitempty"` // Certified car
	Source string `json:"source,omitempty"` // Source domain of the listing
	Dom_active int `json:"dom_active,omitempty"` // Days on Market value for the car based on current and last 30 days listings found in the Marketcheck database for this car
	Vdp_url string `json:"vdp_url,omitempty"` // Vehicle Details Page url of the specific car
	Base_int_color string `json:"base_int_color,omitempty"` // Base interior color of the car
	Financing_options []ListingFinance `json:"financing_options,omitempty"` // Array of all finance offers for this listing
	Title_type string `json:"title_type,omitempty"` // title_type of the listing
	Price int `json:"price,omitempty"` // Asking price for the car
	Leasing_options []ListingLease `json:"leasing_options,omitempty"` // Array of all finance offers for this listing
	Exterior_color string `json:"exterior_color,omitempty"` // Exterior color of the car
	Dom_180 int `json:"dom_180,omitempty"` // Days on Market value for the car based on current and last 6 month listings found in the Marketcheck database for this car
	Carfax_clean_title bool `json:"carfax_clean_title,omitempty"` // Flag to indicate whether listing is carfax_clean_title
}

// CarListingRank represents the CarListingRank schema from the OpenAPI specification
type CarListingRank struct {
	Ranked_listing map[string]interface{} `json:"ranked_listing,omitempty"`
}

// HistoricalListing represents the HistoricalListing schema from the OpenAPI specification
type HistoricalListing struct {
	Msrp int `json:"msrp,omitempty"` // MSRP for the car
	Ref_price_dt int `json:"ref_price_dt,omitempty"` // The date at which the last price was reported online. This is earlier to last_seen_date
	Is_searchable string `json:"is_searchable,omitempty"` // Flag to indicate listing is marked searchable or not
	Dealer_id int `json:"dealer_id,omitempty"` // Unique MC assigned dealers id for the listing
	Zip string `json:"zip,omitempty"` // Zip of the listing
	Latitude string `json:"latitude,omitempty"` // Latitude of the listing
	Interior_color string `json:"interior_color,omitempty"` // Interior color of the car
	Seller_type string `json:"seller_type,omitempty"` // Seller type for the car
	Ref_price string `json:"ref_price,omitempty"` // Last reported price for the car. If the asking price value is not or is available then the last_price could perhaps be used. last_price is the price for the car listed on the source website as of last_price_dt date
	Id string `json:"id,omitempty"` // Unique identifier representing a specific listing from the Marketcheck database
	Ref_miles string `json:"ref_miles,omitempty"` // Last Odometer reading / reported miles usage for the car. If the asking miles value is not or is available then the last_miles could perhaps be used. last_miles is the miles for the car listed on the source website as of last_miles_dt date
	Carfax_1_owner bool `json:"carfax_1_owner,omitempty"` // Flag to indicate whether listing is carfax_1_owner
	Last_seen_at_date string `json:"last_seen_at_date,omitempty"` // Listing last seen at (most recent) date
	First_seen_at_date string `json:"first_seen_at_date,omitempty"` // Listing first seen at first scraped date
	Scraped_at_date string `json:"scraped_at_date,omitempty"` // Listing last seen at date
	Dom int `json:"dom,omitempty"` // Days on Market value for the car based on current and historical listings found in the Marketcheck database for this car
	Dom_180 int `json:"dom_180,omitempty"` // Days on Market value for the car based on current and last 6 month listings found in the Marketcheck database for this car
	Last_seen_at int `json:"last_seen_at,omitempty"` // Listing last seen at (most recent) timestamp
	Miles int `json:"miles,omitempty"` // Odometer reading / reported miles usage for the car
	Trim_r string `json:"trim_r,omitempty"` // Trim of the car
	Vdp_url string `json:"vdp_url,omitempty"` // Vehicle Details Page url of the specific car
	State string `json:"state,omitempty"` // State of the listing
	Price int `json:"price,omitempty"` // Asking price for the car
	Vin string `json:"vin,omitempty"` // VIN for the car
	Carfax_clean_title bool `json:"carfax_clean_title,omitempty"` // Flag to indicate whether listing is carfax_clean_title
	City string `json:"city,omitempty"` // City of the listing
	Heading string `json:"heading,omitempty"` // Listing title as displayed on the source website
	Stock_no string `json:"stock_no,omitempty"` // Stock number of car in dealers inventory
	Is_certified int `json:"is_certified,omitempty"` // Certified car
	Seller_name_o string `json:"seller_name_o,omitempty"` // Seller name of the listing from the Marketcheck database
	First_seen_at int `json:"first_seen_at,omitempty"` // Listing first seen at first scraped timestamp
	Ref_miles_dt int `json:"ref_miles_dt,omitempty"` // The date at which the last miles was reported online. This is earlier to last_seen_date
	Dom_active int `json:"dom_active,omitempty"` // Days on Market value for the car based on current and last 30 days listings found in the Marketcheck database for this car
	Scraped_at int `json:"scraped_at,omitempty"` // Listing last seen at date timestamp
	Street string `json:"street,omitempty"` // Street of the listing
	Inventory_type string `json:"inventory_type,omitempty"` // Inventory type of car
	Leasing_options []ListingLease `json:"leasing_options,omitempty"` // Array of all finance offers for this listing
	Longitude string `json:"longitude,omitempty"` // Longitude of the listing
	Source string `json:"source,omitempty"` // Source domain of the listing
	Seller_name string `json:"seller_name,omitempty"` // Seller name of the listing from the Marketcheck database
	Exterior_color string `json:"exterior_color,omitempty"` // Exterior color of the car
	Data_source string `json:"data_source,omitempty"` // Data source of the listing
	Status_date int `json:"status_date,omitempty"` // Timestamp of status date
	Financing_options []ListingFinance `json:"financing_options,omitempty"` // Array of all finance offers for this listing
}

// ListingExtraAttributes represents the ListingExtraAttributes schema from the OpenAPI specification
type ListingExtraAttributes struct {
	Interior_f []string `json:"interior_f,omitempty"` // List of interior features available with the car
	Features []string `json:"features,omitempty"` // List of Features available with the car
	Options []string `json:"options,omitempty"` // Installed Options of the car
	Safety_f []string `json:"safety_f,omitempty"` // List of safety features available with the car
	Electronics_f []string `json:"electronics_f,omitempty"` // List of electronic features available with the car
	Exterior_f []string `json:"exterior_f,omitempty"` // List of exterior features available with the car
	Id string `json:"id,omitempty"` // Unique identifier representing a specific listing from the Marketcheck database
	Standard_f []string `json:"standard_f,omitempty"` // List of standard features available with the car
	Dealer_added_f []string `json:"dealer_added_f,omitempty"` // List of dealer added features available with the car
	Seller_cmts string `json:"seller_cmts,omitempty"` // Seller comment for the car
	Technical_f []string `json:"technical_f,omitempty"` // List of technical features available with the car
}

// DailyStatsNestedJson represents the DailyStatsNestedJson schema from the OpenAPI specification
type DailyStatsNestedJson struct {
	Interquartile_range float64 `json:"interquartile_range,omitempty"` // Interquartile range
	Mean float64 `json:"mean,omitempty"` // Mean
	Median float64 `json:"median,omitempty"` // Median
	Population_standard_deviation float64 `json:"population_standard_deviation,omitempty"` // population standard deviation
	Standard_deviation float64 `json:"standard_deviation,omitempty"` // Standard Deviation
	Trimmed_mean float64 `json:"trimmed_mean,omitempty"` // Trimmed Mean
	Variance float64 `json:"variance,omitempty"` // variance
}

// Economy represents the Economy schema from the OpenAPI specification
type Economy struct {
	Ghg_emissions_rating float64 `json:"ghg_emissions_rating,omitempty"` // Show vehicle impact on climate change in terms of greeenhouse gases. This rating will be in the range of 1 - 10
	Smog_indicator string `json:"smog_indicator,omitempty"` // Contains text like bad|Good|Very Good| Excellent
	Smog_rating float64 `json:"smog_rating,omitempty"` // Shows vehicle's emissions of pollutants known to cause smog and other forms of air pollution. This rating will be in the range of 1 - 10
	Ghg_emissions_indicator string `json:"ghg_emissions_indicator,omitempty"` // Contains text like bad|Good|Very Good| Excellent
}

// MakeModel represents the MakeModel schema from the OpenAPI specification
type MakeModel struct {
	MakeField string `json:"make,omitempty"`
	Model string `json:"model,omitempty"`
}

// FareValue represents the FareValue schema from the OpenAPI specification
type FareValue struct {
	Avg_days_to_sold_national int `json:"avg_days_to_sold_national,omitempty"` // Average days to sold of car for all available listings
	Fmv_local int `json:"fmv_local,omitempty"` // Fare value calculated for all available listing around given postal code & radius
	Fmv_national int `json:"fmv_national,omitempty"` // Fare value calculated for all available listing
	Avg_days_to_sold_local int `json:"avg_days_to_sold_local,omitempty"` // Average days to sold of car around given postal code & radius
}

// MotorcycleListing represents the MotorcycleListing schema from the OpenAPI specification
type MotorcycleListing struct {
	Miles int `json:"miles,omitempty"` // Odometer reading / reported miles usage for the motorcycle
	Dp_url string `json:"dp_url,omitempty"` // Details Page url of the specific motorcycle
	First_seen_at_date string `json:"first_seen_at_date,omitempty"` // Listing first seen at first scraped date
	Extra ListingNestExtraAttributes `json:"extra,omitempty"`
	Media ListingNestMedia `json:"media,omitempty"`
	First_seen_at int `json:"first_seen_at,omitempty"` // Listing first seen at first scraped timestamp
	Build MotorcycleBuild `json:"build,omitempty"` // Describes the Motorcycle specification
	Vin string `json:"vin,omitempty"` // VIN for the Motorcycle
	Source string `json:"source,omitempty"` // Source domain of the listing
	Seller_type string `json:"seller_type,omitempty"` // Seller type for the motorcycle
	Dealer NestDealer `json:"dealer,omitempty"`
	Last_seen_at int `json:"last_seen_at,omitempty"` // Listing last seen at (most recent) timestamp
	Last_seen_at_date string `json:"last_seen_at_date,omitempty"` // Listing last seen at (most recent) date
	Stock_no string `json:"stock_no,omitempty"` // Stock number of motorcycle in dealers inventory
	Heading string `json:"heading,omitempty"` // Listing title as displayed on the source website
	Price int `json:"price,omitempty"` // Asking price for the motorcycle
	Msrp int `json:"msrp,omitempty"` // MSRP for the motorcycle
	Id string `json:"id,omitempty"` // Unique identifier representing a specific listing from the Marketcheck database
	Scraped_at_date string `json:"scraped_at_date,omitempty"` // Listing last seen at date
	Color string `json:"color,omitempty"` // Color of the motorcycle
	Inventory_type string `json:"inventory_type,omitempty"` // Inventory type of motorcycle
	Scraped_at float64 `json:"scraped_at,omitempty"` // Listing last seen at date timestamp
}

// CarLocation represents the CarLocation schema from the OpenAPI specification
type CarLocation struct {
	Latitude string `json:"latitude,omitempty"` // Latitude component of car location
	Longitude string `json:"longitude,omitempty"` // Longitude component of car location
	Seller_name string `json:"seller_name,omitempty"` // Seller name of the car
	Street string `json:"street,omitempty"` // Street address of the car
	Zip string `json:"zip,omitempty"` // Zip/Postal code of the car
	City string `json:"city,omitempty"` // City of the car
	County string `json:"county,omitempty"` // County of the car
}

// RangeFacetItem represents the RangeFacetItem schema from the OpenAPI specification
type RangeFacetItem struct {
	Start int `json:"start,omitempty"` // The start of range
	After int `json:"after,omitempty"` // Number of listings that have field value after range end
	Before int `json:"before,omitempty"` // Number of listings that have field value before range start
	Between int `json:"between,omitempty"` // Number of listings that have field value between range start and range end
	Counts []map[string]interface{} `json:"counts,omitempty"`
	End int `json:"end,omitempty"` // The end of range
	Gap int `json:"gap,omitempty"` // The range gap applied for range faceting
}

// PopularCars represents the PopularCars schema from the OpenAPI specification
type PopularCars struct {
	Used_top50 []BasePopular `json:"used_top50,omitempty"` // Array of all used top 50 cars
	New_top50 []BasePopular `json:"new_top50,omitempty"` // Array of all new top 50 cars
}

// TrendPoint represents the TrendPoint schema from the OpenAPI specification
type TrendPoint struct {
	Units float64 `json:"units,omitempty"` // Units
	Dom float64 `json:"dom,omitempty"` // Days on Market
	Miles float64 `json:"miles,omitempty"` // Miles
	Msrp float64 `json:"msrp,omitempty"` // MSRP
	Price float64 `json:"price,omitempty"` // Price
}

// Location represents the Location schema from the OpenAPI specification
type Location struct {
	City string `json:"city,omitempty"` // City of the listing
	Latitude float64 `json:"latitude,omitempty"` // Latitude component of the location
	Longitude float64 `json:"longitude,omitempty"` // Longitude component of the location
	State string `json:"state,omitempty"` // State of the listing
	Zip string `json:"zip,omitempty"` // Zip of the listing
}

// RatingComponents represents the RatingComponents schema from the OpenAPI specification
type RatingComponents struct {
	Actual_rating float64 `json:"actual_rating,omitempty"` // rating of car on that condition
	Rating_condition string `json:"rating_condition,omitempty"` // Specifices rating condition parameter
}

// UKRVBuild represents the UKRVBuild schema from the OpenAPI specification
type UKRVBuild struct {
	Transmission string `json:"transmission,omitempty"` // Transmission of the RV
	Width string `json:"width,omitempty"` // width of RV
	Area string `json:"area,omitempty"` // Area of the RV
	Year int `json:"year,omitempty"` // Year of the RV
	Gvwr string `json:"gvwr,omitempty"` // Gross Vehicle Weight Ratio of RV
	Category string `json:"category,omitempty"` // category of the RV
	Engine_size string `json:"engine_size,omitempty"` // engine_size of RV
	MakeField string `json:"make,omitempty"` // RV Make
	Slideouts string `json:"slideouts,omitempty"` // Slideouts of the RV
	Chassis string `json:"chassis,omitempty"` // chassis of RV
	Berths string `json:"berths,omitempty"` // berths of RV
	Drive_type string `json:"drive_type,omitempty"` // drive_type of RV
	Exterior_length string `json:"exterior_length,omitempty"` // exterior_length of RV
	Floorplan string `json:"floorplan,omitempty"` // floorplan of RV
	Floorplan_layout string `json:"floorplan_layout,omitempty"` // floorplan_layout of RV
	Inventory_type string `json:"inventory_type,omitempty"` // Inventory type of the RV
	Model string `json:"model,omitempty"` // RV model
	Mtplm string `json:"mtplm,omitempty"` // mtplm of RV
	Steering string `json:"steering,omitempty"` // steering of RV
	Drivetrain string `json:"drivetrain,omitempty"` // Drivetrain of the RV
	Fuel_type string `json:"fuel_type,omitempty"` // Fuel type of the RV
	Engine string `json:"engine,omitempty"` // Engine of the RV
	Interior_length string `json:"interior_length,omitempty"` // interior_length of RV
	Sub_category string `json:"sub_category,omitempty"` // sub_category of the RV
}

// CRMResponse represents the CRMResponse schema from the OpenAPI specification
type CRMResponse struct {
	For_sale string `json:"for_sale,omitempty"` // Boolean to indicate whether given vin has had listing after given date or not
}

// NestDealer represents the NestDealer schema from the OpenAPI specification
type NestDealer struct {
	City string `json:"city,omitempty"` // City of the dealer
	Latitude string `json:"latitude,omitempty"` // Latutide for the dealer location
	Phone string `json:"phone,omitempty"` // Contact no of the dealer
	Seller_email string `json:"seller_email,omitempty"` // Contact email of the dealer
	Zip string `json:"zip,omitempty"` // Zip of the dealer
	Country string `json:"country,omitempty"` // country of the dealer
	Dealer_type string `json:"dealer_type,omitempty"` // Type of dealer (franchise/independent)
	State string `json:"state,omitempty"` // State of the dealer
	County string `json:"county,omitempty"` // county of the dealer
	Website string `json:"website,omitempty"` // Website of the dealer
	Dealership_group_name string `json:"dealership_group_name,omitempty"` // dealership_group_name of dealer
	Name string `json:"name,omitempty"` // Name of the dealer
	Street string `json:"street,omitempty"` // Street of the dealer
	Id int `json:"id,omitempty"` // The unique id associated with the dealer in the Marketcheck database
	Longitude string `json:"longitude,omitempty"` // Longitude for the dealer location
	Msa_code string `json:"msa_code,omitempty"` // msa_code of the dealer
}

// UKBaseListing represents the UKBaseListing schema from the OpenAPI specification
type UKBaseListing struct {
	Insurance_group string `json:"insurance_group,omitempty"` // Insurance Group of car
	Vdp_url string `json:"vdp_url,omitempty"` // Vehicle Details Page url of the specific car
	Source string `json:"source,omitempty"` // Source domain of the listing
	Heading string `json:"heading,omitempty"` // Listing title as displayed on the source website
	Leasing_options []ListingLease `json:"leasing_options,omitempty"` // Array of all finance offers for this listing
	Price_change_percent float64 `json:"price_change_percent,omitempty"` // Percentage difference between the cars's current price and ref_price i.e. last reported price
	Exterior_color string `json:"exterior_color,omitempty"` // Exterior color of the car
	Miles int `json:"miles,omitempty"` // Odometer reading / reported miles usage for the car
	Seller_type string `json:"seller_type,omitempty"` // Seller type for the car
	Dealer NestDealer `json:"dealer,omitempty"`
	Scraped_at float64 `json:"scraped_at,omitempty"` // Listing last seen at date timestamp
	Id string `json:"id,omitempty"` // Unique identifier representing a specific listing from the Marketcheck database
	Ref_price string `json:"ref_price,omitempty"` // Last reported price for the car. If the asking price value is not or is available then the last_price could perhaps be used. last_price is the price for the car listed on the source website as of last_price_dt date
	Num_owners string `json:"num_owners,omitempty"` // Number of owners
	Financing_options []ListingFinance `json:"financing_options,omitempty"` // Array of all finance offers for this listing
	Stock_no string `json:"stock_no,omitempty"` // Stock number of car in dealers inventory
	Dist float64 `json:"dist,omitempty"` // Distance of the car's location from the specified user lcoation
	Ref_miles_dt int `json:"ref_miles_dt,omitempty"` // The date at which the last miles was reported online. This is earlier to last_seen_date
	Build Build `json:"build,omitempty"` // Describes the Car specification
	Car_location CarLocation `json:"car_location,omitempty"`
	Uvc_id string `json:"uvc_id,omitempty"` // UVC ID of the car
	Last_seen_at int `json:"last_seen_at,omitempty"` // Listing last seen at (most recent) timestamp
	Co2_emissions string `json:"co2_emissions,omitempty"` // CO2 emissions by car
	Is_translated bool `json:"is_translated,omitempty"` // is_translated of the listing
	Carfax_1_owner bool `json:"carfax_1_owner,omitempty"` // Flag to indicate whether listing is carfax_1_owner
	Dom int `json:"dom,omitempty"` // Days on Market value for the car based on current and historical listings found in the Marketcheck database for this car
	Dom_180 int `json:"dom_180,omitempty"` // Days on Market value for the car based on current and last 6 month listings found in the Marketcheck database for this car
	In_transit bool `json:"in_transit,omitempty"` // in_transit of the listing
	Interior_color string `json:"interior_color,omitempty"` // Interior color of the car
	Title_type string `json:"title_type,omitempty"` // title_type of the listing
	Ref_miles string `json:"ref_miles,omitempty"` // Last Odometer reading / reported miles usage for the car. If the asking miles value is not or is available then the last_miles could perhaps be used. last_miles is the miles for the car listed on the source website as of last_miles_dt date
	Price int `json:"price,omitempty"` // Asking price for the car
	Carfax_clean_title bool `json:"carfax_clean_title,omitempty"` // Flag to indicate whether listing is carfax_clean_title
	First_seen_at int `json:"first_seen_at,omitempty"` // Listing first seen at first scraped timestamp
	Is_certified int `json:"is_certified,omitempty"` // Certified car
	Last_seen_at_date string `json:"last_seen_at_date,omitempty"` // Listing last seen at (most recent) date
	Inventory_type string `json:"inventory_type,omitempty"` // Inventory type of car
	Base_int_color string `json:"base_int_color,omitempty"` // Base interior color of the car
	Model_code string `json:"model_code,omitempty"` // model_code of the listing
	Media ListingNestMedia `json:"media,omitempty"`
	Ref_price_dt int `json:"ref_price_dt,omitempty"` // The date at which the last price was reported online. This is earlier to last_seen_date
	Scraped_at_date string `json:"scraped_at_date,omitempty"` // Listing last seen at date
	Dom_active int `json:"dom_active,omitempty"` // Days on Market value for the car based on current and last 30 days listings found in the Marketcheck database for this car
	Msrp int `json:"msrp,omitempty"` // MSRP for the car
	Base_ext_color string `json:"base_ext_color,omitempty"` // Base exterior color of the car
	Data_source string `json:"data_source,omitempty"` // Data source of the listing
	First_seen_at_date string `json:"first_seen_at_date,omitempty"` // Listing first seen at first scraped date
	Vehicle_registration_mark string `json:"vehicle_registration_mark,omitempty"` // Vehicle Registration Mark of the car
}

// RVSearchResponse represents the RVSearchResponse schema from the OpenAPI specification
type RVSearchResponse struct {
	Range_facets map[string]interface{} `json:"range_facets,omitempty"`
	Stats map[string]interface{} `json:"stats,omitempty"`
	Facets map[string]interface{} `json:"facets,omitempty"`
	Listings []RVBaseListing `json:"listings,omitempty"`
	Num_found int `json:"num_found,omitempty"` // The number of listings found
}

// DepreciationStats represents the DepreciationStats schema from the OpenAPI specification
type DepreciationStats struct {
	Five_year_from_now_percent float64 `json:"five_year_from_now_percent,omitempty"` // price depreciation percent after five year from now
	Name string `json:"name,omitempty"` // ymm_comb_name
	One_year_from_now float64 `json:"one_year_from_now,omitempty"` // price after one year from now
	One_year_from_now_percent float64 `json:"one_year_from_now_percent,omitempty"` // price depreciation percent after one year from now
	Two_year_from_now float64 `json:"two_year_from_now,omitempty"` // price after two year from now
	Two_year_from_now_percent float64 `json:"two_year_from_now_percent,omitempty"` // price depreciation percent after two year from now
	Current_value float64 `json:"current_value,omitempty"` // Price of year make model combination
	Five_year_from_now float64 `json:"five_year_from_now,omitempty"` // price after five year from now
}

// SpecsAutoCompleteResponse represents the SpecsAutoCompleteResponse schema from the OpenAPI specification
type SpecsAutoCompleteResponse struct {
	Terms []string `json:"terms,omitempty"`
}

// CompetitorsCarDetails represents the CompetitorsCarDetails schema from the OpenAPI specification
type CompetitorsCarDetails struct {
	Miles string `json:"miles,omitempty"` // mileage of car
	Name string `json:"name,omitempty"` // Current car name
	Price string `json:"price,omitempty"` // price of car
	Avg_market_value string `json:"avg_market_value,omitempty"` // Estimated market value
}

// DailyStats represents the DailyStats schema from the OpenAPI specification
type DailyStats struct {
	Price_stats DailyStatsNestedJson `json:"price_stats,omitempty"` // All nested JSON of stats
	Units_for_sale float64 `json:"units_for_sale,omitempty"` // Number of units of this car for sale on the market
	Dom DailyStatsNestedJson `json:"dom,omitempty"` // All nested JSON of stats
	Miles_stats DailyStatsNestedJson `json:"miles_stats,omitempty"` // All nested JSON of stats
}

// PredictedSpecs represents the PredictedSpecs schema from the OpenAPI specification
type PredictedSpecs struct {
	Cylinders int `json:"cylinders,omitempty"` // Number of cylinders
	Engine_block string `json:"engine_block,omitempty"` // Engine block of the vehicle
	Transmission string `json:"transmission,omitempty"` // transmission
	Miles int `json:"miles,omitempty"` // Miles
	Latitude float64 `json:"latitude,omitempty"` // Latutide for the vehicle location
	City_mpg float64 `json:"city_mpg,omitempty"` // City mileage
	Carfax_clean_title bool `json:"carfax_clean_title,omitempty"` // Indicates whether car has clean ownership records
	Doors int `json:"doors,omitempty"` // Number of doors
	Engine_size float64 `json:"engine_size,omitempty"` // Engine size
	Highway_mpg float64 `json:"highway_mpg,omitempty"` // Highway mileage
	Base_interior_color string `json:"base_interior_color,omitempty"` // base interior color of vehicle
	Longitude float64 `json:"longitude,omitempty"` // Longitude for the vehicle location
	MakeField string `json:"make,omitempty"` // make
	Base_exterior_color string `json:"base_exterior_color,omitempty"` // Base exterior color of vehicle
	Is_certified bool `json:"is_certified,omitempty"` // Certified vehicle
	Model string `json:"model,omitempty"` // model
	Trim string `json:"trim,omitempty"` // trim
	Drivetrain string `json:"drivetrain,omitempty"` // Drivetrain of the vehicle
	Year int `json:"year,omitempty"` // year
	Carfax_1_owner bool `json:"carfax_1_owner,omitempty"` // Indicates whether car has had only one owner or not
}

// FacetItem represents the FacetItem schema from the OpenAPI specification
type FacetItem struct {
	Count int `json:"count,omitempty"` // Facet item count
	Item string `json:"item,omitempty"` // Facet item
}

// MotorcycleSearchRangeFacets represents the MotorcycleSearchRangeFacets schema from the OpenAPI specification
type MotorcycleSearchRangeFacets struct {
	Price map[string]interface{} `json:"price,omitempty"`
	Miles map[string]interface{} `json:"miles,omitempty"`
}

// ListingNestExtraAttributes represents the ListingNestExtraAttributes schema from the OpenAPI specification
type ListingNestExtraAttributes struct {
	Standard_f []string `json:"standard_f,omitempty"` // List of standard features available with the car
	Exterior_f []string `json:"exterior_f,omitempty"` // List of exterior features available with the car
	Features []string `json:"features,omitempty"` // List of Features available with the car
	Safety_f []string `json:"safety_f,omitempty"` // List of safety features available with the car
	Seller_comments string `json:"seller_comments,omitempty"` // Seller comments for the car
	Technical_f []string `json:"technical_f,omitempty"` // List of technical features available with the car
	Dealer_added_f []string `json:"dealer_added_f,omitempty"` // List of dealer added features available with the car
	Electronics_f []string `json:"electronics_f,omitempty"` // List of electronic features available with the car
	Interior_f []string `json:"interior_f,omitempty"` // List of interior features available with the car
	Options []string `json:"options,omitempty"` // Installed Options of the car
}

// CarSearchRangeFacets represents the CarSearchRangeFacets schema from the OpenAPI specification
type CarSearchRangeFacets struct {
	Finance_loan_term map[string]interface{} `json:"finance_loan_term,omitempty"`
	Lease_down_payment map[string]interface{} `json:"lease_down_payment,omitempty"`
	Finance_down_payment map[string]interface{} `json:"finance_down_payment,omitempty"`
	Lease_term map[string]interface{} `json:"lease_term,omitempty"`
	Miles map[string]interface{} `json:"miles,omitempty"`
	Price map[string]interface{} `json:"price,omitempty"`
	Dom map[string]interface{} `json:"dom,omitempty"`
	Dom_active map[string]interface{} `json:"dom_active,omitempty"`
	Finance_emp map[string]interface{} `json:"finance_emp,omitempty"`
	Finance_loan_apr map[string]interface{} `json:"finance_loan_apr,omitempty"`
	Lease_emp map[string]interface{} `json:"lease_emp,omitempty"`
	Msrp map[string]interface{} `json:"msrp,omitempty"`
	Dom_180 map[string]interface{} `json:"dom_180,omitempty"`
}

// PlotPoint represents the PlotPoint schema from the OpenAPI specification
type PlotPoint struct {
	Seller_name string `json:"seller_name,omitempty"` // Seller name
	Vin string `json:"vin,omitempty"` // Vin
	Dom float64 `json:"dom,omitempty"` // DOM
	Id string `json:"id,omitempty"` // Listing id
	Miles float64 `json:"miles,omitempty"` // Miles
	Msrp float64 `json:"msrp,omitempty"` // Msrp
	Price float64 `json:"price,omitempty"` // Price
}

// ListingMedia represents the ListingMedia schema from the OpenAPI specification
type ListingMedia struct {
	Id string `json:"id,omitempty"` // Unique identifier representing a specific listing from the Marketcheck database
	Photo_links []string `json:"photo_links,omitempty"` // A list of photo urls for the car
	Photo_url string `json:"photo_url,omitempty"` // Single photo url of the car
}

// HeavyEquipmentsListing represents the HeavyEquipmentsListing schema from the OpenAPI specification
type HeavyEquipmentsListing struct {
	Last_seen_at_date string `json:"last_seen_at_date,omitempty"` // Listing last seen at (most recent) date
	Vin string `json:"vin,omitempty"` // VIN for the heavy equipment
	First_seen_at int `json:"first_seen_at,omitempty"` // Listing first seen at first scraped timestamp
	Interior_color string `json:"interior_color,omitempty"` // Interior color of the heavy equipment
	Extra ListingNestExtraAttributes `json:"extra,omitempty"`
	Miles int `json:"miles,omitempty"` // Odometer reading / reported miles usage for the heavy equipment
	Scraped_at_date string `json:"scraped_at_date,omitempty"` // Listing last seen at date
	Build RVBuild `json:"build,omitempty"` // Describes the RV specification
	Dealer NestDealer `json:"dealer,omitempty"`
	Id string `json:"id,omitempty"` // Unique identifier representing a specific listing from the Marketcheck database
	Inventory_type string `json:"inventory_type,omitempty"` // Inventory type of heavy equipment
	Msrp int `json:"msrp,omitempty"` // MSRP for the heavy equipment
	Stock_no string `json:"stock_no,omitempty"` // Stock number of heavy equipment in dealers inventory
	Source string `json:"source,omitempty"` // Source domain of the listing
	Exterior_color string `json:"exterior_color,omitempty"` // Exterior color of the heavy equipment
	Scraped_at float64 `json:"scraped_at,omitempty"` // Listing last seen at date timestamp
	Price int `json:"price,omitempty"` // Asking price for the heavy equipment
	Seller_type string `json:"seller_type,omitempty"` // Seller type for the heavy equipment
	Media ListingNestMedia `json:"media,omitempty"`
	Dp_url string `json:"dp_url,omitempty"` // Details Page url of the specific heavy equipment
	First_seen_at_date string `json:"first_seen_at_date,omitempty"` // Listing first seen at first scraped date
	Last_seen_at int `json:"last_seen_at,omitempty"` // Listing last seen at (most recent) timestamp
	Heading string `json:"heading,omitempty"` // Listing title as displayed on the source website
}

// Mds represents the Mds schema from the OpenAPI specification
type Mds struct {
	Sold_vins []string `json:"sold_vins,omitempty"` // Sold vins array
	Total_active_cars_for_ymmt int `json:"total_active_cars_for_ymmt,omitempty"` // Active cars for ymmt combination
	Total_cars_sold_in_last_45_days int `json:"total_cars_sold_in_last_45_days,omitempty"` // Cars sold in last 45 days
	Trim string `json:"trim,omitempty"` // Trim of vin provided
	Year int `json:"year,omitempty"` // Year of vin provided
	MakeField string `json:"make,omitempty"` // Make of vin provided
	Mds int `json:"mds,omitempty"` // Provides Market days supply count
	Model string `json:"model,omitempty"` // Model of vin provided
}

// CarRankResponse represents the CarRankResponse schema from the OpenAPI specification
type CarRankResponse struct {
	Ranked_listings []CarListingRank `json:"ranked_listings,omitempty"`
	Num_ranked int `json:"num_ranked,omitempty"` // The number of listings ranked
}

// SetClientFiltersResponse represents the SetClientFiltersResponse schema from the OpenAPI specification
type SetClientFiltersResponse struct {
	Message string `json:"message,omitempty"` // Description of operation
	Status int `json:"status,omitempty"` // status code of operation
}

// Dealer represents the Dealer schema from the OpenAPI specification
type Dealer struct {
	Seller_phone string `json:"seller_phone,omitempty"` // Contact no of the dealer
	Street string `json:"street,omitempty"` // Street of the dealer
	Id string `json:"id,omitempty"` // The unique id associated with the dealer in the Marketcheck database
	Latitude string `json:"latitude,omitempty"` // Latutide for the dealer location
	State string `json:"state,omitempty"` // State of the dealer
	Data_source string `json:"data_source,omitempty"` // Datasource of the dealer
	Dealership_group_name string `json:"dealership_group_name,omitempty"` // Dealership group name of dealer
	Country string `json:"country,omitempty"` // country of the dealer
	Dealer_type string `json:"dealer_type,omitempty"` // Type of dealer (franchise/independent)
	Status string `json:"status,omitempty"` // Status of the dealer
	Location_ll string `json:"location_ll,omitempty"` // location_ll for the dealer location
	Inventory_url string `json:"inventory_url,omitempty"` // Website of the dealer
	Listing_count int `json:"listing_count,omitempty"` // Listing count of the dealer
	Seller_name string `json:"seller_name,omitempty"` // Name of the dealer
	Longitude string `json:"longitude,omitempty"` // Longitude for the dealer location
	Distance float64 `json:"distance,omitempty"` // Distance of dealer from given location
	Zip string `json:"zip,omitempty"` // Zip of the dealer
	Seller_email string `json:"seller_email,omitempty"` // Contact email of the dealer
	City string `json:"city,omitempty"` // City of the dealer
}

// Averages represents the Averages schema from the OpenAPI specification
type Averages struct {
	Msrp float64 `json:"msrp,omitempty"` // Mean msrp value for the car
	Price float64 `json:"price,omitempty"` // Mean price value for the car
	Units float64 `json:"units,omitempty"` // No of units for this car on the market
	Age float64 `json:"age,omitempty"` // Average age of the car
	Dom float64 `json:"dom,omitempty"` // Average Days on Market value for the car based on all listings found in the Marketcheck database for this car
	Miles float64 `json:"miles,omitempty"` // Mean miles value for the car
}

// CarRankCriteria represents the CarRankCriteria schema from the OpenAPI specification
type CarRankCriteria struct {
	Carfax_1_owner float64 `json:"carfax_1_owner,omitempty"` // Flag to indicate whether listing is carfax_1_owner.Weight for this is ranking process between 0-1.
	Carfax_clean_title float64 `json:"carfax_clean_title,omitempty"` // Flag to indicate whether listing is carfax_clean_title.Weight for this is ranking process between 0-1
	Dom float64 `json:"dom,omitempty"` // Days on Market value for the car based on current and historical listings found in the Marketcheck database for this car.Weight for this is ranking process between 0-1
	Dom_180 float64 `json:"dom_180,omitempty"` // Days on Market value for the car based on current and last 6 month listings found in the Marketcheck database for this car.Weight for this is ranking process between 0-1
	Dom_active float64 `json:"dom_active,omitempty"` // Days on Market value for the car based on current and last 30 days listings found in the Marketcheck database for this car.Weight for this is ranking process between 0-1
	Is_certified float64 `json:"is_certified,omitempty"` // Certified car.Weight for this is ranking process between 0-1
	Miles float64 `json:"miles,omitempty"` // Odometer reading / reported miles usage for the car.Weight for this is ranking process between 0-1
	Price float64 `json:"price,omitempty"` // Asking price for the vehicle. Weight for this is ranking process between 0-1.
}

// CarRankRequest represents the CarRankRequest schema from the OpenAPI specification
type CarRankRequest struct {
	Listing_ids []string `json:"listing_ids,omitempty"`
	Ranking_criteria map[string]interface{} `json:"ranking_criteria,omitempty"`
}

// SafetyRating represents the SafetyRating schema from the OpenAPI specification
type SafetyRating struct {
	Driver_front_rating float64 `json:"driver_front_rating,omitempty"` // Driver front rating of the Listing on scale 1-5
	Driver_side_rating float64 `json:"driver_side_rating,omitempty"` // Driver Side rating of the Listing on scale 1-5
	Overall_rating float64 `json:"overall_rating,omitempty"` // Overall rating of the Listing on scale 1-5
	Passenger_front_rating float64 `json:"passenger_front_rating,omitempty"` // Passenger front rating of the Listing on scale 1-5
	Roll_over_rating float64 `json:"roll_over_rating,omitempty"` // Roll Over rating of the Listing on scale 1-5
}

// MotorcycleSearchStats represents the MotorcycleSearchStats schema from the OpenAPI specification
type MotorcycleSearchStats struct {
	Miles StatsItem `json:"miles,omitempty"`
	Price StatsItem `json:"price,omitempty"`
}

// GetClientFiltersResponse represents the GetClientFiltersResponse schema from the OpenAPI specification
type GetClientFiltersResponse struct {
	List string `json:"list,omitempty"` // Description of operation
	Status int `json:"status,omitempty"` // status code of operation
}

// CompetitorsSimilarCars represents the CompetitorsSimilarCars schema from the OpenAPI specification
type CompetitorsSimilarCars struct {
	Name string `json:"name,omitempty"` // ymm_comb_name
	Avg_market_value float64 `json:"avg_market_value,omitempty"` // Average market value of year make model combination
	Avg_miles float64 `json:"avg_miles,omitempty"` // Average mileage of year make model combination
	Avg_price float64 `json:"avg_price,omitempty"` // Average price of year make model combination
}

// UKCarSearchFacets represents the UKCarSearchFacets schema from the OpenAPI specification
type UKCarSearchFacets struct {
	Seller_name []FacetItem `json:"seller_name,omitempty"`
	Base_exterior_color []FacetItem `json:"base_exterior_color,omitempty"`
	Vrm []FacetItem `json:"vrm,omitempty"`
	Year []FacetItem `json:"year,omitempty"`
	Cylinders []FacetItem `json:"cylinders,omitempty"`
	Doors []FacetItem `json:"doors,omitempty"`
	Engine []FacetItem `json:"engine,omitempty"`
	Engine_aspiration []FacetItem `json:"engine_aspiration,omitempty"`
	Car_location_city []FacetItem `json:"car_location_city,omitempty"`
	Insurance_group []FacetItem `json:"insurance_group,omitempty"`
	Base_interior_color []FacetItem `json:"base_interior_color,omitempty"`
	MakeField []FacetItem `json:"make,omitempty"`
	Model []FacetItem `json:"model,omitempty"`
	In_transit []FacetItem `json:"in_transit,omitempty"`
	Powertrain_type []FacetItem `json:"powertrain_type,omitempty"`
	Source []FacetItem `json:"source,omitempty"`
	Seating_capacity []FacetItem `json:"seating_capacity,omitempty"`
	City []FacetItem `json:"city,omitempty"`
	Body_subtype []FacetItem `json:"body_subtype,omitempty"`
	Engine_size []FacetItem `json:"engine_size,omitempty"`
	Fuel_type []FacetItem `json:"fuel_type,omitempty"`
	Interior_color []FacetItem `json:"interior_color,omitempty"`
	State []FacetItem `json:"state,omitempty"`
	Dealership_group_name []FacetItem `json:"dealership_group_name,omitempty"`
	Carfax_1_owner []FacetItem `json:"carfax_1_owner,omitempty"`
	Dealer_type []FacetItem `json:"dealer_type,omitempty"`
	Engine_block []FacetItem `json:"engine_block,omitempty"`
	Seller_type []FacetItem `json:"seller_type,omitempty"`
	Carfax_clean_title []FacetItem `json:"carfax_clean_title,omitempty"`
	Data_source []FacetItem `json:"data_source,omitempty"`
	Car_type []FacetItem `json:"car_type,omitempty"`
	Trim []FacetItem `json:"trim,omitempty"`
	Exterior_color []FacetItem `json:"exterior_color,omitempty"`
	Drivetrain []FacetItem `json:"drivetrain,omitempty"`
	Car_location_zip []FacetItem `json:"car_location_zip,omitempty"`
	Vehicle_registration_mark []FacetItem `json:"vehicle_registration_mark,omitempty"`
	Co2_emissions []FacetItem `json:"co2_emissions,omitempty"`
	Vehicle_type []FacetItem `json:"vehicle_type,omitempty"`
	Fca_status []FacetItem `json:"fca_status,omitempty"`
	Transmission []FacetItem `json:"transmission,omitempty"`
	Mas_code []FacetItem `json:"mas_code,omitempty"`
	Num_owners []FacetItem `json:"num_owners,omitempty"`
	Body_type []FacetItem `json:"body_type,omitempty"`
	Car_location_seller_name []FacetItem `json:"car_location_seller_name,omitempty"`
	Dealer_id []FacetItem `json:"dealer_id,omitempty"`
	Car_location_county []FacetItem `json:"car_location_county,omitempty"`
	Car_location_street []FacetItem `json:"car_location_street,omitempty"`
}
