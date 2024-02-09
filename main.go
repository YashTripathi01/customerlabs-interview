package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type OriginalRequest struct {
	Ev     string `json:"ev"`
	Et     string `json:"et"`
	Id     string `json:"id"`
	Uid    string `json:"uid"`
	Mid    string `json:"mid"`
	T      string `json:"t"`
	P      string `json:"p"`
	L      string `json:"l"`
	Sc     string `json:"sc"`
	Atrk1  string `json:"atrk1"`
	Atrv1  string `json:"atrv1"`
	Atrt1  string `json:"atrt1"`
	Atrk2  string `json:"atrk2"`
	Atrv2  string `json:"atrv2"`
	Atrt2  string `json:"atrt2"`
	Uatrk1 string `json:"uatrk1"`
	Uatrv1 string `json:"uatrv1"`
	Uatrt1 string `json:"uatrt1"`
	Uatrk2 string `json:"uatrk2"`
	Uatrv2 string `json:"uatrv2"`
	Uatrt2 string `json:"uatrt2"`
	Uatrk3 string `json:"uatrk3"`
	Uatrv3 string `json:"uatrv3"`
	Uatrt3 string `json:"uatrt3"`
}

type ConvertedRequest struct {
	Event           string               `json:"event"`
	EventType       string               `json:"event_type"`
	AppID           string               `json:"app_id"`
	UserID          string               `json:"user_id"`
	MessageID       string               `json:"message_id"`
	PageTitle       string               `json:"page_title"`
	PageURL         string               `json:"page_url"`
	BrowserLanguage string               `json:"browser_language"`
	ScreenSize      string               `json:"screen_size"`
	Attributes      map[string]Attribute `json:"attributes"`
	Traits          map[string]Trait     `json:"traits"`
}

type Attribute struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Trait struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

func main() {
	// Create a channel for receiving requests
	requestChannel := make(chan OriginalRequest)

	// Start worker
	go worker(requestChannel)

	// Handler for receiving requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var originalRequest OriginalRequest

		err := json.NewDecoder(r.Body).Decode(&originalRequest)
		if err != nil {
			http.Error(w, "Error decoding request", http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		// Send the request to the channel
		go func() {
			requestChannel <- originalRequest
		}()

		// Respond to the client
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Request received and processing...")
	})

	// Start the server
	http.ListenAndServe(":8080", nil)
}

func worker(requestChannel chan OriginalRequest) {
	for originalRequest := range requestChannel {
		// Convert the request to the desired format
		convertedRequest := convertRequest(originalRequest)

		// Send the converted request to the webhook
		err := sendToWebhook(convertedRequest)
		if err != nil {
			fmt.Println("Error sending to webhook:", err)
		}
	}
}

func convertRequest(originalRequest OriginalRequest) ConvertedRequest {
	convertedRequest := ConvertedRequest{
		Event:           originalRequest.Ev,
		EventType:       originalRequest.Et,
		AppID:           originalRequest.Id,
		UserID:          originalRequest.Uid,
		MessageID:       originalRequest.Mid,
		PageTitle:       originalRequest.T,
		PageURL:         originalRequest.P,
		BrowserLanguage: originalRequest.L,
		ScreenSize:      originalRequest.Sc,
		Attributes:      make(map[string]Attribute),
		Traits:          make(map[string]Trait),
	}

	// Attributes
	attributes := map[string]string{
		originalRequest.Atrk1: originalRequest.Atrv1,
		originalRequest.Atrk2: originalRequest.Atrv2,
	}

	for key, value := range attributes {
		if key != "" {
			convertedRequest.Attributes[key] = Attribute{
				Value: value,
				Type:  "string",
			}
		}
	}

	// Traits
	traits := map[string]string{
		originalRequest.Uatrk1: originalRequest.Uatrv1,
		originalRequest.Uatrk2: originalRequest.Uatrv2,
		originalRequest.Uatrk3: originalRequest.Uatrv3,
	}

	for key, value := range traits {
		if key != "" {
			convertedRequest.Traits[key] = Trait{
				Value: value,
				Type:  "string",
			}
		}
	}

	return convertedRequest
}

func sendToWebhook(convertedRequest ConvertedRequest) error {
	webhookURL := "https://webhook.site/df3cd4ff-4e3c-4590-8f5d-92f32f49439d"

	requestBody, err := json.Marshal(convertedRequest)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return nil
}
