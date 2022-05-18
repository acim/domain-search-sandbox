# domain-search-sandbox

[![License](https://img.shields.io/badge/license-BSD--2--Clause--Patent-orange.svg)](https://github.com/acim/domain-search-sandbox/blob/main/LICENSE)

Small Go utilitiy to search for available domains

- uses simple DNS queries to check if domain may be unregistered
- additional check is necessary to check for real availability (some bulk search service may be used)
- demonstrates Go concurrency patterns

# todo

Instead of the hardcoded words, use [Unix words](<https://en.m.wikipedia.org/wiki/Words_(Unix)>).
