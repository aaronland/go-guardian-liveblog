package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/anaskhan96/soup"
)

func main() {

	var delay int
	var read_all bool
	var verbose bool

	flag.IntVar(&delay, "delay", 30, "The number of seconds to wait before fetching new updates")
	flag.BoolVar(&read_all, "read-all", false, "If true read all previous posts (written before following has begun)")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Parse one or more Guardian \"live blog\" URLs and read them aloud.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options] url(N) url(N)\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	ctx := context.Background()
	urls := flag.Args()

	cache := new(sync.Map)
	mu := new(sync.RWMutex)

	process(ctx, cache, mu, read_all, urls...)

	ticker := time.NewTicker(time.Duration(delay) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			process(ctx, cache, mu, true, urls...)
		}
	}

}

func process(ctx context.Context, cache *sync.Map, mu *sync.RWMutex, read bool, urls ...string) {

	read_title := false

	if len(urls) > 1 {
		read_title = true
	}
	
	for _, url := range urls {
		go handle_posts(ctx, cache, mu, read, read_title, url)
	}

}

func read_post(ctx context.Context, post string) {
	slog.Info(post)
	cmd := exec.Command("say", post)
	cmd.Run()
}

func handle_posts(ctx context.Context, cache *sync.Map, mu *sync.RWMutex, read bool, read_title bool, url string) {

	slog.Debug("Handle posts", "url", url, "read", read)

	title, posts, err := get_posts(ctx, url)

	if err != nil {
		slog.Error("Failed to retrieve posts", "url", url, "error", err)
		return
	}

	title_read := true

	mu.Lock()
	defer mu.Unlock()

	for _, p := range posts {

		_, exists := cache.LoadOrStore(p, true)

		if exists {
			continue
		}

		if read {

			if read_title && !title_read{
				read_post(ctx, title)
				title_read = false
			}

			read_post(ctx, p)
		}
	}
}

func get_posts(ctx context.Context, url string) (string, []string, error) {

	rsp, err := soup.Get(url)

	if err != nil {
		return "", nil, fmt.Errorf("Failed to retrieve %s, %w", url, err)
	}

	posts := make([]string, 0)

	doc := soup.HTMLParse(rsp)

	title_el := doc.Find("title")
	title := title_el.FullText()

	title_parts := strings.Split(title, " | ")
	title = title_parts[0]

	paras := doc.FindAll("p")

	for _, p := range paras {

		p_class := p.Attrs()["class"]

		if !strings.HasPrefix(p_class, "dcr-") {
			continue
		}

		posts = append(posts, p.FullText())
	}

	return title, posts, nil
}
