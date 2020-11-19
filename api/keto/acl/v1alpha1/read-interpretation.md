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
- Each stage gathers facts and interprets them structurally
- The final stage should represent the complete interpretation of the chapter

So let's stage the facts!

### Stage A

> Our clients read relation tuples to display ACLs or group
  membership to users, or to prepare for a subsequent write.

Clients use cases:
- read tuples for displaying
- read-modify-write
- list groups a user belongs to
  - group = object#relation

### Stage B

> A read request specifies one or multiple tuplesets
  and an optional zookie.

- new term: `tupleset` - a set of tuples
- one read rpc can return **multiple** tuplesets
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
    // Snapshot across all Queries,
    // no ACLs created after given snapshot.
    Zookie string
}
```

### Stage C

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

The following presents the query
concepts introduced in this `Stage C`...

#### Querying

- query `filter parameters` aka. keys
- `filter parameters` resolve to a set of relation tuples
- each query can only read in a single namespace
- different queries can read different namespaces
- a query is one of the two types: single or range


##### "Single" query

The purpose of the "single" query is to only fetch a single
RelationTuple, which means this `Query` resolves to
only a single RelationTuple.

##### "Range" query

The purpose of the "range" query is to fetch multiple
RelationTuples. Additionally, the "range" query _type_
splits into two _kinds_.

- `object(subject)` - object is unknown & subject is known
- `subject(object)` - subject is unknown & object is known

> Both kinds have a function-like name.
> - object(subject) takes "subject parameters" to the tuple "range" query.
> - subject(object) takes "object parameters" to the tuple "range" query.

A RelationTuple as 3 parts:
- `object` - (the left part)
- `relation` - (the middle / connecting part)
- `subject` - (the right part)


**`object(subject)`**

List objects matching `subject` parameters:
- `namespace` required
- `relation` optional
- `subject_id` or `subject_set` required
- if `subject_set`
  - `subject_set_namespace` required
  - `subject_set_object` optional
  - `subject_set_relation` optional

**`subject(object)`**

List subjects matching `object` parameters:
- `namespace` required
- `relation` optional
- `object` required

At this stage, typed request interpretation is as follows:
```go
type ReadRequest struct {
    // A list of queries where each query resolves to 
    // zero or multiple relation tuples.
    // Query => TupleSet
    Queries []struct {
        // Only one field must be provided.
    	// Either Single or Range.

        // A query to fetch a single tuple
        // e.g. read-modify-write a single tuple.
        Single struct {
            Tuple RelationTuple
        }
        // A range query to fetch multiple tuples
        // e.g. read-modify-write multiple tuples
        // e.g. search / display tuples
        Range struct {
            // Only one field must be provided.
            // Either ByObject or ByUserSet.

            // Search all tuples matching
        	// object + optional relation
            // e.g. to list users in a group (reader#member@bob)
            ByObject struct {
                Object string
                // Optional constrain the relation
                Relation string
            }

            ByUserSet struct {
            }
        }
    }
    // Snapshot across all Queries,
    // no ACLs created after given snapshot.
    Zookie string
}
```