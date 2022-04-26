package pool

import (
	"errors"
	"fmt"
	"git.yongche.com/rabbitmq-channel/util"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

type Content struct {
	DeclareExchange bool
	Delay           int32
	Publishing      *amqp.Publishing
}

type Exchange struct {
	Name     string
	Type     string
	Durable  bool
	AutoDel  bool
	Internal bool
	NoWait   bool
}

type Queue struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       map[string]interface{}
}

type channel struct {
	connectId     int
	ch            *amqp.Channel
	notifyClose   chan *amqp.Error
	notifyConfirm chan amqp.Confirmation
	notifyReturn  chan amqp.Return
	id            int
	durable       bool
}

type connection struct {
	channelNum  int //记录同一个连接channel数量
	connect     *amqp.Connection
	id          int  //连接唯一标示
	durable     bool //用来区分创建的连接是池连接还是临时连接
	notifyClose chan *amqp.Error
}

type Options struct {
	ConnectNum    int    `mapstructure:"connect_num"` //配置connection数
	ChannelNum    int    `mapstructure:"channel_num"` //配置channel数
	Username      string `mapstructure:"username"`
	Password      string `mapstructure:"password"`
	Host          string `mapstructure:"host"`
	Port          int    `mapstructure:"port"`
	Vhost         string `mapstructure:"vhost"`
	TimeOut       int32  `mapstructure:"time_out"`
	RetryCount    int    `mapstructure:"retry_count"`
	MaxChannelNum int    `mapstructure:"max_channel_num"` //channel最大数量
}

type PoolService struct {
	AmqpUrl      string
	connections  map[int]*connection
	channels     map[int]*channel
	idleChannel  []int
	mutex        *sync.Mutex
	channelMaxId int
	connectMaxId int
	chanMutex    *sync.Mutex
	Options
}

func NewPoolService(options *Options) *PoolService {
	service := &PoolService{
		/*
			ConnectNum: connectNum,
			ChannelNum: channelNum,
			AmqpUrl:    amqpUrl,
		*/
		connections:  make(map[int]*connection),
		channels:     make(map[int]*channel),
		idleChannel:  []int{},
		mutex:        new(sync.Mutex),
		channelMaxId: 0,
		connectMaxId: 0,
		chanMutex:    new(sync.Mutex),
		Options:      *options,
	}

	service.init()
	return service
}

func (s *PoolService) BuildURL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		s.Username, s.Password, s.Host, s.Port, s.Vhost)
}

func (s *PoolService) init() {
	s.AmqpUrl = s.BuildURL()
	for i := 1; i <= s.ConnectNum; i++ {
		connection, err := s.createConnection(true)
		if err != nil {
			util.FailOnError(err, "connection fail quit init")
			return
		}
		for j := 1; j <= s.ChannelNum; j++ {
			_, err := s.createChannel(connection, true)
			if err != nil {
				util.FailOnError(err, "channel create fail quit init")
			}
		}
	}
}

/**
 * connection、channel健康检查
 */
func (s *PoolService) poolKeepCheck() {
	for {
		for _, connection := range s.connections {
			if connection.durable == false {
				connection.connect.Close()
			}
		}
	}
}

