#ifndef __JAVA_MAPS_H
#define __JAVA_MAPS_H

#include "conn_tuple.h"                // for conn_tuple_t
#include "ktypes.h"                    // for bool
#include "map-defs.h"                  // for BPF_HASH_MAP, BPF_PROG_ARRAY
#include "protocols/tls/java/types.h"  // for MAX_MESSAGE_TYPE, connection_by_peer_key_t

/* A set (map from a key to a const bool value, we care only if the key exists in the map, and not its value) to
   mark if we've seen a specific java tls connection.
   Map size is set to 1 as javaTLS is optional, this will be overwritten to MaxTrackedConnections
   if javaTLS is enabled. */
BPF_HASH_MAP(java_tls_connections, conn_tuple_t, bool, 1)

/* map to correlate peer domain and port with the actual conn_tuple
   Map size is set to 1 as javaTLS is optional, this will be overwritten to MaxTrackedConnections
   if javaTLS is enabled. */
BPF_HASH_MAP(java_conn_tuple_by_peer, connection_by_peer_key_t, conn_tuple_t, 1)

/*
    Map used to store the sub programs used by eRPC mechanism
    This is done to avoid memory limitation when handling different operations sent via ioctl (eRPC) from our dd-java-agent
*/
BPF_PROG_ARRAY(java_tls_erpc_handlers, MAX_MESSAGE_TYPE)

#endif // __JAVA_MAPS_H
