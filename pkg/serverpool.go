package pkg

import "log"

type ServerPool struct {
	Servers []*Backend
}

func (sp *ServerPool) HealthCheck() {
	log.Println("Health check for server pool")
	for _, server := range sp.Servers {
		isAlive := server.CheckAlive()
		if !isAlive {
			log.Printf("Server: %s is inactive\n", server.URL.Host)
		}

		server.SetAlive(isAlive)
	}
	log.Println("Completed health check for server pool")
}
