package main

import (
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"practiceBundleID": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		"reading": [
		  {
			"bucketName": "Grammar Pronunciation",
			"isRecommended": true,
			"practiceActivitySpecs": {
			  "need": {
				"proficiencyID": "urn:aida_english:domain_model:grammar_eo:GLGR0231",
				"dimension": "grammarAndSyntax",
				"context": [
				  {
					"correct": [
					  "I should keep my current position"
					],
					"incorrect": [
					  "I should to keep my current position"
					],
					"tag": "urn:aida_english:domain_model:sub_eo:subject+must+verb_base"
				  }
				]
			  },
			  "activity": {
				"type": "AAIS2",
				"strategy": "grammar_drill",
				"variant": "read_speak"
			  },
			  "maxGSE": 63,
			  "maxInteractions": 5,
			  "accent": "en-US",
			  "practiceRecommendationID": "3fa85f64-5717-4562-b3fc-2c963f66afa6"
			}
		  }
		],
		"listening": [
		  {
			"bucketName": "Grammar Pronunciation",
			"isRecommended": true,
			"practiceActivitySpecs": {
			  "need": {
				"proficiencyID": "urn:aida_english:domain_model:grammar_eo:GLGR0231",
				"dimension": "grammarAndSyntax",
				"context": [
				  {
					"correct": [
					  "I should keep my current position"
					],
					"incorrect": [
					  "I should to keep my current position"
					],
					"tag": "urn:aida_english:domain_model:sub_eo:subject+must+verb_base"
				  }
				]
			  },
			  "activity": {
				"type": "AAIS2",
				"strategy": "grammar_drill",
				"variant": "read_speak"
			  },
			  "maxGSE": 63,
			  "maxInteractions": 5,
			  "accent": "en-US",
			  "practiceRecommendationID": "3fa85f64-5717-4562-b3fc-2c963f66afa6"
			}
		  }
		],
		"speaking": [
		  {
			"bucketName": "Grammar Pronunciation",
			"isRecommended": true,
			"practiceActivitySpecs": {
			  "need": {
				"proficiencyID": "urn:aida_english:domain_model:grammar_eo:GLGR0231",
				"dimension": "grammarAndSyntax",
				"context": [
				  {
					"correct": [
					  "I should keep my current position"
					],
					"incorrect": [
					  "I should to keep my current position"
					],
					"tag": "urn:aida_english:domain_model:sub_eo:subject+must+verb_base"
				  }
				]
			  },
			  "activity": {
				"type": "AAIS2",
				"strategy": "grammar_drill",
				"variant": "read_speak"
			  },
			  "maxGSE": 63,
			  "maxInteractions": 5,
			  "accent": "en-US",
			  "practiceRecommendationID": "3fa85f64-5717-4562-b3fc-2c963f66afa6"
			}
		  }
		],
		"writing": [
		  {
			"bucketName": "Grammar Pronunciation",
			"isRecommended": true,
			"practiceActivitySpecs": {
			  "need": {
				"proficiencyID": "urn:aida_english:domain_model:grammar_eo:GLGR0231",
				"dimension": "grammarAndSyntax",
				"context": [
				  {
					"correct": [
					  "I should keep my current position"
					],
					"incorrect": [
					  "I should to keep my current position"
					],
					"tag": "urn:aida_english:domain_model:sub_eo:subject+must+verb_base"
				  }
				]
			  },
			  "activity": {
				"type": "AAIS2",
				"strategy": "grammar_drill",
				"variant": "read_speak"
			  },
			  "maxGSE": 63,
			  "maxInteractions": 5,
			  "accent": "en-US",
			  "practiceRecommendationID": "3fa85f64-5717-4562-b3fc-2c963f66afa6"
			}
		  }
		],
		"grammar": [
		  {
			"bucketName": "Grammar Pronunciation",
			"isRecommended": true,
			"practiceActivitySpecs": {
			  "need": {
				"proficiencyID": "urn:aida_english:domain_model:grammar_eo:GLGR0231",
				"dimension": "grammarAndSyntax",
				"context": [
				  {
					"correct": [
					  "I should keep my current position"
					],
					"incorrect": [
					  "I should to keep my current position"
					],
					"tag": "urn:aida_english:domain_model:sub_eo:subject+must+verb_base"
				  }
				]
			  },
			  "activity": {
				"type": "AAIS2",
				"strategy": "grammar_drill",
				"variant": "read_speak"
			  },
			  "maxGSE": 63,
			  "maxInteractions": 5,
			  "accent": "en-US",
			  "practiceRecommendationID": "3fa85f64-5717-4562-b3fc-2c963f66afa6"
			}
		  }
		],
		"vocab": [
		  {
			"bucketName": "Grammar Pronunciation",
			"isRecommended": true,
			"practiceActivitySpecs": {
			  "need": {
				"proficiencyID": "urn:aida_english:domain_model:grammar_eo:GLGR0231",
				"dimension": "grammarAndSyntax",
				"context": [
				  {
					"correct": [
					  "I should keep my current position"
					],
					"incorrect": [
					  "I should to keep my current position"
					],
					"tag": "urn:aida_english:domain_model:sub_eo:subject+must+verb_base"
				  }
				]
			  },
			  "activity": {
				"type": "AAIS2",
				"strategy": "grammar_drill",
				"variant": "read_speak"
			  },
			  "maxGSE": 63,
			  "maxInteractions": 5,
			  "accent": "en-US",
			  "practiceRecommendationID": "3fa85f64-5717-4562-b3fc-2c963f66afa6"
			}
		  }
		]
	  }`))
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
