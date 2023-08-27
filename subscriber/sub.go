package subscriber

import (
	"Delivery_Food/common"
	"Delivery_Food/component"
	"Delivery_Food/component/asyncjob"
	"Delivery_Food/pubsub"
	"context"
	"log"
)

type ConsumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type ConsumerEngine struct {
	appCtx component.AppContext
}

func NewConsumerEngine(appCtx component.AppContext) *ConsumerEngine {
	return &ConsumerEngine{appCtx}
}

func (engine *ConsumerEngine) Start() error {
	err := engine.startSubTopic(
		common.TopicUserLikeRestaurant,
		false,
		RunIncreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
	)
	if err != nil {
		return err
	}

	errDislike := engine.startSubTopic(
		common.TopicUserDislikeRestaurant,
		true,
		RunDecreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
	)
	if errDislike != nil {
		return err
	}

	return nil
}

type ConsumerGroup interface {
	Run(ctx context.Context) error
}

func (engine *ConsumerEngine) startSubTopic(topic pubsub.Topic,
	isConcurrent bool, consumerJobs ...ConsumerJob) error {
	c, _ := engine.appCtx.GetPubSub().Subscribe(context.Background(), topic)

	for _, item := range consumerJobs {
		log.Println("Start consumer job: ", item.Title)
	}

	getJobHandler := func(job *ConsumerJob,
		message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("Running job:", job.Title, "with message:", message)
			return job.Hld(ctx, message)
		}
	}

	go func() {
		defer common.AppRecover()
		for {
			select {
			case <-context.Background().Done():
				return
			case msg := <-c:
				if msg == nil {
					return
				}

				jobHdlArr := make([]asyncjob.Job, len(consumerJobs))

				for i := range consumerJobs {
					jobHdlArr[i] = asyncjob.NewJob(getJobHandler(&consumerJobs[i], msg))
				}

				group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

				if err := group.Run(context.Background()); err != nil {
					log.Println(err)
				}
			}
		}
	}()

	return nil
}
