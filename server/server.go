package server

import (
	"errors"
	"fmt"
	pb "github.com/weackd/grpc-pubsub-broker/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"sync"
)

var availableTopics = map[string]string{
	"PFM": "PLATEFORMFILE_MAKER",
	"DFM": "DOCKERFILE_MAKER",
	"DPY": "DEPLOY",
	"GFI": "GETCAASFILE_INTEGRITY",
}

type ClientData struct {
	identity  *pb.Identity
	channel   chan *pb.Message
	mutex     sync.Mutex
	connected bool
}

type ClientRegistry struct {
	mutex   sync.Mutex
	clients []*ClientData
}

//clos le channel du client
func (this *ServerContext) Stop() {
	for _, client := range this.clients {
		client.mutex.Lock()
		close(client.channel)
		client.connected = false
		client.mutex.Unlock()
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func generateRandomString(size int) string {
	b := make([]rune, size)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//Récupère le nom du client - mieux vaut récupérer les infos en bdd
func (this *ClientRegistry) getClient(identity *pb.Identity) (*ClientData, error) {
	this.mutex.Lock()
	for _, client := range this.clients {
		if client.identity.Name == identity.Name {
			this.mutex.Unlock()
			return client, nil
		}
	}
	this.mutex.Unlock()
	return nil, errors.New("Client was not registered")
}

func (this *ClientRegistry) Register(identity *pb.Identity) {
	this.mutex.Lock()
	newClient := new(ClientData)
	newClient.identity = identity
	newClient.channel = make(chan *pb.Message)
	newClient.connected = false
	this.clients = append(this.clients, newClient)
	this.mutex.Unlock()
}

func (this *ClientRegistry) isRegistered(identity *pb.Identity) bool {
	this.mutex.Lock()
	for _, client := range this.clients {
		if client.identity.Name == identity.Name {
			this.mutex.Unlock()
			return true
		}
	}
	this.mutex.Unlock()
	return false
}

func (this *ClientRegistry) Unregister(identity *pb.Identity) error {
	this.mutex.Lock()
	for index, client := range this.clients {
		if client.identity.Name == identity.Name {
			this.clients = append(this.clients[:index], this.clients[index+1:]...)
			this.mutex.Unlock()
			return nil
		}
	}
	this.mutex.Unlock()
	return errors.New("Client was not registered")
}

type SubscriptionRegistry struct {
	topics map[string]*MessageTopic
	mutex  sync.Mutex
}

func isAvailableTopic(key string) bool {
	isAvailable := false
	for topicKey, _ := range availableTopics {
		if topicKey == key {
			isAvailable = true
			break
		}
	}

	return isAvailable
}

func (this *SubscriptionRegistry) getTopic(key string) (topic *MessageTopic) {

	this.mutex.Lock()
	if value, exist := this.topics[key]; exist && isAvailableTopic(key) {
		topic = value
	} else {
		topic = new(MessageTopic)
		this.topics[key] = topic
	}
	this.mutex.Unlock()
	return topic
}

type MessageTopic struct {
	subscriptions []*ClientData
	mutex         sync.Mutex
}

//Diffuse le message dans le channel du client
func (this *MessageTopic) Spread(message *pb.Message) {
	this.mutex.Lock()
	for _, client := range this.subscriptions {
		client.mutex.Lock()
		if client.connected == true {
			client.channel <- message
		}
		client.mutex.Unlock()
	}
	this.mutex.Unlock()
}

func (this *MessageTopic) isSubscribed(client *ClientData) bool {
	this.mutex.Lock()
	for _, subscribedClient := range this.subscriptions {
		if subscribedClient == client {
			this.mutex.Unlock()
			return true
		}
	}
	this.mutex.Unlock()
	return false
}

func (this *MessageTopic) Subscribe(client *ClientData) {
	this.mutex.Lock()
	this.subscriptions = append(this.subscriptions, client)
	this.mutex.Unlock()
}

func (this *MessageTopic) Unsubscribe(client *ClientData) error {
	this.mutex.Lock()
	for index, subscribedClient := range this.subscriptions {
		if subscribedClient == client {
			this.subscriptions = append(this.subscriptions[:index], this.subscriptions[index+1:]...)
			this.mutex.Unlock()
			return nil
		}
	}
	this.mutex.Unlock()
	return errors.New("Client was not subscribed to this topic")
}

type ServerContext struct {
	ClientRegistry
	SubscriptionRegistry
}

//TODO : Checker si le client est présent en bdd.
func (this *ServerContext) Authenticate(ctx context.Context, identity *pb.Identity) (*pb.Identity, error) {
	if identity.Name == "" {
		identity.Name = generateRandomString(8)
	}
	fmt.Printf("Authenticating %s \n", identity.Name)
	if this.isRegistered(identity) == false {
		this.Register(identity)
		fmt.Printf("Created new user %s \n", identity.Name)
	} else {
		fmt.Printf("Existing user %s \n", identity.Name)
	}
	return identity, nil
}

func (this *ServerContext) Subscribe(ctx context.Context, request *pb.SubscribeRequest) (*pb.Subscription, error) {
	client, err := this.getClient(request.Identity)
	if err != nil {
		return nil, errors.New("Client was not authenticated")
	}

	fmt.Printf("Subscribing %s to %s \n", request.Identity.Name, request.Subscription.Key)

	topic := this.getTopic(request.Subscription.Key)
	if topic.isSubscribed(client) == true {
		fmt.Printf("Error already subscribed %s \n", request.Identity.Name)
		return nil, errors.New("Client already subscribed to this key")
	}
	topic.Subscribe(client)
	return request.Subscription, nil
}

func (this *ServerContext) Unsubscribe(ctx context.Context, request *pb.SubscribeRequest) (*pb.Subscription, error) {
	client, err := this.getClient(request.Identity)
	if err != nil {
		return nil, errors.New("Client was not authenticated")
	}

	topic := this.getTopic(request.Subscription.Key)
	if topic.isSubscribed(client) == false {
		fmt.Printf("Error not subscribed %s \n", request.Identity.Name)
		return nil, errors.New("Client was not subscribed to this key")
	}
	topic.Unsubscribe(client)
	return request.Subscription, nil
}

func (this *ServerContext) Pull(identity *pb.Identity, stream pb.Subscriber_PullServer) error {
	client, err := this.getClient(identity)
	if err != nil {
		return errors.New("Client was not authenticated")
	}

	fmt.Printf("Opening stream for %s \n", identity.Name)

	client.mutex.Lock()
	client.connected = true
	client.mutex.Unlock()

	for msg := range client.channel {
		if err := stream.Send(msg); err != nil {
			return err
		}

	}

	client.mutex.Lock()
	client.connected = false
	client.mutex.Unlock()

	fmt.Printf("Closing stream for %s \n", identity.Name)

	return nil
}

//Publie le message aux client du topic
func (this *ServerContext) Publish(ctx context.Context, request *pb.PublishRequest) (*pb.PublishResponse, error) {
	topic := this.getTopic(request.Key)
	for _, message := range request.Messages {
		topic.Spread(message)
		fmt.Printf("Message : %s \n", message)
	}
	return &pb.PublishResponse{}, nil
}

type PubSubServer struct {
	context *ServerContext
	server  *grpc.Server
}

func (this *PubSubServer) Stop() {
	this.context.Stop()
	this.server.GracefulStop()
}

func (this *PubSubServer) Start(port string) {
	var opts []grpc.ServerOption

	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	this.server = grpc.NewServer(opts...)
	pb.RegisterSubscriberServer(this.server, this.context)
	pb.RegisterPublisherServer(this.server, this.context)
	err = this.server.Serve(lis)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
}

func NewPubSubServer() *PubSubServer {
	s := new(PubSubServer)
	s.context = newServerContext()
	return s
}

func newServerContext() *ServerContext {
	s := new(ServerContext)
	s.topics = make(map[string]*MessageTopic)
	return s
}
