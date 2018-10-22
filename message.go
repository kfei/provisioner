package main

type messageRequestCoupon struct {
	Count int64 `json:"count"`
}

type messageAvailableCoupon struct {
	Count int64 `json:"count"`
}

// Message is for Websocket
type Message struct {
	Command         string                  `json:"command"`
	RequestCoupon   *messageRequestCoupon   `json:"requestCoupon,omitempty"`
	AvailableCoupon *messageAvailableCoupon `json:"availableCoupon,omitempty"`
}
