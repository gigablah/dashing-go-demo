package jobs

import (
    "time"
    "math/rand"
    "github.com/gigablah/dashing-go"
)

type Sample struct{}

func (j *Sample) Work(send chan *dashing.Message) {
    ticker := time.NewTicker(1 * time.Second)
    var lastValuation, lastKarma, currentValuation, currentKarma int
    for {
        select {
        case <- ticker.C:
            lastValuation, currentValuation = currentValuation, rand.Intn(100)
            lastKarma, currentKarma = currentKarma, rand.Intn(200000)
            send <- &dashing.Message{map[string]interface{}{
                "id": "valuation",
                "current": currentValuation,
                "last": lastValuation,
            }}
            send <- &dashing.Message{map[string]interface{}{
                "id": "karma",
                "current": currentKarma,
                "last": lastKarma,
            }}
            send <- &dashing.Message{map[string]interface{}{
                "id": "synergy",
                "value": rand.Intn(100),
            }}
        }
    }
}

func init() {
    dashing.Register(&Sample{})
}
