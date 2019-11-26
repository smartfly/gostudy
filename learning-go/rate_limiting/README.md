# golang rate limit

## [golang.org/x/time/rate](https://godoc.org/golang.org/x/time/rate)

Limiter has three main methods, Allow, Reserve, and Wait. Most callers should use Wait

Each of the three methods consumes a single token, They differ in their behavior when no tokens is available.

- Allow: if no token is available, Allow returns false.
- Reserve: if no token is available, Reserve returns a reservation for a future token and the amount of time the caller must wait before using it.
- Wait: if no token is available, Wait blocks until one can be obtained  or its associated ```context.Context``` is canceled.

# 参考资料

[Go如何实现HTTP请求限流](https://xiequan.info/go%E5%A6%82%E4%BD%95%E5%AE%9E%E7%8E%B0http%E8%AF%B7%E6%B1%82%E9%99%90%E6%B5%81/)

[Rate Limiting](https://github.com/golang/go/wiki/RateLimiting)

[Rate limit examples in Go](https://hustcat.github.io/rate-limit-example-in-go/)

[How to Rate Limit HTTP Requests](https://www.alexedwards.net/blog/how-to-rate-limit-http-requests)