func (s *PoolService) createConnection(durable bool) (*connection, error) {
	connection, err := s.buildConnection(durable)
	if err != nil {
		return nil, err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.connectMaxId++
	connection.id = s.connectMaxId
	s.connections[s.connectMaxId] = connection

	return connection, nil
}

func (s *PoolService) buildConnection(durable bool) (*connection, error) {
	connection := new(connection)
	connection.notifyClose = make(chan *amqp.Error)
	conn, err := amqp.Dial(s.AmqpUrl)
	if err != nil {
		return nil, err
	}

	conn.NotifyClose(connection.notifyClose)
	connection.connect = conn
	connection.durable = durable
	go s.connectionClose(connection)
	return connection, nil
}

func (s *PoolService) createChannel(connect *connection, durable bool) (*channel, error) {

	s.mutex.Lock()
	defer s.mutex.Unlock()
	var cha = new(channel)

	cha.connectId = connect.id
	cha.durable = durable

	connect.channelNum++ //当前连接channel+1
	s.channelMaxId++
	cha.id = s.channelMaxId
	cha, err := s.buildChannel(connect, cha)
	if err != nil {
		return nil, err
	}

	s.channels[cha.id] = cha
	s.idleChannel = append(s.idleChannel, cha.id)
	return s.channels[cha.id], nil
}

func (s *PoolService) buildChannel(connection *connection, cha *channel) (*channel, error) {
	cha.notifyClose = make(chan *amqp.Error)
	cha.notifyConfirm = make(chan amqp.Confirmation)

	channel, err := connection.connect.Channel()
	if err != nil {
		util.FailOnError(err, "create channel fail")
		return nil, err
	}

	err = channel.Confirm(false)
	if err != nil {
		util.FailOnError(err, "createChannel confirm")
		return nil, err
	}

	channel.NotifyClose(cha.notifyClose)
	channel.NotifyPublish(cha.notifyConfirm)
	go s.channelClose(cha)

	cha.ch = channel
	return cha, nil
}

func (s *PoolService) getChannel() (*channel, error) {
	s.chanMutex.Lock()
	defer s.chanMutex.Unlock()

	for {
		if len(s.idleChannel) > 0 || len(s.channels) < s.Options.MaxChannelNum {
			break
		}
		time.Sleep(time.Nanosecond * 200)
	}

	//没有空闲channel并且当前最大channel数没有超过阈值
	if len(s.idleChannel) < 1 && len(s.channels) < s.Options.MaxChannelNum {
		var connection = new(connection)
		var err error
		for _, connect := range s.connections {
			//如果一次循环 connection.id == 0
			if connection.id == 0 && connect.connect.IsClosed() == false {
				connection = connect
			}

			//找个channel连接数最少的建立新的channel
			if connect.channelNum < connection.channelNum && connect.connect.IsClosed() == false {
				connection = connect
			}
		}

		if connection.id == 0 || connection.durable == false {
			connection, err = s.createConnection(false)
			if err != nil {
				util.FailOnError(err, "get channel createConnection fail")
				return nil, err
			}
		}
		_, err = s.createChannel(connection, false) //idleChannel 为空创建临时 channel
		if err != nil {
			util.FailOnError(err, "get channel createChannel fail")
			return nil, err
		}
		//fmt.Println("getChannel create channel: ", connection.channelNum)
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	channelId := s.idleChannel[0]
	fmt.Println("get channel idleChannel 2: ", s.idleChannel, " id: ", channelId)
	s.idleChannel = util.DeleteSlice(s.idleChannel, channelId)

	return s.channels[channelId], nil
}

func (s *PoolService) backChannel(channel *channel) {
	if channel == nil {
		return
	}

	//fmt.Println("back channel start lock")
	s.mutex.Lock()
	//fmt.Println("back channel waiting lock")
	defer s.mutex.Unlock()

	if channel.durable == false { //如果是临时连接用完关闭 并踢出全局
		_ = channel.ch.Close()
		if _, ok := s.connections[channel.connectId]; ok && s.connections[channel.connectId].durable == false {
			//delete(s.connections, channel.connectId)
			_ = s.connections[channel.connectId].connect.Close()
		}
	} else {
		s.idleChannel = append(s.idleChannel, channel.id)
	}
}

func (s *PoolService) channelClose(channel *channel) {
	for {
		err := <-channel.notifyClose
		s.mutex.Lock()
		defer s.mutex.Unlock()
		if channel.durable == false {
			return
		}
		//如果是持久的channel关闭后尝试重建channel
		if _, ok := s.connections[channel.connectId]; ok && channel.durable && s.connections[channel.connectId].connect.IsClosed() == false {
			_, channelErr := s.buildChannel(s.connections[channel.connectId], channel)
			if channelErr == nil {
				return
			}
		}

		fmt.Println("channel close delete channel id: ", channel.id)
		delete(s.channels, channel.id)
		util.DeleteSlice(s.idleChannel, channel.id)
		if err == nil {
			return
		}

		fmt.Printf("channel close channelId: %d  connectId: %d errCode: %d errReason: %s errSever: %t errRecover: %t \n",
			channel.id,
			channel.connectId,
			err.Code,
			err.Reason,
			err.Server,
			err.Recover,
		)

		return
	}
}

func (s *PoolService) connectionClose(connection *connection) {
	for {
		err := <-connection.notifyClose
		delete(s.connections, connection.id)
		if connection.durable == false {
			return
		}
		if err != nil {
			fmt.Printf("connection close Id: %d errCode: %d errReason: %s errSever: %t errRecover: %t \n",
				connection.id,
				err.Code,
				err.Reason,
				err.Server,
				err.Recover,
			)
		}
		//如果是持久连接被关闭尝试重建
		newConnect, connectErr := s.buildConnection(true)
		//fmt.Println("connection close recover start id: ", connection.id)
		if connectErr != nil {
			return
		}
		newConnect.id = connection.id
		//fmt.Println("connection close recover success id: ", connection.id)
		for i := 1; i <= s.ChannelNum; i++ {
			_, chanErr := s.createChannel(newConnect, true)
			if chanErr != nil {
				fmt.Println("connection close create channel fail err: ", chanErr)
				return
			}
		}
		s.connections[connection.id] = newConnect
		return
	}
}

func (s *PoolService) Publish(queue *Queue, exchange *Exchange, routeKey string, content *Content) (message interface{}, err error) {
	defer func() {
		var errStr string
		if p := recover(); p != nil {
			errStr = fmt.Sprintf("internal error: %v\n", p)
			err = errors.New(errStr)
			util.FailOnError(err, "publish recover")
		}
	}()

	channel, err := s.getChannel()

	if err != nil {
		util.FailOnError(err, "channel no available ")
		return nil, err
	}

	defer func() {
		s.backChannel(channel)
	}()

	_, err = s.DeclareQueue(channel.ch, queue)
	if err != nil {
		util.FailOnError(err, "declare queue fail")
		return
	}

	var exchangeName string
	//只有需要声明交换机时才对交换机和队列进行绑定
	if content.DeclareExchange {
		exchangeName = exchange.Name
		err = s.DeclareExchange(channel.ch, exchange)
		if err != nil {
			util.FailOnError(err, "declare exchange fail")
			return
		}

		err = s.queueBind(channel.ch, queue.Name, exchange.Name, routeKey)
		if err != nil {
			util.FailOnError(err, "queueBind fail")
			return
		}
	}

	retry := 1
	if s.TimeOut == 0 {
		s.TimeOut = 3
	}

	for {
		err = channel.ch.Publish(
			exchangeName, // exchange
			routeKey,     // routing key
			true,         // mandatory
			false,        // immediate
			*content.Publishing,
		)

		if err != nil {
			if retry <= s.RetryCount { //消息推送失败 重试次数
				retry++
				time.Sleep(time.Millisecond * 50)
				util.FailOnError(err, "pool publish fail continue msg: ")
				continue
			} else {
				util.FailOnError(err, "pool publish fail return msg: ")
				return
			}
		}

		select {
		case returns := <-channel.notifyReturn:
			if returns.ReplyCode != 0 {
				err = errors.New(returns.ReplyText)
				util.FailOnError(err, "notifyReturn")
				return nil, err
			}
			return nil, nil
		case confirm := <-channel.notifyConfirm:
			if confirm.Ack {
				return "success", nil
			}
			return nil, errors.New("publish noAck")

		case <-time.After(time.Duration(s.TimeOut) * time.Second):
			return nil, errors.New("time out ")
		}
	}
}

func (s *PoolService) DeclareExchange(channel *amqp.Channel, exchange *Exchange) error {
	err := channel.ExchangeDeclare(
		exchange.Name,     // name
		exchange.Type,     // type
		exchange.Durable,  // durable
		exchange.AutoDel,  // auto-deleted
		exchange.Internal, // internal
		exchange.NoWait,   // no-wait
		nil,               // arguments
	)

	return err
}

func (s *PoolService) DeclareQueue(channel *amqp.Channel, queue *Queue) (amqp.Queue, error) {
	q, err := channel.QueueDeclare(
		queue.Name,       // name
		queue.Durable,    // durable
		queue.AutoDelete, // delete when unused
		queue.Exclusive,  // exclusive
		queue.NoWait,     // no-wait
		queue.Args,       // arguments
	)

	if err != nil {
		util.FailOnError(err, "queueDeclare fail")
		return q, err
	}

	return q, nil
}

func (s *PoolService) queueBind(channel *amqp.Channel, queueName string, exchangeName string, routeKey string) error {
	err := channel.QueueBind(
		queueName,    // queue name
		routeKey,     // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	return err
}
