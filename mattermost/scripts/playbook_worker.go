package main

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Ganti dengan URL Incoming Webhook Mattermost
const mattermostWebhookURL = "http://localhost:8065/hooks/abcd1234"

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        DB:   0,
    })

    fmt.Println("Worker Playbook Queue started...")

    for {
        // Ambil batch event max 10 tiap 10 detik
        msgs, err := rdb.BLPop(ctx, 10*time.Second, "mm_playbook_events").Result()
        if err == nil && len(msgs) > 0 {
            batchProcess(msgs[1:]) // index 0 = key, index 1 dst = data
        }
    }
}

func batchProcess(msgs []string) {
    fmt.Printf("Processing batch of %d playbook events...\n", len(msgs))

    // Buat pesan gabungan
    message := "📢 **Playbook Events Processed**\n"
    for i, m := range msgs {
        message += fmt.Sprintf("- Event %d: %s\n", i+1, m)
    }

    // Kirim ke Mattermost via webhook
    payload := map[string]string{"text": message}
    jsonData, _ := json.Marshal(payload)

    resp, err := http.Post(mattermostWebhookURL, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Printf("Error sending notification: %v\n", err)
        return
    }
    defer resp.Body.Close()

    fmt.Printf("Notification sent to Mattermost (status: %s)\n", resp.Status)
}