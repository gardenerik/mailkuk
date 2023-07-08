package main

var Router map[string]Sender

func loadRouter(routes []Route) {
	Router = map[string]Sender{}

	for _, route := range routes {
		Router[route.Mail] = Sender{
			url:     route.Url,
			headers: route.Headers,
		}
	}
}
