## Correctly interpret `2.4.1 Read` to design a ProtoBuf Read API

```
2.4.1 Read

Our clients read relation tuples to display ACLs or group
membership to users, or to prepare for a subsequent write.

A read request specifies one or multiple tuplesets and an optional zookie.

Each tupleset specifies keys of a set of relation tuples. The
set can include a single tuple key, or all tuples with a given
object ID or userset in a namespace, optionally constrained
by a relation name.

With the tuplesets, clients can look up
a specific membership entry, read all entries in an ACL or
group, or look up all groups with a given user as a direct
member. All tuplesets in a read request are processed at a
single snapshot.
With the zookie, clients can request a read snapshot no
earlier than a previous write if the zookie from the write response is given in the read request, or at the same snapshot as
a previous read if the zookie from the earlier read response
is given in the subsequent request. If the request doesn’t
contain a zookie, Zanzibar will choose a reasonably recent
snapshot, possibly offering a lower-latency response than if
a zookie were provided.
Read results only depend on contents of relation tuples and
do not reflect userset rewrite rules. For example, even if the
viewer userset always includes the owner userset, reading
tuples with the viewer relation will not return tuples with
the owner relation. Clients that need to understand the effective userset can use the Expand API (§2.4.5).
```

## Interpretation

### How to read
- This interpretation tries to strictly follow the design
proposed in the paper, that covers client's actual use cases
- This interpretation walks through the paper's sentences stage by stage
- Each stage gathers facts and interprets them as Go code
- The final stage represents the complete interpretation as Go code

### Staging the facts

> Our clients read relation tuples to display ACLs or group
  membership to users, or to prepare for a subsequent write.

Clients use cases:
- read tuples for displaying
- read-modify-write
- list groups a user belongs to
  - group = object#relation

> A read request specifies one or multiple tuplesets
  and an optional zookie.

- new term: `tupleset` - a set of tuples
- one read queries for multiple **sets of tuples**
  - means one query for each tubleset
  - means multiple queries send in single read rpc
- one read operates on a single snapshot (zookie) across all queries
  
At this stage, typed request interpretation is as follows:
```go
type ReadRequest struct {
    // A list of queries where each query resolves to 
    // zero or multiple relation tuples.
    // Query => TupleSet
    Queries []struct {
        // to be continued
    }
    // Snapshot read across all Queries
    Zookie string
}
```

> Each tupleset specifies keys of a set of relation tuples. The
  set can include a single tuple key, or all tuples with a given
  object ID or userset in a namespace, optionally constrained
  by a relation name.


- new term: `key` - a key to a relation tuple
  - can be translated to `filter parameters`
- `Each tupleset specifies keys of a set of relation tuples.`
  - Could be translated to:
  - "Each `Query` specifies `filter parameters` to resolve to a `TupleSet`."
  - the `TupleSet` in the response is a deduplicated list (hash set) of tuples
- about TupleSet/`Query` structure:
  - has keys/`filter parameters` resolving to a set of relation tuples
  - `filter parameters` structure (keys):
    - has either single key (query => single tuple)
      - means exact match to single a RelationTuple
      - means a single `Query` resolves to only a single RelationTuple
    - or all tuples where... (range query => multiple tuples)
      - match objectID and optionally relation
        - `WHERE object_id=?`
        - `WHERE object_id=? AND relation=?`
      - match userset in namespace and optionally relation
        - `WHERE namespace=? AND object_id=?`
        - `WHERE namespace=? AND object_id=? AND relation=?`
    
At this stage, typed request interpretation is as follows:
```go
type ReadRequest struct {
    // A list of queries where each query resolves to 
    // zero or multiple relation tuples.
    // Query => TupleSet
    Queries []struct {
        // This is the "keys" as referred to.
        Keys []Key
    }
    // Snapshot read across all Queries
    Zookie string
}
```