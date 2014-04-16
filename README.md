apt-transport-s3-https
======================

APT transport which interacts with S3 buckets over HTTPS with
[s3gof3r](https://github.com/rlmcpherson/s3gof3r/tree/master/gof3r) for multipart/concurrency

TODO
----

- [ ] APT transport Message handling interface
- [ ] APT transport 1.2 capabilities: pipelining
- [ ] Multipart/concurrent S3 requests, using s3gof3r

Why?
---
Other S3 APT transports are available, but many have problems:
- Insecure storage of S3 credentials
- Unable to make use of EC2 IAM role credentials (s3access metadata endpoint)
- Use plain HTTP
- Non feature complete / non-functional

Notable Projects
---
- [depot](https://github.com/coderanger/depot)
- [repo_server](https://github.com/qur/repo_server)
- [s3gof3r](https://github.com/rlmcpherson/s3gof3r)

Honorable Mentions
---
- [@coderanger](http://github.com/coderanger) -- showed me an
  internal experimental Python implementation of this and provided me
  with the understanding of the stdio protocol an APT transport implements


Baby Gopher Project
---
[![baby-gopher](https://raw2.github.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)
