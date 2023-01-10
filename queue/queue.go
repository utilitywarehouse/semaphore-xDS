package queue

import (
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"github.com/utilitywarehouse/semaphore-xds/log"
	"github.com/utilitywarehouse/semaphore-xds/metrics"
)

// queueReconcileFunc reconciles the object indicated by the name and namespace
type queueReconcileFunc func(name, namespace string) error

// queue provides a rate-limited queue that processes items with a provided
// reconcile function
type Queue struct {
	name          string
	reconcileFunc queueReconcileFunc
	queue         workqueue.RateLimitingInterface
	requeued      []string
}

// newQueue returns a new queue
func NewQueue(name string, reconcileFunc queueReconcileFunc) *Queue {
	return &Queue{
		name:          name,
		reconcileFunc: reconcileFunc,
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), name),
	}
}

// Add an item to the queue, where that item is an object that
// implements meta.Interface.
func (q *Queue) Add(obj interface{}) {
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		log.Logger.Error("couldn't create object key", "queue", q.name, "err", err)
		return
	}
	q.queue.Add(key)
}

// Run processes items from the queue as they're added
func (q *Queue) Run() {
	q.updateMetrics()
	for q.processItem() {
		q.updateMetrics()
	}
}

// Stop causes the queue to shut down
func (q *Queue) Stop() {
	q.queue.ShutDown()
}

// processItem processes the next item in the queue
func (q *Queue) processItem() bool {
	key, shutdown := q.queue.Get()
	if shutdown {
		log.Logger.Info("queue shutdown", "queue", q.name)
		return false
	}
	defer q.queue.Done(key)

	namespace, name, err := cache.SplitMetaNamespaceKey(key.(string))
	if err != nil {
		log.Logger.Error(
			"error parsing key",
			"queue", q.name,
			"key", key.(string),
			"err", err,
		)
		q.forget(key)
		return true
	}

	log.Logger.Info(
		"reconciling item",
		"queue", q.name,
		"namespace", namespace,
		"name", name,
	)
	if err := q.reconcileFunc(name, namespace); err != nil {
		log.Logger.Error(
			"reconcile error",
			"queue", q.name,
			"namespace", namespace,
			"name", name,
			"err", err,
		)
		q.requeue(key)
		log.Logger.Info(
			"requeued item",
			"queue", q.name,
			"namespace", namespace,
			"name", name,
		)
	} else {
		log.Logger.Info(
			"successfully reconciled item",
			"queue", q.name,
			"namespace", namespace,
			"name", name,
		)
		q.forget(key)
	}

	return true
}

func (q *Queue) requeue(key interface{}) {
	q.queue.AddRateLimited(key)
	q.addRequeued(key.(string))
}

func (q *Queue) forget(key interface{}) {
	q.queue.Forget(key)
	q.removeRequeued(key.(string))
}

func (q *Queue) addRequeued(key string) {
	for _, k := range q.requeued {
		if k == key {
			return
		}
	}
	q.requeued = append(q.requeued, key)
}

func (q *Queue) removeRequeued(key string) {
	for i, k := range q.requeued {
		if k == key {
			q.requeued = append(q.requeued[:i], q.requeued[i+1:]...)
			break
		}
	}
}

func (q *Queue) updateMetrics() {
	metrics.SetRequeued(q.name, float64(len(q.requeued)))
}
