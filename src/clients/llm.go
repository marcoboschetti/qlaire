package clients

import "github.com/marcoboschetti/qlaire/src/entities"

type LLMClient interface {
	LLMGenerateQlooSeed(inputs entities.JobInputs) (*entities.GeneratedSeed, error)
	LLMGenerateAdsCampaign(job *entities.AdsInsightsJob) (*entities.AdsCampaign, error)
}

const (
	seedPromptTemplate = `You are an AI assistant tasked with finding the single most culturally relevant reference for a user’s ad campaign.

	User’s ad platform: %s  
	Product description: "%s"  
	Product title: "%s"
	
	Requirements:
	1. Choose **one** seed entity (movie, videogame, book, artist, brand, etc.) **most closely aligned** with the product’s genre, tone, and target audience. It must be a well‑known example that directly reflects the user’s description.
	2. Return only the exact name of that entity.
	3. Also return its Qloo entity type from the following list:
	
	- urn:entity:actor  
	- urn:entity:album  
	- urn:entity:artist  
	- urn:entity:author  
	- urn:entity:book  
	- urn:entity:brand  
	- urn:entity:destination  
	- urn:entity:director  
	- urn:entity:locality  
	- urn:entity:movie  
	- urn:entity:person  
	- urn:entity:place  
	- urn:entity:podcast  
	- urn:entity:tv_show  
	- urn:entity:videogame  
	
	**Output format** (raw JSON only, no markdown):
	
	{
	  "query": "seed entity name here",
	  "type": "urn:entity:..."
	}`
	campaignPromptTemplate = `You are a senior digital marketing strategist.
	Client: generate a Meta Ads campaign for "%s" (%s).
	
	User Inputs:
	- Platform: %s
	- Product: %s
	- Title: %s
	
	Seed Entity:
	- Query: %s
	- Type: %s
	
	Top Related Entities (from Qloo Insights):
	%s
	
	Demographic Insights:
	%s
	
	Using the data above, produce a JSON object under the key "ads_campaign_result" with the following **strict** schema, without additional content:
	{
	  "ads_campaign_result": {
		"ad_copy": [
		  {
			"headline": "<string>",
			"description": "<string>"
		  }
		  // two or three relevant entries
		],
		"creative_concepts": [
		  {
			"type": "Image|Video",
			"description": "<string>",
			"elements": "<string>"
		  }
		  // two or three relevant entries
		],
		"persona_summary": {
		  "age": "<string>",
		  "gender": "<string>",
		  "behavior": "<string>",
		  "interests": "<string>"
		},
		"segmentation": {
		  "age": "<string>",
		  "gender": "<string>",
		  "behavior": "<string>",
		  "devices": "<string>",
		  "interests": "<string>",
		  "location": "<string>"
		},
		"campaign_config": {
		  "objective": "<string>",
		  "placements": "<string>",
		  "budget": "<string>",
		  "a_b_testing": [
			{
			  "test_name": "<string>",
			  "variants": "<string>"
			}
			// at least two relevant entries
		  ]
		},
		"key_insights": [
			"<string>"
		] // As many or few as you find relevant
	  }
	}`
)
