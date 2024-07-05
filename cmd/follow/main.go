package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/anaskhan96/soup"
)

func main() {

	var url string
	var delay int
	var read_all bool
	
	flag.StringVar(&url, "url", "", "The URL of the event being live blogged")
	flag.IntVar(&delay, "delay", 30, "The number of seconds to wait before fetching new updates")
	flag.BoolVar(&read_all, "read-all", false, "If true read all previous posts (written before following has begun)")
	
	flag.Parse()

	ctx := context.Background()

	cache := new(sync.Map)

	handle_posts(ctx, url, cache, read_all)

	ticker := time.NewTicker(time.Duration(delay) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			handle_posts(ctx, url, cache, true)
		}
	}

}

func read_post(ctx context.Context, post string) {
	slog.Info(post)
	cmd := exec.Command("say", post)
	cmd.Run()
}

func handle_posts(ctx context.Context, url string, cache *sync.Map, read bool) {

	posts, err := get_posts(ctx, url)

	if err != nil {
		slog.Error("Failed to retrieve posts", "url", url, "error", err)
		return
	}

	for _, p := range posts {

		_, exists := cache.LoadOrStore(p, true)

		if exists {
			continue
		}

		if read {
			read_post(ctx, p)
		}
	}
}

func get_posts(ctx context.Context, url string) ([]string, error) {

	slog.Info("Get posts", "url", url)
	rsp, err := soup.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve %s, %w", url, err)
	}

	posts := make([]string, 0)

	doc := soup.HTMLParse(rsp)
	paras := doc.FindAll("p")

	for _, p := range paras {

		p_class := p.Attrs()["class"]

		if !strings.HasPrefix(p_class, "dcr-") {
			continue
		}

		posts = append(posts, p.FullText())
	}

	return posts, nil
}
