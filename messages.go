//
// Copyright (C) 2017 Daisho Group - All Rights Reserved
//
// List of messages.

package derrors

// InvalidEntity message indicating that the associated entity cannot be validated.
const InvalidEntity = "invalid entity, check mandatory fields"
// MarshalError message to indicate errors with the json.Marshal operation.
const MarshalError = "cannot marshal structure"
// UnmarshalError message to indicate errors with the json.Unmarshal operation.
const UnmarshalError = "cannot unmarshal structure"
// IOError message to indicate errors reading or writing data to sockets or persistent storage.
const IOError = "I/O error"
// OpFail message to indicate that a complex operation has failed.
const OpFail = "operation failed"
// MissingRESTParameter message to indicate that a required parameter is missing.
const MissingRESTParameter = "missing rest parameter"
// InvalidCondition message to indicate that a expected condition is not satisfied.
const InvalidCondition = "invalid condition"
// HTTPConnectionError message to indicate that the communication with an external entity has failed.
const HTTPConnectionError = "HTTP connection error"

// MissingNetwork message indicating that the request requires a target network.
const MissingNetwork = "missing target network"
// MissingCluster message indicating that the request requires a target cluster.
const MissingCluster = "missing target cluster"
// MissingNode message indicating that the request requires a target node.
const MissingNode = "missing target node"
// MissingAppDesc message indicating that the request requires a target application descriptor.
const MissingAppDesc = "missing application descriptor"
// MissingAppInst message indicating that the request requires a target application instance.
const MissingAppInst = "missing application instance"

// NetworkDoesNotExists message indicating that the network cannot be found in the system.
const NetworkDoesNotExists = "network does not exists"
// NetworkAlreadyExists message indicating that the network already exists.
const NetworkAlreadyExists = "network already exists"

// ClusterDoesNotExists message indicating that the cluster cannot be found in the system.
const ClusterDoesNotExists = "cluster does not exits"
// ClusterAlreadyExists message indicating that the cluster already exists.
const ClusterAlreadyExists = "cluster already exits"
// ClusterAlreadyAttached message indicating that the cluster is already attached to a given network.
const ClusterAlreadyAttached = "cluster is already attached to network"
// ClusterNotAttachedToNetwork message indicating that the cluster is not attached to the given network.
const ClusterNotAttachedToNetwork = "cluster not attached to network"

// NodeDoesNotExists message indicating that the node cannot be found in the system.
const NodeDoesNotExists = "node does not exits"
// NodeAlreadyExists message indicating that the node already exists in the system.
const NodeAlreadyExists = "node already exits"
// NodeAlreadyAttachedToCluster message indicating that the node is already attached to a cluster.
const NodeAlreadyAttachedToCluster = "node already attached to cluster"
// NodeNotAttachedToCluster message indicating that the node is not attached to the given cluster.
const NodeNotAttachedToCluster = "node not attached to cluster"

// AppDescDoesNotExists message indicating that the application descriptor does not exists in the system.
const AppDescDoesNotExists = "application descriptor does not exists"
// AppDescAlreadyExists message indicating that the application descriptor already exists in the system.
const AppDescAlreadyExists = "application descriptor already exists"
// AppDescAlreadyAttached message indicating that the application descriptor is already attached to the given network.
const AppDescAlreadyAttached = "application descriptor already attached to network"
// AppDescNotAttached message indicating that the application descriptor is not attached to the given network.
const AppDescNotAttached = "application descriptor not attached to network"

// AppInstDoesNotExists message indicating that the application instance does not exists in the system.
const AppInstDoesNotExists = "application instance does not exists"
// AppInstAlreadyExists message indicating that the application instance already exists in the system.
const AppInstAlreadyExists = "application instance already exists"
// AppInstAlreadyAttached message indicating that the application instance is already attached to the given network.
const AppInstAlreadyAttached = "application instance already attached to network"
// AppInstNotAttachedToNetwork message indicating that the application instance is not attached to the given network.
const AppInstNotAttachedToNetwork = "application instance not attached to network"
