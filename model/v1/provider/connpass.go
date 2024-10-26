package model

import "time"

type Event struct {
    EventID     int       `json:"event_id"`
    Title       string    `json:"title"`
    Catch       string    `json:"catch"`
    Description string    `json:"description"`
    EventURL    string    `json:"event_url"`
    StartedAt   time.Time `json:"started_at"`
    EndedAt     time.Time `json:"ended_at"`
    Limit       int       `json:"limit"`
    HashTag     string    `json:"hash_tag"`
    EventType   string    `json:"event_type"`
    Accepted    int       `json:"accepted"`
    Waiting     int       `json:"waiting"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type ConnpassResponse struct {
    ResultsReturned int     `json:"results_returned"`
    Events          []Event `json:"events"`
}