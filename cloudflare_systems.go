package main

import (
    "fmt"
    "io/ioutil"
    "net"
    "crypto/tls"
    "os"
    "strings"
    "flag"
    "time"
    "math"
)

func main() {
    urlPtr := flag.String("url", "cloudflare-test.snehashish.workers.dev", "Enter the url. eg. www.google.com")
    nPtr := flag.Int("profile", 1, "Enter number of times to hit the url.")
    flag.Parse()

    url := *urlPtr
    n := *nPtr

    // Parsing url
    if url[:7] == "http://" {
        url = url[7:]
    }
    if url[:8] == "https://" {
        url = url[8:]
    }
    idx := strings.Index(url, "/")

    var server, filePath string
    if idx != -1 {
        server = url[:idx]
        filePath = url[idx:]
    } else {
        server = url
        filePath = "/"
    }

    var dialer = net.Dialer{
	       Timeout: time.Minute,
    }

    errors := make(map[string]struct{})
    medianTime := []int64{}
    missed := 0
    var totalTime, slowestTime, fastestTime int64
    totalTime = 0
    fastestTime = math.MaxInt64
    slowestTime = -1
    maxSize := -1
    minSize := math.MaxInt64

    for i := 0; i < n; i++ {
        // opening a tcp socket using tls handshake
        start := time.Now()
        conn, err := tls.DialWithDialer(&dialer, "tcp", server + ":https", nil)
        checkError(err)
        // sending the GET request using HTTP 1.1 protocol
        _, err = conn.Write([]byte("GET " + filePath + " HTTP/1.1\r\nHost: " + server + "\r\nConnection: close\r\n\r\n"))
        checkError(err)
        result, err := ioutil.ReadAll(conn)
        end := time.Now()
        checkError(err)
        elapsed := end.Sub(start).Milliseconds()
        totalTime += elapsed
        // for median we need the middle 2 times at max when n is even
        if i == n / 2 || i == (n / 2) + 1 {
            medianTime = append(medianTime, elapsed)
        }
        if slowestTime < elapsed {
            slowestTime = elapsed
        }
        if fastestTime > elapsed {
            fastestTime = elapsed
        }
        if maxSize < len(result) {
            maxSize = len(result)
        }
        if minSize > len(result) {
            minSize = len(result)
        }
        res := string(result)
        msg := strings.Split(res, "\n")[0]
        if !strings.Contains(strings.ToLower(msg), "200 ok") {
            missed += 1
            errors[msg[9:]] = struct{}{}
        }
        if i == 0 {
            fmt.Println(res)
        }
    }
    mean := float64(totalTime) / float64(n)
    // calculating median
    var median float64
    if n % 2 == 0 {
        median = (float64(medianTime[0]) + float64(medianTime[1])) / 2
    } else {
        median = float64(medianTime[0])
    }
    succeeded := ((n - missed) / n) * 100
    fmt.Println("Webpage:", *urlPtr)
    fmt.Println("The number of requests:", n)
    fmt.Println("The fastest time:", fastestTime, "ms")
    fmt.Println("The slowest time:", slowestTime, "ms")
    fmt.Println("The mean & median times: mean:", mean, "ms", ", median:", median, "ms")
    fmt.Println("The percentage of requests that succeeded:", succeeded, "%")
    if missed > 0 {
        fmt.Println("Error codes:")
        for key := range(errors) {
            fmt.Println(key)
        }
    }
    fmt.Println("The size in bytes of the smallest response:", minSize)
    fmt.Println("The size in bytes of the largest response:", maxSize)

    os.Exit(0)
}
func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
