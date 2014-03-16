package jobs

import (
    "time"
    "math/rand"
    "github.com/gigablah/dashing-go"
)

type Buzzwords struct{
    words []map[string]interface{}
}

func (j *Buzzwords) Work(send chan *dashing.Message) {
    ticker := time.NewTicker(1 * time.Second)
    for {
        select {
        case <- ticker.C:
            for i := 0; i < len(j.words); i++ {
                if 1 < rand.Intn(3) {
                    value := j.words[i]["value"].(int)
                    j.words[i]["value"] = (value + 1) % 30
                }
            }
            send <- &dashing.Message{map[string]interface{}{
                "id": "buzzwords",
                "items": j.words,
            }}
        }
    }
}

func init() {
    dashing.Register(&Buzzwords{[]map[string]interface{}{
        map[string]interface{}{"label": "Paradigm shift", "value": 0},
        map[string]interface{}{"label": "Leverage", "value": 0},
        map[string]interface{}{"label": "Pivoting", "value": 0},
        map[string]interface{}{"label": "Turn-key", "value": 0},
        map[string]interface{}{"label": "Streamlininess", "value": 0},
        map[string]interface{}{"label": "Exit strategy", "value": 0},
        map[string]interface{}{"label": "Synergy", "value": 0},
        map[string]interface{}{"label": "Enterprise", "value": 0},
        map[string]interface{}{"label": "Web 2.0", "value": 0},
    }})
}
