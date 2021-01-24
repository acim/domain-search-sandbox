# domain-search-sandbox

Small Go utilitiy to search for available domains

- uses simple DNS queries to check if domain may be unregistered
- additional check is necessary to check for real availability (some bulk search service may be used)
- demonstrates Go concurrency patterns

# todo

Instead of the hardcoded words, use [Unix words](<https://en.m.wikipedia.org/wiki/Words_(Unix)>).
