package main

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/api-gateway/grpc_clients"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/messaging"
	"ride-sharing/shared/proto/driver"
)

var (
	connManager = messaging.NewConnectionManager()
)

func handleRidersWebsocket(w http.ResponseWriter, r *http.Request, rb *messaging.RabbitMQ) {
	conn, err := connManager.Upgrade(w, r)
	if err != nil {
		log.Printf("[ME] Websocket upgrade failed: %v", err)
		return
	}

	defer conn.Close()

	userID := r.URL.Query().Get("userID")
	if userID == "" {
		log.Println("[ME] No userID provided")
		return
	}

	// Add connection to manager
	connManager.Add(userID, conn)
	defer connManager.Remove(userID)

	// Initialize queue consumers
	queues := []string{
		messaging.NotifyDriverNoDriversFoundQueue,
		messaging.NotifyDriverAssignQueue,
		messaging.NotifyPaymentSessionCreatedQueue,
	}

	for _, q := range queues {
		consumer := messaging.NewQueueConsumer(rb, connManager, q)

		if err := consumer.Start(); err != nil {
			log.Printf("[ME] Failed to start consumer for queue: %s: err: %v", q, err)
			return
		}
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("[ME] Error reading message: %v", err)
			break
		}

		log.Printf("[ME] Received message: %s", message)
	}

}

func handleDriversWebsocket(w http.ResponseWriter, r *http.Request, rb *messaging.RabbitMQ) {
	conn, err := connManager.Upgrade(w, r)
	if err != nil {
		log.Printf("[ME] Websocket upgrade failed: %v", err)
		return
	}

	defer conn.Close()

	userID := r.URL.Query().Get("userID")
	if userID == "" {
		log.Println("[ME] No userID provided")
		return
	}

	packageSlug := r.URL.Query().Get("packageSlug")
	if packageSlug == "" {
		log.Println("[ME] No package slug provided")
		return
	}

	// Add connection to manager
	connManager.Add(userID, conn)

	ctx := r.Context()

	driverService, err := grpc_clients.NewDriverServiceClient()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		connManager.Remove(userID)
		driverService.Client.UnregisterDriver(ctx, &driver.RegisterDriverRequest{
			DriverID:    userID,
			PackageSlug: packageSlug,
		})

		driverService.Close()

		log.Println("[ME] Driver unregistered: ", userID)
	}()

	driverData, err := driverService.Client.RegisterDriver(ctx, &driver.RegisterDriverRequest{
		DriverID:    userID,
		PackageSlug: packageSlug,
	})

	if err != nil {
		log.Printf("[ME] Error registering driver: %v", err)
		return
	}

	// type Driver struct {
	// 	Id             string `json:"id"`
	// 	Name           string `json:"name"`
	// 	ProfilePicture string `json:"profilePicture"`
	// 	CarPlate       string `json:"carPlate"`
	// 	PackageSlug    string `json:"packageSlug"`
	// }

	// msg := contracts.WSMessage{
	// 	Type: "driver.cmd.register",
	// 	Data: driverData.Driver,
	// }

	// msg := contracts.WSMessage{
	// 	Type: "driver.cmd.register",
	// 	Data: Driver{
	// 		Id:             userID,
	// 		Name:           "John Doe",
	// 		ProfilePicture: util.GetRandomAvatar(1),
	// 		CarPlate:       "ABC123",
	// 		PackageSlug:    packageSlug,
	// 	},
	// }

	if err := connManager.SendMessage(userID, contracts.WSMessage{
		Type: contracts.DriverCmdRegister,
		Data: driverData.Driver,
	}); err != nil {
		log.Printf("[ME] Error sending message: %v", err)
		return
	}

	// Initialize queue consumers
	queues := []string{
		messaging.DriverCmdTripRequestQueue,
	}

	for _, q := range queues {
		consumer := messaging.NewQueueConsumer(rb, connManager, q)

		if err := consumer.Start(); err != nil {
			log.Printf("[ME] Failed to start consumer for queue: %s: err: %v", q, err)
			return
		}
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("[ME] Error reading message: %v", err)
			break
		}

		type driverMessage struct {
			Type string          `json:"type"`
			Data json.RawMessage `json:"data"`
		}

		var driverMsg driverMessage

		if err := json.Unmarshal(message, &driverMsg); err != nil {
			log.Printf("[ME] Error unmarshalling driver message: %v", err)
			continue
		}

		log.Printf("[ME] Received message: %s", message)

		// Handle the different message type
		switch driverMsg.Type {
		case contracts.DriverCmdLocation:
			// Handle driver location update in the future
			continue
		case contracts.DriverCmdTripAccept, contracts.DriverCmdTripDecline:
			// Forward the message to the RabbitMQ
			if err := rb.PublishMessage(ctx, driverMsg.Type, contracts.AmqpMessage{
				OwnerID: userID,
				Data:    driverMsg.Data,
			}); err != nil {
				log.Printf("[ME] Error publishing message to RabbitMQ: %v", err)
			}
		default:
			log.Printf("[Me] Unknown message type: %s", driverMsg.Type)
		}
	}

}
