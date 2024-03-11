# Builder

When piecewise object construction is complicated, provide an API fo doing it succinctly

## Summary

- A builder is a separate component used for building an object
- To make builder fluent, return the receiver - allows chaining
- Different facets of an object can be built with different builders working in tandem via a common struct